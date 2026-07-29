package directworker

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

type Manager struct {
	mu             sync.Mutex
	workers        map[string]*accountWorker
	statuses       map[string]AccountStatus
	nextGeneration uint64
	out            chan<- model.DirectMessage
	readOut        chan<- model.DirectReadReceipt
	sent           chan<- model.DirectSent
	logger         *log.Logger
	clientFactory  directClientFactory
	sleepBackoff   func(context.Context, *time.Duration)
	now            func() time.Time
}

type RuntimeAccount struct {
	Config config.AccountConfig
	Token  string
}

type accountWorker struct {
	account    RuntimeAccount
	generation uint64
	queue      chan model.DirectOutbound
	cancel     context.CancelFunc
	restart    chan struct{}
	status     func(func(*AccountStatus))
	names      *nameResolver
}

type AccountStatus struct {
	AccountID    string    `json:"account_id"`
	Generation   uint64    `json:"generation"`
	Running      bool      `json:"running"`
	Connected    bool      `json:"connected"`
	Ready        bool      `json:"ready"`
	StartedAt    time.Time `json:"started_at"`
	UnreadySince time.Time `json:"unready_since,omitempty"`
	UpdatedAt    time.Time `json:"updated_at"`
	ReadyAt      time.Time `json:"ready_at,omitempty"`
	RestartedAt  time.Time `json:"restarted_at,omitempty"`
	Restarts     int       `json:"restarts"`
	AuthInvalid  bool      `json:"auth_invalid"`
	LastError    string    `json:"last_error,omitempty"`
}

type directClient interface {
	Connect() error
	ConnectWithContext(context.Context) error
	Close() error
	On(string, direct.EventHandler)
	OnMessage(func(direct.ReceivedMessage))
	CreateTextMessageWithContext(context.Context, string, string) (string, error)
	CreateUploadAuth(context.Context, string, string, int64, string) (*direct.UploadAuth, error)
	CallWithContext(context.Context, string, []interface{}) (interface{}, error)
	Done() <-chan struct{}
}

type directClientFactory func(direct.Options) directClient

type productionDirectClient struct {
	*direct.Client
}

func (c productionDirectClient) Done() <-chan struct{} {
	return c.Client.Done
}

func NewManager(out chan<- model.DirectMessage, readOut chan<- model.DirectReadReceipt, sent chan<- model.DirectSent, logger *log.Logger) *Manager {
	return &Manager{
		workers:       map[string]*accountWorker{},
		out:           out,
		readOut:       readOut,
		sent:          sent,
		logger:        logger,
		clientFactory: newProductionDirectClient,
		sleepBackoff:  sleepWithBackoff,
		statuses:      map[string]AccountStatus{},
		now:           time.Now,
	}
}

func newProductionDirectClient(opts direct.Options) directClient {
	return productionDirectClient{Client: direct.NewClient(opts)}
}

func (m *Manager) Send(ctx context.Context, msg model.DirectOutbound) error {
	m.mu.Lock()
	worker, ok := m.workers[msg.AccountID]
	var ch chan model.DirectOutbound
	if ok {
		ch = worker.queue
	}
	m.mu.Unlock()
	if !ok {
		return fmt.Errorf("unknown account %q", msg.AccountID)
	}
	select {
	case ch <- msg:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (m *Manager) Statuses() []AccountStatus {
	m.mu.Lock()
	defer m.mu.Unlock()
	out := make([]AccountStatus, 0, len(m.statuses))
	for _, status := range m.statuses {
		out = append(out, status)
	}
	return out
}

func (m *Manager) Healthy() (bool, []AccountStatus) {
	statuses := m.Statuses()
	for _, status := range statuses {
		if !status.Running || !status.Connected || !status.Ready {
			return false, statuses
		}
	}
	return true, statuses
}

func (m *Manager) StartWatchdog(ctx context.Context, interval, startupTimeout time.Duration) {
	if interval <= 0 {
		interval = 10 * time.Second
	}
	if startupTimeout <= 0 {
		startupTimeout = 2 * time.Minute
	}
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			m.restartStaleWorkers(startupTimeout)
		}
	}
}

