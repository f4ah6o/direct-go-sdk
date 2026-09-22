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
	// SubjectClaim names the JWT claim used to identify the caller for account
	// authorization ("sub" by default). A string claim produces one principal;
	// a list claim (e.g. groups) produces one principal per element.
	SubjectClaim string `yaml:"subject_claim"`
	// AccountSubjects maps a principal (subject_claim value) to the account IDs
	// it may use. The "*" key grants listed accounts to any authenticated
	// principal. When empty and SingleTenant is false, every account call is
	// denied (fail-closed).
	AccountSubjects map[string][]string `yaml:"account_subjects"`
	// SingleTenant is an explicit compatibility mode for single-tenant
	// deployments: every authenticated principal may use all accounts.
	SingleTenant bool `yaml:"single_tenant"`
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
	if c.MCP.SubjectClaim == "" {
		c.MCP.SubjectClaim = "sub"
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
	for principal, ids := range c.MCP.AccountSubjects {
		if strings.TrimSpace(principal) == "" {
			return errors.New("account_subjects contains an empty principal")
		}
		if len(ids) == 0 {
			return fmt.Errorf("account_subjects[%q] lists no accounts", principal)
		}
		for _, id := range ids {
			if !seen[id] {
				return fmt.Errorf("account_subjects[%q] references unknown account %q", principal, id)
			}
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

// AllowedAccountIDs returns the set of configured account IDs the given
// principals may access. In single-tenant mode all accounts are allowed.
// Otherwise the "*" entry grants its accounts to any authenticated principal
// and each principal grants its mapped accounts. An empty mapping denies all.
func (c *Config) AllowedAccountIDs(principals []string) map[string]bool {
	allowed := map[string]bool{}
	if c.MCP.SingleTenant {
		for _, account := range c.Accounts {
			allowed[account.ID] = true
		}
		return allowed
	}
	for _, id := range c.MCP.AccountSubjects["*"] {
		allowed[id] = true
	}
	for _, principal := range principals {
		for _, id := range c.MCP.AccountSubjects[principal] {
			allowed[id] = true
		}
	}
	return allowed
}

func (c *Config) ResourceURL() string {
	base := strings.TrimRight(c.MCP.PublicBaseURL, "/")
	path := c.MCP.EndpointPath
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}
	return base + path
}
