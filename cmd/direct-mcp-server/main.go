package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/mcpserver"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds)
	if err := run(os.Args[1:], logger); err != nil {
		logger.Fatal(err)
	}
}

func run(args []string, logger *log.Logger) error {
	fs := flag.NewFlagSet("direct-mcp-server", flag.ExitOnError)
	configPath := fs.String("config", "mcp.config.yaml", "config file")
	if err := fs.Parse(args); err != nil {
		return err
	}
	cfg, err := mcpserver.LoadConfig(*configPath)
	if err != nil {
		return err
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	server, err := mcpserver.New(ctx, cfg, logger)
	if err != nil {
		return err
	}
	return server.Run(ctx)
}
