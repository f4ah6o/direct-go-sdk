// Example: Ping bot using daabgo framework
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
)

func main() {
	// Enable debug server if running
	debugServer := os.Getenv("DEBUG_SERVER")
	if debugServer == "" {
		debugServer = "http://localhost:9999"
	}
	direct.EnableDebugServer(debugServer)

	robot := bot.New(
		bot.WithName("pingbot"),
	)

	// Respond to "ping" command
	robot.Respond("ping$", func(ctx context.Context, res bot.Response) {
		if err := res.Send("PONG"); err != nil {
			log.Printf("Error sending PONG: %v", err)
		}
	})

	// Respond to "echo <text>" command
	robot.Respond("echo (.+)$", func(ctx context.Context, res bot.Response) {
		if len(res.Match) > 1 {
			if err := res.Send(res.Match[1]); err != nil {
				log.Printf("Error sending echo: %v", err)
			}
		}
	})

	// Respond to "time" command
	robot.Respond("time$", func(ctx context.Context, res bot.Response) {
		msg := fmt.Sprintf("Server time is: %s", time.Now().Format(time.RFC1123))
		if err := res.Send(msg); err != nil {
			log.Printf("Error sending time: %v", err)
		}
	})

	// Respond to "shout <text>" command
	robot.Respond("shout (.+)$", func(ctx context.Context, res bot.Response) {
		if len(res.Match) > 1 {
			text := res.Match[1]
			// Send to the same room where the command was received
			if err := res.Send(text); err != nil {
				log.Printf("Error shouting: %v", err)
			}
		}
	})

	// Hear all messages (optional logging)
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		fmt.Printf("[%s] %s: %s\n", res.RoomID(), res.UserID(), res.Text())
	})

	// Run the bot with context
	ctx := context.Background()
	if err := robot.Run(ctx); err != nil {
		log.Fatalf("Bot error: %v", err)
	}
}