func (m *Manager) Apply(ctx context.Context, accounts []RuntimeAccount) {
	next := map[string]RuntimeAccount{}
	for _, account := range accounts {
		next[account.Config.ID] = account
	}

	m.mu.Lock()
	defer m.mu.Unlock()
	for id, worker := range m.workers {
		account, ok := next[id]
		if !ok || !sameRuntimeAccount(worker.account, account) {
			m.logger.Printf("[%s] stopping direct worker", debuglog.RedactID(id))
			worker.cancel()
			delete(m.workers, id)
			delete(m.statuses, id)
		}
	}
	for id, account := range next {
		if _, ok := m.workers[id]; ok {
			continue
		}
		m.nextGeneration++
		generation := m.nextGeneration
		workerCtx, cancel := context.WithCancel(ctx)
		worker := &accountWorker{
			account:    account,
			generation: generation,
			queue:      make(chan model.DirectOutbound, 100),
			cancel:     cancel,
			restart:    make(chan struct{}, 1),
		}
		worker.status = func(update func(*AccountStatus)) {
			m.updateStatus(id, generation, update)
		}
		m.workers[id] = worker
		now := m.now()
		m.statuses[id] = AccountStatus{AccountID: id, Generation: generation, StartedAt: now, UnreadySince: now, UpdatedAt: now}
		m.logger.Printf("[%s] starting direct worker", debuglog.RedactID(id))
		go runAccountWorker(workerCtx, worker, m.out, m.readOut, m.sent, m.logger, m.clientFactory, m.sleepBackoff, m.now)
	}
}

func (m *Manager) restartStaleWorkers(startupTimeout time.Duration) {
	now := m.now()
	m.mu.Lock()
	defer m.mu.Unlock()
	for id, worker := range m.workers {
		status := m.statuses[id]
		if status.AuthInvalid {
			continue
		}
		if status.Ready {
			continue
		}
		unreadySince := status.UnreadySince
		if unreadySince.IsZero() {
			unreadySince = status.StartedAt
		}
		if unreadySince.IsZero() || now.Sub(unreadySince) < startupTimeout {
			continue
		}
		m.logger.Printf("[%s] direct worker is not ready after %s; requesting restart", debuglog.RedactID(id), startupTimeout)
		select {
		case worker.restart <- struct{}{}:
			status.RestartedAt = now
			status.UpdatedAt = now
			status.LastError = fmt.Sprintf("not ready after %s", startupTimeout)
			m.statuses[id] = status
		default:
		}
	}
}

func (m *Manager) updateStatus(accountID string, generation uint64, update func(*AccountStatus)) {
	m.mu.Lock()
	defer m.mu.Unlock()
	worker, ok := m.workers[accountID]
	if !ok || worker.generation != generation {
		return
	}
	status := m.statuses[accountID]
	if status.AccountID == "" {
		status.AccountID = accountID
	}
	status.Generation = generation
	update(&status)
	status.UpdatedAt = m.now()
	m.statuses[accountID] = status
}

func sameRuntimeAccount(a, b RuntimeAccount) bool {
	return a.Token == b.Token &&
		a.Config.ID == b.Config.ID &&
		a.Config.Endpoint == b.Config.Endpoint &&
		a.Config.ProxyURL == b.Config.ProxyURL &&
		a.Config.TeamsChannel == b.Config.TeamsChannel
}

func markUnready(status *AccountStatus, now time.Time) {
	if status.UnreadySince.IsZero() {
		status.UnreadySince = now
	}
}

