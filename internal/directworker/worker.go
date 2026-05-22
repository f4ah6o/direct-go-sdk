package directworker

import (
	"bytes"
	"context"
	"crypto/sha256"
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
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

type Manager struct {
	mu            sync.Mutex
	workers       map[string]*accountWorker
	out           chan<- model.DirectMessage
	sent          chan<- model.DirectSent
	logger        *log.Logger
	clientFactory directClientFactory
	sleepBackoff  func(context.Context, *time.Duration)
}

type RuntimeAccount struct {
	Config config.AccountConfig
	Token  string
}

type accountWorker struct {
	account RuntimeAccount
	queue   chan model.DirectOutbound
	cancel  context.CancelFunc
}

type directClient interface {
	Connect() error
	Close() error
	On(string, direct.EventHandler)
	OnMessage(func(direct.ReceivedMessage))
	CreateTextMessageWithContext(context.Context, string, string) (string, error)
	CreateUploadAuth(context.Context, string, string, int64, string) (*direct.UploadAuth, error)
	Call(string, []interface{}) (interface{}, error)
	Done() <-chan struct{}
}

type directClientFactory func(direct.Options) directClient

type productionDirectClient struct {
	*direct.Client
}

func (c productionDirectClient) Done() <-chan struct{} {
	return c.Client.Done
}

func NewManager(out chan<- model.DirectMessage, sent chan<- model.DirectSent, logger *log.Logger) *Manager {
	return &Manager{
		workers:       map[string]*accountWorker{},
		out:           out,
		sent:          sent,
		logger:        logger,
		clientFactory: newProductionDirectClient,
		sleepBackoff:  sleepWithBackoff,
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
			m.logger.Printf("[%s] stopping direct worker", id)
			worker.cancel()
			delete(m.workers, id)
		}
	}
	for id, account := range next {
		if _, ok := m.workers[id]; ok {
			continue
		}
		workerCtx, cancel := context.WithCancel(ctx)
		worker := &accountWorker{
			account: account,
			queue:   make(chan model.DirectOutbound, 100),
			cancel:  cancel,
		}
		m.workers[id] = worker
		m.logger.Printf("[%s] starting direct worker", id)
		go runAccountWorker(workerCtx, account, worker.queue, m.out, m.sent, m.logger, m.clientFactory, m.sleepBackoff)
	}
}

func sameRuntimeAccount(a, b RuntimeAccount) bool {
	return a.Token == b.Token &&
		a.Config.ID == b.Config.ID &&
		a.Config.Endpoint == b.Config.Endpoint &&
		a.Config.ProxyURL == b.Config.ProxyURL &&
		a.Config.TeamsChannel == b.Config.TeamsChannel
}

