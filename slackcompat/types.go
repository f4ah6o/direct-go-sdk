package slackcompat

import (
	"context"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

type DirectAPI interface {
	AccountID() string
	GetMe(context.Context) (*direct.UserInfo, error)
	GetTalks(context.Context) ([]direct.Talk, error)
	GetUsers(context.Context) ([]direct.UserInfo, error)
	GetMessages(context.Context, interface{}, interface{}, *direct.GetMessagesOptions) ([]direct.ReceivedMessage, error)
	SendText(context.Context, string, string) (string, error)
}

type SlackResponse struct {
	OK    bool   `json:"ok"`
	Error string `json:"error,omitempty"`
}

type AuthTestResponse struct {
	OK     bool   `json:"ok"`
	URL    string `json:"url,omitempty"`
	Team   string `json:"team,omitempty"`
	TeamID string `json:"team_id,omitempty"`
	User   string `json:"user,omitempty"`
	UserID string `json:"user_id,omitempty"`
	BotID  string `json:"bot_id,omitempty"`
	Error  string `json:"error,omitempty"`
}

type PostMessageResponse struct {
	OK      bool         `json:"ok"`
	Channel string       `json:"channel,omitempty"`
	TS      string       `json:"ts,omitempty"`
	Message SlackMessage `json:"message,omitempty"`
	Error   string       `json:"error,omitempty"`
}

type ConversationsListResponse struct {
	OK       bool           `json:"ok"`
	Channels []SlackChannel `json:"channels,omitempty"`
	Error    string         `json:"error,omitempty"`
	Meta     ResponseMeta   `json:"response_metadata,omitempty"`
}

type ConversationsHistoryResponse struct {
	OK               bool           `json:"ok"`
	Messages         []SlackMessage `json:"messages,omitempty"`
	HasMore          bool           `json:"has_more"`
	PinCount         int            `json:"pin_count"`
	ChannelActionsTS string         `json:"channel_actions_ts,omitempty"`
	Error            string         `json:"error,omitempty"`
}

type UsersListResponse struct {
	OK      bool         `json:"ok"`
	Members []SlackUser  `json:"members,omitempty"`
	Error   string       `json:"error,omitempty"`
	Meta    ResponseMeta `json:"response_metadata,omitempty"`
}

type ResponseMeta struct {
	NextCursor string `json:"next_cursor,omitempty"`
}

type SlackChannel struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IsChannel  bool   `json:"is_channel"`
	IsGroup    bool   `json:"is_group"`
	IsIM       bool   `json:"is_im"`
	IsPrivate  bool   `json:"is_private"`
	IsArchived bool   `json:"is_archived"`
	NumMembers int    `json:"num_members,omitempty"`
}

type SlackUser struct {
	ID        string           `json:"id"`
	Name      string           `json:"name"`
	Deleted   bool             `json:"deleted"`
	IsBot     bool             `json:"is_bot"`
	IsAppUser bool             `json:"is_app_user"`
	Profile   SlackUserProfile `json:"profile"`
}

type SlackUserProfile struct {
	RealName string `json:"real_name,omitempty"`
	Email    string `json:"email,omitempty"`
	Image48  string `json:"image_48,omitempty"`
}

type SlackMessage struct {
	Type string `json:"type"`
	User string `json:"user,omitempty"`
	Text string `json:"text"`
	TS   string `json:"ts"`
}

type EventsEnvelope struct {
	Token       string     `json:"token,omitempty"`
	TeamID      string     `json:"team_id,omitempty"`
	APIAppID    string     `json:"api_app_id,omitempty"`
	Type        string     `json:"type"`
	EventID     string     `json:"event_id,omitempty"`
	EventTime   int64      `json:"event_time,omitempty"`
	AuthedUsers []string   `json:"authed_users,omitempty"`
	Event       SlackEvent `json:"event"`
}

type SlackEvent struct {
	Type    string `json:"type"`
	Channel string `json:"channel"`
	User    string `json:"user,omitempty"`
	Text    string `json:"text"`
	TS      string `json:"ts"`
}

type Config struct {
	OP       OPConfig        `yaml:"op"`
	Server   ServerConfig    `yaml:"server"`
	Slack    SlackConfig     `yaml:"slack"`
	Accounts []AccountConfig `yaml:"accounts"`
}

type OPConfig struct {
	Binary string `yaml:"binary"`
}

type ServerConfig struct {
	ListenAddr     string `yaml:"listen_addr"`
	BearerTokenEnv string `yaml:"bearer_token_env"`
	BearerTokenRef string `yaml:"bearer_token_ref"`
}

type SlackConfig struct {
	TeamID           string `yaml:"team_id"`
	TeamName         string `yaml:"team_name"`
	BotUserID        string `yaml:"bot_user_id"`
	EventCallbackURL string `yaml:"event_callback_url"`
}

type AccountConfig struct {
	ID       string `yaml:"id"`
	TokenEnv string `yaml:"token_env"`
	TokenRef string `yaml:"token_ref"`
	Endpoint string `yaml:"endpoint"`
	ProxyURL string `yaml:"proxy_url"`
}

type EventSink interface {
	Publish(context.Context, EventsEnvelope) error
}

type Event struct {
	AccountID string
	Message   direct.ReceivedMessage
	CreatedAt time.Time
}
