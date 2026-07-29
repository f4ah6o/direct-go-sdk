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
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

type Client struct {
	cfg           config.BotConfig
	publicBaseURL string
	fileProxyTTL  string
	httpClient    *http.Client
	mu            sync.Mutex
	token         string
	expiresAt     time.Time
}

const (
	ReactionEyes                = "1f440_eyes"
	ReactionBallotBoxWithBallot = "1f5f3_ballotboxwithballot"
)

func NewClient(cfg config.BotConfig, publicBaseURL ...string) *Client {
	baseURL := ""
	if len(publicBaseURL) > 0 {
		baseURL = publicBaseURL[0]
	}
	ttl := "24h"
	if len(publicBaseURL) > 1 && strings.TrimSpace(publicBaseURL[1]) != "" {
		ttl = publicBaseURL[1]
	}
	return &Client{
		cfg:           cfg,
		publicBaseURL: strings.TrimRight(baseURL, "/"),
		fileProxyTTL:  ttl,
		httpClient:    &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) CreateRootMessage(ctx context.Context, serviceURL, conversationID string, msg model.DirectMessage) (string, error) {
	activity := NewMessageActivity(c.formatDirectRootMessage(msg))
	return c.sendActivity(ctx, serviceURL, conversationID, "", activity)
}

func (c *Client) CreateRootThread(ctx context.Context, serviceURL string, binding ChannelThreadBinding, msg model.DirectMessage) (string, error) {
	activity := NewMessageActivity(c.formatDirectRootMessage(msg))
	return c.createConversation(ctx, serviceURL, binding.ConversationParameters(activity))
}

func (c *Client) CreateRootThreadText(ctx context.Context, serviceURL string, binding ChannelThreadBinding, text string) (string, error) {
	activity := NewMessageActivity(text)
	return c.createConversation(ctx, serviceURL, binding.ConversationParameters(activity))
}

func (c *Client) ReplyToThread(ctx context.Context, serviceURL, conversationID, rootID string, msg model.DirectMessage) (string, error) {
	activity := NewMessageActivity(c.formatDirectReplyMessage(msg))
	return c.sendActivity(ctx, serviceURL, teamsThreadConversationID(conversationID, rootID), "", activity)
}

func (c *Client) SendText(ctx context.Context, serviceURL, conversationID, replyToID, text string) (string, error) {
	return c.sendActivity(ctx, serviceURL, conversationID, replyToID, NewMessageActivity(text))
}

func (c *Client) AddReaction(ctx context.Context, serviceURL, conversationID, activityID, reactionType string) error {
	if reactionType == "" {
		reactionType = ReactionEyes
	}
	return c.addReaction(ctx, serviceURL, conversationID, activityID, reactionType, 3)
}

type ChannelThreadBinding struct {
	TeamID         string
	ChannelID      string
	ConversationID string
	TenantID       string
	BotID          string
}

func (b ChannelThreadBinding) ConversationParameters(activity Activity) ConversationParameters {
	conversationID := firstNonEmptyString(b.ChannelID, b.ConversationID)
	return ConversationParameters{
		IsGroup:  true,
		Bot:      ChannelAccount{ID: b.BotID},
		Activity: activity,
		ChannelData: ChannelData{
			Team:    TeamInfo{ID: b.TeamID},
			Channel: ChannelInfo{ID: conversationID},
			Tenant:  TenantInfo{ID: b.TenantID},
		},
		TenantID: b.TenantID,
		Conversation: ConversationAccount{
			ID:               conversationID,
			ConversationType: "channel",
			TenantID:         b.TenantID,
		},
	}
}

func (c *Client) DownloadAttachment(ctx context.Context, activity Attachment, maxBytes int64) ([]byte, string, error) {
	contentURL := firstNonEmptyString(activity.DownloadURL(), activity.ContentURL)
	if contentURL == "" {
		return nil, "", fmt.Errorf("attachment has no contentUrl")
	}
	if !trustedTeamsAttachmentURL(contentURL) {
		return nil, "", fmt.Errorf("attachment contentUrl host is not trusted")
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, contentURL, nil)
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
		return nil, "", fmt.Errorf("attachment download status=%d response_body_bytes=%d", resp.StatusCode, len(b))
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
		return "", fmt.Errorf("bot connector status=%d response_body_bytes=%d", resp.StatusCode, len(b))
	}
	var out struct {
		ID string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	return out.ID, nil
}

func (c *Client) addReaction(ctx context.Context, serviceURL, conversationID, activityID, reactionType string, attempts int) error {
	path := fmt.Sprintf("%s/v3/conversations/%s/activities/%s/reactions/%s",
		strings.TrimRight(serviceURL, "/"),
		url.PathEscape(conversationID),
		url.PathEscape(activityID),
		url.PathEscape(reactionType),
	)
	token, err := c.accessToken(ctx)
	if err != nil {
		return err
	}
	for attempt := 0; attempt < attempts; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodPut, path, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Authorization", "Bearer "+token)
		resp, err := c.httpClient.Do(req)
		if err != nil {
			return err
		}
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		_ = resp.Body.Close()
		if resp.StatusCode >= 200 && resp.StatusCode < 300 {
			return nil
		}
		if resp.StatusCode == http.StatusTooManyRequests && attempt+1 < attempts {
			delay := retryAfter(resp.Header.Get("Retry-After"))
			timer := time.NewTimer(delay)
			select {
			case <-ctx.Done():
				timer.Stop()
				return ctx.Err()
			case <-timer.C:
			}
			continue
		}
		return fmt.Errorf("bot connector reaction status=%d response_body_bytes=%d", resp.StatusCode, len(body))
	}
	return nil
}

