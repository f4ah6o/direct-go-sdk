// Package main provides a Teams bridge bot that forwards messages
// between direct (1:1 pair talks) and Microsoft Teams via n8n.
package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/daab-go/bot"
)

// N8NPayload is the JSON payload sent to n8n webhook
type N8NPayload struct {
	UserID  string `json:"userId"`
	TalkID  string `json:"talkId"`
	Message string `json:"message"`
}

// TeamsReply is the JSON payload received from n8n when Teams replies
type TeamsReply struct {
	UserID  string `json:"userId"`
	TalkID  string `json:"talkId"`
	Message string `json:"message"`
}

var robot *bot.Robot

func main() {
	// Enable debug server if running
	debugServer := os.Getenv("DEBUG_SERVER")
	if debugServer == "" {
		debugServer = "http://localhost:9999"
	}
	direct.EnableDebugServer(debugServer)

	n8nWebhookURL := os.Getenv("N8N_WEBHOOK_URL")
	if n8nWebhookURL == "" {
		log.Fatal("N8N_WEBHOOK_URL environment variable is required")
	}

	callbackPort := os.Getenv("CALLBACK_PORT")
	if callbackPort == "" {
		callbackPort = "8080"
	}

	robot = bot.New(
		bot.WithName("support"),
	)

	// Start HTTP server for receiving replies from n8n
	go startCallbackServer(callbackPort)

	// Handle all messages in 1:1 pair talks
	robot.Hear(".*", func(ctx context.Context, res bot.Response) {
		// Only process 1:1 pair talks
		// In direct, pair talks have exactly 2 members (bot + user)
		// For now, we process all messages - filtering can be added later

		userID := res.UserID()
		talkID := res.RoomID()
		message := res.Text()

		log.Printf("[BRIDGE] Received message from user=%s talk=%s: %s", userID, talkID, message)

		// Forward to n8n
		payload := N8NPayload{
			UserID:  userID,
			TalkID:  talkID,
			Message: message,
		}

		if err := sendToN8N(n8nWebhookURL, payload); err != nil {
			log.Printf("[BRIDGE] Error sending to n8n: %v", err)
		}
	})

	// Run the bot in a goroutine
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	errCh := make(chan error, 1)
	go func() {
		if err := robot.Run(ctx); err != nil {
			errCh <- err
		}
	}()

	// Wait for interrupt or error
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-sigCh:
		log.Printf("Received signal %s, shutting down...", sig)
		cancel()
	case err := <-errCh:
		log.Printf("Bot error: %v", err)
	}
}

// sendToN8N forwards a message to the n8n webhook
func sendToN8N(webhookURL string, payload N8NPayload) error {
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %w", err)
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to post to n8n: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("n8n returned status %d", resp.StatusCode)
	}

	log.Printf("[BRIDGE] Message forwarded to n8n successfully")
	return nil
}

// startCallbackServer starts the HTTP server for receiving replies from n8n
func startCallbackServer(port string) {
	http.HandleFunc("/webhook/teams-reply", handleTeamsReply)

	log.Printf("[BRIDGE] Starting callback server on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start callback server: %v", err)
	}
}

// handleTeamsReply handles incoming replies from n8n (Teams → direct)
func handleTeamsReply(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var reply TeamsReply
	if err := json.NewDecoder(r.Body).Decode(&reply); err != nil {
		log.Printf("[BRIDGE] Failed to decode reply: %v", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	log.Printf("[BRIDGE] Received reply from Teams for user=%s talk=%s: %s",
		reply.UserID, reply.TalkID, reply.Message)

	// Send reply to direct
	if err := robot.SendText(reply.TalkID, reply.Message); err != nil {
		log.Printf("[BRIDGE] Failed to send reply to direct: %v", err)
		http.Error(w, "Failed to send to direct", http.StatusInternalServerError)
		return
	}

	log.Printf("[BRIDGE] Reply sent to direct successfully")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"ok"}`))
}
