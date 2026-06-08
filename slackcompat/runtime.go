package slackcompat

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	opsecret "github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/secrets/op"
)

func ResolveTokens(ctx context.Context, cfg *Config) (map[string]string, error) {
	runner := opsecret.Runner{Binary: cfg.OP.Binary}
	tokens := map[string]string{}
	for _, account := range cfg.Accounts {
		if token := strings.TrimSpace(os.Getenv(account.TokenEnv)); token != "" {
			tokens[account.ID] = token
			continue
		}
		if account.TokenRef == "" {
			return nil, errors.New("account " + account.ID + " token env is empty and token_ref is not set")
		}
		token, err := runner.Read(ctx, account.TokenRef)
		if err != nil {
			return nil, err
		}
		if strings.TrimSpace(token) == "" {
			return nil, errors.New("account " + account.ID + " token ref is empty")
		}
		tokens[account.ID] = token
	}
	return tokens, nil
}

func ResolveServerBearerToken(ctx context.Context, cfg *Config) (string, error) {
	if cfg.Server.BearerTokenEnv != "" {
		if token := strings.TrimSpace(os.Getenv(cfg.Server.BearerTokenEnv)); token != "" {
			return token, nil
		}
	}
	if cfg.Server.BearerTokenRef == "" {
		return "", nil
	}
	token, err := opsecret.Runner{Binary: cfg.OP.Binary}.Read(ctx, cfg.Server.BearerTokenRef)
	if err != nil {
		return "", err
	}
	if strings.TrimSpace(token) == "" {
		return "", errors.New("server bearer token ref is empty")
	}
	return strings.TrimSpace(token), nil
}

type HTTPEventSink struct {
	URL    string
	Client *http.Client
}

func (s HTTPEventSink) Publish(ctx context.Context, event EventsEnvelope) error {
	if strings.TrimSpace(s.URL) == "" {
		return nil
	}
	client := s.Client
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(event); err != nil {
		return err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, s.URL, &body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("event callback returned status %d", resp.StatusCode)
	}
	return nil
}
