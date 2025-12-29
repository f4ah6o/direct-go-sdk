package direct

import (
	"os"
	"testing"
	"time"
)

func TestConfigValidate(t *testing.T) {
	tests := []struct {
		name    string
		config  Config
		wantErr bool
	}{
		{
			name: "valid config",
			config: Config{
				Token: "test-token",
			},
			wantErr: false,
		},
		{
			name:    "missing token",
			config:  Config{},
			wantErr: true,
		},
		{
			name: "invalid endpoint URL",
			config: Config{
				Token:    "test-token",
				Endpoint: ":invalid-url",
			},
			wantErr: true,
		},
		{
			name: "invalid proxy URL",
			config: Config{
				Token:    "test-token",
				ProxyURL: ":invalid-proxy",
			},
			wantErr: true,
		},
		{
			name: "timeout too short",
			config: Config{
				Token:   "test-token",
				Timeout: 50 * time.Millisecond,
			},
			wantErr: true,
		},
		{
			name: "timeout too long",
			config: Config{
				Token:   "test-token",
				Timeout: 10 * time.Minute,
			},
			wantErr: true,
		},
		{
			name: "valid timeout",
			config: Config{
				Token:   "test-token",
				Timeout: 30 * time.Second,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.Validate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConfigDefaults(t *testing.T) {
	cfg := &Config{}
	cfg.Defaults()

	if cfg.Endpoint != DefaultEndpoint {
		t.Errorf("Defaults() Endpoint = %s, want %s", cfg.Endpoint, DefaultEndpoint)
	}
	if cfg.Timeout != DefaultRequestTimeout {
		t.Errorf("Defaults() Timeout = %v, want %v", cfg.Timeout, DefaultRequestTimeout)
	}
	if cfg.Name != "direct-go" {
		t.Errorf("Defaults() Name = %s, want direct-go", cfg.Name)
	}
}

func TestConfigToOptions(t *testing.T) {
	cfg := &Config{
		Token:    "test-token",
		Endpoint: "wss://example.com",
		ProxyURL: "http://proxy:8080",
		Name:     "testbot",
	}

	opts := cfg.ToOptions()

	if opts.AccessToken != cfg.Token {
		t.Errorf("ToOptions() AccessToken = %s, want %s", opts.AccessToken, cfg.Token)
	}
	if opts.Endpoint != cfg.Endpoint {
		t.Errorf("ToOptions() Endpoint = %s, want %s", opts.Endpoint, cfg.Endpoint)
	}
	if opts.ProxyURL != cfg.ProxyURL {
		t.Errorf("ToOptions() ProxyURL = %s, want %s", opts.ProxyURL, cfg.ProxyURL)
	}
	if opts.Name != cfg.Name {
		t.Errorf("ToOptions() Name = %s, want %s", opts.Name, cfg.Name)
	}
}

func TestLoadConfigFromEnv(t *testing.T) {
	// Save original env values
	envToken := envOrDefault("HUBOT_DIRECT_TOKEN", "")
	envEndpoint := envOrDefault("HUBOT_DIRECT_ENDPOINT", "")
	envProxy := envOrDefault("DIRECT_PROXY_URL", "")
	envTimeout := envOrDefault("DIRECT_TIMEOUT", "")

	// Clean up after test
	defer func() {
		if envToken == "" {
			os.Unsetenv("HUBOT_DIRECT_TOKEN")
		}
		if envEndpoint == "" {
			os.Unsetenv("HUBOT_DIRECT_ENDPOINT")
		}
		if envProxy == "" {
			os.Unsetenv("DIRECT_PROXY_URL")
		}
		if envTimeout == "" {
			os.Unsetenv("DIRECT_TIMEOUT")
		}
	}()

	// Test loading config from environment
	os.Setenv("HUBOT_DIRECT_TOKEN", "env-token-123")
	os.Setenv("DIRECT_NAME", "env-bot")
	os.Setenv("DIRECT_DEBUG", "true")

	auth := NewAuth()
	cfg, err := LoadConfigFromEnv(auth)
	if err != nil {
		t.Fatalf("LoadConfigFromEnv() error = %v", err)
	}

	if cfg.Token != "env-token-123" {
		t.Errorf("Token = %s, want env-token-123", cfg.Token)
	}
	if cfg.Name != "env-bot" {
		t.Errorf("Name = %s, want env-bot", cfg.Name)
	}
	if !cfg.Debug {
		t.Error("Debug = false, want true")
	}
}

func TestLoadConfigFromEnvMissingToken(t *testing.T) {
	auth := NewAuthWithFile("/nonexistent/.env")
	_, err := LoadConfigFromEnv(auth)
	if err == nil {
		t.Error("LoadConfigFromEnv() should return error when token is missing")
	}
}
