package slackcompat

import (
	"errors"
	"fmt"
	"os"
	"strings"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"gopkg.in/yaml.v3"
)

const DefaultListenAddr = "127.0.0.1:8091"

func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := yaml.Unmarshal([]byte(os.ExpandEnv(string(data))), &cfg); err != nil {
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
		c.Server.ListenAddr = DefaultListenAddr
	}
	if c.Slack.TeamID == "" {
		c.Slack.TeamID = "Tdirect"
	}
	if c.Slack.TeamName == "" {
		c.Slack.TeamName = "Direct4B"
	}
	for i := range c.Accounts {
		if c.Accounts[i].Endpoint == "" {
			c.Accounts[i].Endpoint = direct.DefaultEndpoint
		}
		if c.Accounts[i].TokenEnv == "" && c.Accounts[i].ID != "" {
			c.Accounts[i].TokenEnv = "DIRECT_TOKEN_" + strings.ToUpper(strings.NewReplacer("-", "_", ".", "_").Replace(c.Accounts[i].ID))
		}
	}
}

func (c *Config) Validate() error {
	if len(c.Accounts) == 0 {
		return errors.New("at least one account is required")
	}
	seen := map[string]bool{}
	for _, account := range c.Accounts {
		if strings.TrimSpace(account.ID) == "" {
			return errors.New("account id is required")
		}
		if seen[account.ID] {
			return fmt.Errorf("duplicate account id %q", account.ID)
		}
		seen[account.ID] = true
		if strings.TrimSpace(account.TokenEnv) == "" && strings.TrimSpace(account.TokenRef) == "" {
			return fmt.Errorf("account %q requires token_env or token_ref", account.ID)
		}
	}
	return nil
}
