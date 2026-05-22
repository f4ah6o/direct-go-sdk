package teams

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
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
