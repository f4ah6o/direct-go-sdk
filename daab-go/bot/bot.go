// Package bot provides a hubot-like chatbot framework for the direct messaging service.
//
// This package provides a high-level Robot type that simplifies building chatbots.
// It handles connection management, message routing, and event handling, allowing
// developers to focus on defining message patterns and handlers.
//
// Basic Usage:
//
// Create and start a robot:
//
//	robot := bot.New(
//		bot.WithName("mybot"),
//		bot.WithToken(accessToken),
//	)
//
// Register message handlers:
//
//	// Matches any message containing "hello" (case-insensitive)
//	robot.Hear("hello", func(ctx context.Context, res bot.Response) {
//		res.Send("Hello there!")
//	})
//
//	// Matches messages directed at the bot: "@mybot help" or "mybot: help"
//	robot.Respond("help", func(ctx context.Context, res bot.Response) {
//		res.Send("I'm a helpful bot!")
//	})
//
// Start the robot (blocks until interrupted):
//
//	if err := robot.Run(context.Background()); err != nil {
//		log.Fatal(err)
//	}
//
// The framework is inspired by Hubot (https://hubot.github.io/) and provides
// similar pattern-matching and message handling semantics.
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
	// ErrNotConnected is returned when calling methods on a Robot that is not connected.
	// This typically means Run() has not been called yet or the connection was closed.
	ErrNotConnected = errors.New("daab: robot not connected")

	// ErrNoToken is returned when no access token is available for authentication.
	// Set token via WithToken option or HUBOT_DIRECT_TOKEN environment variable.
	ErrNoToken = errors.New("daab: no access token available")
)

// EventType represents robot lifecycle events.
// These are used with the On method to register event handlers.
type EventType string

// Robot lifecycle event constants.
const (
	// EventConnected is emitted when the robot successfully connects to the direct service.
	// This occurs after the session is created but before all data is recovered.
	EventConnected EventType = "connected"

	// EventDisconnected is emitted when the robot disconnects from the service.
	// This can occur due to network issues, context cancellation, or normal shutdown.
	EventDisconnected EventType = "disconnected"

	// EventReady is emitted when the robot has fully initialized and is ready to receive messages.
	// This occurs after data recovery is complete. Start using the robot in this event.
	EventReady EventType = "ready"
)

// Handler is a callback function for messages matched by a listener pattern.
// The context allows for timeouts and cancellation.
// The Response provides methods to send replies and access message details.
type Handler func(ctx context.Context, res Response)

// Listener represents a registered message pattern listener.
// It combines a regex pattern with a handler function.
type Listener struct {
	// Pattern is the compiled regular expression to match against message text.
	Pattern *regexp.Regexp

	// Handler is the callback function invoked when a message matches the pattern.
	Handler Handler

	// IsDirect indicates whether this listener only responds when directly addressed.
	// For Hear() patterns, this is false (matches any message).
	// For Respond() patterns, this is true (matches direct addresses like "@botname pattern").
	IsDirect bool
}

// Response provides context for responding to a received message.
// It provides convenience methods for sending replies and accessing message metadata.
type Response struct {
	// Message is the received message that triggered this response.
	Message direct.ReceivedMessage

	// Match contains the regex match groups from the pattern.
	// Match[0] is the complete match, Match[1] is the first capture group, etc.
	Match []string

	// Robot is a reference to the Robot instance for advanced operations.
	Robot *Robot
}

// Text returns the text content of the matched message.
// Returns an empty string if the message type is not text.
func (r Response) Text() string {
	return r.Message.Text
}

// RoomID returns the room (talk) ID where the message was sent.
// This is useful for sending follow-up messages to the same room.
func (r Response) RoomID() string {
	return r.Message.TalkID
}

// UserID returns the ID of the user who sent the message.
// This can be used to send direct messages or look up user information.
func (r Response) UserID() string {
	return r.Message.UserID
}

// Send sends a text message to the same room where the triggering message was received.
// Returns an error if the send fails.
//
// Example:
//
//	robot.Hear("hello", func(ctx context.Context, res bot.Response) {
//		if err := res.Send("Hello back!"); err != nil {
//			log.Printf("Failed to send message: %v", err)
//		}
//	})
func (r Response) Send(text string) error {
	return r.Robot.client.SendText(r.Message.TalkID, text)
}

