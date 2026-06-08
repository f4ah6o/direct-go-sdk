package mcpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHTTPHandlerServesProtectedResourceMetadataAnd401Challenge(t *testing.T) {
	t.Setenv("DIRECT_TOKEN_ACCOUNT_A", "token-a")
	cfg := testConfig()
	srv, err := New(context.Background(), cfg, log.New(testWriter{t}, "", 0))
	if err != nil {
		t.Fatal(err)
	}
	httpSrv := httptest.NewServer(srv.HTTPHandler())
	defer httpSrv.Close()

	resp, err := http.Get(httpSrv.URL + "/.well-known/oauth-protected-resource")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("metadata status = %d", resp.StatusCode)
	}
	var meta struct {
		Resource        string   `json:"resource"`
		ScopesSupported []string `json:"scopes_supported"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&meta); err != nil {
		t.Fatal(err)
	}
	if meta.Resource != "http://localhost:8090/mcp" {
		t.Fatalf("metadata resource = %q", meta.Resource)
	}
	if len(meta.ScopesSupported) != 2 {
		t.Fatalf("metadata scopes = %#v", meta.ScopesSupported)
	}

	req, err := http.NewRequest(http.MethodPost, httpSrv.URL+"/mcp", nil)
	if err != nil {
		t.Fatal(err)
	}
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusUnauthorized {
		t.Fatalf("mcp status = %d", resp.StatusCode)
	}
	if got := resp.Header.Get("WWW-Authenticate"); got == "" {
		t.Fatalf("missing WWW-Authenticate challenge")
	}
}

func TestListAccountsRequiresReadScope(t *testing.T) {
	t.Setenv("DIRECT_TOKEN_ACCOUNT_A", "token-a")
	srv, err := New(context.Background(), testConfig(), log.New(testWriter{t}, "", 0))
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.WithValue(context.Background(), authInfoKey, &AuthInfo{Subject: "user-1", Scopes: map[string]bool{DefaultReadScope: true}})
	_, out, err := srv.listAccounts(ctx, nil, struct{}{})
	if err != nil {
		t.Fatalf("listAccounts() = %v", err)
	}
	accounts := out.Accounts
	if len(accounts) != 1 || accounts[0].ID != "account-a" || !accounts[0].Available {
		t.Fatalf("accounts = %#v", accounts)
	}
	ctx = context.WithValue(context.Background(), authInfoKey, &AuthInfo{Subject: "user-1", Scopes: map[string]bool{}})
	if _, _, err := srv.listAccounts(ctx, nil, struct{}{}); err != ErrForbidden {
		t.Fatalf("listAccounts without scope = %v, want ErrForbidden", err)
	}
}

func testConfig() *Config {
	cfg := &Config{
		MCP: MCPConfig{
			PublicBaseURL: "http://localhost:8090",
			JWTIssuer:     "https://auth.example.com",
			JWTAudience:   "http://localhost:8090/mcp",
			JWKSURL:       "https://auth.example.com/jwks",
		},
		Accounts: []AccountConfig{{ID: "account-a", TokenEnv: "DIRECT_TOKEN_ACCOUNT_A"}},
	}
	cfg.Defaults()
	return cfg
}

type testWriter struct {
	t *testing.T
}

func (w testWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
