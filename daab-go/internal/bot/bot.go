// Package bot provides a hubot-like bot framework for direct.
package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

// Handler is a callback for matched messages.
type Handler func(Response)

// Listener represents a registered message listener.
type Listener struct {
	Pattern  *regexp.Regexp
	Handler  Handler
	IsDirect bool // If true, only responds when directly addressed
}

// Response provides context for responding to a message.
type Response struct {
	Message direct.ReceivedMessage
	Match   []string
	Robot   *Robot
}

// Text returns the text of the message.
func (r Response) Text() string {
	return r.Message.Text
}

// RoomID returns the room ID of the message.
func (r Response) RoomID() string {
	return r.Message.TalkID
}

// UserID returns the user ID who sent the message.
func (r Response) UserID() string {
	return r.Message.UserID
}

// Send sends a text message to the same room.
func (r Response) Send(text string) error {
	return r.Robot.client.SendText(r.Message.TalkID, text)
}

// SendSelect sends a select action stamp to the same room and returns the created message ID.
func (r Response) SendSelect(question string, options []string) (string, error) {
	content := direct.SelectMessage{
		Question:    question,
		Options:     options,
		Listing:     true,
		ClosingType: 1, // default to "all must answer" per daab spec
	}
	return r.Robot.sendActionMessage(r.Message.TalkID, direct.MsgTypeSelect, content)
}

// Reply sends a reply mentioning the user.
func (r Response) Reply(text string) error {
	return r.Robot.client.SendText(r.Message.TalkID, fmt.Sprintf("@%s %s", r.Message.UserID, text))
}

// Robot is the main bot instance.
type Robot struct {
	Name      string
	Token     string // Access token (optional, overrides env)
	client    *direct.Client
	listeners []*Listener
	auth      *direct.Auth
}

// New creates a new Robot instance.
func New() *Robot {
	return &Robot{
		Name:      "daabgo",
		listeners: make([]*Listener, 0),
		auth:      direct.NewAuth(),
	}
}

// Hear registers a listener that matches any message containing the pattern.
func (r *Robot) Hear(pattern string, handler Handler) {
	re := regexp.MustCompile("(?i)" + pattern)
	r.listeners = append(r.listeners, &Listener{
		Pattern:  re,
		Handler:  handler,
		IsDirect: false,
	})
}

// Respond registers a listener that only matches messages directed at the bot.
func (r *Robot) Respond(pattern string, handler Handler) {
	re := regexp.MustCompile(fmt.Sprintf("(?i)^@?%s[,:]?\\s*%s", r.Name, pattern))
	r.listeners = append(r.listeners, &Listener{
		Pattern:  re,
		Handler:  handler,
		IsDirect: true,
	})
}

// Run starts the bot and blocks until interrupted.
func (r *Robot) Run() error {
	// Load environment
	if err := r.auth.LoadEnv(); err != nil {
		log.Printf("Warning: could not load .env: %v", err)
	}

	// Get token
	token := r.Token
	if token == "" {
		token = r.auth.GetToken()
	}
	if token == "" {
		return fmt.Errorf("no access token. Run 'daabgo login' first or set Robot.Token")
	}

	// Get configuration from environment
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

	// Create client
	r.client = direct.NewClient(direct.Options{
		Endpoint:    endpoint,
		AccessToken: token,
		ProxyURL:    proxyURL,
		Name:        r.Name,
	})

	// Register event handlers
	r.client.On(direct.EventSessionCreated, func(data interface{}) {
		fmt.Printf("%s: Session created\n", r.Name)
	})

	r.client.On(direct.EventDataRecovered, func(data interface{}) {
		fmt.Printf("%s: Ready to receive messages\n", r.Name)
	})

	// Register message handler
	r.client.OnMessage(r.handleMessage)

	// Connect
	fmt.Printf("%s is starting...\n", r.Name)
	if err := r.client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer r.client.Close()

	fmt.Printf("%s is running! Press Ctrl+C to stop.\n", r.Name)

	// Wait for interrupt
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigCh:
		fmt.Printf("\n%s is shutting down...\n", r.Name)
	case <-r.client.Done:
		fmt.Printf("\n%s: connection closed.\n", r.Name)
	}

	return nil
}

// handleMessage processes incoming messages.
func (r *Robot) handleMessage(msg direct.ReceivedMessage) {
	for _, listener := range r.listeners {
		matches := listener.Pattern.FindStringSubmatch(msg.Text)
		if matches != nil {
			log.Printf("[DEBUG] Matched pattern: %s with text: %s", listener.Pattern.String(), msg.Text)
			response := Response{
				Message: msg,
				Match:   matches,
				Robot:   r,
			}
			go listener.Handler(response)
		} else {
			// log.Printf("[DEBUG] Did not match pattern: %s with text: %s", listener.Pattern.String(), msg.Text)
		}
	}
}

// SendText sends a text message to a room.
func (r *Robot) SendText(roomID, text string) error {
	if r.client == nil {
		return fmt.Errorf("not connected")
	}
	return r.client.SendText(roomID, text)
}

// Call exposes direct-go Client.Call for advanced use cases such as fetching action stamp answers.
func (r *Robot) Call(method string, params []interface{}) (interface{}, error) {
	if r.client == nil {
		return nil, fmt.Errorf("not connected")
	}
	return r.client.Call(method, params)
}

func (r *Robot) sendActionMessage(roomID string, msgType int, content interface{}) (string, error) {
	if r.client == nil {
		return "", fmt.Errorf("not connected")
	}

	talkID := normalizeRoomID(roomID)
	result, err := r.client.Call(direct.MethodCreateMessage, []interface{}{talkID, msgType, content})
	if err != nil {
		return "", err
	}

	messageID := extractMessageID(result)
	if messageID == "" {
		return "", fmt.Errorf("create_message returned empty id")
	}
	return messageID, nil
}

func normalizeRoomID(roomID string) interface{} {
	if id, err := strconv.ParseUint(roomID, 10, 64); err == nil {
		return id
	}
	return roomID
}

func extractMessageID(result interface{}) string {
	switch v := result.(type) {
	case map[string]interface{}:
		if id, ok := v["message_id"]; ok {
			return fmt.Sprintf("%v", id)
		}
		if id, ok := v["id"]; ok {
			return fmt.Sprintf("%v", id)
		}
	case string:
		return v
	default:
		if result != nil {
			return fmt.Sprintf("%v", result)
		}
	}
	return ""
}
