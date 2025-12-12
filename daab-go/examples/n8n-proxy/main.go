// Package main provides an n8n webhook proxy that forwards all direct events to n8n.
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
	"github.com/f4ah6o/direct-go-sdk/daab-go/internal/webhook"
)

func main() {
	// Enable debug server if running
	debugServer := os.Getenv("DEBUG_SERVER")
	if debugServer == "" {
		debugServer = "http://localhost:9999"
	}
	direct.EnableDebugServer(debugServer)

	// Check for N8N_WEBHOOK_URL (will be loaded from .env by robot.Run())
	// We do an early check here to fail fast with a clear error message
	auth := direct.NewAuth()
	if err := auth.LoadEnv(); err != nil {
		log.Printf("Warning: could not load .env: %v", err)
	}
	
	n8nWebhookURL := os.Getenv("N8N_WEBHOOK_URL")
	if n8nWebhookURL == "" {
		log.Fatal("N8N_WEBHOOK_URL environment variable is required")
	}

	robot := bot.New(
		bot.WithName("n8nproxy"),
	)

	// Create webhook client
	webhookClient := webhook.NewClient(n8nWebhookURL, "n8nproxy")

	// Listen to all messages and forward to n8n
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		handleMessage(ctx, res, webhookClient)
	})

	// Run the bot
	if err := robot.Run(context.Background()); err != nil {
		log.Fatalf("Bot error: %v", err)
	}
}


func handleMessage(ctx context.Context, res bot.Response, client *webhook.Client) {
	msg := res.Message

	// Convert to webhook payload
	msgData := webhook.MessageData{
		ID:       msg.ID,
		TalkID:   msg.TalkID,
		UserID:   msg.UserID,
		Type:     int(msg.Type),
		TypeName: webhook.MessageTypeToName(int(msg.Type)),
		Text:     msg.Text,
		Content:  msg.Content,
		Created:  msg.Created,
	}

	payload := webhook.NewPayload("message_created", client.BotName, msgData)

	log.Printf("[N8N PROXY] Forwarding message: type=%s user=%s talk=%s",
		msgData.TypeName, msgData.UserID, msgData.TalkID)

	// Send to n8n
	resp, err := client.Send(payload)
	if err != nil {
		log.Printf("[N8N PROXY] Error sending to n8n: %v", err)
		return
	}

	// Validate response
	if errCode := resp.Validate(); errCode != webhook.ErrorCodeOK {
		log.Printf("[N8N PROXY] Invalid response from n8n: %s", errCode)
		return
	}

	// Execute action from n8n
	if err := executeAction(ctx, res, resp); err != nil {
		log.Printf("[N8N PROXY] Error executing action: %v", err)
	}
}

func executeAction(ctx context.Context, res bot.Response, resp *webhook.WebhookResponse) error {
	log.Printf("[N8N PROXY] Executing action: %s", resp.Action)

	switch resp.Action {
	case "none":
		// Do nothing
		return nil

	case "reply":
		return res.Send(resp.Text)

	case "send":
		return res.Robot.SendText(resp.RoomID, resp.Text)

	case "send_select":
		// TODO: Implement SendSelect in bot package with Question and Options
		log.Printf("[N8N PROXY] send_select not fully implemented yet: question=%s options=%v",
			resp.Question, resp.Options)
		return nil

	case "send_yesno":
		// TODO: Implement SendYesNo in bot package
		log.Printf("[N8N PROXY] send_yesno not fully implemented yet: question=%s", resp.Question)
		return nil

	case "send_task":
		// TODO: Implement SendTask in bot package
		log.Printf("[N8N PROXY] send_task not fully implemented yet: title=%s", resp.Title)
		return nil

	case "reply_select":
		// TODO: Implement ReplySelect in bot package
		log.Printf("[N8N PROXY] reply_select not fully implemented yet: inReplyTo=%s response=%v",
			resp.InReplyTo, resp.Response)
		return nil

	case "reply_yesno":
		// TODO: Implement ReplyYesNo in bot package
		log.Printf("[N8N PROXY] reply_yesno not fully implemented yet: inReplyTo=%s response=%v",
			resp.InReplyTo, resp.ResponseBool)
		return nil

	case "reply_task":
		// TODO: Implement ReplyTask in bot package
		log.Printf("[N8N PROXY] reply_task not fully implemented yet: inReplyTo=%s done=%v",
			resp.InReplyTo, resp.Done)
		return nil

	case "close_select":
		// TODO: Implement CloseSelect in bot package
		log.Printf("[N8N PROXY] close_select not fully implemented yet: messageId=%s", resp.MessageID)
		return nil

	case "close_yesno":
		// TODO: Implement CloseYesNo in bot package
		log.Printf("[N8N PROXY] close_yesno not fully implemented yet: messageId=%s", resp.MessageID)
		return nil

	default:
		return fmt.Errorf("unknown action: %s", resp.Action)
	}
}
