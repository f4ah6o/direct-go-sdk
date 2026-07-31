package slackcompat

import (
	"context"
	"fmt"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

type ProductionDirectClient struct {
	accountID string
	client    *direct.Client
}

func NewProductionDirectClient(account AccountConfig, token string) *ProductionDirectClient {
	return &ProductionDirectClient{
		accountID: account.ID,
		client: direct.NewClient(direct.Options{
			Endpoint:    account.Endpoint,
			AccessToken: token,
			ProxyURL:    account.ProxyURL,
			Name:        "direct-slack-compat",
		}),
	}
}

func (c *ProductionDirectClient) Connect() error {
	return c.client.Connect()
}

func (c *ProductionDirectClient) ConnectWithContext(ctx context.Context) error {
	return c.client.ConnectWithContext(ctx)
}

func (c *ProductionDirectClient) Close() error {
	c.client.Close()
	return nil
}

func (c *ProductionDirectClient) AccountID() string {
	return c.accountID
}

func (c *ProductionDirectClient) GetMe(ctx context.Context) (*direct.UserInfo, error) {
	return c.client.GetMeWithContext(ctx)
}

func (c *ProductionDirectClient) GetTalks(ctx context.Context) ([]direct.Talk, error) {
	return c.client.GetTalksWithContext(ctx)
}

func (c *ProductionDirectClient) GetUsers(ctx context.Context) ([]direct.UserInfo, error) {
	// Direct does not expose a Slack-style workspace-wide user directory here.
	// Friends are the smallest stable offline-compatible users.list source.
	return c.client.GetFriends(ctx)
}

func (c *ProductionDirectClient) GetMessages(ctx context.Context, domainID, talkID interface{}, opts *direct.GetMessagesOptions) ([]direct.ReceivedMessage, error) {
	return c.client.GetMessages(ctx, domainID, talkID, opts)
}

func (c *ProductionDirectClient) SendText(ctx context.Context, talkID, text string) (string, error) {
	return c.client.CreateTextMessageWithContext(ctx, talkID, text)
}

func (c *ProductionDirectClient) OnMessage(handler func(direct.ReceivedMessage)) {
	c.client.OnMessage(handler)
}

func NewConnectedProductionClients(ctx context.Context, cfg *Config, tokens map[string]string) ([]*ProductionDirectClient, error) {
	clients := make([]*ProductionDirectClient, 0, len(cfg.Accounts))
	for _, account := range cfg.Accounts {
		token := tokens[account.ID]
		if token == "" {
			return nil, fmt.Errorf("missing resolved token for account %q", account.ID)
		}
		client := NewProductionDirectClient(account, token)
		if err := client.ConnectWithContext(ctx); err != nil {
			for _, c := range clients {
				_ = c.Close()
			}
			return nil, fmt.Errorf("connect account %q: %w", account.ID, err)
		}
		clients = append(clients, client)
	}
	return clients, nil
}
