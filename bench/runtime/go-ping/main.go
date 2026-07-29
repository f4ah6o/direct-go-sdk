package main

import (
	"context"
	"log"
	"os"
	"regexp"

	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
)

func main() {
	name := getenv("BOT_NAME", "runtime-ping")
	pingText := getenv("BENCH_PING_TEXT", "ping")

	robot := bot.New(bot.WithName(name))
	robot.On(bot.EventReady, func() {
		log.Printf("READY runtime-go-ping")
	})
	robot.Hear("^"+regexp.QuoteMeta(pingText)+"$", func(ctx context.Context, res bot.Response) {
		if err := res.Send("PONG"); err != nil {
			log.Printf("send failed: %s", debuglog.SummarizePayload(err))
			return
		}
		log.Printf("RUNTIME pong sent text=%s room=%s user=%s", debuglog.SummarizePayload(res.Text()), debuglog.RedactID(res.RoomID()), debuglog.RedactID(res.UserID()))
	})

	if err := robot.Run(context.Background()); err != nil {
		log.Fatalf("bot failed: %s", debuglog.SummarizePayload(err))
	}
}

func getenv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}
