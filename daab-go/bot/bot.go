// Package bot provides a hubot-like bot framework for direct.
package bot

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

// Errors returned by bot operations.
var (
	// ErrNotConnected is returned when calling methods on an unconnected Robot.
	ErrNotConnected = errors.New("daab: robot not connected")

	// ErrNoToken is returned when no access token is available.
	ErrNoToken = errors.New("daab: no access token available")
)

// EventType represents robot lifecycle events.
type EventType string

const (
	// EventConnected is emitted when the robot connects to Direct.
	EventConnected EventType = "connected"
	// EventDisconnected is emitted when the robot disconnects.
	EventDisconnected EventType = "disconnected"
	// EventReady is emitted when the robot is ready to receive messages.
	EventReady EventType = "ready"
)

// Handler is a callback for matched messages.
type Handler func(ctx context.Context, res Response)

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
	// Use map format instead of struct to ensure proper msgpack serialization
	content := map[string]interface{}{
		"question":     question,
		"options":      options,
		"listing":      true,
		"closing_type": 1, // default to "all must answer" per daab spec
	}
	// Use wire type (502) not internal enum value (15) for action stamps
	return r.Robot.sendActionMessage(r.Message.TalkID, direct.WireTypeSelect, content)
}

// Reply sends a reply mentioning the user.
func (r Response) Reply(text string) error {
	return r.Robot.client.SendText(r.Message.TalkID, fmt.Sprintf("@%s %s", r.Message.UserID, text))
}

// Robot is the main bot instance.
type Robot struct {
	Name          string
	Token         string // Access token (optional, overrides env)
	client        *direct.Client
	listeners     []*Listener
	auth          *direct.Auth
	endpoint      string
	proxyURL      string
	eventHandlers map[EventType][]func()
}

// Option configures Robot behavior.
type Option func(*Robot)

// WithName sets the bot name.
func WithName(name string) Option {
	return func(r *Robot) {
		r.Name = name
	}
}

// WithToken sets the access token directly.
func WithToken(token string) Option {
	return func(r *Robot) {
		r.Token = token
	}
}

// WithEndpoint sets custom API endpoint.
func WithEndpoint(endpoint string) Option {
	return func(r *Robot) {
		r.endpoint = endpoint
	}
}

// WithProxy sets the proxy URL for connections.
func WithProxy(proxyURL string) Option {
	return func(r *Robot) {
		r.proxyURL = proxyURL
	}
}

// New creates a new Robot with the given options.
func New(opts ...Option) *Robot {
	r := &Robot{
		Name:          "daabgo",
		listeners:     make([]*Listener, 0),
		auth:          direct.NewAuth(),
		eventHandlers: make(map[EventType][]func()),
	}
	for _, opt := range opts {
		opt(r)
	}
	return r
}

// On registers a lifecycle event handler.
func (r *Robot) On(event EventType, handler func()) {
	r.eventHandlers[event] = append(r.eventHandlers[event], handler)
}

func (r *Robot) emit(event EventType) {
	for _, handler := range r.eventHandlers[event] {
		go handler()
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

// Run starts the bot and blocks until the context is cancelled or interrupted.
func (r *Robot) Run(ctx context.Context) error {
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
		return ErrNoToken
	}

	// Get configuration from environment (can be overridden by options)
	endpoint := r.endpoint
	if endpoint == "" {
		endpoint = os.Getenv("HUBOT_DIRECT_ENDPOINT")
	}
	if endpoint == "" {
		endpoint = direct.DefaultEndpoint
	}

	proxyURL := r.proxyURL
	if proxyURL == "" {
		proxyURL = os.Getenv("HUBOT_DIRECT_PROXY_URL")
	}
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
		r.emit(EventConnected)
	})

	r.client.On(direct.EventDataRecovered, func(data interface{}) {
		fmt.Printf("%s: Ready to receive messages\n", r.Name)
		r.emit(EventReady)
	})

	// Register message handler
	r.client.OnMessage(func(msg direct.ReceivedMessage) {
		r.handleMessage(ctx, msg)
	})

	// Connect
	fmt.Printf("%s is starting...\n", r.Name)
	if err := r.client.Connect(); err != nil {
		return fmt.Errorf("failed to connect: %w", err)
	}
	defer func() {
		r.client.Close()
		r.emit(EventDisconnected)
	}()

	fmt.Printf("%s is running! Press Ctrl+C to stop.\n", r.Name)

	// Wait for interrupt or context cancellation
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		fmt.Printf("\n%s: context cancelled, shutting down...\n", r.Name)
	case <-sigCh:
		fmt.Printf("\n%s is shutting down...\n", r.Name)
	case <-r.client.Done:
		fmt.Printf("\n%s: connection closed.\n", r.Name)
	}

	return nil
}

// handleMessage processes incoming messages.
func (r *Robot) handleMessage(ctx context.Context, msg direct.ReceivedMessage) {
	for _, listener := range r.listeners {
		matches := listener.Pattern.FindStringSubmatch(msg.Text)
		if matches != nil {
			log.Printf("[DEBUG] Matched pattern: %s with text: %s", listener.Pattern.String(), msg.Text)
			response := Response{
				Message: msg,
				Match:   matches,
				Robot:   r,
			}
			go listener.Handler(ctx, response)
		}
	}
}

// SendText sends a text message to a room.
func (r *Robot) SendText(roomID, text string) error {
	if r.client == nil {
		return ErrNotConnected
	}
	return r.client.SendText(roomID, text)
}

// Call exposes direct-go Client.Call for advanced use cases such as fetching action stamp answers.
func (r *Robot) Call(method string, params []interface{}) (interface{}, error) {
	if r.client == nil {
		return nil, ErrNotConnected
	}
	return r.client.Call(method, params)
}

func (r *Robot) sendActionMessage(roomID string, msgType int, content interface{}) (string, error) {
	if r.client == nil {
		return "", ErrNotConnected
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
