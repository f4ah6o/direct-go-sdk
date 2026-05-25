package teams

import (
	"encoding/json"
	"html"
	"regexp"
	"strings"
)

type Activity struct {
	Type         string              `json:"type"`
	ID           string              `json:"id,omitempty"`
	Timestamp    string              `json:"timestamp,omitempty"`
	ServiceURL   string              `json:"serviceUrl,omitempty"`
	ChannelID    string              `json:"channelId,omitempty"`
	From         ChannelAccount      `json:"from,omitempty"`
	Recipient    ChannelAccount      `json:"recipient,omitempty"`
	Conversation ConversationAccount `json:"conversation,omitempty"`
	Text         string              `json:"text,omitempty"`
	TextFormat   string              `json:"textFormat,omitempty"`
	TopicName    string              `json:"topicName,omitempty"`
	ReplyToID    string              `json:"replyToId,omitempty"`
	MembersAdded []ChannelAccount    `json:"membersAdded,omitempty"`
	Entities     []Entity            `json:"entities,omitempty"`
	Attachments  []Attachment        `json:"attachments,omitempty"`
	ChannelData  ChannelData         `json:"channelData,omitempty"`
}

type ChannelAccount struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ConversationAccount struct {
	ID               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`
	ConversationType string `json:"conversationType,omitempty"`
	TenantID         string `json:"tenantId,omitempty"`
}

type ConversationParameters struct {
	IsGroup      bool                `json:"isGroup,omitempty"`
	Bot          ChannelAccount      `json:"bot,omitempty"`
	Activity     Activity            `json:"activity,omitempty"`
	ChannelData  ChannelData         `json:"channelData,omitempty"`
	TenantID     string              `json:"tenantId,omitempty"`
	Conversation ConversationAccount `json:"conversation,omitempty"`
}

type Entity struct {
	Type      string         `json:"type,omitempty"`
	Text      string         `json:"text,omitempty"`
	Mentioned ChannelAccount `json:"mentioned,omitempty"`
}

type Attachment struct {
	ContentType string      `json:"contentType,omitempty"`
	ContentURL  string      `json:"contentUrl,omitempty"`
	Name        string      `json:"name,omitempty"`
	Content     interface{} `json:"content,omitempty"`
}

func (a Attachment) DownloadURL() string {
	switch content := a.Content.(type) {
	case map[string]interface{}:
		if v, ok := content["downloadUrl"].(string); ok {
			return v
		}
	case map[string]string:
		return content["downloadUrl"]
	case json.RawMessage:
		var m map[string]interface{}
		if err := json.Unmarshal(content, &m); err == nil {
			if v, ok := m["downloadUrl"].(string); ok {
				return v
			}
		}
	}
	return ""
}

type ChannelData struct {
	Team    TeamInfo    `json:"team,omitempty"`
	Channel ChannelInfo `json:"channel,omitempty"`
	Tenant  TenantInfo  `json:"tenant,omitempty"`
}

type TeamInfo struct {
	ID string `json:"id,omitempty"`
}

type ChannelInfo struct {
	ID string `json:"id,omitempty"`
}

type TenantInfo struct {
	ID string `json:"id,omitempty"`
}

func NewMessageActivity(text string) Activity {
	return Activity{
		Type:       "message",
		TextFormat: "markdown",
		Text:       text,
	}
}

func MentionsRecipient(a Activity) bool {
	for _, entity := range a.Entities {
		if entity.Type == "mention" && entity.Mentioned.ID == a.Recipient.ID {
			return true
		}
	}
	return false
}

func StripRecipientMention(a Activity) string {
	text := a.Text
	for _, entity := range a.Entities {
		if entity.Type == "mention" && entity.Mentioned.ID == a.Recipient.ID {
			text = strings.ReplaceAll(text, entity.Text, "")
		}
	}
	text = stripTags(text)
	text = html.UnescapeString(text)
	return strings.TrimSpace(text)
}

func ParseBindAlias(a Activity) (string, bool) {
	text := strings.ToLower(StripRecipientMention(a))
	fields := strings.Fields(text)
	if len(fields) == 2 && fields[0] == "bind" && fields[1] != "" {
		return fields[1], true
	}
	return "", false
}

func ParseUnbindAlias(a Activity) (string, bool) {
	text := strings.ToLower(StripRecipientMention(a))
	fields := strings.Fields(text)
	if len(fields) == 2 && fields[0] == "unbind" && fields[1] != "" {
		return fields[1], true
	}
	return "", false
}

func ParseCommand(a Activity) string {
	text := strings.ToLower(StripRecipientMention(a))
	fields := strings.Fields(text)
	if len(fields) == 0 {
		return ""
	}
	return fields[0]
}

func IsNewThreadCommand(command string) bool {
	switch strings.ToLower(command) {
	case "new-thread", "newthread", "reset-thread", "reset":
		return true
	default:
		return false
	}
}

func BotWasAdded(a Activity) bool {
	if len(a.MembersAdded) == 0 {
		return false
	}
	for _, member := range a.MembersAdded {
		if member.ID != "" && member.ID == a.Recipient.ID {
			return true
		}
	}
	return false
}

var tagRE = regexp.MustCompile(`<[^>]+>`)

func stripTags(s string) string {
	s = strings.ReplaceAll(s, "<br>", "\n")
	s = strings.ReplaceAll(s, "<br/>", "\n")
	s = strings.ReplaceAll(s, "<br />", "\n")
	return tagRE.ReplaceAllString(s, "")
}
