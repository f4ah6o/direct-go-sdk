package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OP            OPConfig                      `yaml:"op"`
	TeamsChannels map[string]TeamsChannelConfig `yaml:"teams_channels"`
	Accounts      []AccountConfig               `yaml:"accounts"`
	Graph         GraphConfig                   `yaml:"graph"`
	State         StateConfig                   `yaml:"state"`
	Server        ServerConfig                  `yaml:"server"`
	Queues        QueueConfig                   `yaml:"queues"`
	Retry         RetryConfig                   `yaml:"retry"`
	Attachments   AttachmentConfig              `yaml:"attachments"`
}

type TeamsChannelConfig struct {
	TeamID        string `yaml:"team_id"`
	ChannelID     string `yaml:"channel_id"`
	MentionUserID string `yaml:"mention_user_id"`
}

type AccountConfig struct {
	ID           string `yaml:"id"`
	TokenEnv     string `yaml:"token_env"`
	TokenRef     string `yaml:"token_ref"`
	Endpoint     string `yaml:"endpoint"`
	ProxyURL     string `yaml:"proxy_url"`
	TeamsChannel string `yaml:"teams_channel"`
}

type OPConfig struct {
	Binary string `yaml:"binary"`
}

type GraphConfig struct {
	TenantID        string   `yaml:"tenant_id"`
	ClientID        string   `yaml:"client_id"`
	ClientSecret    string   `yaml:"client_secret"`
	ClientSecretEnv string   `yaml:"client_secret_env"`
	AccessTokenEnv  string   `yaml:"access_token_env"`
	Scopes          []string `yaml:"scopes"`
	ClientState     string   `yaml:"client_state"`
	TokenURL        string   `yaml:"token_url"`
	APIBaseURL      string   `yaml:"api_base_url"`
}

type StateConfig struct {
	Path string `yaml:"path"`
}

type ServerConfig struct {
	ListenAddr    string `yaml:"listen_addr"`
	PublicBaseURL string `yaml:"public_base_url"`
}

type QueueConfig struct {
	DirectToTeams int `yaml:"direct_to_teams"`
	TeamsToDirect int `yaml:"teams_to_direct"`
}

type RetryConfig struct {
	MaxAttempts int    `yaml:"max_attempts"`
	MaxBackoff  string `yaml:"max_backoff"`
}

type AttachmentConfig struct {
	MaxBytes int64  `yaml:"max_bytes"`
	TempDir  string `yaml:"temp_dir"`
}

func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	expanded := os.ExpandEnv(string(data))
	var cfg Config
	if err := yaml.Unmarshal([]byte(expanded), &cfg); err != nil {
		return nil, err
	}
	cfg.Defaults()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (c *Config) Defaults() {
	if c.Server.ListenAddr == "" {
		c.Server.ListenAddr = ":8080"
	}
	if c.State.Path == "" {
		c.State.Path = "./state/direct-teams-bridge.json"
	}
	if c.Queues.DirectToTeams == 0 {
		c.Queues.DirectToTeams = 1000
	}
	if c.Queues.TeamsToDirect == 0 {
		c.Queues.TeamsToDirect = 1000
	}
	if c.Retry.MaxAttempts == 0 {
		c.Retry.MaxAttempts = 5
	}
	if c.Retry.MaxBackoff == "" {
		c.Retry.MaxBackoff = "30s"
	}
	if c.Attachments.MaxBytes == 0 {
		c.Attachments.MaxBytes = 25 << 20
	}
	if c.Graph.APIBaseURL == "" {
		c.Graph.APIBaseURL = "https://graph.microsoft.com/v1.0"
	}
	if c.Graph.TokenURL == "" && c.Graph.TenantID != "" {
		c.Graph.TokenURL = fmt.Sprintf("https://login.microsoftonline.com/%s/oauth2/v2.0/token", c.Graph.TenantID)
	}
}

func (c *Config) Validate() error {
	if len(c.Accounts) == 0 {
		return errors.New("at least one account is required")
	}
	if len(c.TeamsChannels) == 0 {
		return errors.New("at least one teams channel is required")
	}
	if c.Graph.ClientState == "" {
		return errors.New("graph.client_state is required")
	}
	if c.Graph.ClientID == "" {
		return errors.New("graph.client_id is required")
	}
	if c.Graph.TokenURL == "" {
		return errors.New("graph.token_url or graph.tenant_id is required")
	}
	for name, ch := range c.TeamsChannels {
		if strings.TrimSpace(ch.TeamID) == "" || strings.TrimSpace(ch.ChannelID) == "" {
			return fmt.Errorf("teams channel %q requires team_id and channel_id", name)
		}
		if strings.TrimSpace(ch.MentionUserID) == "" {
			return fmt.Errorf("teams channel %q requires mention_user_id", name)
		}
	}
	seen := map[string]bool{}
	for _, account := range c.Accounts {
		if account.ID == "" {
			return errors.New("account id is required")
		}
		if seen[account.ID] {
			return fmt.Errorf("duplicate account id %q", account.ID)
		}
		seen[account.ID] = true
		if account.TokenEnv == "" && account.TokenRef == "" {
			return fmt.Errorf("account %q requires token_env or token_ref", account.ID)
		}
		if _, ok := c.TeamsChannels[account.TeamsChannel]; !ok {
			return fmt.Errorf("account %q references unknown teams channel %q", account.ID, account.TeamsChannel)
		}
	}
	return nil
}

func (c *Config) Account(id string) (AccountConfig, bool) {
	for _, account := range c.Accounts {
		if account.ID == id {
			return account, true
		}
	}
	return AccountConfig{}, false
}
