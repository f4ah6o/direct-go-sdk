package cli

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the daabgo bot",
	Long:  `Run the bot defined in the current directory. Press Ctrl+C to stop.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return runBot()
	},
}

func runBot() error {
	auth := direct.NewAuth()

	// Load environment
	if err := auth.LoadEnv(); err != nil {
		log.Printf("Warning: could not load .env: %v", err)
	}

	// Check if logged in
	if !auth.HasToken() {
		fmt.Println("Not logged in. Run 'daabgo login' first.")
		return nil
	}

	token := auth.GetToken()

	endpoint := os.Getenv("HUBOT_DIRECT_ENDPOINT")
	if endpoint == "" {
		endpoint = direct.DefaultEndpoint
	}

	proxyURL := os.Getenv("HUBOT_DIRECT_PROXY_URL")
	if proxyURL == "" {
		proxyURL = os.Getenv("HTTPS_PROXY")
	}
	if proxyURL == "" {
		proxyURL = os.Getenv("HTTP_PROXY")
	}

	client := direct.NewClient(direct.Options{
		Endpoint:    endpoint,
		AccessToken: token,
		ProxyURL:    proxyURL,
	})

	// Register event handlers
	client.On(direct.EventSessionCreated, func(data interface{}) {
		fmt.Println("Session created successfully!")
	})

	client.On(direct.EventDataRecovered, func(data interface{}) {
		fmt.Println("Ready to receive messages.")
	})

	client.On(direct.EventError, func(data interface{}) {
		fmt.Printf("Error: %v\n", data)
	})

	// Simple echo handler
	client.OnMessage(func(msg direct.ReceivedMessage) {
		fmt.Printf("[%s] Message: %s\n", msg.TalkID, msg.Text)

		// Echo back
		if msg.Text != "" {
			client.SendText(msg.TalkID, "Echo: "+msg.Text)
		}
	})

	fmt.Println("Connecting to direct...")
	if err := client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer client.Close()

	fmt.Println("Bot is running! Press Ctrl+C to stop.")

	// Wait for interrupt
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigCh:
		fmt.Println("\nShutting down...")
	case <-client.Done:
		fmt.Println("\nConnection closed.")
	}

	return nil
}