func runAccountWorker(ctx context.Context, worker *accountWorker, out chan<- model.DirectMessage, readOut chan<- model.DirectReadReceipt, sent chan<- model.DirectSent, logger *log.Logger, clientFactory directClientFactory, sleepBackoff func(context.Context, *time.Duration), now func() time.Time) {
	account := worker.account
	in := worker.queue
	cfg := account.Config
	backoff := time.Second
	var pending *model.DirectOutbound
	if worker.names == nil {
		worker.names = newNameResolver()
	}
	worker.status(func(status *AccountStatus) {
		status.Running = true
		status.StartedAt = now()
		status.Connected = false
		status.Ready = false
		markUnready(status, now())
		status.AuthInvalid = false
		status.LastError = ""
	})
	defer worker.status(func(status *AccountStatus) {
		status.Running = false
		status.Connected = false
		status.Ready = false
		markUnready(status, now())
	})
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		token := account.Token
		if token == "" {
			logger.Printf("[%s] direct token is empty; retrying", debuglog.RedactID(cfg.ID))
			worker.status(func(status *AccountStatus) {
				status.LastError = "direct token is empty"
			})
			sleepBackoff(ctx, &backoff)
			continue
		}
		worker.status(func(status *AccountStatus) {
			// StartedAt tracks the current connection attempt; UnreadySince tracks
			// the continuous period without a ready notification.
			status.StartedAt = now()
			status.Connected = false
			status.Ready = false
			markUnready(status, now())
			status.AuthInvalid = false
			status.LastError = ""
		})
		logger.Printf("[%s] using direct token configured", debuglog.RedactID(cfg.ID))

		endpoint := cfg.Endpoint
		if endpoint == "" {
			endpoint = direct.DefaultEndpoint
		}
		client := clientFactory(direct.Options{
			AccessToken: token,
			Endpoint:    endpoint,
			ProxyURL:    cfg.ProxyURL,
			Name:        cfg.ID,
		})
		reconnectNow := make(chan struct{}, 1)
		readyNow := make(chan struct{}, 1)
		requestReconnect := func() {
			select {
			case reconnectNow <- struct{}{}:
			default:
			}
		}
		markReady := func() {
			select {
			case readyNow <- struct{}{}:
			default:
			}
		}

		client.On(direct.EventError, func(data interface{}) {
			logger.Printf("[%s] direct error: %s", debuglog.RedactID(cfg.ID), debuglog.SummarizePayload(data))
			worker.status(func(status *AccountStatus) {
				status.LastError = debuglog.SummarizePayload(data)
			})
			requestReconnect()
		})
		client.On(direct.EventSessionError, func(data interface{}) {
			logger.Printf("[%s] direct session error: %s", debuglog.RedactID(cfg.ID), debuglog.SummarizePayload(data))
			worker.status(func(status *AccountStatus) {
				status.LastError = debuglog.SummarizePayload(data)
			})
			if isInvalidTokenError(data) {
				logger.Printf("[%s] direct token is invalid; waiting for token update before reconnect", debuglog.RedactID(cfg.ID))
				worker.status(func(status *AccountStatus) {
					status.AuthInvalid = true
					status.Connected = false
					status.Ready = false
					markUnready(status, now())
				})
				return
			}
			requestReconnect()
		})
		client.On(direct.EventSessionCreated, func(data interface{}) {
			logger.Printf("[%s] direct session created", debuglog.RedactID(cfg.ID))
		})
		client.On(direct.EventDataRecovered, func(data interface{}) {
			logger.Printf("[%s] direct notification ready", debuglog.RedactID(cfg.ID))
			worker.status(func(status *AccountStatus) {
				status.Ready = true
				status.ReadyAt = now()
				status.UnreadySince = time.Time{}
				status.LastError = ""
			})
			markReady()
		})
		client.On(direct.EventNotificationError, func(data interface{}) {
			logger.Printf("[%s] direct notification error: %s", debuglog.RedactID(cfg.ID), debuglog.SummarizePayload(data))
			worker.status(func(status *AccountStatus) {
				status.LastError = debuglog.SummarizePayload(data)
			})
			requestReconnect()
		})
		client.OnMessage(func(msg direct.ReceivedMessage) {
			if msg.ID == "" {
				logger.Printf("[%s] direct message ignored without id: talk=%s user=%s type=%d", debuglog.RedactID(cfg.ID), debuglog.RedactID(msg.TalkID), debuglog.RedactID(msg.UserID), msg.Type)
				return
			}
			bm := model.DirectMessage{
				AccountID: cfg.ID,
				TalkID:    msg.TalkID,
				UserID:    msg.UserID,
				UserName:  worker.names.userName(ctx, client, msg.DomainID, msg.UserID, msg.TalkID, logger, cfg.ID),
				RoomName:  worker.names.roomName(ctx, client, msg.TalkID, logger, cfg.ID),
				Text:      msg.Text,
				MessageID: msg.ID,
				CreatedAt: messageTime(msg),
				Raw:       msg,
			}
			bm.Attachments = attachmentsFromDirectMessage(msg)
			logger.Printf("[%s] direct message received: id=%s talk=%s user=%s type=%d attachments=%d", debuglog.RedactID(cfg.ID), debuglog.RedactID(msg.ID), debuglog.RedactID(msg.TalkID), debuglog.RedactID(msg.UserID), msg.Type, len(bm.Attachments))
			select {
			case out <- bm:
			case <-ctx.Done():
			}
		})
		client.On(direct.EventNotifyUpdateReadStatuses, func(data interface{}) {
			update := direct.ParseReadStatusesUpdate(data)
			if update.TalkID == "" || len(update.MessageIDs) == 0 {
				return
			}
			receipt := model.DirectReadReceipt{
				AccountID:   cfg.ID,
				TalkID:      update.TalkID,
				MessageIDs:  update.MessageIDs,
				ReadUserIDs: update.ReadUserIDs,
			}
			logger.Printf("[%s] direct read status updated: talk=%s messages=%d read_users=%d", debuglog.RedactID(cfg.ID), debuglog.RedactID(receipt.TalkID), len(receipt.MessageIDs), len(receipt.ReadUserIDs))
			if readOut == nil {
				return
			}
			select {
			case readOut <- receipt:
			case <-ctx.Done():
			}
		})

		if err := client.ConnectWithContext(ctx); err != nil {
			logger.Printf("[%s] connect failed: %s", debuglog.RedactID(cfg.ID), debuglog.SummarizePayload(err))
			worker.status(func(status *AccountStatus) {
				status.Connected = false
				status.Ready = false
				markUnready(status, now())
				status.AuthInvalid = isInvalidTokenError(err)
				status.LastError = debuglog.SummarizePayload(err)
			})
			_ = client.Close()
			if isInvalidTokenError(err) {
				logger.Printf("[%s] direct token is invalid; waiting for token update before retrying", debuglog.RedactID(cfg.ID))
				<-ctx.Done()
				return
			}
			sleepBackoff(ctx, &backoff)
			continue
		}
		logger.Printf("[%s] connected", debuglog.RedactID(cfg.ID))
		selfUserID := directSelfUserID(ctx, client)
		if selfUserID == "" {
			logger.Printf("[%s] direct self user id unavailable; read reactions will ignore self-read filtering", debuglog.RedactID(cfg.ID))
		}
		worker.status(func(status *AccountStatus) {
			status.Connected = true
			status.LastError = ""
		})

		disconnected := false
		for !disconnected {
			if pending != nil {
				if selfUserID == "" {
					selfUserID = directSelfUserID(ctx, client)
				}
				messageID, err := sendDirect(ctx, client, *pending)
				if err != nil {
					logger.Printf("[%s] direct send failed: talk=%s err=%s", debuglog.RedactID(cfg.ID), debuglog.RedactID(pending.TalkID), debuglog.SummarizePayload(err))
					worker.status(func(status *AccountStatus) {
						status.LastError = debuglog.SummarizePayload(err)
					})
					if isRecoverableDirectError(err) {
						logger.Printf("[%s] recreating client after recoverable direct send error", debuglog.RedactID(cfg.ID))
						_ = client.Close()
						disconnected = true
						continue
					}
					select {
					case sent <- model.DirectSent{Outbound: *pending, SenderID: selfUserID, Err: err}:
					case <-ctx.Done():
					}
					pending = nil
					continue
				}
				worker.status(func(status *AccountStatus) {
					status.LastError = ""
				})
				select {
				case sent <- model.DirectSent{Outbound: *pending, MessageID: messageID, SenderID: selfUserID}:
				case <-ctx.Done():
				}
				pending = nil
				continue
			}

			select {
			case <-ctx.Done():
				logger.Printf("[%s] shutting down", debuglog.RedactID(cfg.ID))
				_ = client.Close()
				return
			case <-client.Done():
				logger.Printf("[%s] disconnected; recreating client", debuglog.RedactID(cfg.ID))
				worker.status(func(status *AccountStatus) {
					status.Connected = false
					status.Ready = false
					markUnready(status, now())
					status.LastError = "disconnected"
				})
				_ = client.Close()
				disconnected = true
			case <-reconnectNow:
				logger.Printf("[%s] recreating client after direct startup error", debuglog.RedactID(cfg.ID))
				worker.status(func(status *AccountStatus) {
					status.Connected = false
					status.Ready = false
					markUnready(status, now())
				})
				_ = client.Close()
				disconnected = true
			case <-worker.restart:
				logger.Printf("[%s] recreating client after watchdog restart request", debuglog.RedactID(cfg.ID))
				worker.status(func(status *AccountStatus) {
					status.Connected = false
					status.Ready = false
					markUnready(status, now())
					status.Restarts++
					status.RestartedAt = now()
				})
				_ = client.Close()
				disconnected = true
			case <-readyNow:
				if selfUserID == "" {
					selfUserID = directSelfUserID(ctx, client)
				}
				backoff = time.Second
			case msg := <-in:
				pending = &msg
			}
		}
		sleepBackoff(ctx, &backoff)
	}
}

