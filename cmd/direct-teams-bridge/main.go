package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	appbridge "github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/bridge"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/directworker"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	opsecret "github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/secrets/op"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/teams"
	"golang.org/x/term"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	if err := run(os.Args, logger); err != nil {
		logger.Fatal(err)
	}
}

func run(args []string, logger *log.Logger) error {
	if len(args) < 2 {
		return usage()
	}
	switch args[1] {
	case "run":
		return runBridge(args[2:], logger)
	case "login-direct":
		return loginDirect(args[2:], logger)
	case "mappings":
		return mappings(args[2:])
	case "channels":
		return channels(args[2:])
	default:
		return usage()
	}
}

func runBridge(args []string, logger *log.Logger) error {
	fs := flag.NewFlagSet("run", flag.ExitOnError)
	configPath := fs.String("config", "config.yaml", "config file")
	if err := fs.Parse(args); err != nil {
		return err
	}
	cfg, err := config.Load(*configPath)
	if err != nil {
		return err
	}
	st, err := store.Open(cfg.State.Path)
	if err != nil {
		return err
	}
	runtimeState, err := newRuntimeState(context.Background(), cfg)
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	directToTeams := make(chan model.DirectMessage, cfg.Queues.DirectToTeams)
	teamsToDirect := make(chan model.DirectOutbound, cfg.Queues.TeamsToDirect)
	directSent := make(chan model.DirectSent, cfg.Queues.TeamsToDirect)
	teamsClient := teams.NewClient(cfg.Bot, cfg.Server.PublicBaseURL, cfg.Attachments.FileProxyTTL)
	directManager := directworker.NewManager(directToTeams, directSent, logger)
	service := appbridge.NewService(runtimeState.Account, st, teamsClient, directManager, directToTeams, teamsToDirect, directSent, logger)
	server := teams.NewServer(cfg, teamsClient, st, teamsToDirect, logger, teams.WithRuntimeLookups(runtimeState.HasChannel, runtimeState.Account, runtimeState.Token))

	directManager.Apply(ctx, runtimeState.DirectAccounts())
	go watchConfig(ctx, *configPath, runtimeState, directManager, logger)
	service.Run(ctx)
	return server.Run(ctx)
}

func loginDirect(args []string, logger *log.Logger) error {
	fs := flag.NewFlagSet("login-direct", flag.ExitOnError)
	configPath := fs.String("config", "config.yaml", "config file")
	accountID := fs.String("account", "", "account id")
	resetDeviceID := fs.Bool("reset-device-id", false, "reset the stored direct device id before login")
	if err := fs.Parse(args); err != nil {
		return err
	}
	if *accountID == "" {
		return fmt.Errorf("--account is required")
	}
	cfg, err := config.LoadPartial(*configPath)
	if err != nil {
		return err
	}
	account, ok := cfg.Account(*accountID)
	if !ok {
		return fmt.Errorf("unknown account %q", *accountID)
	}
	if account.TokenRef == "" {
		return fmt.Errorf("account %q requires token_ref for login-direct", account.ID)
	}
	st, err := store.Open(cfg.State.Path)
	if err != nil {
		return err
	}
	deviceID, err := st.EnsureDirectDevice(account.ID)
	if err != nil {
		return err
	}
	if *resetDeviceID {
		deviceID, err = st.ResetDirectDevice(account.ID)
		if err != nil {
			return err
		}
	}
	email, password, err := promptCredentials()
	if err != nil {
		return err
	}
	endpoint := account.Endpoint
	if endpoint == "" {
		endpoint = direct.DefaultEndpoint
	}
	client := direct.NewClient(direct.Options{Endpoint: endpoint, ProxyURL: account.ProxyURL, Name: "direct-teams-bridge-login"})
	if err := client.Connect(); err != nil {
		return err
	}
	defer client.Close()
	token, err := client.CreateAccessToken(email, password, deviceID, direct.DefaultBotOS)
	if err != nil {
		return err
	}
	runner := opsecret.Runner{Binary: cfg.OP.Binary}
	if err := runner.Write(context.Background(), account.TokenRef, token); err != nil {
		return err
	}
	logger.Printf("[%s] direct token saved to 1Password field %s", account.ID, redactRef(account.TokenRef))
	return nil
}

