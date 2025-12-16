// messages.go defines message types and data structures for the direct API.
//
// This file contains definitions for various message content types (TextMessage,
// StampMessage, FileMessage, etc.), message type constants, and related data structures
// like Room, User, Talk, and Domain.
package direct

import (
	"encoding/json"
	"time"
)

// TextMessage represents a text message payload.
// It is used when sending text messages via the Client.Send method.
//
// Example:
//
//	msg := direct.TextMessage{Text: "Hello, World!"}
//	// Send with Client.Send(roomID, MsgTypeText, "Hello, World!")
type TextMessage struct {
	Text string `json:"text" msgpack:"text"`
}

// StampMessage represents a stamp/emoji message payload.
// Stamps are emoticon reactions that can be sent in conversations.
type StampMessage struct {
	// StampSet identifies the stamp set (collection) the stamp belongs to.
	StampSet string `json:"stamp_set" msgpack:"stamp_set"`

	// StampIndex identifies the specific stamp within the set.
	StampIndex string `json:"stamp_index" msgpack:"stamp_index"`

	// Text is optional additional text accompanying the stamp.
	Text string `json:"text,omitempty" msgpack:"text,omitempty"`
}

// YesNoMessage represents a Yes/No action stamp (interactive poll).
// Recipients can respond with Yes or No to the question.
type YesNoMessage struct {
	// Question is the text of the yes/no question.
	Question string `json:"question" msgpack:"question"`

	// Listing shows whether responses should be listed.
	Listing bool `json:"listing,omitempty" msgpack:"listing,omitempty"`

	// CloseYes determines if the poll closes when all recipients answer yes.
	CloseYes bool `json:"close_yes,omitempty" msgpack:"close_yes,omitempty"`

	// CloseNo determines if the poll closes when all recipients answer no.
	CloseNo bool `json:"close_no,omitempty" msgpack:"close_no,omitempty"`
}

// SelectMessage represents a select action stamp (multiple choice poll).
// Recipients can select one option from the provided options.
type SelectMessage struct {
	// Question is the text of the selection question.
	Question string `json:"question" msgpack:"question"`

	// Options is the list of selectable options.
	Options []string `json:"options" msgpack:"options"`

	// Listing shows whether selections should be listed.
	Listing bool `json:"listing,omitempty" msgpack:"listing,omitempty"`

	// ClosingType determines how/when the poll closes.
	ClosingType int `json:"closing_type,omitempty" msgpack:"closing_type,omitempty"`
}

// TaskMessage represents a task action stamp (task assignment).
// Recipients can mark the task as done when they have completed it.
type TaskMessage struct {
	// Title is the task description.
	Title string `json:"title" msgpack:"title"`

	// ClosingType determines how/when the task closes.
	ClosingType int `json:"closing_type,omitempty" msgpack:"closing_type,omitempty"`

	// ClosingUsers is the number of users required to mark the task as done.
	ClosingUsers int `json:"closing_users,omitempty" msgpack:"closing_users,omitempty"`

	// TargetUserIDs is the list of user IDs assigned to this task.
	TargetUserIDs []string `json:"target_user_ids,omitempty" msgpack:"target_user_ids,omitempty"`
}

// FileMessage represents a file attachment message payload.
type FileMessage struct {
	// FileID identifies the file in the server's file storage.
	FileID interface{} `json:"file_id" msgpack:"file_id"`

	// Name is the filename.
	Name string `json:"name" msgpack:"name"`

	// MimeType is the MIME type of the file (e.g., "image/png", "application/pdf").
	MimeType string `json:"mime_type" msgpack:"mime_type"`

	// Text is optional caption text for the file.
	Text string `json:"text,omitempty" msgpack:"text,omitempty"`
}

// NoteMessage represents a note/document message.
// Notes are longer-form content typically displayed in a separate note view.
type NoteMessage struct {
	// Title is the note title.
	Title string `json:"title" msgpack:"title"`

	// Content is the note body text.
	Content string `json:"content" msgpack:"content"`
}

// LocationMessage represents a location share message payload.
type LocationMessage struct {
	// Address is the human-readable location address.
	Address string `json:"address" msgpack:"address"`

	// Latitude is the geographical latitude coordinate.
	Latitude float64 `json:"latitude" msgpack:"latitude"`

	// Longitude is the geographical longitude coordinate.
	Longitude float64 `json:"longitude" msgpack:"longitude"`
}

// MessageType represents the type of message content.
// Different message types have different content structures and behaviors.
type MessageType int

