package slackcompat

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

type fakeDirect struct {
	accountID string
	me        *direct.UserInfo
	talks     []direct.Talk
	users     []direct.UserInfo
	messages  []direct.ReceivedMessage
	sentTalk  string
	sentText  string
}

func (f *fakeDirect) AccountID() string { return f.accountID }

func (f *fakeDirect) GetMe(context.Context) (*direct.UserInfo, error) {
	return f.me, nil
}

func (f *fakeDirect) GetTalks(context.Context) ([]direct.Talk, error) {
	return f.talks, nil
}

func (f *fakeDirect) GetUsers(context.Context) ([]direct.UserInfo, error) {
	return f.users, nil
}

func (f *fakeDirect) GetMessages(context.Context, interface{}, interface{}, *direct.GetMessagesOptions) ([]direct.ReceivedMessage, error) {
	return f.messages, nil
}

func (f *fakeDirect) SendText(_ context.Context, talkID, text string) (string, error) {
	f.sentTalk = talkID
	f.sentText = text
	return "msg-123", nil
}

func TestChatPostMessage(t *testing.T) {
	fake := fixtureDirect()
	server := NewServer([]DirectAPI{fake})
	channel := NewMapper().ChannelID("account-a", "talk-1")
	req := httptest.NewRequest(http.MethodPost, "/api/chat.postMessage", strings.NewReader("channel="+channel+"&text=hello"))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d", rec.Code)
	}
	var body PostMessageResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if !body.OK || body.Channel != channel || body.Message.Text != "hello" {
		t.Fatalf("response = %+v", body)
	}
	if fake.sentTalk != "talk-1" || fake.sentText != "hello" {
		t.Fatalf("sent talk=%q text=%q", fake.sentTalk, fake.sentText)
	}
}

func TestAuthTest(t *testing.T) {
	server := NewServer([]DirectAPI{fixtureDirect()}, WithTeam("T1", "Direct Team"))
	req := httptest.NewRequest(http.MethodGet, "/api/auth.test", nil)
	rec := httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	var body AuthTestResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if !body.OK || body.TeamID != "T1" || body.User != "Bot User" {
		t.Fatalf("response = %+v", body)
	}
}

func TestBearerTokenAuth(t *testing.T) {
	server := NewServer([]DirectAPI{fixtureDirect()}, WithBearerToken("secret"))
	req := httptest.NewRequest(http.MethodGet, "/api/auth.test", nil)
	rec := httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	var denied SlackResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &denied); err != nil {
		t.Fatal(err)
	}
	if denied.OK || denied.Error != "not_authed" {
		t.Fatalf("denied response = %+v", denied)
	}

	req = httptest.NewRequest(http.MethodGet, "/api/auth.test", nil)
	req.Header.Set("Authorization", "Bearer secret")
	rec = httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	var allowed AuthTestResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &allowed); err != nil {
		t.Fatal(err)
	}
	if !allowed.OK {
		t.Fatalf("allowed response = %+v", allowed)
	}
}

func TestConversationsList(t *testing.T) {
	server := NewServer([]DirectAPI{fixtureDirect()})
	req := httptest.NewRequest(http.MethodGet, "/api/conversations.list", nil)
	rec := httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	var body ConversationsListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if !body.OK || len(body.Channels) != 1 {
		t.Fatalf("response = %+v", body)
	}
	if body.Channels[0].Name != "support-room" {
		t.Fatalf("channel name = %q", body.Channels[0].Name)
	}
}

func TestUsersList(t *testing.T) {
	server := NewServer([]DirectAPI{fixtureDirect()})
	req := httptest.NewRequest(http.MethodPost, "/api/users.list", bytes.NewReader([]byte(`{}`)))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	var body UsersListResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if !body.OK || len(body.Members) != 1 || body.Members[0].Profile.Email != "user@example.com" {
		t.Fatalf("response = %+v", body)
	}
}

func TestConversationsHistory(t *testing.T) {
	server := NewServer([]DirectAPI{fixtureDirect()})
	channel := NewMapper().ChannelID("account-a", "talk-1")
	req := httptest.NewRequest(http.MethodGet, "/api/conversations.history?channel="+channel+"&limit=1", nil)
	rec := httptest.NewRecorder()

	server.HTTPHandler().ServeHTTP(rec, req)

	var body ConversationsHistoryResponse
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if !body.OK || len(body.Messages) != 1 || body.Messages[0].Text != "old hello" {
		t.Fatalf("response = %+v", body)
	}
}

func fixtureDirect() *fakeDirect {
	return &fakeDirect{
		accountID: "account-a",
		me:        &direct.UserInfo{ID: "bot-user", DisplayName: "Bot User"},
		talks: []direct.Talk{{
			ID:       "talk-1",
			DomainID: "domain-1",
			Type:     int(direct.RoomTypeGroup),
			Name:     "Support Room",
			UserIDs:  []interface{}{"user-1", "bot-user"},
		}},
		users: []direct.UserInfo{{
			ID:          "user-1",
			Name:        "User One",
			DisplayName: "User One",
			Email:       "user@example.com",
		}},
		messages: []direct.ReceivedMessage{{
			ID:        "msg-old",
			TalkID:    "talk-1",
			UserID:    "user-1",
			Text:      "old hello",
			Timestamp: time.Unix(1710000000, 0).UTC(),
		}},
	}
}
