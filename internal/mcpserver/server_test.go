package mcpserver

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
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
			SingleTenant:  true,
		},
		Accounts: []AccountConfig{{ID: "account-a", TokenEnv: "DIRECT_TOKEN_ACCOUNT_A"}},
	}
	cfg.Defaults()
	return cfg
}

func authCtx(subject string, principals []string, scopes ...string) context.Context {
	return context.WithValue(context.Background(), authInfoKey, &AuthInfo{
		Subject:    subject,
		Scopes:     parseScopes(scopes...),
		Principals: principals,
	})
}

func multiAccountConfig(t *testing.T) *Config {
	t.Helper()
	t.Setenv("DIRECT_TOKEN_ACCOUNT_A", "token-a")
	t.Setenv("DIRECT_TOKEN_ACCOUNT_B", "token-b")
	cfg := &Config{
		MCP: MCPConfig{
			PublicBaseURL: "http://localhost:8090",
			JWTIssuer:     "https://auth.example.com",
			JWTAudience:   "http://localhost:8090/mcp",
			JWKSURL:       "https://auth.example.com/jwks",
			AccountSubjects: map[string][]string{
				"alice": {"account-a"},
				"*":     {"account-b"},
			},
		},
		Accounts: []AccountConfig{
			{ID: "account-a", TokenEnv: "DIRECT_TOKEN_ACCOUNT_A"},
			{ID: "account-b", TokenEnv: "DIRECT_TOKEN_ACCOUNT_B"},
		},
	}
	cfg.Defaults()
	if err := cfg.Validate(); err != nil {
		t.Fatal(err)
	}
	return cfg
}

func TestListAccountsFiltersByPrincipal(t *testing.T) {
	srv, err := New(context.Background(), multiAccountConfig(t), log.New(testWriter{t}, "", 0))
	if err != nil {
		t.Fatal(err)
	}

	// alice sees her mapped account plus the wildcard account.
	_, out, err := srv.listAccounts(authCtx("alice", []string{"alice"}, DefaultReadScope), nil, struct{}{})
	if err != nil {
		t.Fatalf("listAccounts() = %v", err)
	}
	if len(out.Accounts) != 2 {
		t.Fatalf("alice accounts = %#v", out.Accounts)
	}

	// carol has no mapping and only sees the wildcard account.
	_, out, err = srv.listAccounts(authCtx("carol", []string{"carol"}, DefaultReadScope), nil, struct{}{})
	if err != nil {
		t.Fatalf("listAccounts() = %v", err)
	}
	if len(out.Accounts) != 1 || out.Accounts[0].ID != "account-b" {
		t.Fatalf("carol accounts = %#v", out.Accounts)
	}
}

func TestAccountToolsEnforceAuthorization(t *testing.T) {
	srv, err := New(context.Background(), multiAccountConfig(t), log.New(testWriter{t}, "", 0))
	if err != nil {
		t.Fatal(err)
	}
	srv.clientFactory = func(direct.Options) directClient { return &stubDirectClient{} }

	ctx := authCtx("carol", []string{"carol"}, DefaultReadScope)
	_, _, err = srv.getMe(ctx, nil, accountArgs{AccountID: "account-a"})
	if err == nil || !strings.Contains(err.Error(), "not authorized") {
		t.Fatalf("getMe denied = %v", err)
	}
	if strings.Contains(err.Error(), "account-a") {
		t.Fatalf("denied error leaks account id: %v", err)
	}

	// Scope alone never grants account access: a token holding both scopes
	// still cannot reach an account it is not mapped to.
	_, _, err = srv.getMe(authCtx("carol", []string{"carol"}, DefaultReadScope, DefaultWriteScope), nil, accountArgs{AccountID: "account-a"})
	if err == nil || !strings.Contains(err.Error(), "not authorized") {
		t.Fatalf("getMe with write scope = %v", err)
	}

	// Wildcard-mapped account is reachable.
	if _, _, err = srv.getMe(ctx, nil, accountArgs{AccountID: "account-b"}); err != nil {
		t.Fatalf("getMe wildcard account = %v", err)
	}
}

func TestAccountToolsFailClosedWithoutMapping(t *testing.T) {
	t.Setenv("DIRECT_TOKEN_ACCOUNT_A", "token-a")
	cfg := testConfig()
	cfg.MCP.SingleTenant = false
	srv, err := New(context.Background(), cfg, log.New(testWriter{t}, "", 0))
	if err != nil {
		t.Fatal(err)
	}
	ctx := authCtx("alice", []string{"alice"}, DefaultReadScope)
	if _, _, err := srv.getMe(ctx, nil, accountArgs{AccountID: "account-a"}); err == nil {
		t.Fatalf("expected denial when no account_subjects configured")
	}
	_, out, err := srv.listAccounts(ctx, nil, struct{}{})
	if err != nil {
		t.Fatal(err)
	}
	if len(out.Accounts) != 0 {
		t.Fatalf("listAccounts leaked %d accounts", len(out.Accounts))
	}
}

func TestSingleTenantAllowsAllAccounts(t *testing.T) {
	t.Setenv("DIRECT_TOKEN_ACCOUNT_A", "token-a")
	srv, err := New(context.Background(), testConfig(), log.New(testWriter{t}, "", 0))
	if err != nil {
		t.Fatal(err)
	}
	srv.clientFactory = func(direct.Options) directClient { return &stubDirectClient{} }
	ctx := authCtx("anyone", nil, DefaultReadScope)
	if _, _, err := srv.getMe(ctx, nil, accountArgs{AccountID: "account-a"}); err != nil {
		t.Fatalf("single tenant getMe = %v", err)
	}
}

type stubDirectClient struct{}

func (*stubDirectClient) Connect() error                           { return nil }
func (*stubDirectClient) ConnectWithContext(context.Context) error { return nil }
func (*stubDirectClient) Close() error                             { return nil }
func (*stubDirectClient) GetMeWithContext(context.Context) (*direct.UserInfo, error) {
	return &direct.UserInfo{}, nil
}
func (*stubDirectClient) GetDomainsWithContext(context.Context) ([]direct.DomainInfo, error) {
	return nil, nil
}
func (*stubDirectClient) GetTalksWithContext(context.Context) ([]direct.Talk, error) {
	return nil, nil
}
func (*stubDirectClient) GetMessages(context.Context, interface{}, interface{}, *direct.GetMessagesOptions) ([]direct.ReceivedMessage, error) {
	return nil, nil
}
func (*stubDirectClient) SearchMessages(context.Context, interface{}, interface{}, string, interface{}, int) (*direct.SearchMessagesResult, error) {
	return nil, nil
}
func (*stubDirectClient) CreateTextMessageWithContext(context.Context, string, string) (string, error) {
	return "msg-1", nil
}

type testWriter struct {
	t *testing.T
}

func (w testWriter) Write(p []byte) (int, error) {
	w.t.Log(string(p))
	return len(p), nil
}
