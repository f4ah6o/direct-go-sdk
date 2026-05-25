package teams

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
)

func TestAuthValidationBypassedOnlyForLoopbackRemote(t *testing.T) {
	s := &Server{cfg: &config.Config{Bot: config.BotConfig{DisableAuthValidation: true}}}

	local := httptest.NewRequest(http.MethodPost, "/api/messages", nil)
	local.RemoteAddr = "127.0.0.1:12345"
	if !s.authValidationBypassed(local) {
		t.Fatalf("expected local request to bypass auth validation")
	}

	remote := httptest.NewRequest(http.MethodPost, "/api/messages", nil)
	remote.RemoteAddr = "203.0.113.10:12345"
	if s.authValidationBypassed(remote) {
		t.Fatalf("expected remote request not to bypass auth validation")
	}
}

func TestHealthzReportsDirectWorkerStatus(t *testing.T) {
	s := NewServer(
		&config.Config{Bot: config.BotConfig{EndpointPath: "/api/messages"}},
		nil,
		nil,
		nil,
		discardLogger(),
		WithHealthCheck(func() (bool, interface{}) {
			return false, []map[string]interface{}{{"account_id": "account-a", "ready": false}}
		}),
	)
	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()

	s.httpServer.Handler.ServeHTTP(rec, req)

	if rec.Code != http.StatusServiceUnavailable {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusServiceUnavailable)
	}
	var body map[string]interface{}
	if err := json.Unmarshal(rec.Body.Bytes(), &body); err != nil {
		t.Fatalf("healthz response is not json: %v", err)
	}
	if body["ok"] != false {
		t.Fatalf("ok = %v, want false", body["ok"])
	}
	if _, ok := body["direct_accounts"]; !ok {
		t.Fatalf("direct_accounts missing in response: %v", body)
	}
}

func discardLogger() *log.Logger {
	return log.New(io.Discard, "", 0)
}

func TestHandleDirectFileRejectsUnsignedURL(t *testing.T) {
	s := &Server{cfg: &config.Config{Bot: config.BotConfig{AppPassword: "secret"}}}
	req := httptest.NewRequest(http.MethodGet, "/files/direct?account=a&url=https://api.direct4b.com/albero-app-server/files/x", nil)
	rec := httptest.NewRecorder()

	s.handleDirectFile(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestHelpTextIsJapaneseAndDocumentsCommands(t *testing.T) {
	text := rootOnlyHelpText()
	for _, want := range []string{
		"Direct と Teams をつなぐブリッジです。",
		"`@direct bind <alias>`",
		"`@direct unbind <alias>`",
		"`@direct reply <本文>`",
		"`@direct new-thread`",
		"`@direct help`",
	} {
		if !strings.Contains(text, want) {
			t.Fatalf("help text missing %q:\n%s", want, text)
		}
	}
}

func TestThreadReplyRequiresReplyCommand(t *testing.T) {
	responseText, out := processMappedThreadMessage(t, `<at>direct</at> これは誤送信しない`, nil)
	if !strings.Contains(responseText, "`@direct reply <本文>`") {
		t.Fatalf("unexpected usage response: %q", responseText)
	}
	select {
	case got := <-out:
		t.Fatalf("unexpected direct outbound: %+v", got)
	default:
	}
}

func TestThreadReplyCommandForwardsToDirect(t *testing.T) {
	responseText, out := processMappedThreadMessage(t, `<at>direct</at> reply 返信します`, nil)
	if responseText != "" {
		t.Fatalf("unexpected Teams response: %q", responseText)
	}
	select {
	case got := <-out:
		if got.AccountID != "account-a" || got.TalkID != "talk-a" || got.Text != "返信します（Teams User）" {
			t.Fatalf("unexpected outbound: %+v", got)
		}
	case <-time.After(time.Second):
		t.Fatal("timed out waiting for direct outbound")
	}
}

func TestThreadUnknownCommandDoesNotForwardToDirect(t *testing.T) {
	responseText, out := processMappedThreadMessage(t, `<at>direct</at> repyl typo`, nil)
	if !strings.Contains(responseText, "`@direct reply <本文>`") {
		t.Fatalf("unexpected usage response: %q", responseText)
	}
	select {
	case got := <-out:
		t.Fatalf("unexpected direct outbound: %+v", got)
	default:
	}
}

func processMappedThreadMessage(t *testing.T, text string, attachments []Attachment) (string, <-chan model.DirectOutbound) {
	t.Helper()
	var responseText string
	connector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		case strings.Contains(r.URL.Path, "/activities"):
			var activity Activity
			if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
				t.Fatalf("decode activity: %v", err)
			}
			responseText = activity.Text
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":"response-id"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	t.Cleanup(connector.Close)

	st, err := store.Open(t.TempDir() + "/state.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutMapping(store.ThreadMapping{
		AccountID:      "account-a",
		TalkID:         "talk-a",
		ConversationID: "conversation-id",
		ServiceURL:     connector.URL,
		RootID:         "root-id",
	}); err != nil {
		t.Fatal(err)
	}
	client := NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       connector.URL + "/token",
		ConnectorScope: "scope",
	})
	out := make(chan model.DirectOutbound, 1)
	s := &Server{
		cfg:    &config.Config{},
		client: client,
		store:  st,
		out:    out,
		logger: discardLogger(),
	}
	activity := Activity{
		Type:        "message",
		ID:          "activity-id",
		ServiceURL:  connector.URL,
		Text:        text,
		From:        ChannelAccount{ID: "teams-user", Name: "Teams User"},
		Recipient:   ChannelAccount{ID: "direct-bot", Name: "direct"},
		Attachments: attachments,
		Conversation: ConversationAccount{
			ID: "conversation-id;messageid=root-id",
		},
		ChannelData: ChannelData{Channel: ChannelInfo{ID: "conversation-id"}},
		Entities: []Entity{{
			Type:      "mention",
			Text:      "<at>direct</at>",
			Mentioned: ChannelAccount{ID: "direct-bot", Name: "direct"},
		}},
	}
	s.processActivity(t.Context(), activity)
	return responseText, out
}

