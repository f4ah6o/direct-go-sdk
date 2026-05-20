package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadExpandsEnvAndDefaults(t *testing.T) {
	t.Setenv("BOT_APP_ID", "bot-1")
	path := filepath.Join(t.TempDir(), "config.yaml")
	err := os.WriteFile(path, []byte(`
bot:
  app_id: ${BOT_APP_ID}
  app_password_env: MICROSOFT_APP_PASSWORD
teams_channels:
  support: {}
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
	if cfg.Bot.AppID != "bot-1" {
		t.Fatalf("bot app id was not expanded")
	}
	if cfg.Server.ListenAddr != ":8080" {
		t.Fatalf("default listen addr = %q", cfg.Server.ListenAddr)
	}
	if cfg.Bot.EndpointPath != "/api/messages" {
		t.Fatalf("endpoint path default = %q", cfg.Bot.EndpointPath)
	}
}

func TestValidateRejectsUnknownTeamsChannel(t *testing.T) {
	cfg := Config{
		Bot: BotConfig{AppID: "bot", AppPassword: "secret"},
		TeamsChannels: map[string]TeamsChannelConfig{
			"support": {},
		},
		Accounts: []AccountConfig{{ID: "account-a", TokenEnv: "TOKEN", TeamsChannel: "missing"}},
	}
	cfg.Defaults()
	if err := cfg.Validate(); err == nil {
		t.Fatalf("expected validation error")
	}
}