type nameResolver struct {
	mu    sync.Mutex
	users map[string]string
	rooms map[string]directRoomInfo
}

type directRoomInfo struct {
	name     string
	domainID string
}

func newNameResolver() *nameResolver {
	return &nameResolver{
		users: map[string]string{},
		rooms: map[string]directRoomInfo{},
	}
}

func (r *nameResolver) userName(ctx context.Context, client directClient, domainID, userID, talkID string, logger *log.Logger, accountID string) string {
	if userID == "" {
		return ""
	}
	if domainID == "" && talkID != "" {
		domainID = r.roomInfo(ctx, client, talkID, logger, accountID).domainID
	}
	if domainID == "" {
		logger.Printf("[%s] direct user name lookup skipped without domain: talk=%s user=%s", debuglog.RedactID(accountID), debuglog.RedactID(talkID), debuglog.RedactID(userID))
		return ""
	}
	key := domainID + ":" + userID
	r.mu.Lock()
	name, ok := r.users[key]
	r.mu.Unlock()
	if ok {
		return name
	}

	result, err := client.CallWithContext(ctx, direct.MethodGetUsers, []interface{}{normalizeRPCID(domainID), []interface{}{normalizeRPCID(userID)}})
	if err != nil {
		logger.Printf("[%s] direct user name lookup failed: domain=%s user=%s err=%s", debuglog.RedactID(accountID), debuglog.RedactID(domainID), debuglog.RedactID(userID), debuglog.SummarizePayload(err))
		return ""
	}
	name = displayNameFromGetUsers(result, userID)
	r.mu.Lock()
	r.users[key] = name
	r.mu.Unlock()
	return name
}

