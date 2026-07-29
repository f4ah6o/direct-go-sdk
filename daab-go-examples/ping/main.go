// Example: Ping bot using daabgo framework
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
)

func main() {
	// Enable safe diagnostics only when explicitly configured.
	if debugServer := os.Getenv("DEBUG_SERVER"); debugServer != "" {
		direct.EnableDebugServer(debugServer)
	}

	robot := bot.New(
		bot.WithName("pingbot"),
	)

	// Respond to "ping" command
	robot.Respond("ping$", func(ctx context.Context, res bot.Response) {
		if err := res.Send("PONG"); err != nil {
			log.Printf("Error sending PONG: %s", debuglog.SummarizePayload(err))
		}
	})

	// Respond to "echo <text>" command
	robot.Respond("echo (.+)$", func(ctx context.Context, res bot.Response) {
		if len(res.Match) > 1 {
			if err := res.Send(res.Match[1]); err != nil {
				log.Printf("Error sending echo: %s", debuglog.SummarizePayload(err))
			}
		}
	})

	// Respond to "time" command
	robot.Respond("time$", func(ctx context.Context, res bot.Response) {
		msg := fmt.Sprintf("Server time is: %s", time.Now().Format(time.RFC1123))
		if err := res.Send(msg); err != nil {
			log.Printf("Error sending time: %s", debuglog.SummarizePayload(err))
		}
	})

	// Respond to "shout <text>" command
	robot.Respond("shout (.+)$", func(ctx context.Context, res bot.Response) {
		if len(res.Match) > 1 {
			text := res.Match[1]
			// Send to the same room where the command was received
			if err := res.Send(text); err != nil {
				log.Printf("Error shouting: %s", debuglog.SummarizePayload(err))
			}
		}
	})

	// Hear all messages (optional logging)
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		fmt.Printf("[talk=%s] user=%s text=%s\n", debuglog.RedactID(res.RoomID()), debuglog.RedactID(res.UserID()), debuglog.SummarizePayload(res.Text()))
	})

	// Run the bot with context
	ctx := context.Background()
	if err := robot.Run(ctx); err != nil {
		log.Fatalf("Bot error: %s", debuglog.SummarizePayload(err))
	}
}
