package model

import (
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

type Attachment struct {
	Name        string
	ContentType string
	Size        int64
	URL         string
	Data        []byte
}

type DirectMessage struct {
	AccountID   string
	TalkID      string
	UserID      string
	Text        string
	MessageID   string
	CreatedAt   time.Time
	Attachments []Attachment
	Raw         direct.ReceivedMessage
}

type TeamsMessage struct {
	TeamID      string
	ChannelID   string
	RootID      string
	MessageID   string
	AccountID   string
	TalkID      string
	UserID      string
	Text        string
	Attachments []Attachment
	CreatedAt   time.Time
}

type DirectOutbound struct {
	AccountID   string
	TalkID      string
	Text        string
	Attachments []Attachment
	Echo        bool
	TeamsSource *TeamsSource
}

type DirectSent struct {
	Outbound  DirectOutbound
	MessageID string
	SenderID  string
	Err       error
}

type TeamsSource struct {
	ServiceURL     string
	ConversationID string
	ActivityID     string
}

type DirectReadReceipt struct {
	AccountID   string
	TalkID      string
	MessageIDs  []string
	ReadUserIDs []string
}