func runAccountWorker(ctx context.Context, account RuntimeAccount, in <-chan model.DirectOutbound, out chan<- model.DirectMessage, sent chan<- model.DirectSent, logger *log.Logger, clientFactory directClientFactory, sleepBackoff func(context.Context, *time.Duration)) {
	cfg := account.Config
	backoff := time.Second
	var pending *model.DirectOutbound
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		token := account.Token
		if token == "" {
			logger.Printf("[%s] direct token is empty; retrying", cfg.ID)
			sleepBackoff(ctx, &backoff)
			continue
		}
		logger.Printf("[%s] using direct token len=%d sha=%s", cfg.ID, len(token), tokenFingerprint(token))

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
			logger.Printf("[%s] direct error: %+v", cfg.ID, data)
			requestReconnect()
		})
		client.On(direct.EventSessionError, func(data interface{}) {
			logger.Printf("[%s] direct session error: %+v", cfg.ID, data)
			if isInvalidTokenError(data) {
				logger.Printf("[%s] direct token is invalid; waiting for token update before reconnect", cfg.ID)
				return
			}
			requestReconnect()
		})
		client.On(direct.EventSessionCreated, func(data interface{}) {
			logger.Printf("[%s] direct session created", cfg.ID)
		})
		client.On(direct.EventDataRecovered, func(data interface{}) {
			logger.Printf("[%s] direct notification ready", cfg.ID)
			markReady()
		})
		client.On(direct.EventNotificationError, func(data interface{}) {
			logger.Printf("[%s] direct notification error: %+v", cfg.ID, data)
			requestReconnect()
		})
		client.OnMessage(func(msg direct.ReceivedMessage) {
			if msg.ID == "" {
				logger.Printf("[%s] direct message ignored without id: talk=%s user=%s type=%d", cfg.ID, msg.TalkID, msg.UserID, msg.Type)
				return
			}
			bm := model.DirectMessage{
				AccountID: cfg.ID,
				TalkID:    msg.TalkID,
				UserID:    msg.UserID,
				Text:      msg.Text,
				MessageID: msg.ID,
				CreatedAt: messageTime(msg),
				Raw:       msg,
			}
			bm.Attachments = attachmentsFromDirectMessage(msg)
			logger.Printf("[%s] direct message received: id=%s talk=%s user=%s type=%d attachments=%d", cfg.ID, msg.ID, msg.TalkID, msg.UserID, msg.Type, len(bm.Attachments))
			select {
			case out <- bm:
			case <-ctx.Done():
			}
		})

		if err := client.Connect(); err != nil {
			logger.Printf("[%s] connect failed: %v", cfg.ID, err)
			_ = client.Close()
			sleepBackoff(ctx, &backoff)
			continue
		}
		logger.Printf("[%s] connected", cfg.ID)

		disconnected := false
		for !disconnected {
			if pending != nil {
				messageID, err := sendDirect(ctx, client, *pending)
				if err != nil {
					logger.Printf("[%s] direct send failed: talk=%s err=%v", cfg.ID, pending.TalkID, err)
					if isRecoverableDirectError(err) {
						logger.Printf("[%s] recreating client after recoverable direct send error", cfg.ID)
						_ = client.Close()
						disconnected = true
						continue
					}
					select {
					case sent <- model.DirectSent{Outbound: *pending, Err: err}:
					case <-ctx.Done():
					}
					pending = nil
					continue
				}
				select {
				case sent <- model.DirectSent{Outbound: *pending, MessageID: messageID}:
				case <-ctx.Done():
				}
				pending = nil
				continue
			}

			select {
			case <-ctx.Done():
				logger.Printf("[%s] shutting down", cfg.ID)
				_ = client.Close()
				return
			case <-client.Done():
				logger.Printf("[%s] disconnected; recreating client", cfg.ID)
				_ = client.Close()
				disconnected = true
			case <-reconnectNow:
				logger.Printf("[%s] recreating client after direct startup error", cfg.ID)
				_ = client.Close()
				disconnected = true
			case <-readyNow:
				backoff = time.Second
			case msg := <-in:
				pending = &msg
			}
		}
		sleepBackoff(ctx, &backoff)
	}
}

func tokenFingerprint(token string) string {
	sum := sha256.Sum256([]byte(token))
	return fmt.Sprintf("%x", sum[:])[:12]
}

func isInvalidTokenError(data interface{}) bool {
	m, ok := data.(map[string]interface{})
	if !ok {
		return false
	}
	code := fmt.Sprint(m["code"])
	message := fmt.Sprint(m["message"])
	return code == "401" && (message == "invalid token" || message == "bad token")
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
		result, err := client.Call(direct.MethodCreateMessage, []interface{}{normalizeTalkID(msg.TalkID), direct.MsgTypeFile, files[0]})
		return directMessageID(result), err
	}
	content := map[string]interface{}{"files": files}
	if msg.Text != "" {
		content["text"] = msg.Text
	}
	result, err := client.Call(direct.MethodCreateMessage, []interface{}{normalizeTalkID(msg.TalkID), direct.MsgTypeTextMultipleFile, content})
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
