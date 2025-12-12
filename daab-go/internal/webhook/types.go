// Package webhook provides n8n webhook integration for daab-go.
package webhook

import "time"

// WebhookPayload is the JSON payload sent to n8n webhook.
type WebhookPayload struct {
	Version   string       `json:"version"`
	EventType string       `json:"eventType"`
	Timestamp string       `json:"timestamp"`
	Bot       BotInfo      `json:"bot"`
	Message   *MessageData `json:"message,omitempty"`
	Event     interface{}  `json:"event,omitempty"`
	Raw       interface{}  `json:"raw,omitempty"`
}

// BotInfo contains bot information.
type BotInfo struct {
	Name string `json:"name"`
}

// UserData contains detailed user information.
type UserData struct {
	ID          string `json:"id"`
	Email       string `json:"email,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	Name        string `json:"name,omitempty"`
}

// MessageData contains message information.
type MessageData struct {
	ID       string      `json:"id"`
	TalkID   string      `json:"talkId"`
	UserID   string      `json:"userId"`
	User     *UserData   `json:"user,omitempty"` // Optional full user info
	Type     int         `json:"type"`
	TypeName string      `json:"typeName"`
	Text     string      `json:"text"`
	Content  interface{} `json:"content,omitempty"`
	Created  int64       `json:"created"`
}

// WebhookResponse is the JSON response from n8n.
type WebhookResponse struct {
	Action      string   `json:"action"`
	RoomID      string   `json:"roomId,omitempty"`
	Text        string   `json:"text,omitempty"`
	Question    string   `json:"question,omitempty"`
	Options     []string `json:"options,omitempty"`
	Title       string   `json:"title,omitempty"`
	InReplyTo   string   `json:"inReplyTo,omitempty"`
	Response    *int     `json:"response,omitempty"` // For select: index
	ResponseBool *bool   `json:"responseBool,omitempty"` // For yesno
	Done        *bool    `json:"done,omitempty"` // For task
	MessageID   string   `json:"messageId,omitempty"`
	ErrorCode   string   `json:"errorCode,omitempty"`
}

// ErrorCode represents structured error codes for "Parse, don't validate" pattern.
type ErrorCode string

const (
	ErrorCodeOK              ErrorCode = "ok"
	ErrorCodeInvalidJSON     ErrorCode = "invalid_json"
	ErrorCodeMissingAction   ErrorCode = "missing_action"
	ErrorCodeInvalidAction   ErrorCode = "invalid_action"
	ErrorCodeMissingRoomID   ErrorCode = "missing_room_id"
	ErrorCodeMissingText     ErrorCode = "missing_text"
	ErrorCodeMissingQuestion ErrorCode = "missing_question"
	ErrorCodeMissingOptions  ErrorCode = "missing_options"
	ErrorCodeMissingInReplyTo ErrorCode = "missing_in_reply_to"
	ErrorCodeMissingResponse ErrorCode = "missing_response"
	ErrorCodeMissingTitle    ErrorCode = "missing_title"
	ErrorCodeMissingMessageID ErrorCode = "missing_message_id"
)

// NewPayload creates a new WebhookPayload for a message event.
func NewPayload(eventType, botName string, msg MessageData) *WebhookPayload {
	return &WebhookPayload{
		Version:   "1.0",
		EventType: eventType,
		Timestamp: time.Now().Format(time.RFC3339),
		Bot: BotInfo{
			Name: botName,
		},
		Message: &msg,
	}
}

// MessageTypeToName converts message type integer to human-readable name.
func MessageTypeToName(msgType int) string {
	switch msgType {
	case 0:
		return "system"
	case 1:
		return "text"
	case 2:
		return "stamp"
	case 3:
		return "location"
	case 4:
		return "file"
	case 5:
		return "text_multiple_file"
	case 7:
		return "deleted"
	case 8:
		return "note_shared"
	case 9:
		return "note_deleted"
	case 10:
		return "note_created"
	case 11:
		return "note_updated"
	case 12:
		return "original_stamp"
	case 13:
		return "yesno"
	case 14:
		return "yesno_reply"
	case 15:
		return "select"
	case 16:
		return "select_reply"
	case 17:
		return "task"
	case 18:
		return "task_done"
	case 19:
		return "yesno_closed"
	case 20:
		return "select_closed"
	case 21:
		return "task_closed"
	default:
		return "unknown"
	}
}