func retryAfter(value string) time.Duration {
	if value == "" {
		return time.Second
	}
	if seconds, err := strconv.Atoi(value); err == nil && seconds > 0 {
		return time.Duration(seconds) * time.Second
	}
	if when, err := http.ParseTime(value); err == nil {
		if delay := time.Until(when); delay > 0 {
			return delay
		}
	}
	return time.Second
}

func (c *Client) createConversation(ctx context.Context, serviceURL string, params ConversationParameters) (string, error) {
	path := fmt.Sprintf("%s/v3/conversations", strings.TrimRight(serviceURL, "/"))
	token, err := c.accessToken(ctx)
	if err != nil {
		return "", err
	}
	body, err := json.Marshal(params)
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
		return "", fmt.Errorf("bot connector create conversation status=%d response_body_bytes=%d", resp.StatusCode, len(b))
	}
	var out struct {
		ActivityID string `json:"activityId"`
		ID         string `json:"id"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&out); err != nil {
		return "", err
	}
	if out.ActivityID != "" {
		return out.ActivityID, nil
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
		return "", fmt.Errorf("bot token request failed status=%d response_body_bytes=%d", resp.StatusCode, len(b))
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

func (c *Client) formatDirectRootMessage(msg model.DirectMessage) string {
	return appendAttachmentLinks("# "+formatDirectRootTopic(msg)+"\n\n"+msg.Text, msg, c.cfg, c.publicBaseURL, c.fileProxyTTL)
}

func formatDirectRootTopic(msg model.DirectMessage) string {
	user := strings.TrimSpace(msg.UserName)
	if user == "" {
		user = msg.UserID
	}
	room := strings.TrimSpace(msg.RoomName)
	if room == "" {
		return user
	}
	return fmt.Sprintf("%s / %s", user, room)
}

func (c *Client) formatDirectReplyMessage(msg model.DirectMessage) string {
	sender := strings.TrimSpace(msg.UserName)
	if sender == "" {
		sender = msg.UserID
	}
	return appendAttachmentLinks(fmt.Sprintf("送信: %s  \n%s", sender, msg.Text), msg, c.cfg, c.publicBaseURL, c.fileProxyTTL)
}

func appendAttachmentLinks(text string, msg model.DirectMessage, cfg config.BotConfig, publicBaseURL, fileProxyTTL string) string {
	for _, att := range msg.Attachments {
		name := strings.TrimSpace(att.Name)
		if name == "" {
			name = "attachment"
		}
		if att.URL != "" && publicBaseURL != "" {
			u, err := signedDirectFileURL(publicBaseURL, cfg, fileProxyTTL, msg.AccountID, att.URL, time.Now())
			if err != nil {
				u = att.URL
			}
			text += "\n" + fmt.Sprintf("[attachment: %s](%s)", escapeMarkdownLinkText(name), u)
		} else if att.URL != "" {
			text += "\n" + fmt.Sprintf("[attachment: %s](%s)", escapeMarkdownLinkText(name), att.URL)
		} else {
			text += "\n" + fmt.Sprintf("[attachment: %s]", name)
		}
	}
	return text
}

func escapeMarkdownLinkText(s string) string {
	return strings.NewReplacer("[", "\\[", "]", "\\]").Replace(s)
}

func firstNonEmptyString(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}

func teamsThreadConversationID(conversationID, rootID string) string {
	if rootID == "" || strings.Contains(conversationID, ";messageid=") {
		return conversationID
	}
	return conversationID + ";messageid=" + rootID
}

func trustedTeamsAttachmentURL(raw string) bool {
	u, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || u.Scheme != "https" {
		return false
	}
	host := strings.ToLower(u.Hostname())
	return host == "smba.trafficmanager.net" ||
		strings.HasSuffix(host, ".trafficmanager.net") ||
		strings.HasSuffix(host, ".skype.com") ||
		strings.HasSuffix(host, ".botframework.com") ||
		strings.HasSuffix(host, ".teams.microsoft.com")
}
