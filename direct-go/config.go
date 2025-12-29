package direct

import (
	"errors"
	"net/url"
	"os"
	"time"
)

// Config holds the configuration for the direct client.
// It provides a centralized way to manage client configuration with validation.
type Config struct {
	// Token is the authentication token for the API.
	// Required.
	Token string

	// Endpoint is the WebSocket API endpoint.
	// If empty, DefaultEndpoint is used.
	Endpoint string

	// ProxyURL is an optional HTTP proxy URL for WebSocket connections.
	ProxyURL string

	// Timeout is the timeout for RPC requests.
	// If zero, DefaultRequestTimeout is used.
	Timeout time.Duration

	// Name is the bot name, used in log messages.
	Name string

	// Debug enables debug logging when true.
	Debug bool
}

// Validate checks that the configuration is valid.
// Returns an error if any required fields are missing or invalid.
func (c *Config) Validate() error {
	if c.Token == "" {
		return errors.New("config: Token is required")
	}

	// Validate endpoint format if provided
	if c.Endpoint != "" {
		if _, err := url.Parse(c.Endpoint); err != nil {
			return errors.New("config: invalid Endpoint URL")
		}
	}

	// Validate proxy URL format if provided
	if c.ProxyURL != "" {
		if _, err := url.Parse(c.ProxyURL); err != nil {
			return errors.New("config: invalid ProxyURL")
		}
	}

	// Validate timeout range if provided
	if c.Timeout != 0 {
		if c.Timeout < 100*time.Millisecond {
			return errors.New("config: Timeout must be at least 100ms")
		}
		if c.Timeout > 5*time.Minute {
			return errors.New("config: Timeout must be at most 5 minutes")
		}
	}

	return nil
}

// Defaults fills in default values for empty fields.
func (c *Config) Defaults() *Config {
	if c.Endpoint == "" {
		c.Endpoint = DefaultEndpoint
	}
	if c.Timeout == 0 {
		c.Timeout = DefaultRequestTimeout
	}
	if c.Name == "" {
		c.Name = "direct-go"
	}
	return c
}

// ToOptions converts the Config to an Options struct for use with NewClient.
func (c *Config) ToOptions() Options {
	return Options{
		Endpoint:    c.Endpoint,
		AccessToken: c.Token,
		ProxyURL:    c.ProxyURL,
		Name:        c.Name,
	}
}

// LoadConfig creates a Config from environment variables.
// It reads the following environment variables:
//   - HUBOT_DIRECT_TOKEN: authentication token (required)
//   - HUBOT_DIRECT_ENDPOINT: API endpoint (optional)
//   - DIRECT_PROXY_URL: HTTP proxy URL (optional)
//   - DIRECT_TIMEOUT: request timeout in duration format (optional)
//   - DIRECT_DEBUG: set to "true" to enable debug logging (optional)
func LoadConfig() (*Config, error) {
	return LoadConfigFromEnv(NewAuth())
}

// LoadConfigFromEnv creates a Config using the given Auth instance.
func LoadConfigFromEnv(auth *Auth) (*Config, error) {
	cfg := &Config{
		Token: auth.GetToken(),
	}

	// Load optional settings from environment
	cfg.Endpoint = envOrDefault("HUBOT_DIRECT_ENDPOINT", "")
	cfg.ProxyURL = envOrDefault("DIRECT_PROXY_URL", "")
	cfg.Name = envOrDefault("DIRECT_NAME", "")
	cfg.Debug = envOrDefault("DIRECT_DEBUG", "") == "true"

	// Parse timeout if provided
	if timeoutStr := envOrDefault("DIRECT_TIMEOUT", ""); timeoutStr != "" {
		if d, err := time.ParseDuration(timeoutStr); err == nil {
			cfg.Timeout = d
		}
		// If parsing fails, just use the default
	}

	cfg.Defaults()
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

// envOrDefault returns the environment variable value or the default if not set.
func envOrDefault(key, defaultValue string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultValue
}
