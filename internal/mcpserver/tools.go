package mcpserver

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type directClient interface {
	Connect() error
	Close() error
	GetMeWithContext(context.Context) (*direct.UserInfo, error)
	GetDomainsWithContext(context.Context) ([]direct.DomainInfo, error)
	GetTalksWithContext(context.Context) ([]direct.Talk, error)
	GetMessages(context.Context, interface{}, interface{}, *direct.GetMessagesOptions) ([]direct.ReceivedMessage, error)
	SearchMessages(context.Context, interface{}, interface{}, string, interface{}, int) (*direct.SearchMessagesResult, error)
	CreateTextMessageWithContext(context.Context, string, string) (string, error)
}

type directClientFactory func(direct.Options) directClient

type productionDirectClient struct {
	*direct.Client
}

func newProductionDirectClient(opts direct.Options) directClient {
	return productionDirectClient{Client: direct.NewClient(opts)}
}

type accountArgs struct {
	AccountID string `json:"account_id" jsonschema:"configured direct account id"`
}

type getMessagesArgs struct {
	AccountID string      `json:"account_id" jsonschema:"configured direct account id"`
	DomainID  interface{} `json:"domain_id" jsonschema:"direct domain id"`
	TalkID    interface{} `json:"talk_id" jsonschema:"direct talk id"`
	SinceID   interface{} `json:"since_id,omitempty" jsonschema:"message id lower bound"`
	MaxID     interface{} `json:"max_id,omitempty" jsonschema:"message id upper bound"`
	Order     string      `json:"order,omitempty" jsonschema:"asc or desc"`
}

type searchMessagesArgs struct {
	AccountID string      `json:"account_id" jsonschema:"configured direct account id"`
	DomainID  interface{} `json:"domain_id" jsonschema:"direct domain id"`
	TalkID    interface{} `json:"talk_id,omitempty" jsonschema:"direct talk id; omit to search the domain"`
	Keyword   string      `json:"keyword" jsonschema:"search keyword"`
	Marker    interface{} `json:"marker,omitempty" jsonschema:"pagination marker"`
	Limit     int         `json:"limit,omitempty" jsonschema:"maximum number of results"`
}

type sendTextArgs struct {
	AccountID string `json:"account_id" jsonschema:"configured direct account id"`
	TalkID    string `json:"talk_id" jsonschema:"direct talk id"`
	Text      string `json:"text" jsonschema:"message text"`
}

type accountOutput struct {
	ID        string `json:"id"`
	Endpoint  string `json:"endpoint"`
	Available bool   `json:"available"`
}

type listAccountsOutput struct {
	Accounts []accountOutput `json:"accounts"`
}

type userOutput struct {
	User *direct.UserInfo `json:"user,omitempty"`
}

type listDomainsOutput struct {
	Domains []direct.DomainInfo `json:"domains"`
}

type listTalksOutput struct {
	Talks []direct.Talk `json:"talks"`
}

type getMessagesOutput struct {
	Messages []direct.ReceivedMessage `json:"messages"`
}

type searchMessagesOutput struct {
	Result *direct.SearchMessagesResult `json:"result,omitempty"`
}

type sendTextOutput struct {
	MessageID string `json:"message_id"`
}

func (s *Server) registerTools() {
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_list_accounts",
		Description: "List configured Direct4B accounts without exposing access tokens.",
	}, s.listAccounts)
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_get_me",
		Description: "Get the authenticated Direct4B user profile for an account.",
	}, s.getMe)
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_list_domains",
		Description: "List Direct4B domains visible to an account.",
	}, s.listDomains)
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_list_talks",
		Description: "List Direct4B talks visible to an account.",
	}, s.listTalks)
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_get_messages",
		Description: "Get messages from a Direct4B talk.",
	}, s.getMessages)
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_search_messages",
		Description: "Search Direct4B messages in a domain or talk.",
	}, s.searchMessages)
	mcp.AddTool(s.mcpServer, &mcp.Tool{
		Name:        "direct_send_text",
		Description: "Send a text message to a Direct4B talk.",
	}, s.sendText)
}

func (s *Server) listAccounts(ctx context.Context, _ *mcp.CallToolRequest, _ struct{}) (*mcp.CallToolResult, listAccountsOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.ReadScope); err != nil {
		return nil, listAccountsOutput{}, err
	}
	out := make([]accountOutput, 0, len(s.cfg.Accounts))
	for _, account := range s.cfg.Accounts {
		_, available := s.tokens[account.ID]
		out = append(out, accountOutput{ID: account.ID, Endpoint: account.Endpoint, Available: available})
	}
	return nil, listAccountsOutput{Accounts: out}, nil
}

func (s *Server) getMe(ctx context.Context, _ *mcp.CallToolRequest, args accountArgs) (*mcp.CallToolResult, userOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.ReadScope); err != nil {
		return nil, userOutput{}, err
	}
	client, closeClient, err := s.openClient(args.AccountID)
	if err != nil {
		return nil, userOutput{}, err
	}
	defer closeClient()
	out, err := client.GetMeWithContext(ctx)
	return nil, userOutput{User: out}, err
}

