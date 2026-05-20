package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadExpandsEnvAndDefaults(t *testing.T) {
	t.Setenv("TEAM_ID", "team-1")
	t.Setenv("CHANNEL_ID", "channel-1")
	t.Setenv("MENTION_ID", "user-1")
	t.Setenv("GRAPH_CLIENT_ID", "client-1")
	t.Setenv("GRAPH_CLIENT_STATE", "state-1")
	path := filepath.Join(t.TempDir(), "config.yaml")
	err := os.WriteFile(path, []byte(`
graph:
  tenant_id: tenant-1
  client_id: ${GRAPH_CLIENT_ID}
  client_state: ${GRAPH_CLIENT_STATE}
teams_channels:
  support:
    team_id: ${TEAM_ID}
    channel_id: ${CHANNEL_ID}
    mention_user_id: ${MENTION_ID}
accounts:
  - id: account-a
    token_env: DIRECT_TOKEN_ACCOUNT_A
    teams_channel: support
`), 0600)
	if err != nil {
		t.Fatal(err)
	}
	cfg, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if cfg.TeamsChannels["support"].TeamID != "team-1" {
		t.Fatalf("team id was not expanded")
	}
	if cfg.Server.ListenAddr != ":8080" {
		t.Fatalf("default listen addr = %q", cfg.Server.ListenAddr)
	}
	if cfg.Graph.TokenURL == "" {
		t.Fatalf("token url default was not set")
	}
}

func TestValidateRejectsUnknownTeamsChannel(t *testing.T) {
	cfg := Config{
		Graph: GraphConfig{ClientID: "client", ClientState: "state", TokenURL: "https://example.invalid/token"},
		TeamsChannels: map[string]TeamsChannelConfig{
			"support": {TeamID: "team", ChannelID: "channel", MentionUserID: "user"},
		},
		Accounts: []AccountConfig{{ID: "account-a", TokenEnv: "TOKEN", TeamsChannel: "missing"}},
	}
	cfg.Defaults()
	if err := cfg.Validate(); err == nil {
		t.Fatalf("expected validation error")
	}
}