func TestHelpCommandSendsJapaneseHelp(t *testing.T) {
	var responseText string
	connector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		case strings.Contains(r.URL.Path, "/activities"):
			var activity Activity
			if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
				t.Fatalf("decode activity: %v", err)
			}
			responseText = activity.Text
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":"response-id"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer connector.Close()

	client := NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       connector.URL + "/token",
		ConnectorScope: "scope",
	})
	s := &Server{client: client, logger: discardLogger()}
	activity := Activity{
		Type:       "message",
		ID:         "activity-id",
		ServiceURL: connector.URL,
		Text:       `<at>direct</at> help`,
		From:       ChannelAccount{ID: "teams-user", Name: "Teams User"},
		Recipient:  ChannelAccount{ID: "direct-bot", Name: "direct"},
		Conversation: ConversationAccount{
			ID: "conversation-id",
		},
		Entities: []Entity{{
			Type:      "mention",
			Text:      "<at>direct</at>",
			Mentioned: ChannelAccount{ID: "direct-bot", Name: "direct"},
		}},
	}

	s.processActivity(t.Context(), activity)

	if !strings.Contains(responseText, "Direct と Teams をつなぐブリッジです。") ||
		!strings.Contains(responseText, "`@direct new-thread`") {
		t.Fatalf("unexpected help response: %q", responseText)
	}
}

func TestNewThreadCommandForgetsMappedThread(t *testing.T) {
	var responseText string
	connector := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		case strings.Contains(r.URL.Path, "/activities"):
			var activity Activity
			if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
				t.Fatalf("decode activity: %v", err)
			}
			responseText = activity.Text
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":"response-id"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	defer connector.Close()

	st, err := store.Open(t.TempDir() + "/state.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutMapping(store.ThreadMapping{
		AccountID:      "account-a",
		TalkID:         "talk-a",
		ConversationID: "conversation-id",
		ServiceURL:     connector.URL,
		RootID:         "root-id",
	}); err != nil {
		t.Fatal(err)
	}
	client := NewClient(config.BotConfig{
		AppID:          "app-id",
		AppPassword:    "secret",
		TokenURL:       connector.URL + "/token",
		ConnectorScope: "scope",
	})
	s := &Server{client: client, store: st, logger: discardLogger()}
	activity := Activity{
		Type:       "message",
		ID:         "activity-id",
		ServiceURL: connector.URL,
		Text:       `<at>direct</at> new-thread`,
		From:       ChannelAccount{ID: "teams-user", Name: "Teams User"},
		Recipient:  ChannelAccount{ID: "direct-bot", Name: "direct"},
		Conversation: ConversationAccount{
			ID: "conversation-id;messageid=root-id",
		},
		ChannelData: ChannelData{Channel: ChannelInfo{ID: "conversation-id"}},
		Entities: []Entity{{
			Type:      "mention",
			Text:      "<at>direct</at>",
			Mentioned: ChannelAccount{ID: "direct-bot", Name: "direct"},
		}},
	}

	s.processActivity(t.Context(), activity)

	if _, ok := st.GetByTalk("account-a", "talk-a"); ok {
		t.Fatalf("mapping remained after new-thread command")
	}
	if _, ok := st.GetByThread("conversation-id", "root-id"); ok {
		t.Fatalf("thread index remained after new-thread command")
	}
	if !strings.Contains(responseText, "next Direct message") {
		t.Fatalf("unexpected response text: %q", responseText)
	}
}
