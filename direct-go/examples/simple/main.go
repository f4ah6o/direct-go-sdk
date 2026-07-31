// Example: Simple direct client usage
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
)

func main() {
	// Load .env file
	auth := direct.NewAuth()
	if err := auth.LoadEnv(); err != nil {
		log.Printf("Warning: could not load .env: %s", debuglog.SummarizePayload(err))
	}

	// Get token from environment or .env
	token := auth.GetToken()
	if token == "" {
		log.Fatal("No access token found. Set HUBOT_DIRECT_TOKEN environment variable or run daabgo login.")
	}

	// Create client
	client := direct.NewClient(direct.Options{
		AccessToken: token,
	})

	// Register event handlers
	client.On(direct.EventSessionCreated, func(data interface{}) {
		fmt.Println("Session created successfully!")
	})

	client.On(direct.EventDataRecovered, func(data interface{}) {
		fmt.Println("Data recovered, ready to receive messages.")
	})

	client.On(direct.EventError, func(data interface{}) {
		fmt.Printf("Error: %s\n", debuglog.SummarizePayload(data))
	})

	// Register message handler
	client.OnMessage(func(msg direct.ReceivedMessage) {
		fmt.Printf("[talk=%s] user=%s text=%s\n", debuglog.RedactID(msg.TalkID), debuglog.RedactID(msg.UserID), debuglog.SummarizePayload(msg.Text))
	})

	// Connect
	fmt.Println("Connecting to direct...")
	if err := client.ConnectWithContext(context.Background()); err != nil {
		log.Fatalf("Failed to connect: %s", debuglog.SummarizePayload(err))
	}
	defer client.Close()

	fmt.Println("Connected! Waiting for messages... Press Ctrl+C to exit.")

	// Wait for interrupt
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigCh:
		fmt.Println("\nShutting down...")
	case <-client.Done:
		fmt.Println("\nConnection closed.")
	}
}
