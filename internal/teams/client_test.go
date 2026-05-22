package teams

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

func TestTeamsThreadConversationID(t *testing.T) {
	got := teamsThreadConversationID("19:channel@thread.tacv2", "123")
	want := "19:channel@thread.tacv2;messageid=123"
	if got != want {
		t.Fatalf("teamsThreadConversationID() = %q, want %q", got, want)
	}
}

func TestTeamsThreadConversationIDAlreadyThread(t *testing.T) {
	id := "19:channel@thread.tacv2;messageid=123"
	if got := teamsThreadConversationID(id, "456"); got != id {
		t.Fatalf("teamsThreadConversationID() = %q, want %q", got, id)
	}
}

func TestFormatDirectRootAndReplyMessages(t *testing.T) {
	client := NewClient(config.BotConfig{AppPassword: "secret"}, "https://bridge.example.com")
	msg := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		UserID:    "1792959268018716672",
		Text:      "こんにちは",
		Attachments: []model.Attachment{{
			Name: "image.png",
			URL:  "https://api.direct4b.com/albero-app-server/files/file/token?message_id=1",
		}},
	}

	root := client.formatDirectRootMessage(msg)
	if !strings.HasPrefix(root, "# [direct:bot-trial] room=1792967566075891712 user=1792959268018716672\n\nこんにちは") {
		t.Fatalf("unexpected root message: %q", root)
	}
	if !strings.Contains(root, "[attachment: image.png](https://bridge.example.com/files/direct?") ||
		!strings.Contains(root, "account=bot-trial") ||
		!strings.Contains(root, "url=https%3A%2F%2Fapi.direct4b.com%2Falbero-app-server%2Ffiles%2Ffile%2Ftoken%3Fmessage_id%3D1") ||
		!strings.Contains(root, "sig=") {
		t.Fatalf("expected proxied markdown link: %q", root)
	}

	reply := client.formatDirectReplyMessage(msg)
	if strings.Contains(reply, "[direct:") || strings.Contains(reply, "room=") {
		t.Fatalf("reply should not repeat thread title details: %q", reply)
	}
	if !strings.HasPrefix(reply, "user=1792959268018716672  \nこんにちは") {
		t.Fatalf("unexpected reply message: %q", reply)
	}

	if got := formatDirectRootTopic(msg); got != "[direct:bot-trial] room=1792967566075891712 user=1792959268018716672" {
		t.Fatalf("formatDirectRootTopic() = %q", got)
	}
}

func TestReplyToThreadSendsUserHeaderAsMarkdownHardBreak(t *testing.T) {
	var request Activity
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		case "/v3/conversations/19:channel@thread.tacv2;messageid=root-id/activities":
			if r.Method != http.MethodPost {
				t.Fatalf("method = %s, want POST", r.Method)
			}
			if got := r.Header.Get("Authorization"); got != "Bearer token" {
				t.Fatalf("authorization = %q", got)
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Fatalf("decode request: %v", err)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":"reply-message-id"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       server.URL + "/token",
		ConnectorScope: "scope",
	})
	msg := model.DirectMessage{
		UserID: "1792959268018716672",
		Text:   "こんにちは",
	}

	replyID, err := client.ReplyToThread(t.Context(), server.URL, "19:channel@thread.tacv2", "root-id", msg)
	if err != nil {
		t.Fatalf("ReplyToThread() error = %v", err)
	}
	if replyID != "reply-message-id" {
		t.Fatalf("replyID = %q", replyID)
	}
	if request.TextFormat != "markdown" {
		t.Fatalf("textFormat = %q, want markdown", request.TextFormat)
	}
	if request.Text != "user=1792959268018716672  \nこんにちは" {
		t.Fatalf("activity text = %q", request.Text)
	}
}

func TestTrustedTeamsAttachmentURL(t *testing.T) {
	if !trustedTeamsAttachmentURL("https://smba.trafficmanager.net/amer/v3/attachments/1/views/original") {
		t.Fatalf("expected Bot Framework attachment URL to be trusted")
	}
	if trustedTeamsAttachmentURL("https://example.com/file.png") {
		t.Fatalf("expected unknown host to be rejected")
	}
	if trustedTeamsAttachmentURL("http://smba.trafficmanager.net/amer/v3/attachments/1") {
		t.Fatalf("expected non-HTTPS URL to be rejected")
	}
}

func TestCreateRootThreadSendsHeadingInActivityText(t *testing.T) {
	var request ConversationParameters
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		case "/v3/conversations":
			if r.Method != http.MethodPost {
				t.Fatalf("method = %s, want POST", r.Method)
			}
			if got := r.Header.Get("Authorization"); got != "Bearer token" {
				t.Fatalf("authorization = %q", got)
			}
			if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
				t.Fatalf("decode request: %v", err)
			}
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"activityId":"root-message-id"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer server.Close()

	client := NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       server.URL + "/token",
		ConnectorScope: "scope",
	})
	msg := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		UserID:    "1792959268018716672",
		Text:      "こんにちは",
	}

	rootID, err := client.CreateRootThread(t.Context(), server.URL, ChannelThreadBinding{
		TeamID:         "team-id",
		ChannelID:      "channel-id",
		ConversationID: "conversation-id",
		TenantID:       "tenant-id",
		BotID:          "bot-id",
	}, msg)
	if err != nil {
		t.Fatalf("CreateRootThread() error = %v", err)
	}
	if rootID != "root-message-id" {
		t.Fatalf("rootID = %q", rootID)
	}
	if request.Activity.TopicName != "" {
		t.Fatalf("topicName should not be required for display, got %q", request.Activity.TopicName)
	}
	if request.Activity.Text != "# [direct:bot-trial] room=1792967566075891712 user=1792959268018716672\n\nこんにちは" {
		t.Fatalf("activity text = %q", request.Activity.Text)
	}
	if request.ChannelData.Team.ID != "team-id" ||
		request.ChannelData.Channel.ID != "channel-id" ||
		request.ChannelData.Tenant.ID != "tenant-id" ||
		request.Bot.ID != "bot-id" ||
		request.Conversation.ID != "channel-id" ||
		request.Conversation.ConversationType != "channel" {
		t.Fatalf("unexpected conversation parameters: %+v", request)
	}
}

func TestAddReactionUsesBotConnectorReactionEndpoint(t *testing.T) {
	var reactionPath string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		default:
			if r.Method != http.MethodPut {
				t.Fatalf("method = %s, want PUT", r.Method)
			}
			if got := r.Header.Get("Authorization"); got != "Bearer token" {
				t.Fatalf("authorization = %q", got)
			}
			reactionPath = r.URL.EscapedPath()
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	defer server.Close()

	client := NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       server.URL + "/token",
		ConnectorScope: "scope",
	})
	err := client.AddReaction(t.Context(), server.URL, "conversation;messageid=root", "activity/id", ReactionEyes)
	if err != nil {
		t.Fatalf("AddReaction() error = %v", err)
	}
	want := "/v3/conversations/conversation%3Bmessageid=root/activities/activity%2Fid/reactions/1f440_eyes"
	if reactionPath != want {
		t.Fatalf("reaction path = %q, want %q", reactionPath, want)
	}
}
