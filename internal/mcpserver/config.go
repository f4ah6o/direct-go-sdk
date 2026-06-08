package mcpserver

import (
	"errors"
	"fmt"
	"os"
	"strings"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"gopkg.in/yaml.v3"
)

const (
	DefaultListenAddr   = ":8090"
	DefaultEndpointPath = "/mcp"
	DefaultReadScope    = "direct:read"
	DefaultWriteScope   = "direct:write"
)

type Config struct {
	OP       OPConfig        `yaml:"op"`
	MCP      MCPConfig       `yaml:"mcp"`
	Accounts []AccountConfig `yaml:"accounts"`
}

type OPConfig struct {
	Binary string `yaml:"binary"`
}

type MCPConfig struct {
	ListenAddr           string   `yaml:"listen_addr"`
	EndpointPath         string   `yaml:"endpoint_path"`
	PublicBaseURL        string   `yaml:"public_base_url"`
	AuthorizationServers []string `yaml:"authorization_servers"`
	JWTIssuer            string   `yaml:"jwt_issuer"`
	JWTAudience          string   `yaml:"jwt_audience"`
	JWKSURL              string   `yaml:"jwks_url"`
	ReadScope            string   `yaml:"read_scope"`
	WriteScope           string   `yaml:"write_scope"`
}

type AccountConfig struct {
	ID       string `yaml:"id"`
	TokenEnv string `yaml:"token_env"`
	TokenRef string `yaml:"token_ref"`
	Endpoint string `yaml:"endpoint"`
	ProxyURL string `yaml:"proxy_url"`
}

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
	if c.MCP.ListenAddr == "" {
		c.MCP.ListenAddr = DefaultListenAddr
	}
	if c.MCP.EndpointPath == "" {
		c.MCP.EndpointPath = DefaultEndpointPath
	}
	if c.MCP.ReadScope == "" {
		c.MCP.ReadScope = DefaultReadScope
	}
	if c.MCP.WriteScope == "" {
		c.MCP.WriteScope = DefaultWriteScope
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
	if strings.TrimSpace(c.MCP.PublicBaseURL) == "" {
		return errors.New("mcp.public_base_url is required")
	}
	if strings.TrimSpace(c.MCP.JWTIssuer) == "" {
		return errors.New("mcp.jwt_issuer is required")
	}
	if strings.TrimSpace(c.MCP.JWTAudience) == "" {
		return errors.New("mcp.jwt_audience is required")
	}
	if strings.TrimSpace(c.MCP.JWKSURL) == "" {
		return errors.New("mcp.jwks_url is required")
	}
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
		if account.TokenEnv == "" && account.TokenRef == "" {
			return fmt.Errorf("account %q requires token_env or token_ref", account.ID)
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

func (c *Config) ResourceURL() string {
	base := strings.TrimRight(c.MCP.PublicBaseURL, "/")
	path := c.MCP.EndpointPath
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return base + path
}
