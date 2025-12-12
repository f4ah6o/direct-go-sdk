package direct

import (
	"encoding/json"
	"time"
)

// TextMessage represents a text message.
type TextMessage struct {
	Text string `json:"text" msgpack:"text"`
}

// StampMessage represents a stamp message.
type StampMessage struct {
	StampSet   string `json:"stamp_set" msgpack:"stamp_set"`
	StampIndex string `json:"stamp_index" msgpack:"stamp_index"`
	Text       string `json:"text,omitempty" msgpack:"text,omitempty"`
}

// YesNoMessage represents a Yes/No action stamp.
type YesNoMessage struct {
	Question string `json:"question" msgpack:"question"`
	Listing  bool   `json:"listing,omitempty" msgpack:"listing,omitempty"`
	CloseYes bool   `json:"close_yes,omitempty" msgpack:"close_yes,omitempty"`
	CloseNo  bool   `json:"close_no,omitempty" msgpack:"close_no,omitempty"`
}

// SelectMessage represents a select action stamp.
type SelectMessage struct {
	Question    string   `json:"question" msgpack:"question"`
	Options     []string `json:"options" msgpack:"options"`
	Listing     bool     `json:"listing,omitempty" msgpack:"listing,omitempty"`
	ClosingType int      `json:"closing_type,omitempty" msgpack:"closing_type,omitempty"`
}

// TaskMessage represents a task action stamp.
type TaskMessage struct {
	Title         string   `json:"title" msgpack:"title"`
	ClosingType   int      `json:"closing_type,omitempty" msgpack:"closing_type,omitempty"`
	ClosingUsers  int      `json:"closing_users,omitempty" msgpack:"closing_users,omitempty"`
	TargetUserIDs []string `json:"target_user_ids,omitempty" msgpack:"target_user_ids,omitempty"`
}

// FileMessage represents a file attachment.
type FileMessage struct {
	FileID   interface{} `json:"file_id" msgpack:"file_id"`
	Name     string      `json:"name" msgpack:"name"`
	MimeType string      `json:"mime_type" msgpack:"mime_type"`
	Text     string      `json:"text,omitempty" msgpack:"text,omitempty"`
}

// NoteMessage represents a note.
type NoteMessage struct {
	Title   string `json:"title" msgpack:"title"`
	Content string `json:"content" msgpack:"content"`
}

// LocationMessage represents a location share.
type LocationMessage struct {
	Address   string  `json:"address" msgpack:"address"`
	Latitude  float64 `json:"latitude" msgpack:"latitude"`
	Longitude float64 `json:"longitude" msgpack:"longitude"`
}

// MessageType represents the type of message.
type MessageType int

const (
	MessageTypeSystem MessageType = iota
	MessageTypeText
	MessageTypeStamp
	MessageTypeLocation
	MessageTypeFile
	MessageTypeTextMultipleFile
	MessageTypeUnused
	MessageTypeDeleted
	MessageTypeNoteShared
	MessageTypeNoteDeleted
	MessageTypeNoteCreated
	MessageTypeNoteUpdated
	MessageTypeOriginalStamp
	MessageTypeYesNo
	MessageTypeYesNoReply
	MessageTypeSelect
	MessageTypeSelectReply
	MessageTypeTask
	MessageTypeTaskDone
	MessageTypeYesNoClosed
	MessageTypeSelectClosed
	MessageTypeTaskClosed
)

// Message is the legacy interface for compatibility.
type Message = ReceivedMessage

// ReceivedMessage is a parsed incoming message with all fields.
type ReceivedMessage struct {
	ID        string      `json:"id" msgpack:"id"`
	TalkID    string      `json:"talk_id" msgpack:"talk_id"`
	RoomID    string      `json:"room_id" msgpack:"-"`
	UserID    string      `json:"user_id" msgpack:"user_id"`
	DomainID  string      `json:"domain_id,omitempty" msgpack:"domain_id"`
	Text      string      `json:"text,omitempty" msgpack:"-"`
	Type      MessageType `json:"type" msgpack:"type"`
	Timestamp time.Time   `json:"timestamp,omitempty" msgpack:"-"`
	Created   int64       `json:"created,omitempty" msgpack:"created"`
	Content   interface{} `json:"content,omitempty" msgpack:"content"`

	// Raw data for custom parsing
	Raw json.RawMessage `json:"-" msgpack:"-"`
}

// Room represents a talk room.
type Room struct {
	ID       interface{}   `json:"id" msgpack:"id"`
	Name     string        `json:"name" msgpack:"name"`
	Type     RoomType      `json:"type" msgpack:"type"`
	UserIDs  []interface{} `json:"user_ids" msgpack:"user_ids"`
	DomainID interface{}   `json:"domain_id,omitempty" msgpack:"domain_id,omitempty"`
}

// RoomType represents the type of a room.
type RoomType int

const (
	RoomTypePair  RoomType = 1 // 1:1 chat
	RoomTypeGroup RoomType = 2 // Group chat
)

// User represents a user/contact.
type User struct {
	ID           interface{} `json:"id" msgpack:"id"`
	Name         string      `json:"name" msgpack:"name"`
	DisplayName  string      `json:"display_name,omitempty" msgpack:"display_name,omitempty"`
	Email        string      `json:"email,omitempty" msgpack:"email,omitempty"`
	PhoneticName string      `json:"phonetic_name,omitempty" msgpack:"phonetic_name,omitempty"`
	IconURL      string      `json:"icon_url,omitempty" msgpack:"icon_url,omitempty"`
}

// Domain represents an organization/domain.
type Domain struct {
	ID   interface{} `json:"id" msgpack:"id"`
	Name string      `json:"name" msgpack:"name"`
}

// DomainInvite represents a pending domain invitation.
type DomainInvite struct {
	ID                      interface{} `json:"id" msgpack:"id"`
	Name                    string      `json:"name" msgpack:"name"`
	AccountControlRequestID interface{} `json:"accountControlRequestId,omitempty" msgpack:"accountControlRequestId,omitempty"`
}

// Talk represents a talk room from the API.
type Talk struct {
	ID                       interface{}   `json:"id" msgpack:"id"`
	DomainID                 interface{}   `json:"domain_id" msgpack:"domain_id"`
	Type                     int           `json:"type" msgpack:"type"`
	Name                     string        `json:"name,omitempty" msgpack:"name,omitempty"`
	UserIDs                  []interface{} `json:"user_ids" msgpack:"user_ids"`
	AllowDisplayPastMessages bool          `json:"allow_display_past_messages" msgpack:"allow_display_past_messages"`
}

// TalkStatus represents the status of a talk.
type TalkStatus struct {
	TalkID      interface{} `json:"talk_id" msgpack:"talk_id"`
	UnreadCount int         `json:"unread_count" msgpack:"unread_count"`
	LatestMsgID interface{} `json:"latest_msg_id,omitempty" msgpack:"latest_msg_id,omitempty"`
}

// SessionResponse represents the response from create_session.
type SessionResponse struct {
	UserID             interface{} `json:"user_id" msgpack:"user_id"`
	DeviceID           interface{} `json:"device_id" msgpack:"device_id"`
	PasswordExpiration interface{} `json:"password_expiration" msgpack:"password_expiration"`
}