// MessageType constants define all possible message types in the direct API.
const (
	// MessageTypeSystem is a system-generated message (rarely used in user messages).
	MessageTypeSystem MessageType = iota

	// MessageTypeText is a plain text message.
	MessageTypeText

	// MessageTypeStamp is a stamp/emoji reaction message.
	MessageTypeStamp

	// MessageTypeLocation is a location share message.
	MessageTypeLocation

	// MessageTypeFile is a file attachment message.
	MessageTypeFile

	// MessageTypeTextMultipleFile is a text message with multiple file attachments.
	MessageTypeTextMultipleFile

	// MessageTypeUnused is reserved but not currently used.
	MessageTypeUnused

	// MessageTypeDeleted indicates the original message was deleted.
	MessageTypeDeleted

	// MessageTypeNoteShared indicates a note was shared in the room.
	MessageTypeNoteShared

	// MessageTypeNoteDeleted indicates a shared note was deleted.
	MessageTypeNoteDeleted

	// MessageTypeNoteCreated indicates a note was created and shared.
	MessageTypeNoteCreated

	// MessageTypeNoteUpdated indicates a shared note was updated.
	MessageTypeNoteUpdated

	// MessageTypeOriginalStamp is an original stamp message (legacy type).
	MessageTypeOriginalStamp

	// MessageTypeYesNo is a yes/no poll question message.
	MessageTypeYesNo

	// MessageTypeYesNoReply is a response to a yes/no poll.
	MessageTypeYesNoReply

	// MessageTypeSelect is a multiple choice poll message.
	MessageTypeSelect

	// MessageTypeSelectReply is a response to a multiple choice poll.
	MessageTypeSelectReply

	// MessageTypeTask is a task assignment message.
	MessageTypeTask

	// MessageTypeTaskDone indicates a task was marked as done by a recipient.
	MessageTypeTaskDone

	// MessageTypeYesNoClosed indicates a yes/no poll was closed.
	MessageTypeYesNoClosed

	// MessageTypeSelectClosed indicates a multiple choice poll was closed.
	MessageTypeSelectClosed

	// MessageTypeTaskClosed indicates a task was closed/completed.
	MessageTypeTaskClosed
)

// Message is a type alias for ReceivedMessage provided for backwards compatibility.
// Deprecated: Use ReceivedMessage directly.
type Message = ReceivedMessage

// ReceivedMessage is a parsed incoming message received from the server.
// It represents a message notification and contains all relevant message metadata.
//
// The Content field's structure depends on the Type field:
// - MessageTypeText: string containing the text
// - MessageTypeStamp: map with stamp_set and stamp_index
// - MessageTypeFile: map with file_id, name, mime_type
// - MessageTypeLocation: map with address, latitude, longitude
// - MessageTypeYesNo: map with question and response data
// - MessageTypeSelect: map with question and response data
// - MessageTypeTask: map with task data
//
// Example:
//
//	client.OnMessage(func(msg direct.ReceivedMessage) {
//		log.Printf("Message from %s in room %s: %s (Type: %d)",
//			msg.UserID, msg.RoomID, msg.Text, msg.Type)
//	})
type ReceivedMessage struct {
	// ID is the unique message identifier.
	ID string `json:"id" msgpack:"id"`

	// TalkID is the conversation/room ID where the message was sent.
	TalkID string `json:"talk_id" msgpack:"talk_id"`

	// RoomID is an alias for TalkID (both refer to the conversation).
	RoomID string `json:"room_id" msgpack:"-"`

	// UserID is the ID of the user who sent the message.
	UserID string `json:"user_id" msgpack:"user_id"`

	// DomainID is the organization/domain ID the message belongs to.
	DomainID string `json:"domain_id,omitempty" msgpack:"domain_id"`

	// Text is the text content for text messages.
	// For non-text messages, see Content for type-specific data.
	Text string `json:"text,omitempty" msgpack:"-"`

	// Type is the message type (text, stamp, file, location, etc.).
	// See MessageType constants for possible values.
	Type MessageType `json:"type" msgpack:"type"`

	// Timestamp is when the message was created (parsed from Created timestamp).
	Timestamp time.Time `json:"timestamp,omitempty" msgpack:"-"`

	// Created is the Unix timestamp when the message was created.
	Created int64 `json:"created,omitempty" msgpack:"created"`

	// Content is the type-specific message content.
	// Structure depends on the Type field.
	Content interface{} `json:"content,omitempty" msgpack:"content"`

	// Raw is the unparsed JSON representation of the message for custom parsing.
	Raw json.RawMessage `json:"-" msgpack:"-"`
}

