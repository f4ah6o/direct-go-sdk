package webhook

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Client handles webhook HTTP requests to n8n.
type Client struct {
	WebhookURL string
	HTTPClient *http.Client
	BotName    string
}

// NewClient creates a new webhook client.
func NewClient(webhookURL, botName string) *Client {
	return &Client{
		WebhookURL: webhookURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		BotName: botName,
	}
}

// Send sends a webhook payload to n8n and returns the response.
func (c *Client) Send(payload *WebhookPayload) (*WebhookResponse, error) {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := c.HTTPClient.Post(c.WebhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to post to webhook: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("webhook returned status %d", resp.StatusCode)
	}

	var webhookResp WebhookResponse
	if err := json.NewDecoder(resp.Body).Decode(&webhookResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return &webhookResp, nil
}

// Validate validates the webhook response and returns an error code if invalid.
func (r *WebhookResponse) Validate() ErrorCode {
	if r.Action == "" {
		return ErrorCodeMissingAction
	}

	switch r.Action {
	case "none":
		return ErrorCodeOK
	case "reply":
		if r.Text == "" {
			return ErrorCodeMissingText
		}
	case "send":
		if r.RoomID == "" {
			return ErrorCodeMissingRoomID
		}
		if r.Text == "" {
			return ErrorCodeMissingText
		}
	case "send_select":
		if r.RoomID == "" {
			return ErrorCodeMissingRoomID
		}
		if r.Question == "" {
			return ErrorCodeMissingQuestion
		}
		if len(r.Options) == 0 {
			return ErrorCodeMissingOptions
		}
	case "send_yesno":
		if r.RoomID == "" {
			return ErrorCodeMissingRoomID
		}
		if r.Question == "" {
			return ErrorCodeMissingQuestion
		}
	case "send_task":
		if r.RoomID == "" {
			return ErrorCodeMissingRoomID
		}
		if r.Title == "" {
			return ErrorCodeMissingTitle
		}
	case "reply_select":
		if r.InReplyTo == "" {
			return ErrorCodeMissingInReplyTo
		}
		if r.Response == nil {
			return ErrorCodeMissingResponse
		}
	case "reply_yesno":
		if r.InReplyTo == "" {
			return ErrorCodeMissingInReplyTo
		}
		if r.ResponseBool == nil {
			return ErrorCodeMissingResponse
		}
	case "reply_task":
		if r.InReplyTo == "" {
			return ErrorCodeMissingInReplyTo
		}
		if r.Done == nil {
			return ErrorCodeMissingResponse
		}
	case "close_select", "close_yesno":
		if r.MessageID == "" {
			return ErrorCodeMissingMessageID
		}
	default:
		return ErrorCodeInvalidAction
	}

	return ErrorCodeOK
}
