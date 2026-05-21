package teams

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

type Client struct {
	cfg        config.BotConfig
	httpClient *http.Client
	mu         sync.Mutex
	token      string
	expiresAt  time.Time
}

func NewClient(cfg config.BotConfig) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) CreateRootMessage(ctx context.Context, serviceURL, conversationID string, msg model.DirectMessage) (string, error) {
	activity := NewMessageActivity(formatDirectMessage(msg))
	return c.sendActivity(ctx, serviceURL, conversationID, "", activity)
}

func (c *Client) ReplyToThread(ctx context.Context, serviceURL, conversationID, rootID string, msg model.DirectMessage) (string, error) {
	text := formatDirectMessage(msg)
	for _, att := range msg.Attachments {
		if att.URL != "" {
			text += "\n" + fmt.Sprintf("[attachment: %s] %s", att.Name, att.URL)
		} else if att.Name != "" {
			text += "\n" + fmt.Sprintf("[attachment: %s]", att.Name)
		}
	}
	activity := NewMessageActivity(text)
	return c.sendActivity(ctx, serviceURL, teamsThreadConversationID(conversationID, rootID), "", activity)
}

func (c *Client) SendText(ctx context.Context, serviceURL, conversationID, replyToID, text string) (string, error) {
	return c.sendActivity(ctx, serviceURL, conversationID, replyToID, NewMessageActivity(text))
}

func (c *Client) DownloadAttachment(ctx context.Context, activity Attachment, maxBytes int64) ([]byte, string, error) {
	if activity.ContentURL == "" {
		return nil, "", fmt.Errorf("attachment has no contentUrl")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, activity.ContentURL, nil)
	if err != nil {
		return nil, "", err
	}
	token, err := c.accessToken(ctx)
	if err == nil && token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, "", fmt.Errorf("attachment download status=%d body=%s", resp.StatusCode, string(b))
	}
	reader := io.Reader(resp.Body)
	if maxBytes > 0 {
		reader = io.LimitReader(resp.Body, maxBytes+1)
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}
	if maxBytes > 0 && int64(len(data)) > maxBytes {
		return nil, "", fmt.Errorf("attachment exceeds max size")
	}
	return data, resp.Header.Get("Content-Type"), nil
}

func (c *Client) sendActivity(ctx context.Context, serviceURL, conversationID, replyToID string, activity Activity) (string, error) {
	path := fmt.Sprintf("%s/v3/conversations/%s/activities", strings.TrimRight(serviceURL, "/"), url.PathEscape(conversationID))
	if replyToID != "" {
		path += "/" + url.PathEscape(replyToID)
	}
	token, err := c.accessToken(ctx)
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(activity)
	if err != nil {
		return "", err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, path, bytes.NewReader(body))
	if err != nil {
		return "", err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return "", fmt.Errorf("bot connector status=%d body=%s", resp.StatusCode, string(b))
	}
	var out struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.ID, nil
}

func (c *Client) accessToken(ctx context.Context) (string, error) {
	c.mu.Lock()
	if c.token != "" && time.Now().Before(c.expiresAt.Add(-time.Minute)) {
		token := c.token
		c.mu.Unlock()
		return token, nil
	}
	c.mu.Unlock()

	secret := c.cfg.AppPassword
	if secret == "" && c.cfg.AppPasswordEnv != "" {
		secret = os.Getenv(c.cfg.AppPasswordEnv)
	}
	if secret == "" {
		return "", fmt.Errorf("bot app password is empty")
	}
	form := url.Values{}
	form.Set("client_id", c.cfg.AppID)
	form.Set("client_secret", secret)
	form.Set("grant_type", "client_credentials")
	form.Set("scope", c.cfg.ConnectorScope)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.cfg.TokenURL, strings.NewReader(form.Encode()))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return "", fmt.Errorf("bot token url=%s status=%d body=%s", c.cfg.TokenURL, resp.StatusCode, string(b))
	}
	var tr struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", err
	}
	if tr.AccessToken == "" {
		return "", fmt.Errorf("bot token response missing access_token")
	}
	c.mu.Lock()
	c.token = tr.AccessToken
	c.expiresAt = time.Now().Add(time.Duration(tr.ExpiresIn) * time.Second)
	c.mu.Unlock()
	return tr.AccessToken, nil
}

func formatDirectMessage(msg model.DirectMessage) string {
	return fmt.Sprintf("[direct:%s]\nroom=%s user=%s\n%s", msg.AccountID, msg.TalkID, msg.UserID, msg.Text)
}

func teamsThreadConversationID(conversationID, rootID string) string {
	if rootID == "" || strings.Contains(conversationID, ";messageid=") {
		return conversationID
	}
	return conversationID + ";messageid=" + rootID
}