// Room represents a talk/conversation room.
type Room struct {
	// ID is the unique room identifier.
	ID interface{} `json:"id" msgpack:"id"`

	// Name is the display name of the room.
	Name string `json:"name" msgpack:"name"`

	// Type indicates whether this is a pair or group room.
	// See RoomType constants for possible values.
	Type RoomType `json:"type" msgpack:"type"`

	// UserIDs is the list of user IDs in this room.
	UserIDs []interface{} `json:"user_ids" msgpack:"user_ids"`

	// DomainID is the organization/domain this room belongs to.
	DomainID interface{} `json:"domain_id,omitempty" msgpack:"domain_id,omitempty"`
}

// RoomType represents whether a room is a pair (1:1) or group conversation.
type RoomType int

// Room type constants.
const (
	// RoomTypePair is a one-on-one direct message conversation.
	RoomTypePair RoomType = 1

	// RoomTypeGroup is a group conversation with multiple participants.
	RoomTypeGroup RoomType = 2
)

// User represents a user account or contact.
type User struct {
	// ID is the unique user identifier.
	ID interface{} `json:"id" msgpack:"id"`

	// Name is the user's login name or identifier.
	Name string `json:"name" msgpack:"name"`

	// DisplayName is the user's display name shown in the UI.
	DisplayName string `json:"display_name,omitempty" msgpack:"display_name,omitempty"`

	// Email is the user's email address.
	Email string `json:"email,omitempty" msgpack:"email,omitempty"`

	// PhoneticName is the phonetic reading of the user's name (used in some locales).
	PhoneticName string `json:"phonetic_name,omitempty" msgpack:"phonetic_name,omitempty"`

	// IconURL is the URL to the user's profile icon/avatar.
	IconURL string `json:"icon_url,omitempty" msgpack:"icon_url,omitempty"`
}

// Domain represents an organization or workspace within direct.
type Domain struct {
	// ID is the unique domain identifier.
	ID interface{} `json:"id" msgpack:"id"`

	// Name is the display name of the organization.
	Name string `json:"name" msgpack:"name"`
}

// DomainInvite represents a pending invitation to join a domain/organization.
type DomainInvite struct {
	// ID is the unique invitation identifier.
	ID interface{} `json:"id" msgpack:"id"`

	// Name is the name of the domain being invited to.
	Name string `json:"name" msgpack:"name"`

	// AccountControlRequestID is the associated account control request ID if applicable.
	AccountControlRequestID interface{} `json:"accountControlRequestId,omitempty" msgpack:"accountControlRequestId,omitempty"`
}

// Talk represents a talk/conversation room from the API.
// This is the API representation of a conversation.
type Talk struct {
	// ID is the unique talk identifier.
	ID interface{} `json:"id" msgpack:"id"`

	// DomainID is the organization this talk belongs to.
	DomainID interface{} `json:"domain_id" msgpack:"domain_id"`

	// Type indicates whether this is a pair (1) or group (2) conversation.
	Type int `json:"type" msgpack:"type"`

	// Name is the display name of the talk (mainly for groups).
	Name string `json:"name,omitempty" msgpack:"name,omitempty"`

	// UserIDs is the list of user IDs participating in this talk.
	UserIDs []interface{} `json:"user_ids" msgpack:"user_ids"`

	// AllowDisplayPastMessages indicates whether past messages can be retrieved.
	AllowDisplayPastMessages bool `json:"allow_display_past_messages" msgpack:"allow_display_past_messages"`
}

// TalkStatus represents the current status of a talk/conversation.
type TalkStatus struct {
	// TalkID identifies which talk this status information refers to.
	TalkID interface{} `json:"talk_id" msgpack:"talk_id"`

	// UnreadCount is the number of unread messages in this talk.
	UnreadCount int `json:"unread_count" msgpack:"unread_count"`

	// LatestMsgID is the ID of the most recent message in this talk.
	LatestMsgID interface{} `json:"latest_msg_id,omitempty" msgpack:"latest_msg_id,omitempty"`
}

// SessionResponse represents the response from the create_session RPC method.
// It contains information about the established session.
type SessionResponse struct {
	// UserID is the authenticated user's ID.
	UserID interface{} `json:"user_id" msgpack:"user_id"`

	// DeviceID is the device identifier for this session.
	DeviceID interface{} `json:"device_id" msgpack:"device_id"`

	// PasswordExpiration indicates when the password will expire (if applicable).
	PasswordExpiration interface{} `json:"password_expiration" msgpack:"password_expiration"`
}