type runtimeState struct {
	mu       sync.RWMutex
	cfg      *config.Config
	tokens   map[string]string
	staticID string
}

func newRuntimeState(ctx context.Context, cfg *config.Config) (*runtimeState, error) {
	tokens, err := resolveAccountTokens(ctx, cfg)
	if err != nil {
		return nil, err
	}
	if err := resolveBotSecret(ctx, cfg); err != nil {
		return nil, err
	}
	return &runtimeState{cfg: cfg, tokens: tokens, staticID: staticConfigID(cfg)}, nil
}

func (r *runtimeState) Apply(cfg *config.Config, tokens map[string]string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.cfg.Accounts = cfg.Accounts
	r.cfg.TeamsChannels = cfg.TeamsChannels
	r.tokens = tokens
}

func (r *runtimeState) Account(id string) (config.AccountConfig, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.cfg.Account(id)
}

func (r *runtimeState) HasChannel(alias string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	_, ok := r.cfg.TeamsChannels[alias]
	return ok
}

func (r *runtimeState) Token(accountID string) (string, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	token, ok := r.tokens[accountID]
	return token, ok
}

func (r *runtimeState) DirectAccounts() []directworker.RuntimeAccount {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]directworker.RuntimeAccount, 0, len(r.cfg.Accounts))
	for _, account := range r.cfg.Accounts {
		out = append(out, directworker.RuntimeAccount{Config: account, Token: r.tokens[account.ID]})
	}
	return out
}

func watchConfig(ctx context.Context, path string, runtimeState *runtimeState, directManager *directworker.Manager, logger *log.Logger) {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	var lastMod time.Time
	if info, err := os.Stat(path); err == nil {
		lastMod = info.ModTime()
	}
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
		}
		info, err := os.Stat(path)
		if err != nil {
			logger.Printf("[config] stat failed: %v", err)
			continue
		}
		if !info.ModTime().After(lastMod) {
			continue
		}
		lastMod = info.ModTime()
		cfg, err := config.Load(path)
		if err != nil {
			logger.Printf("[config] reload ignored: %v", err)
			continue
		}
		if id := staticConfigID(cfg); id != runtimeState.staticID {
			logger.Printf("[config] static settings changed; restart required for bot/server/queue/state changes")
		}
		tokens, err := resolveAccountTokens(ctx, cfg)
		if err != nil {
			logger.Printf("[config] reload ignored while resolving tokens: %v", err)
			continue
		}
		runtimeState.Apply(cfg, tokens)
		directManager.Apply(ctx, runtimeState.DirectAccounts())
		logger.Printf("[config] reloaded accounts=%d channels=%d", len(cfg.Accounts), len(cfg.TeamsChannels))
	}
}

func staticConfigID(cfg *config.Config) string {
	return strings.Join([]string{
		cfg.Bot.AppID,
		cfg.Bot.AppPassword,
		cfg.Bot.AppPasswordEnv,
		cfg.Bot.AppPasswordRef,
		cfg.Bot.TenantID,
		cfg.Bot.EndpointPath,
		cfg.Bot.TokenURL,
		cfg.Bot.ConnectorScope,
		cfg.Server.ListenAddr,
		cfg.Server.PublicBaseURL,
		cfg.State.Path,
		fmt.Sprint(cfg.Queues.DirectToTeams),
		fmt.Sprint(cfg.Queues.TeamsToDirect),
		cfg.Attachments.FileProxyTTL,
	}, "\x00")
}

func mappings(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("mappings requires list or forget")
	}
	switch args[0] {
	case "list":
		fs := flag.NewFlagSet("mappings list", flag.ExitOnError)
		configPath := fs.String("config", "config.yaml", "config file")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		cfg, err := config.LoadPartial(*configPath)
		if err != nil {
			return err
		}
		st, err := store.Open(cfg.State.Path)
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(st.ListMappings())
	case "forget":
		fs := flag.NewFlagSet("mappings forget", flag.ExitOnError)
		configPath := fs.String("config", "config.yaml", "config file")
		accountID := fs.String("account", "", "account id")
		talkID := fs.String("talk-id", "", "direct talk id")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		if *accountID == "" || *talkID == "" {
			return fmt.Errorf("--account and --talk-id are required")
		}
		cfg, err := config.LoadPartial(*configPath)
		if err != nil {
			return err
		}
		st, err := store.Open(cfg.State.Path)
		if err != nil {
			return err
		}
		return st.Forget(*accountID, *talkID)
	default:
		return fmt.Errorf("unknown mappings command %q", args[0])
	}
}