// SendSelect sends a multiple-choice poll (select action stamp) to the same room.
// Recipients can click one of the options to respond.
// Returns the created message ID or an error if the send fails.
//
// Parameters:
// - question: The poll question text
// - options: Array of selectable options (e.g., []string{"Yes", "No", "Maybe"})
//
// Example:
//
//	robot.Hear("decide", func(ctx context.Context, res bot.Response) {
//		msgID, err := res.SendSelect("What should we do?", []string{"Option A", "Option B", "Option C"})
//		if err != nil {
//			log.Printf("Failed to send poll: %v", err)
//		}
//	})
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

// Reply sends a text message mentioning the user who sent the original message.
// The message is prefixed with "@userid " for proper mention.
// Returns an error if the send fails.
//
// Example:
//
//	robot.Respond("hi", func(ctx context.Context, res bot.Response) {
//		res.Reply("Hi there, nice to meet you!")
//	})
func (r Response) Reply(text string) error {
	return r.Robot.client.SendText(r.Message.TalkID, fmt.Sprintf("@%s %s", r.Message.UserID, text))
}

// Robot is the main chatbot instance.
// It manages the direct client connection, message routing, and event handling.
// Create instances via New() and start with Run().
type Robot struct {
	// Name is the bot's name used for direct addressing and logging.
	// Messages like "@Name help" or "Name: help" will match Respond patterns.
	Name string

	// Token is the direct API access token.
	// If set via WithToken, it overrides the environment variable.
	// If empty, the environment variable HUBOT_DIRECT_TOKEN is used.
	Token string

	client        *direct.Client
	listeners     []*Listener
	auth          *direct.Auth
	endpoint      string
	proxyURL      string
	eventHandlers map[EventType][]func()
}

// Option is a configuration function for Robot behavior.
// Use with New() to customize the bot.
type Option func(*Robot)

// WithName sets the bot's name.
// The name is used in Respond patterns and for logging.
// Default: "daabgo"
//
// Example:
//
//	robot := bot.New(bot.WithName("mybot"))
func WithName(name string) Option {
	return func(r *Robot) {
		r.Name = name
	}
}

// WithToken sets the direct API access token directly.
// If provided, this overrides the HUBOT_DIRECT_TOKEN environment variable.
// At least one of WithToken or HUBOT_DIRECT_TOKEN environment variable must be set.
//
// Example:
//
//	robot := bot.New(bot.WithToken("your-access-token"))
func WithToken(token string) Option {
	return func(r *Robot) {
		r.Token = token
	}
}

// WithEndpoint sets a custom API endpoint.
// If not set, DefaultEndpoint is used.
// Can also be set via HUBOT_DIRECT_ENDPOINT environment variable.
//
// Example:
//
//	robot := bot.New(bot.WithEndpoint("wss://api.custom.com/..."))
func WithEndpoint(endpoint string) Option {
	return func(r *Robot) {
		r.endpoint = endpoint
	}
}

// WithProxy sets the HTTP proxy URL for connections.
// Can also be set via HUBOT_DIRECT_PROXY_URL, HTTPS_PROXY, or HTTP_PROXY environment variables.
//
// Example:
//
//	robot := bot.New(bot.WithProxy("http://proxy.example.com:8080"))
func WithProxy(proxyURL string) Option {
	return func(r *Robot) {
		r.proxyURL = proxyURL
	}
}

// New creates a new Robot instance with the given options.
// The robot is not connected until Run() is called.
// Configuration can be set via options or environment variables:
// - HUBOT_DIRECT_TOKEN: Access token
// - HUBOT_DIRECT_ENDPOINT: API endpoint
// - HUBOT_DIRECT_PROXY_URL: HTTP proxy
// - HTTPS_PROXY/HTTP_PROXY: HTTP proxy (fallback)
//
// Parameters:
// - opts: Configuration options (WithName, WithToken, WithEndpoint, WithProxy)
//
// Example:
//
//	robot := bot.New(
//		bot.WithName("mybot"),
//		bot.WithToken(token),
//	)
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

// On registers a callback for a robot lifecycle event.
// Handlers for the same event are called in registration order.
// Handlers run in separate goroutines.
//
// Events:
// - EventConnected: Bot connected to service
// - EventReady: Bot ready to receive messages
// - EventDisconnected: Bot disconnected
//
// Example:
//
//	robot.On(bot.EventReady, func() {
//		log.Println("Bot is ready!")
//	})
func (r *Robot) On(event EventType, handler func()) {
	r.eventHandlers[event] = append(r.eventHandlers[event], handler)
}

