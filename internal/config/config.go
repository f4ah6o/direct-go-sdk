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
	Bot           BotConfig                     `yaml:"bot"`
	State         StateConfig                   `yaml:"state"`
	Server        ServerConfig                  `yaml:"server"`
	Queues        QueueConfig                   `yaml:"queues"`
	Retry         RetryConfig                   `yaml:"retry"`
	Attachments   AttachmentConfig              `yaml:"attachments"`
}

type TeamsChannelConfig struct {
	// Alias-only config. The actual Teams conversation is bound at runtime by
	// sending "@bot bind <alias>" in the target channel.
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

type BotConfig struct {
	AppID                 string   `yaml:"app_id"`
	AppPassword           string   `yaml:"app_password"`
	AppPasswordEnv        string   `yaml:"app_password_env"`
	AppPasswordRef        string   `yaml:"app_password_ref"`
	EndpointPath          string   `yaml:"endpoint_path"`
	TokenURL              string   `yaml:"token_url"`
	ConnectorScope        string   `yaml:"connector_scope"`
	AllowedServiceURLs    []string `yaml:"allowed_service_urls"`
	DisableAuthValidation bool     `yaml:"disable_auth_validation"`
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
	cfg, err := LoadPartial(path)
	if err != nil {
		return nil, err
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return cfg, nil
}

func LoadPartial(path string) (*Config, error) {
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
	if c.Bot.EndpointPath == "" {
		c.Bot.EndpointPath = "/api/messages"
	}
	if c.Bot.TokenURL == "" {
		c.Bot.TokenURL = "https://login.microsoftonline.com/botframework.com/oauth2/v2.0/token"
	}
	if c.Bot.ConnectorScope == "" {
		c.Bot.ConnectorScope = "https://api.botframework.com/.default"
	}
}

func (c *Config) Validate() error {
	if len(c.Accounts) == 0 {
		return errors.New("at least one account is required")
	}
	if len(c.TeamsChannels) == 0 {
		return errors.New("at least one teams channel is required")
	}
	if c.Bot.AppID == "" {
		return errors.New("bot.app_id is required")
	}
	if c.Bot.AppPassword == "" && c.Bot.AppPasswordEnv == "" && c.Bot.AppPasswordRef == "" {
		return errors.New("bot.app_password, bot.app_password_env, or bot.app_password_ref is required")
	}
	for name := range c.TeamsChannels {
		if strings.TrimSpace(name) == "" {
			return errors.New("teams channel alias cannot be empty")
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