func (r *nameResolver) roomName(ctx context.Context, client directClient, talkID string, logger *log.Logger, accountID string) string {
	return r.roomInfo(ctx, client, talkID, logger, accountID).name
}

func (r *nameResolver) roomInfo(ctx context.Context, client directClient, talkID string, logger *log.Logger, accountID string) directRoomInfo {
	if talkID == "" {
		return directRoomInfo{}
	}
	r.mu.Lock()
	info, ok := r.rooms[talkID]
	r.mu.Unlock()
	if ok {
		return info
	}

	result, err := client.CallWithContext(ctx, direct.MethodGetTalks, []interface{}{})
	if err != nil {
		logger.Printf("[%s] direct room name lookup failed: talk=%s err=%s", debuglog.RedactID(accountID), debuglog.RedactID(talkID), debuglog.SummarizePayload(err))
		return directRoomInfo{}
	}
	rooms := roomInfoFromGetTalks(result)
	r.mu.Lock()
	for id, room := range rooms {
		r.rooms[id] = room
	}
	info = r.rooms[talkID]
	if _, ok := r.rooms[talkID]; !ok {
		r.rooms[talkID] = directRoomInfo{}
	}
	r.mu.Unlock()
	if info.name == "" && info.domainID == "" {
		logger.Printf("[%s] direct room lookup returned no match: talk=%s talks=%d", debuglog.RedactID(accountID), debuglog.RedactID(talkID), len(rooms))
	}
	return info
}

func displayNameFromGetUsers(result interface{}, userID string) string {
	arr, ok := result.([]interface{})
	if !ok {
		return ""
	}
	for _, item := range arr {
		user, ok := stringKeyMap(item)
		if !ok {
			continue
		}
		if id, ok := user["id"]; ok && normalizeID(id) != userID {
			continue
		}
		if id, ok := user["user_id"]; ok && normalizeID(id) != userID {
			continue
		}
		if name := stringMapValue(user, "display_name"); name != "" {
			return name
		}
		if name := stringMapValue(user, "name"); name != "" {
			return name
		}
	}
	return ""
}

func roomInfoFromGetTalks(result interface{}) map[string]directRoomInfo {
	out := map[string]directRoomInfo{}
	arr, ok := result.([]interface{})
	if !ok {
		return out
	}
	for _, item := range arr {
		talk, ok := stringKeyMap(item)
		if !ok {
			continue
		}
		id := firstStringMapValue(talk, "id", "talk_id", "talkId")
		if id == "" {
			continue
		}
		out[id] = directRoomInfo{
			name:     firstStringMapValue(talk, "name", "display_name", "displayName"),
			domainID: firstStringMapValue(talk, "domain_id", "domainId"),
		}
	}
	return out
}

