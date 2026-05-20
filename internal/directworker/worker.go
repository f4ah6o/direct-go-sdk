package directworker

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

type Manager struct {
	accounts map[string]config.AccountConfig
	queues   map[string]chan model.DirectOutbound
	out      chan<- model.DirectMessage
	logger   *log.Logger
}

func NewManager(accounts []config.AccountConfig, out chan<- model.DirectMessage, logger *log.Logger) *Manager {
	m := &Manager{
		accounts: map[string]config.AccountConfig{},
		queues:   map[string]chan model.DirectOutbound{},
		out:      out,
		logger:   logger,
	}
	for _, account := range accounts {
		m.accounts[account.ID] = account
		m.queues[account.ID] = make(chan model.DirectOutbound, 100)
	}
	return m
}

func (m *Manager) Run(ctx context.Context) {
	for _, account := range m.accounts {
		account := account
		queue := m.queues[account.ID]
		go runAccountWorker(ctx, account, queue, m.out, m.logger)
	}
}

func (m *Manager) Send(ctx context.Context, msg model.DirectOutbound) error {
	ch, ok := m.queues[msg.AccountID]
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

func runAccountWorker(ctx context.Context, cfg config.AccountConfig, in <-chan model.DirectOutbound, out chan<- model.DirectMessage, logger *log.Logger) {
	backoff := time.Second
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		token := os.Getenv(cfg.TokenEnv)
		if token == "" {
			logger.Printf("[%s] token env %s is empty; retrying", cfg.ID, cfg.TokenEnv)
			sleepWithBackoff(ctx, &backoff)
			continue
		}

		endpoint := cfg.Endpoint
		if endpoint == "" {
			endpoint = direct.DefaultEndpoint
		}
		client := direct.NewClient(direct.Options{
			AccessToken: token,
			Endpoint:    endpoint,
			ProxyURL:    cfg.ProxyURL,
			Name:        cfg.ID,
		})

		client.On(direct.EventError, func(data interface{}) {
			logger.Printf("[%s] direct error: %+v", cfg.ID, data)
		})
		client.On(direct.EventSessionCreated, func(data interface{}) {
			logger.Printf("[%s] direct session created", cfg.ID)
		})
		client.On(direct.EventDataRecovered, func(data interface{}) {
			logger.Printf("[%s] direct notification ready", cfg.ID)
		})
		client.OnMessage(func(msg direct.ReceivedMessage) {
			if msg.ID == "" {
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
			select {
			case out <- bm:
			case <-ctx.Done():
			}
		})

		if err := client.Connect(); err != nil {
			logger.Printf("[%s] connect failed: %v", cfg.ID, err)
			_ = client.Close()
			sleepWithBackoff(ctx, &backoff)
			continue
		}
		logger.Printf("[%s] connected", cfg.ID)
		backoff = time.Second

		disconnected := false
		for !disconnected {
			select {
			case <-ctx.Done():
				logger.Printf("[%s] shutting down", cfg.ID)
				_ = client.Close()
				return
			case <-client.Done:
				logger.Printf("[%s] disconnected; recreating client", cfg.ID)
				_ = client.Close()
				disconnected = true
			case msg := <-in:
				if err := sendDirect(ctx, client, msg); err != nil {
					logger.Printf("[%s] direct send failed: talk=%s err=%v", cfg.ID, msg.TalkID, err)
				}
			}
		}
		sleepWithBackoff(ctx, &backoff)
	}
}

func sendDirect(ctx context.Context, client *direct.Client, msg model.DirectOutbound) error {
	if len(msg.Attachments) == 0 {
		return client.SendTextWithContext(ctx, msg.TalkID, msg.Text)
	}
	files := make([]map[string]interface{}, 0, len(msg.Attachments))
	for _, att := range msg.Attachments {
		fileInfo, err := uploadAttachment(ctx, client, att)
		if err != nil {
			if msg.Text != "" {
				return client.SendTextWithContext(ctx, msg.TalkID, msg.Text+"\n"+fallbackAttachmentText(att))
			}
			return client.SendTextWithContext(ctx, msg.TalkID, fallbackAttachmentText(att))
		}
		files = append(files, fileInfo)
	}
	if len(files) == 1 && msg.Text == "" {
		_, err := client.Call(direct.MethodCreateMessage, []interface{}{normalizeTalkID(msg.TalkID), direct.MsgTypeFile, files[0]})
		return err
	}
	content := map[string]interface{}{"files": files}
	if msg.Text != "" {
		content["text"] = msg.Text
	}
	_, err := client.Call(direct.MethodCreateMessage, []interface{}{normalizeTalkID(msg.TalkID), direct.MsgTypeTextMultipleFile, content})
	return err
}

func uploadAttachment(ctx context.Context, client *direct.Client, att model.Attachment) (map[string]interface{}, error) {
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
	if att.URL != "" {
		return fmt.Sprintf("[attachment: %s] %s", att.Name, att.URL)
	}
	return fmt.Sprintf("[attachment: %s]", att.Name)
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
