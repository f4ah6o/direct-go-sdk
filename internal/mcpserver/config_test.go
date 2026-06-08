package mcpserver

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfigDefaultsAndEnvExpansion(t *testing.T) {
	t.Setenv("MCP_PUBLIC", "http://localhost:8090")
	t.Setenv("JWT_ISSUER", "http://localhost:9999")
	path := filepath.Join(t.TempDir(), "mcp.yaml")
	err := os.WriteFile(path, []byte(`
mcp:
  public_base_url: ${MCP_PUBLIC}
  jwt_issuer: ${JWT_ISSUER}
  jwt_audience: http://localhost:8090/mcp
  jwks_url: http://localhost:9999/jwks
accounts:
  - id: account-a
    token_ref: op://vault/item/token
`), 0600)
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := LoadConfig(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.MCP.ListenAddr != DefaultListenAddr {
		t.Fatalf("listen addr = %q", cfg.MCP.ListenAddr)
	}
	if cfg.MCP.EndpointPath != DefaultEndpointPath {
		t.Fatalf("endpoint path = %q", cfg.MCP.EndpointPath)
	}
	if cfg.MCP.ReadScope != DefaultReadScope || cfg.MCP.WriteScope != DefaultWriteScope {
		t.Fatalf("scope defaults = %q/%q", cfg.MCP.ReadScope, cfg.MCP.WriteScope)
	}
	if cfg.Accounts[0].TokenEnv != "DIRECT_TOKEN_ACCOUNT_A" {
		t.Fatalf("token env = %q", cfg.Accounts[0].TokenEnv)
	}
	if got := cfg.ResourceURL(); got != "http://localhost:8090/mcp" {
		t.Fatalf("resource url = %q", got)
	}
}

func TestValidateRequiresOAuthResourceFields(t *testing.T) {
	cfg := Config{Accounts: []AccountConfig{{ID: "account-a", TokenEnv: "DIRECT_TOKEN_ACCOUNT_A"}}}
	cfg.Defaults()
	if err := cfg.Validate(); err == nil {
		t.Fatalf("expected validation error")
	}
	cfg.MCP.PublicBaseURL = "http://localhost:8090"
	cfg.MCP.JWTIssuer = "http://localhost:9999"
	cfg.MCP.JWTAudience = "http://localhost:8090/mcp"
	cfg.MCP.JWKSURL = "http://localhost:9999/jwks"
	if err := cfg.Validate(); err != nil {
		t.Fatalf("Validate() = %v", err)
	}
}
