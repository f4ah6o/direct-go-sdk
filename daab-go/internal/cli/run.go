package cli

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the daabgo bot",
	Long:  `Run the bot using the high-level daab-go framework. Press Ctrl+C to stop.`,
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

	// Create a new bot instance
	robot := bot.New(
		bot.WithName("daabgo"),
	)

	// Register a handler that responds when the bot is directly mentioned with "ping"
	robot.Respond("ping", func(ctx context.Context, res bot.Response) {
		if err := res.Send("PONG"); err != nil {
			log.Printf("Failed to send response: %v", err)
		}
	})

	// Register a handler that logs all messages
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		log.Printf("[%s] %s: %s", res.RoomID(), res.UserID(), res.Text())
	})

	// Run the bot
	if err := robot.Run(context.Background()); err != nil {
		return fmt.Errorf("failed to run bot: %v", err)
	}

	return nil
}
