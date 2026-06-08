package mcpserver

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	opsecret "github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/secrets/op"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/modelcontextprotocol/go-sdk/oauthex"
)

type Server struct {
	cfg           *Config
	tokens        map[string]string
	authenticator *Authenticator
	mcpServer     *mcp.Server
	clientFactory directClientFactory
	logger        *log.Logger
}

func New(ctx context.Context, cfg *Config, logger *log.Logger) (*Server, error) {
	tokens, err := resolveTokens(ctx, cfg)
	if err != nil {
		return nil, err
	}
	s := &Server{
		cfg:           cfg,
		tokens:        tokens,
		authenticator: NewAuthenticator(cfg.MCP),
		mcpServer: mcp.NewServer(&mcp.Implementation{
			Name:    "direct-mcp-server",
			Version: "0.1.0",
		}, nil),
		clientFactory: newProductionDirectClient,
		logger:        logger,
	}
	s.registerTools()
	return s, nil
}

func (s *Server) HTTPHandler() http.Handler {
	mcpHandler := mcp.NewStreamableHTTPHandler(func(*http.Request) *mcp.Server {
		return s.mcpServer
	}, nil)
	mux := http.NewServeMux()
	// Register both exact and subtree forms so clients that append a resource
	// path under the well-known endpoint still receive the same metadata.
	mux.HandleFunc("/.well-known/oauth-protected-resource", s.handleProtectedResourceMetadata)
	mux.HandleFunc("/.well-known/oauth-protected-resource/", s.handleProtectedResourceMetadata)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	mux.Handle(s.cfg.MCP.EndpointPath, s.authMiddleware(mcpHandler))
	return mux
}

func (s *Server) Run(ctx context.Context) error {
	httpServer := &http.Server{
		Addr:              s.cfg.MCP.ListenAddr,
		Handler:           s.HTTPHandler(),
		ReadHeaderTimeout: 10 * time.Second,
	}
	errCh := make(chan error, 1)
	go func() {
		s.logger.Printf("[mcp] listening on %s%s", s.cfg.MCP.ListenAddr, s.cfg.MCP.EndpointPath)
		errCh <- httpServer.ListenAndServe()
	}()
	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return httpServer.Shutdown(shutdownCtx)
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx, err := s.authenticator.Authenticate(r.Context(), r)
		if err != nil {
			s.writeAuthError(w, err)
			return
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Server) writeAuthError(w http.ResponseWriter, err error) {
	w.Header().Set("WWW-Authenticate", `Bearer resource_metadata="`+s.metadataURL()+`"`)
	if errors.Is(err, ErrForbidden) {
		http.Error(w, "forbidden", http.StatusForbidden)
		return
	}
	http.Error(w, "unauthorized", http.StatusUnauthorized)
}

func (s *Server) handleProtectedResourceMetadata(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	meta := oauthex.ProtectedResourceMetadata{
		Resource:               s.cfg.ResourceURL(),
		AuthorizationServers:   s.cfg.MCP.AuthorizationServers,
		ScopesSupported:        []string{s.cfg.MCP.ReadScope, s.cfg.MCP.WriteScope},
		BearerMethodsSupported: []string{"header"},
		ResourceName:           "Direct4B MCP Server",
	}
	_ = json.NewEncoder(w).Encode(meta)
}

func (s *Server) metadataURL() string {
	return strings.TrimRight(s.cfg.MCP.PublicBaseURL, "/") + "/.well-known/oauth-protected-resource"
}

func resolveTokens(ctx context.Context, cfg *Config) (map[string]string, error) {
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
