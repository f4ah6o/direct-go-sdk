package teams

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

type Client struct {
	cfg        config.GraphConfig
	httpClient *http.Client
	mu         sync.Mutex
	token      string
	expiresAt  time.Time
}

type ChatMessage struct {
	ID          string              `json:"id"`
	ReplyToID   string              `json:"replyToId"`
	CreatedDate string              `json:"createdDateTime"`
	Body        ItemBody            `json:"body"`
	From        MessageFrom         `json:"from"`
	Mentions    []Mention           `json:"mentions"`
	Attachments []MessageAttachment `json:"attachments"`
}

type ItemBody struct {
	ContentType string `json:"contentType"`
	Content     string `json:"content"`
}

type MessageFrom struct {
	User *UserIdentity `json:"user"`
}

type UserIdentity struct {
	ID          string `json:"id"`
	DisplayName string `json:"displayName"`
}

type Mention struct {
	ID          int       `json:"id"`
	MentionText string    `json:"mentionText"`
	Mentioned   Mentioned `json:"mentioned"`
}

type Mentioned struct {
	User *UserIdentity `json:"user"`
}

type MessageAttachment struct {
	ID          string `json:"id"`
	ContentType string `json:"contentType"`
	ContentURL  string `json:"contentUrl"`
	Name        string `json:"name"`
	Content     string `json:"content"`
}

