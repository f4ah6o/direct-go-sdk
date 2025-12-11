// Example: Simple direct client usage
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	direct "github.com/f4ah6o/direct-go"
)

func main() {
	// Load .env file
	auth := direct.NewAuth()
	if err := auth.LoadEnv(); err != nil {
		log.Printf("Warning: could not load .env: %v", err)
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
		fmt.Printf("Error: %v\n", data)
	})

	// Register message handler
	client.OnMessage(func(msg direct.ReceivedMessage) {
		fmt.Printf("[%s] User %s: %s\n", msg.TalkID, msg.UserID, msg.Text)
	})

	// Connect
	fmt.Println("Connecting to direct...")
	if err := client.Connect(); err != nil {
		log.Fatalf("Failed to connect: %v", err)
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
