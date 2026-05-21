package config

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	OP            OPConfig         `yaml:"op"`
	TeamsChannels TeamsChannels    `yaml:"teams_channels"`
	Accounts      []AccountConfig  `yaml:"accounts"`
	Bot           BotConfig        `yaml:"bot"`
	State         StateConfig      `yaml:"state"`
	Server        ServerConfig     `yaml:"server"`
	Queues        QueueConfig      `yaml:"queues"`
	Retry         RetryConfig      `yaml:"retry"`
	Attachments   AttachmentConfig `yaml:"attachments"`
}

type TeamsChannels map[string]TeamsChannelConfig

type TeamsChannelConfig struct {
	// Alias-only config. The actual Teams conversation is bound at runtime by
	// sending "@bot bind <alias>" in the target channel.
}

func (tc *TeamsChannels) UnmarshalYAML(value *yaml.Node) error {
	out := TeamsChannels{}
	switch value.Kind {
	case yaml.MappingNode:
		var m map[string]TeamsChannelConfig
		if err := value.Decode(&m); err != nil {
			return err
		}
		for k, v := range m {
			out[k] = v
		}
	case yaml.SequenceNode:
		for _, item := range value.Content {
			var m map[string]TeamsChannelConfig
			if err := item.Decode(&m); err != nil {
				return err
			}
			for k, v := range m {
				out[k] = v
			}
		}
	default:
		return fmt.Errorf("teams_channels must be a map or list of maps")
	}
	*tc = out
	return nil
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
	AppID                     string   `yaml:"app_id"`
	AppPassword               string   `yaml:"app_password"`
	AppPasswordEnv            string   `yaml:"app_password_env"`
	AppPasswordRef            string   `yaml:"app_password_ref"`
	TenantID                  string   `yaml:"tenant_id"`
	EndpointPath              string   `yaml:"endpoint_path"`
	TokenURL                  string   `yaml:"token_url"`
	ConnectorScope            string   `yaml:"connector_scope"`
	OpenIDMetadataURL         string   `yaml:"openid_metadata_url"`
	EmulatorOpenIDMetadataURL string   `yaml:"emulator_openid_metadata_url"`
	AllowedServiceURLs        []string `yaml:"allowed_service_urls"`
	AllowEmulator             bool     `yaml:"allow_emulator"`
	DisableAuthValidation     bool     `yaml:"disable_auth_validation"`
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
	MaxBytes     int64  `yaml:"max_bytes"`
	TempDir      string `yaml:"temp_dir"`
	FileProxyTTL string `yaml:"file_proxy_ttl"`
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
	if c.Attachments.FileProxyTTL == "" {
		c.Attachments.FileProxyTTL = "24h"
	}
	if c.Bot.EndpointPath == "" {
		c.Bot.EndpointPath = "/api/messages"
	}
	if c.Bot.TokenURL == "" {
		tenant := strings.TrimSpace(c.Bot.TenantID)
		if tenant == "" {
			tenant = "botframework.com"
		}
		c.Bot.TokenURL = "https://login.microsoftonline.com/" + tenant + "/oauth2/v2.0/token"
	}
	if c.Bot.ConnectorScope == "" {
		c.Bot.ConnectorScope = "https://api.botframework.com/.default"
	}
	if c.Bot.OpenIDMetadataURL == "" {
		c.Bot.OpenIDMetadataURL = "https://login.botframework.com/v1/.well-known/openidconfiguration"
	}
	if c.Bot.EmulatorOpenIDMetadataURL == "" {
		c.Bot.EmulatorOpenIDMetadataURL = "https://login.microsoftonline.com/common/v2.0/.well-known/openid-configuration"
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
	if c.Bot.DisableAuthValidation && !c.allowsLocalAuthBypass() {
		return errors.New("bot.disable_auth_validation is only allowed with local listen/public URLs")
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

func (c *Config) allowsLocalAuthBypass() bool {
	host := listenHost(c.Server.ListenAddr)
	if host != "" && !isLocalHost(host) {
		return false
	}
	if strings.TrimSpace(c.Server.PublicBaseURL) == "" {
		return true
	}
	return isLocalURL(c.Server.PublicBaseURL)
}

func listenHost(addr string) string {
	addr = strings.TrimSpace(addr)
	if addr == "" || strings.HasPrefix(addr, ":") {
		return ""
	}
	if strings.HasPrefix(addr, "[") {
		end := strings.Index(addr, "]")
		if end >= 0 {
			return strings.Trim(addr[1:end], "[]")
		}
	}
	host, _, ok := strings.Cut(addr, ":")
	if ok {
		return host
	}
	return addr
}

func isLocalURL(raw string) bool {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return true
	}
	withoutScheme := strings.TrimPrefix(strings.TrimPrefix(raw, "http://"), "https://")
	host, _, _ := strings.Cut(withoutScheme, "/")
	host, _, _ = strings.Cut(host, ":")
	return isLocalHost(host)
}

func isLocalHost(host string) bool {
	host = strings.Trim(strings.ToLower(strings.TrimSpace(host)), "[]")
	return host == "" || host == "localhost" || host == "127.0.0.1" || host == "::1"
}

func (c *Config) Account(id string) (AccountConfig, bool) {
	for _, account := range c.Accounts {
		if account.ID == id {
			return account, true
		}
	}
	return AccountConfig{}, false
}