func stringKeyMap(value interface{}) (map[string]interface{}, bool) {
	switch m := value.(type) {
	case map[string]interface{}:
		return m, true
	case map[interface{}]interface{}:
		out := make(map[string]interface{}, len(m))
		for key, value := range m {
			out[fmt.Sprint(key)] = value
		}
		return out, true
	default:
		return nil, false
	}
}

func stringMapValue(m map[string]interface{}, key string) string {
	value, ok := m[key]
	if !ok || value == nil {
		return ""
	}
	return strings.TrimSpace(normalizeID(value))
}

func firstStringMapValue(m map[string]interface{}, keys ...string) string {
	for _, key := range keys {
		if value := stringMapValue(m, key); value != "" {
			return value
		}
	}
	return ""
}

func normalizeID(value interface{}) string {
	switch v := value.(type) {
	case float32:
		return strconv.FormatInt(int64(v), 10)
	case float64:
		return strconv.FormatInt(int64(v), 10)
	default:
		return fmt.Sprint(v)
	}
}

func normalizeRPCID(value string) interface{} {
	if id, err := strconv.ParseUint(value, 10, 64); err == nil {
		return id
	}
	return value
}

func isInvalidTokenError(data interface{}) bool {
	m, ok := data.(map[string]interface{})
	if ok {
		code := fmt.Sprint(m["code"])
		message := fmt.Sprint(m["message"])
		return code == "401" && (message == "invalid token" || message == "bad token")
	}
	if err, ok := data.(error); ok {
		message := strings.ToLower(err.Error())
		return strings.Contains(message, "invalid token") || strings.Contains(message, "bad token")
	}
	return false
}

func isRecoverableDirectError(err error) bool {
	if err == nil {
		return false
	}
	var connErr *direct.ConnectionError
	if errors.Is(err, direct.ErrNotConnected) ||
		errors.Is(err, direct.ErrTimeout) ||
		strings.Contains(err.Error(), direct.ErrNotConnected.Error()) ||
		strings.Contains(err.Error(), direct.ErrTimeout.Error()) ||
		strings.Contains(err.Error(), "not connected") ||
		strings.Contains(err.Error(), "websocket: close") ||
		strings.Contains(err.Error(), "use of closed network connection") ||
		strings.Contains(err.Error(), "broken pipe") ||
		strings.Contains(err.Error(), "connection reset") ||
		strings.Contains(err.Error(), "i/o timeout") {
		return true
	}
	return errors.As(err, &connErr)
}

func directSelfUserID(ctx context.Context, client directClient) string {
	result, err := client.CallWithContext(ctx, direct.MethodGetMe, []interface{}{})
	if err != nil {
		return ""
	}
	m, ok := result.(map[string]interface{})
	if !ok {
		return ""
	}
	if id, ok := m["id"]; ok {
		return fmt.Sprint(id)
	}
	if id, ok := m["user_id"]; ok {
		return fmt.Sprint(id)
	}
	return ""
}

func sendDirect(ctx context.Context, client directClient, msg model.DirectOutbound) (string, error) {
	if len(msg.Attachments) == 0 {
		return client.CreateTextMessageWithContext(ctx, msg.TalkID, msg.Text)
	}
	files := make([]map[string]interface{}, 0, len(msg.Attachments))
	for _, att := range msg.Attachments {
		fileInfo, err := uploadAttachment(ctx, client, att)
		if err != nil {
			if msg.Text != "" {
				return client.CreateTextMessageWithContext(ctx, msg.TalkID, msg.Text+"\n"+fallbackAttachmentText(att))
			}
			return client.CreateTextMessageWithContext(ctx, msg.TalkID, fallbackAttachmentText(att))
		}
		files = append(files, fileInfo)
	}
	if len(files) == 1 && msg.Text == "" {
		result, err := client.CallWithContext(ctx, direct.MethodCreateMessage, []interface{}{normalizeTalkID(msg.TalkID), direct.MsgTypeFile, files[0]})
		return directMessageID(result), err
	}
	content := map[string]interface{}{"files": files}
	if msg.Text != "" {
		content["text"] = msg.Text
	}
	result, err := client.CallWithContext(ctx, direct.MethodCreateMessage, []interface{}{normalizeTalkID(msg.TalkID), direct.MsgTypeTextMultipleFile, content})
	return directMessageID(result), err
}