func (r *Robot) emit(event EventType) {
	for _, handler := range r.eventHandlers[event] {
		go handler()
	}
}

// Hear registers a listener that matches any message containing the pattern.
// The pattern is treated as a regular expression (case-insensitive).
// Listeners are checked against all incoming messages in registration order.
//
// Parameters:
// - pattern: Regular expression pattern to match (case-insensitive)
// - handler: Callback function invoked when a message matches
//
// Example:
//
//	robot.Hear("hello", func(ctx context.Context, res bot.Response) {
//		res.Send("Hello!")
//	})
//
//	robot.Hear("time", func(ctx context.Context, res bot.Response) {
//		res.Send(time.Now().String())
//	})
func (r *Robot) Hear(pattern string, handler Handler) {
	re := regexp.MustCompile("(?i)" + pattern)
	r.listeners = append(r.listeners, &Listener{
		Pattern:  re,
		Handler:  handler,
		IsDirect: false,
	})
}

// Respond registers a listener that only matches messages directed at the bot.
// Responds to messages in the format:
// - "@botname pattern" or
// - "botname: pattern" or
// - "botname, pattern"
//
// The pattern and bot name matching is case-insensitive.
//
// Parameters:
// - pattern: Regular expression pattern to match (case-insensitive)
// - handler: Callback function invoked when a directed message matches
//
// Example:
//
//	robot.Respond("help", func(ctx context.Context, res bot.Response) {
//		res.Reply("I can help with that!")
//	})
//
//	robot.Respond("(hello|hi)", func(ctx context.Context, res bot.Response) {
//		res.Reply("Nice to meet you!")
//	})
func (r *Robot) Respond(pattern string, handler Handler) {
	re := regexp.MustCompile(fmt.Sprintf("(?i)^@?%s[,:]?\\s*%s", r.Name, pattern))
	r.listeners = append(r.listeners, &Listener{
		Pattern:  re,
		Handler:  handler,
		IsDirect: true,
	})
}

// Run starts the bot and blocks until the context is cancelled or the bot is interrupted.
// This method:
// 1. Loads environment variables from .env file
// 2. Authenticates with the access token
// 3. Connects to the direct service
// 4. Starts listening for messages and routes them to registered handlers
// 5. Blocks until context cancellation or Ctrl+C (SIGINT/SIGTERM)
//
// Returns ErrNoToken if no access token is available.
// Returns a connection error if the WebSocket connection fails.
//
// Example:
//
//	ctx, cancel := context.WithCancel(context.Background())
//	defer cancel()
//
//	robot := bot.New(bot.WithToken(token))
//	robot.Hear("hello", func(ctx context.Context, res bot.Response) {
//		res.Send("Hi there!")
//	})
//
//	if err := robot.Run(ctx); err != nil {
//		log.Fatalf("Bot error: %v", err)
//	}
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

// SendText sends a text message to the specified room.
// This method can be called directly on the Robot to send messages to any room,
// not just replies to incoming messages.
//
// Parameters:
// - roomID: The talk/room ID to send the message to
// - text: The message text
//
// Returns ErrNotConnected if the robot is not running (Run() not called or already closed).
// Returns an error if the message send fails.
//
// Example:
//
//	if err := robot.SendText("room-id", "Message content"); err != nil {
//		log.Printf("Failed to send message: %v", err)
//	}
func (r *Robot) SendText(roomID, text string) error {
	if r.client == nil {
		return ErrNotConnected
	}
	return r.client.SendText(roomID, text)
}

// Call exposes the underlying direct-go Client.Call method for advanced use cases.
// This allows calling any direct API method directly (e.g., GetMessages, GetUsers, etc.).
// See the direct-go documentation for available methods and parameters.
//
// Parameters:
// - method: The RPC method name (e.g., "get_talks", "get_messages")
// - params: An array of parameters for the method
//
// Returns:
// - The result from the API server
// - ErrNotConnected if the robot is not running
// - An error if the RPC call fails
//
// Common use cases:
// - Fetching action stamp (poll) responses
// - Getting message history
// - Managing users and rooms
// - Advanced message operations
//
// Example:
//
//	result, err := robot.Call("get_talks", []interface{}{})
//	if err != nil {
//		log.Printf("API error: %v", err)
//	}
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