func NewClient(cfg config.GraphConfig) *Client {
	return &Client{
		cfg:        cfg,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (c *Client) CreateRootMessage(ctx context.Context, teamID, channelID string, msg model.DirectMessage) (string, error) {
	payload := map[string]interface{}{
		"subject": nil,
		"body": map[string]string{
			"contentType": "html",
			"content":     html.EscapeString(formatDirectMessage(msg)),
		},
	}
	path := fmt.Sprintf("/teams/%s/channels/%s/messages", url.PathEscape(teamID), url.PathEscape(channelID))
	var out ChatMessage
	if err := c.doJSON(ctx, http.MethodPost, path, payload, &out); err != nil {
		return "", err
	}
	return out.ID, nil
}

func (c *Client) ReplyToThread(ctx context.Context, teamID, channelID, rootID string, msg model.DirectMessage) (string, error) {
	content := html.EscapeString(formatDirectMessage(msg))
	for _, att := range msg.Attachments {
		if att.URL != "" {
			content += "<br/>" + html.EscapeString(fmt.Sprintf("[attachment: %s] %s", att.Name, att.URL))
		} else if att.Name != "" {
			content += "<br/>" + html.EscapeString(fmt.Sprintf("[attachment: %s]", att.Name))
		}
	}
	payload := map[string]interface{}{
		"body": map[string]string{
			"contentType": "html",
			"content":     content,
		},
	}
	path := fmt.Sprintf("/teams/%s/channels/%s/messages/%s/replies", url.PathEscape(teamID), url.PathEscape(channelID), url.PathEscape(rootID))
	var out ChatMessage
	if err := c.doJSON(ctx, http.MethodPost, path, payload, &out); err != nil {
		return "", err
	}
	return out.ID, nil
}

func (c *Client) GetReply(ctx context.Context, teamID, channelID, rootID, replyID string) (*ChatMessage, error) {
	path := fmt.Sprintf("/teams/%s/channels/%s/messages/%s/replies/%s", url.PathEscape(teamID), url.PathEscape(channelID), url.PathEscape(rootID), url.PathEscape(replyID))
	var out ChatMessage
	if err := c.doJSON(ctx, http.MethodGet, path, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *Client) DownloadHostedContent(ctx context.Context, teamID, channelID, messageID, hostedContentID string) ([]byte, string, error) {
	path := fmt.Sprintf("/teams/%s/channels/%s/messages/%s/hostedContents/%s/$value", url.PathEscape(teamID), url.PathEscape(channelID), url.PathEscape(messageID), url.PathEscape(hostedContentID))
	return c.doBytes(ctx, path)
}

func (c *Client) DownloadDriveItemByURL(ctx context.Context, sharingURL string) ([]byte, string, error) {
	encoded := base64.RawURLEncoding.EncodeToString([]byte(sharingURL))
	path := "/shares/u!" + encoded + "/driveItem/content"
	return c.doBytes(ctx, path)
}

func (c *Client) doJSON(ctx context.Context, method, path string, payload interface{}, out interface{}) error {
	var body io.Reader
	if payload != nil {
		data, err := json.Marshal(payload)
		if err != nil {
			return err
		}
		body = bytes.NewReader(data)
	}
	req, err := c.newRequest(ctx, method, path, body)
	if err != nil {
		return err
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("graph status=%d body=%s", resp.StatusCode, string(b))
	}
	if out == nil {
		return nil
	}
	return json.NewDecoder(resp.Body).Decode(out)
}

func (c *Client) doBytes(ctx context.Context, path string) ([]byte, string, error) {
	req, err := c.newRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, "", err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, "", fmt.Errorf("graph status=%d body=%s", resp.StatusCode, string(b))
	}
	data, err := io.ReadAll(resp.Body)
	return data, resp.Header.Get("Content-Type"), err
}

func (c *Client) newRequest(ctx context.Context, method, path string, body io.Reader) (*http.Request, error) {
	token, err := c.accessToken(ctx)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, method, strings.TrimRight(c.cfg.APIBaseURL, "/")+path, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)
	return req, nil
}

func (c *Client) accessToken(ctx context.Context) (string, error) {
	if c.cfg.AccessTokenEnv != "" {
		if token := getenv(c.cfg.AccessTokenEnv); token != "" {
			return token, nil
		}
	}
	c.mu.Lock()
	if c.token != "" && time.Now().Before(c.expiresAt.Add(-time.Minute)) {
		token := c.token
		c.mu.Unlock()
		return token, nil
	}
	c.mu.Unlock()

	secret := c.cfg.ClientSecret
	if secret == "" && c.cfg.ClientSecretEnv != "" {
		secret = getenv(c.cfg.ClientSecretEnv)
	}
	if secret == "" {
		return "", fmt.Errorf("graph client secret is empty")
	}
	form := url.Values{}
	form.Set("client_id", c.cfg.ClientID)
	form.Set("client_secret", secret)
	form.Set("grant_type", "client_credentials")
	scope := strings.Join(c.cfg.Scopes, " ")
	if scope == "" {
		scope = "https://graph.microsoft.com/.default"
	}
	form.Set("scope", scope)
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
		return "", fmt.Errorf("graph token status=%d body=%s", resp.StatusCode, string(b))
	}
	var tr struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		return "", err
	}
	if tr.AccessToken == "" {
		return "", fmt.Errorf("graph token response missing access_token")
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

func MessageText(msg *ChatMessage) string {
	text := msg.Body.Content
	text = strings.ReplaceAll(text, "<br>", "\n")
	text = strings.ReplaceAll(text, "<br/>", "\n")
	text = stripTags(text)
	text = html.UnescapeString(text)
	return strings.TrimSpace(text)
}

func MentionsUser(msg *ChatMessage, userID string) bool {
	for _, mention := range msg.Mentions {
		if mention.Mentioned.User != nil && mention.Mentioned.User.ID == userID {
			return true
		}
	}
	return false
}

func StripMentions(msg *ChatMessage) string {
	text := MessageText(msg)
	for _, mention := range msg.Mentions {
		text = strings.ReplaceAll(text, mention.MentionText, "")
	}
	return strings.TrimSpace(text)
}

func AttachmentsFromMessage(ctx context.Context, client *Client, teamID, channelID string, msg *ChatMessage, maxBytes int64) []model.Attachment {
	var out []model.Attachment
	for _, att := range msg.Attachments {
		name := att.Name
		if name == "" {
			name = "attachment"
		}
		item := model.Attachment{Name: name, ContentType: att.ContentType, URL: att.ContentURL}
		if strings.Contains(att.ContentType, "reference") && att.ContentURL != "" {
			if data, ct, err := client.DownloadDriveItemByURL(ctx, att.ContentURL); err == nil && int64(len(data)) <= maxBytes {
				item.Data = data
				item.ContentType = firstNonEmpty(ct, mime.TypeByExtension(filepath.Ext(name)), "application/octet-stream")
			}
		}
		out = append(out, item)
	}
	return out
}

func stripTags(s string) string {
	var b strings.Builder
	inTag := false
	for _, r := range s {
		switch r {
		case '<':
			inTag = true
		case '>':
			inTag = false
		default:
			if !inTag {
				b.WriteRune(r)
			}
		}
	}
	return b.String()
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}
	return ""
}

func getenv(key string) string {
	if key == "" {
		return ""
	}
	return strings.TrimSpace(getenvRaw(key))
}

var getenvRaw = os.Getenv
