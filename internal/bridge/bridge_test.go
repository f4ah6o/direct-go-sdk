package bridge

import (
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/teams"
)

func TestConsumePendingDirectMessageSuppressesTeamsToDirectEcho(t *testing.T) {
	s := NewService(nil, nil, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	s.markPendingDirectMessage(outbound)

	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		UserID:    "1792959268018716672",
		Text:      "了解です（Taro Yamada）",
		MessageID: "direct-message-id",
	}
	if _, ok := s.consumePendingDirectMessage(inbound); !ok {
		t.Fatalf("expected pending direct message to be consumed")
	}
	if _, ok := s.consumePendingDirectMessage(inbound); ok {
		t.Fatalf("pending direct message should only be consumed once")
	}
}

func TestClearPendingDirectMessageAllowsLaterUserMessage(t *testing.T) {
	s := NewService(nil, nil, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	s.markPendingDirectMessage(outbound)
	s.clearPendingDirectMessage(outbound)

	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	if _, ok := s.consumePendingDirectMessage(inbound); ok {
		t.Fatalf("cleared pending direct message should not be consumed")
	}
}

func TestSuccessfulDirectSentKeepsPendingUntilDirectNotification(t *testing.T) {
	s := NewService(nil, nil, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	s.markPendingDirectMessage(outbound)

	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "  了解です（Taro Yamada）\r\n",
		MessageID: "direct-message-id",
	}
	if _, ok := s.consumePendingDirectMessage(inbound); !ok {
		t.Fatalf("successful send pending marker should remain until direct notification is consumed")
	}
}

func TestConsumedPendingDirectMessageStoresTeamsSourceRef(t *testing.T) {
	st, err := store.Open(filepath.Join(t.TempDir(), "state.json"))
	if err != nil {
		t.Fatal(err)
	}
	s := NewService(nil, st, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "talk-a",
		Text:      "hello",
		TeamsSource: &model.TeamsSource{
			ServiceURL:     "https://service.example",
			ConversationID: "conversation-id",
			ActivityID:     "teams-activity-id",
		},
	}
	s.markPendingDirectMessage(outbound)
	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "talk-a",
		UserID:    "direct-self",
		Text:      "hello",
		MessageID: "direct-message-id",
	}
	consumed, ok := s.consumePendingDirectMessage(inbound)
	if !ok {
		t.Fatalf("expected pending direct message to be consumed")
	}
	s.storeTeamsToDirectMessageRef(t.Context(), consumed, inbound)
	ref, ok := st.GetTeamsMessageRef("bot-trial", "direct-message-id")
	if !ok {
		t.Fatalf("expected teams ref to be stored")
	}
	if ref.ActivityID != "teams-activity-id" || ref.DirectSenderID != "direct-self" {
		t.Fatalf("unexpected ref: %+v", ref)
	}
}

func TestHandleDirectReadReceiptAddsTeamsReactionOnce(t *testing.T) {
	reactions := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		default:
			if r.Method != http.MethodPut {
				t.Fatalf("method = %s, want PUT", r.Method)
			}
			reactions++
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	defer server.Close()

	st, err := store.Open(filepath.Join(t.TempDir(), "state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutTeamsMessageRef(store.TeamsMessageRef{
		AccountID:       "account-a",
		TalkID:          "talk-a",
		DirectMessageID: "direct-msg",
		ServiceURL:      server.URL,
		ConversationID:  "conversation-id",
		ActivityID:      "teams-activity",
		DirectSenderID:  "direct-self",
	}); err != nil {
		t.Fatal(err)
	}
	client := teams.NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       server.URL + "/token",
		ConnectorScope: "scope",
	})
	s := NewService(nil, st, client, nil, nil, nil, nil, nil, log.Default())
	receipt := model.DirectReadReceipt{
		AccountID:   "account-a",
		TalkID:      "talk-a",
		MessageIDs:  []string{"direct-msg"},
		ReadUserIDs: []string{"direct-self", "direct-reader"},
	}
	if err := s.handleDirectReadReceipt(t.Context(), receipt); err != nil {
		t.Fatal(err)
	}
	if err := s.handleDirectReadReceipt(t.Context(), receipt); err != nil {
		t.Fatal(err)
	}
	if reactions != 1 {
		t.Fatalf("reactions = %d, want 1", reactions)
	}
}

func TestPendingReadReceiptReactsWhenTeamsRefIsStoredLater(t *testing.T) {
	reactions := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		default:
			reactions++
			w.WriteHeader(http.StatusNoContent)
		}
	}))
	defer server.Close()

	st, err := store.Open(filepath.Join(t.TempDir(), "state.json"))
	if err != nil {
		t.Fatal(err)
	}
	client := teams.NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       server.URL + "/token",
		ConnectorScope: "scope",
	})
	s := NewService(nil, st, client, nil, nil, nil, nil, nil, log.Default())
	receipt := model.DirectReadReceipt{
		AccountID:   "account-a",
		TalkID:      "talk-a",
		MessageIDs:  []string{"direct-msg"},
		ReadUserIDs: []string{"direct-reader"},
	}
	if err := s.handleDirectReadReceipt(t.Context(), receipt); err != nil {
		t.Fatal(err)
	}
	if reactions != 0 {
		t.Fatalf("reaction before ref = %d, want 0", reactions)
	}
	s.storeDirectToTeamsMessageRef(t.Context(), store.TeamsMessageRef{
		AccountID:       "account-a",
		TalkID:          "talk-a",
		DirectMessageID: "direct-msg",
		ServiceURL:      server.URL,
		ConversationID:  "conversation-id",
		ActivityID:      "teams-activity",
		DirectSenderID:  "direct-self",
	})
	if reactions != 1 {
		t.Fatalf("reactions = %d, want 1", reactions)
	}
}

func TestHandleDirectReadReceiptIgnoresSelfOnlyRead(t *testing.T) {
	st, err := store.Open(filepath.Join(t.TempDir(), "state.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutTeamsMessageRef(store.TeamsMessageRef{
		AccountID:       "account-a",
		TalkID:          "talk-a",
		DirectMessageID: "direct-msg",
		ServiceURL:      "https://service.example",
		ConversationID:  "conversation-id",
		ActivityID:      "teams-activity",
		DirectSenderID:  "direct-self",
	}); err != nil {
		t.Fatal(err)
	}
	s := NewService(nil, st, nil, nil, nil, nil, nil, nil, log.Default())
	err = s.handleDirectReadReceipt(t.Context(), model.DirectReadReceipt{
		AccountID:   "account-a",
		TalkID:      "talk-a",
		MessageIDs:  []string{"direct-msg"},
		ReadUserIDs: []string{"direct-self"},
	})
	if err != nil {
		t.Fatal(err)
	}
	ref, ok := st.GetTeamsMessageRef("account-a", "direct-msg")
	if !ok || !ref.ReactedAt.IsZero() {
		t.Fatalf("self-only read should not react: %+v ok=%v", ref, ok)
	}
}
