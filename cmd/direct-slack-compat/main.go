package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/slackcompat"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	if err := run(os.Args[1:], logger); err != nil {
		logger.Fatal(err)
	}
}

func run(args []string, logger *log.Logger) error {
	fs := flag.NewFlagSet("direct-slack-compat", flag.ExitOnError)
	configPath := fs.String("config", "slackcompat.config.yaml", "config file")
	if err := fs.Parse(args); err != nil {
		return err
	}
	cfg, err := slackcompat.LoadConfig(*configPath)
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	tokens, err := slackcompat.ResolveTokens(ctx, cfg)
	if err != nil {
		return err
	}
	serverToken, err := slackcompat.ResolveServerBearerToken(ctx, cfg)
	if err != nil {
		return err
	}
	prodClients, err := slackcompat.NewConnectedProductionClients(ctx, cfg, tokens)
	if err != nil {
		return err
	}
	defer func() {
		for _, client := range prodClients {
			_ = client.Close()
		}
	}()
	clients := make([]slackcompat.DirectAPI, 0, len(prodClients))
	for _, client := range prodClients {
		clients = append(clients, client)
	}
	server := slackcompat.NewServer(clients,
		slackcompat.WithTeam(cfg.Slack.TeamID, cfg.Slack.TeamName),
		slackcompat.WithBotUserID(cfg.Slack.BotUserID),
		slackcompat.WithBearerToken(serverToken),
		slackcompat.WithLogger(logger),
	)
	if cfg.Slack.EventCallbackURL != "" {
		sink := slackcompat.HTTPEventSink{URL: cfg.Slack.EventCallbackURL}
		mapper := slackcompat.NewMapper()
		for _, client := range prodClients {
			c := client
			c.OnMessage(func(msg direct.ReceivedMessage) {
				event := slackcompat.ConvertMessageEvent(mapper, cfg.Slack.TeamID, c.AccountID(), msg)
				if err := sink.Publish(context.Background(), event); err != nil {
					logger.Printf("[slackcompat] event callback failed account=%s message=%s err=%v", c.AccountID(), msg.ID, err)
				}
			})
		}
	}
	httpServer := &http.Server{
		Addr:              cfg.Server.ListenAddr,
		Handler:           server.HTTPHandler(),
		ReadHeaderTimeout: 10 * time.Second,
	}
	errCh := make(chan error, 1)
	go func() {
		logger.Printf("[slackcompat] listening on %s", cfg.Server.ListenAddr)
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