func (s *Server) listDomains(ctx context.Context, _ *mcp.CallToolRequest, args accountArgs) (*mcp.CallToolResult, listDomainsOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.ReadScope); err != nil {
		return nil, listDomainsOutput{}, err
	}
	client, closeClient, err := s.openClient(args.AccountID)
	if err != nil {
		return nil, listDomainsOutput{}, err
	}
	defer closeClient()
	out, err := client.GetDomainsWithContext(ctx)
	return nil, listDomainsOutput{Domains: out}, err
}

func (s *Server) listTalks(ctx context.Context, _ *mcp.CallToolRequest, args accountArgs) (*mcp.CallToolResult, listTalksOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.ReadScope); err != nil {
		return nil, listTalksOutput{}, err
	}
	client, closeClient, err := s.openClient(args.AccountID)
	if err != nil {
		return nil, listTalksOutput{}, err
	}
	defer closeClient()
	out, err := client.GetTalksWithContext(ctx)
	return nil, listTalksOutput{Talks: out}, err
}

func (s *Server) getMessages(ctx context.Context, _ *mcp.CallToolRequest, args getMessagesArgs) (*mcp.CallToolResult, getMessagesOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.ReadScope); err != nil {
		return nil, getMessagesOutput{}, err
	}
	client, closeClient, err := s.openClient(args.AccountID)
	if err != nil {
		return nil, getMessagesOutput{}, err
	}
	defer closeClient()
	opts := &direct.GetMessagesOptions{
		SinceID: normalizeID(args.SinceID),
		MaxID:   normalizeID(args.MaxID),
		Order:   parseOrder(args.Order),
	}
	out, err := client.GetMessages(ctx, normalizeID(args.DomainID), normalizeID(args.TalkID), opts)
	return nil, getMessagesOutput{Messages: out}, err
}

func (s *Server) searchMessages(ctx context.Context, _ *mcp.CallToolRequest, args searchMessagesArgs) (*mcp.CallToolResult, searchMessagesOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.ReadScope); err != nil {
		return nil, searchMessagesOutput{}, err
	}
	if strings.TrimSpace(args.Keyword) == "" {
		return nil, searchMessagesOutput{}, fmt.Errorf("keyword is required")
	}
	client, closeClient, err := s.openClient(args.AccountID)
	if err != nil {
		return nil, searchMessagesOutput{}, err
	}
	defer closeClient()
	limit := args.Limit
	if limit <= 0 {
		limit = 20
	}
	out, err := client.SearchMessages(ctx, normalizeID(args.DomainID), normalizeID(args.TalkID), args.Keyword, normalizeID(args.Marker), limit)
	return nil, searchMessagesOutput{Result: out}, err
}

func (s *Server) sendText(ctx context.Context, _ *mcp.CallToolRequest, args sendTextArgs) (*mcp.CallToolResult, sendTextOutput, error) {
	if err := RequireScope(ctx, s.cfg.MCP.WriteScope); err != nil {
		return nil, sendTextOutput{}, err
	}
	if strings.TrimSpace(args.TalkID) == "" {
		return nil, sendTextOutput{}, fmt.Errorf("talk_id is required")
	}
	if strings.TrimSpace(args.Text) == "" {
		return nil, sendTextOutput{}, fmt.Errorf("text is required")
	}
	client, closeClient, err := s.openClient(args.AccountID)
	if err != nil {
		return nil, sendTextOutput{}, err
	}
	defer closeClient()
	messageID, err := client.CreateTextMessageWithContext(ctx, args.TalkID, args.Text)
	if err != nil {
		return nil, sendTextOutput{}, err
	}
	return nil, sendTextOutput{MessageID: messageID}, nil
}

func (s *Server) openClient(accountID string) (directClient, func(), error) {
	account, ok := s.cfg.Account(accountID)
	if !ok {
		return nil, nil, fmt.Errorf("unknown account %q", accountID)
	}
	token, ok := s.tokens[account.ID]
	if !ok || token == "" {
		return nil, nil, fmt.Errorf("account %q token is not available", account.ID)
	}
	// Keep v0.1 stateless and simple. If MCP traffic becomes high-frequency,
	// replace this with a per-account connection pool.
	client := s.clientFactory(direct.Options{
		Endpoint:    account.Endpoint,
		AccessToken: token,
		ProxyURL:    account.ProxyURL,
		Name:        "direct-mcp-server-" + account.ID,
	})
	if err := client.Connect(); err != nil {
		return nil, nil, err
	}
	return client, func() { client.Close() }, nil
}

func parseOrder(order string) direct.MessageOrder {
	switch strings.ToLower(strings.TrimSpace(order)) {
	case "asc", "ascending":
		return direct.MessageOrderAsc
	default:
		return direct.MessageOrderDesc
	}
}

func normalizeID(v interface{}) interface{} {
	const maxSafeJSONInteger = 1<<53 - 1
	switch x := v.(type) {
	case nil:
		return nil
	case string:
		if x == "" {
			return nil
		}
		if id, err := strconv.ParseUint(x, 10, 64); err == nil {
			return id
		}
		return x
	case float64:
		if x == 0 {
			return nil
		}
		// JSON numbers above 2^53 may already be rounded by decoding; callers
		// should pass large Direct IDs as strings.
		if x > 0 && x <= maxSafeJSONInteger && x == float64(uint64(x)) {
			return uint64(x)
		}
		return x
	default:
		return x
	}
}