func directMessageID(result interface{}) string {
	if m, ok := result.(map[string]interface{}); ok {
		if id, ok := m["id"].(string); ok {
			return id
		}
		if id, ok := m["id"]; ok {
			return fmt.Sprint(id)
		}
	}
	return ""
}

func uploadAttachment(ctx context.Context, client directClient, att model.Attachment) (map[string]interface{}, error) {
	if len(att.Data) == 0 {
		return nil, fmt.Errorf("attachment %q has no data", att.Name)
	}
	if att.ContentType == "" {
		att.ContentType = "application/octet-stream"
	}
	auth, err := client.CreateUploadAuth(ctx, att.Name, att.ContentType, int64(len(att.Data)), "message")
	if err != nil {
		return nil, err
	}
	if auth.PutURL == "" {
		return nil, fmt.Errorf("direct upload auth did not include put_url")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPut, auth.PutURL, bytes.NewReader(att.Data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", att.ContentType)
	req.Header.Set("Content-Length", strconv.Itoa(len(att.Data)))
	for k, v := range auth.PostForm {
		if k == "Content-Type" || k == "Content-Disposition" {
			req.Header.Set(k, v)
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("direct upload failed status=%d body=%s", resp.StatusCode, string(b))
	}
	return map[string]interface{}{
		"url":          firstNonEmpty(att.URL, auth.PutURL),
		"content_type": att.ContentType,
		"content_size": int64(len(att.Data)),
		"name":         att.Name,
		"file_id":      auth.FileID,
	}, nil
}

func attachmentsFromDirectMessage(msg direct.ReceivedMessage) []model.Attachment {
	switch msg.Type {
	case direct.MessageTypeFile:
		if m, ok := msg.Content.(map[string]interface{}); ok {
			return []model.Attachment{attachmentFromMap(m)}
		}
	case direct.MessageTypeTextMultipleFile:
		if m, ok := msg.Content.(map[string]interface{}); ok {
			if text, ok := m["text"].(string); ok && msg.Text == "" {
				msg.Text = text
			}
			if arr, ok := m["files"].([]interface{}); ok {
				out := make([]model.Attachment, 0, len(arr))
				for _, item := range arr {
					if fm, ok := item.(map[string]interface{}); ok {
						out = append(out, attachmentFromMap(fm))
					}
				}
				return out
			}
		}
	}
	return nil
}

func attachmentFromMap(m map[string]interface{}) model.Attachment {
	return model.Attachment{
		Name:        fmt.Sprint(m["name"]),
		ContentType: fmt.Sprint(m["content_type"]),
		Size:        toInt64(m["content_size"]),
		URL:         fmt.Sprint(m["url"]),
	}
}

func messageTime(msg direct.ReceivedMessage) time.Time {
	if !msg.Timestamp.IsZero() {
		return msg.Timestamp
	}
	if msg.Created != 0 {
		return time.Unix(msg.Created, 0)
	}
	return time.Now()
}

func normalizeTalkID(talkID string) interface{} {
	if id, err := strconv.ParseUint(talkID, 10, 64); err == nil {
		return id
	}
	return talkID
}

func sleepWithBackoff(ctx context.Context, backoff *time.Duration) {
	d := *backoff
	if d > 60*time.Second {
		d = 60 * time.Second
	}
	if d > 0 {
		d += time.Duration(rand.Int63n(int64(d/3 + 1)))
	}
	timer := time.NewTimer(d)
	defer timer.Stop()
	select {
	case <-ctx.Done():
		return
	case <-timer.C:
	}
	*backoff *= 2
	if *backoff > 60*time.Second {
		*backoff = 60 * time.Second
	}
}

func fallbackAttachmentText(att model.Attachment) string {
	name := strings.TrimSpace(att.Name)
	if name == "" {
		name = "attachment"
	}
	if att.URL != "" {
		return fmt.Sprintf("[attachment: %s] %s", name, att.URL)
	}
	return fmt.Sprintf("[attachment: %s]", name)
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func toInt64(v interface{}) int64 {
	switch n := v.(type) {
	case int:
		return int64(n)
	case int64:
		return n
	case uint64:
		return int64(n)
	case float64:
		return int64(n)
	default:
		return 0
	}
}