func channels(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("channels requires list or forget")
	}
	switch args[0] {
	case "list":
		fs := flag.NewFlagSet("channels list", flag.ExitOnError)
		configPath := fs.String("config", "config.yaml", "config file")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		cfg, err := config.LoadPartial(*configPath)
		if err != nil {
			return err
		}
		st, err := store.Open(cfg.State.Path)
		if err != nil {
			return err
		}
		return json.NewEncoder(os.Stdout).Encode(st.ListChannelBindings())
	case "forget":
		fs := flag.NewFlagSet("channels forget", flag.ExitOnError)
		configPath := fs.String("config", "config.yaml", "config file")
		alias := fs.String("alias", "", "teams channel alias")
		if err := fs.Parse(args[1:]); err != nil {
			return err
		}
		if *alias == "" {
			return fmt.Errorf("--alias is required")
		}
		cfg, err := config.LoadPartial(*configPath)
		if err != nil {
			return err
		}
		st, err := store.Open(cfg.State.Path)
		if err != nil {
			return err
		}
		return st.ForgetChannelBinding(*alias)
	default:
		return fmt.Errorf("unknown channels command %q", args[0])
	}
}

func resolveBotSecret(ctx context.Context, cfg *config.Config) error {
	runner := opsecret.Runner{Binary: cfg.OP.Binary}
	if cfg.Bot.AppPassword == "" && cfg.Bot.AppPasswordEnv != "" && os.Getenv(cfg.Bot.AppPasswordEnv) == "" && cfg.Bot.AppPasswordRef != "" {
		secret, err := runner.Read(ctx, cfg.Bot.AppPasswordRef)
		if err != nil {
			return err
		}
		if err := os.Setenv(cfg.Bot.AppPasswordEnv, secret); err != nil {
			return err
		}
	}
	return nil
}

func resolveAccountTokens(ctx context.Context, cfg *config.Config) (map[string]string, error) {
	runner := opsecret.Runner{Binary: cfg.OP.Binary}
	tokens := map[string]string{}
	for i := range cfg.Accounts {
		account := &cfg.Accounts[i]
		if account.TokenEnv == "" {
			account.TokenEnv = "DIRECT_TOKEN_" + strings.ToUpper(strings.NewReplacer("-", "_", ".", "_").Replace(account.ID))
		}
		if token := os.Getenv(account.TokenEnv); token != "" {
			tokens[account.ID] = token
			continue
		}
		if account.TokenRef == "" {
			return nil, fmt.Errorf("account %q token env %s is empty", account.ID, account.TokenEnv)
		}
		token, err := runner.Read(ctx, account.TokenRef)
		if err != nil {
			return nil, err
		}
		if token == "" {
			return nil, fmt.Errorf("account %q token ref is empty", account.ID)
		}
		tokens[account.ID] = token
	}
	return tokens, nil
}

func promptCredentials() (string, string, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	fmt.Print("Password: ")
	pw, err := term.ReadPassword(int(os.Stdin.Fd()))
	fmt.Println()
	if err != nil {
		return "", "", err
	}
	return strings.TrimSpace(email), strings.TrimSpace(string(pw)), nil
}

func extractToken(result interface{}) string {
	if token, ok := result.(string); ok {
		return token
	}
	if m, ok := result.(map[string]interface{}); ok {
		if token, ok := m["access_token"].(string); ok {
			return token
		}
	}
	if arr, ok := result.([]interface{}); ok && len(arr) > 0 {
		if token, ok := arr[0].(string); ok {
			return token
		}
	}
	return ""
}

func redactRef(ref string) string {
	parts := strings.Split(ref, "/")
	if len(parts) <= 1 {
		return "op://..."
	}
	return strings.Join(parts[:len(parts)-1], "/") + "/..."
}

func usage() error {
	return fmt.Errorf("usage: direct-teams-bridge <run|login-direct|mappings|channels>")
}
