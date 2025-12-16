// Package direct provides a Go client for the direct chat service.
//
// This package implements the MessagePack RPC protocol used by direct-js
// to communicate with the direct API server over WebSocket.
//
// Basic Usage:
//
// Create a client with credentials:
//
//	client := direct.NewClient(direct.Options{
//		Endpoint:    direct.DefaultEndpoint,
//		AccessToken: token,
//	})
//	defer client.Close()
//
// Connect and register event handlers:
//
//	if err := client.Connect(); err != nil {
//		log.Fatal(err)
//	}
//
//	client.OnMessage(func(msg direct.ReceivedMessage) {
//		log.Printf("Received message: %s", msg.Text)
//	})
//
// Send messages to a room:
//
//	err := client.SendTextWithContext(ctx, "room-id", "Hello!")
//	if err != nil {
//		log.Printf("Failed to send message: %v", err)
//	}
//
// The package provides methods for managing users, talks (conversations), domains,
// messages, files, and other direct API features.
package direct

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

// EnableDebugServer enables sending logs to a debug server for remote debugging.
// The url parameter should be the full address of the debug log server.
// This is useful for debugging client behavior in production or testing environments.
func EnableDebugServer(url string) {
	debuglog.SetServer(url)
}

// dlog is a helper for debug logging (level 1 = normal)
func dlog(format string, v ...interface{}) {
	debuglog.Printf(format, v...)
}

// vlog is a helper for verbose debug logging (level 2 = verbose, includes ping/pong)
func vlog(format string, v ...interface{}) {
	debuglog.Verbose(format, v...)
}

// Protocol constants for MessagePack RPC communication.
const (
	// RpcRequest is the message type for RPC requests from client to server.
	RpcRequest = 0

	// RpcResponse is the message type for RPC responses from server to client.
	RpcResponse = 1

	// APIVersion specifies the protocol version for the direct API.
	APIVersion = "1.128"

	// DefaultEndpoint is the default WebSocket API endpoint for the direct service.
	DefaultEndpoint = "wss://api.direct4b.com/albero-app-server/api"
)

// Options configures the direct client.
//
// Example:
//
//	opts := direct.Options{
//		Endpoint:    direct.DefaultEndpoint,
//		AccessToken: "your-auth-token",
//		ProxyURL:    "http://proxy.example.com:8080",
//		Name:        "MyBot",
//	}
type Options struct {
	// Endpoint is the WebSocket API endpoint.
	// If empty, DefaultEndpoint is used.
	Endpoint string

	// AccessToken is the authentication token obtained from the direct service.
	// If provided, the client automatically creates a session upon connection.
	AccessToken string

	// ProxyURL is an optional HTTP proxy URL for WebSocket connections.
	// Format: "http://proxy.example.com:8080" or "https://proxy.example.com:8080"
	ProxyURL string

	// Host is the API host extracted from Endpoint.
	// This is automatically derived from the Endpoint if not set explicitly.
	Host string

	// Name is the bot/client name used for logging and identification.
	// Useful for distinguishing between multiple client instances in logs.
	Name string
}

// ResponseHandler handles the success and error callbacks for RPC responses.
// It is used internally to match responses to their corresponding requests.
type ResponseHandler struct {
	// Method is the RPC method name that was called.
	Method string

	// OnSuccess is called when the RPC request succeeds.
	// The result parameter contains the response value.
	OnSuccess func(result interface{})

	// OnError is called when the RPC request fails.
	// The err parameter contains the error information from the server.
	OnError func(err interface{})
}

// Client is a direct API client that communicates with the direct chat service.
// It manages WebSocket connections, RPC calls, and event handling.
// Use NewClient to create a new instance.
type Client struct {
	options          Options
	conn             *websocket.Conn
	mu               sync.RWMutex
	handlers         map[string][]EventHandler
	responseHandlers map[int64]*ResponseHandler
	msgID            int64
	closed           bool
	connected        bool

	// talkDomains maps talk_id to domain_id for user lookups
	talkDomains map[string]string

	// Messages is a channel that receives incoming messages from the server.
	// Messages are buffered with a capacity of 100.
	Messages chan ReceivedMessage

	// Done is a channel that is closed when the client shuts down.
	Done chan struct{}
}

// EventHandler is a callback function for event listeners.
// It is called whenever an event is emitted.
// Data parameter contains event-specific information.
// Handlers are called asynchronously in separate goroutines.
type EventHandler func(data interface{})

// NewClient creates a new direct client with the given options.
// If no endpoint is provided, DefaultEndpoint is used.
// The Host field is automatically derived from the Endpoint if not set explicitly.
// The client must be connected via Connect() before use.
//
// Example:
//
//	client := direct.NewClient(direct.Options{
//		Endpoint:    direct.DefaultEndpoint,
//		AccessToken: token,
//	})
//	defer client.Close()
//	if err := client.Connect(); err != nil {
//		log.Fatal(err)
//	}
func NewClient(opts Options) *Client {
	if opts.Endpoint == "" {
		opts.Endpoint = DefaultEndpoint
	}
	if opts.Host == "" {
		if u, err := url.Parse(opts.Endpoint); err == nil {
			opts.Host = u.Host
		}
	}

	return &Client{
		options:          opts,
		handlers:         make(map[string][]EventHandler),
		responseHandlers: make(map[int64]*ResponseHandler),
		talkDomains:      make(map[string]string),
		Messages:         make(chan ReceivedMessage, 100),
		Done:             make(chan struct{}),
	}
}

// Connect establishes a WebSocket connection to the direct API.
// It starts the message reader and ping keepalive loops.
// If an access token is provided in Options, it automatically creates a session
// and initializes data (talks, domains, etc.).
//
// Returns an error if already connected, if the WebSocket connection fails,
// or if the proxy configuration is invalid.
//
// The method is non-blocking and starts background goroutines for communication.
// Use the EventSessionCreated event or call GetTalksWithContext to verify successful
// session creation.
//
// Example:
//
//	if err := client.Connect(); err != nil {
//		log.Fatalf("Failed to connect: %v", err)
//	}
func (c *Client) Connect() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn != nil {
		return fmt.Errorf("already connected")
	}

	dialer := websocket.Dialer{
		HandshakeTimeout: 10 * time.Second,
	}

	// Set up proxy if configured
	if c.options.ProxyURL != "" {
		proxyURL, err := url.Parse(c.options.ProxyURL)
		if err != nil {
			return fmt.Errorf("invalid proxy URL: %w", err)
		}
		dialer.Proxy = http.ProxyURL(proxyURL)
	}

	header := http.Header{}
	// Don't set Origin - let the library handle it
	// header.Set("Origin", "https://"+c.options.Host)

	conn, _, err := dialer.Dial(c.options.Endpoint, header)
	if err != nil {
		return fmt.Errorf("websocket dial failed: %w", err)
	}

	c.conn = conn
	c.closed = false

	// Set up pong handler
	c.conn.SetPongHandler(func(appData string) error {
		vlog("[DEBUG] Received pong: %s", appData)
		return nil
	})

	// Start message reader
	go c.readLoop()

	// Start ping keepalive (every 45 seconds like direct-js)
	go c.pingLoop()

	// Create session if access token is provided
	if c.options.AccessToken != "" {
		go c.createSession()
	}

	return nil
}

// pingLoop sends periodic pings to keep the connection alive
func (c *Client) pingLoop() {
	ticker := time.NewTicker(45 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.RLock()
			closed := c.closed
			conn := c.conn
			c.mu.RUnlock()

			if closed || conn == nil {
				return
			}

			vlog("[DEBUG] Sending ping...")
			if err := conn.WriteMessage(websocket.PingMessage, []byte("PING")); err != nil {
				vlog("[DEBUG] Ping error: %v", err)
				return
			}
		case <-c.Done:
			return
		}
	}
}

// createSession authenticates with the server.
func (c *Client) createSession() {
	dlog("[DEBUG] Creating session with token: %s...", c.options.AccessToken[:min(20, len(c.options.AccessToken))])
	osString := "Go"
	params := []interface{}{c.options.AccessToken, APIVersion, osString}

	c.call("create_session", params, func(result interface{}) {
		dlog("[DEBUG] Session created successfully: %+v", result)
		c.mu.Lock()
		c.connected = true
		c.mu.Unlock()

		c.emit("session_created", result)

		// Start notification after session is created
		c.startNotification()
	}, func(err interface{}) {
		dlog("[DEBUG] Session error: %+v", err)
		c.emit("session_error", err)
	})
}

// startNotification tells the server to start sending notifications.
func (c *Client) startNotification() {
	dlog("[DEBUG] Starting notification...")

	// First, get domains to initialize data
	c.call("get_domains", []interface{}{}, func(result interface{}) {
		dlog("[DEBUG] get_domains success: %d domains", countItems(result))

		// Then get talks
		c.call("get_talks", []interface{}{}, func(result interface{}) {
			dlog("[DEBUG] get_talks success: %d talks", countItems(result))

			// Log talk details and cache talk->domain mapping
			if talks, ok := result.([]interface{}); ok && len(talks) > 0 {
				for i, talk := range talks {
					if talkMap, ok := talk.(map[string]interface{}); ok {
						// Print all keys in the map
						keys := make([]string, 0, len(talkMap))
						for k := range talkMap {
							keys = append(keys, k)
						}
						dlog("[DEBUG] Talk %d keys: %v", i, keys)
						dlog("[DEBUG] Talk %d: %+v", i, talkMap)

						// Cache talk_id -> domain_id mapping
						var talkID, domainID string
						if id, ok := talkMap["talk_id"]; ok {
							talkID = fmt.Sprintf("%v", id)
						}
						if domID, ok := talkMap["domain_id"]; ok {
							domainID = fmt.Sprintf("%v", domID)
						}
						if talkID != "" && domainID != "" {
							c.mu.Lock()
							c.talkDomains[talkID] = domainID
							c.mu.Unlock()
							dlog("[DEBUG] Cached talk->domain: %s -> %s", talkID, domainID)
						}
					} else {
						dlog("[DEBUG] Talk %d: unexpected type %T: %v", i, talk, talk)
					}
				}

				// Try to send a test message to the first talk
				if firstTalk, ok := talks[0].(map[string]interface{}); ok {
					// Find the talk ID - might be "id" or encoded differently
					var talkID interface{}
					for k, v := range firstTalk {
						dlog("[DEBUG] First talk field: %s = %v (type %T)", k, v, v)
						if k == "id" || k == "talk_id" || k == "talkId" {
							talkID = v
						}
					}

					if talkID != nil {
						dlog("[DEBUG] Sending test message to talk: %v", talkID)
						c.call("create_message", []interface{}{}, func(result interface{}) {
							dlog("[DEBUG] create_message success: %+v", result)
						}, func(err interface{}) {
							dlog("[DEBUG] create_message error: %+v", err)
						})
					} else {
						dlog("[DEBUG] Could not find talk ID in first talk")
					}
				}
			} else {
				dlog("[DEBUG] get_talks result is not []interface{}, type=%T", result)
			}

			// Then get talk statuses
			c.call("get_talk_statuses", []interface{}{}, func(result interface{}) {
				dlog("[DEBUG] get_talk_statuses success")

				// Try start_notification first
				c.call("start_notification", []interface{}{}, func(result interface{}) {
					dlog("[DEBUG] start_notification result: %+v", result)

					// If false, try reset_notification and then start_notification again
					if result == false {
						dlog("[DEBUG] start_notification returned false, trying reset_notification...")
						c.call("reset_notification", []interface{}{}, func(result interface{}) {
							dlog("[DEBUG] reset_notification result: %+v", result)

							// After reset, call start_notification again
							c.call("start_notification", []interface{}{}, func(result interface{}) {
								dlog("[DEBUG] start_notification (after reset) result: %+v", result)

								// Call update_last_used_at to mark session as active
								c.call("update_last_used_at", []interface{}{}, func(result interface{}) {
									dlog("[DEBUG] update_last_used_at result: %+v", result)
									c.emit("data_recovered", result)
								}, func(err interface{}) {
									dlog("[DEBUG] update_last_used_at error: %+v", err)
									c.emit("data_recovered", nil)
								})
							}, func(err interface{}) {
								dlog("[DEBUG] start_notification (after reset) error: %+v", err)
								c.emit("notification_error", err)
							})
						}, func(err interface{}) {
							dlog("[DEBUG] reset_notification error: %+v", err)
						})
					} else {
						c.emit("data_recovered", result)
					}
				}, func(err interface{}) {
					dlog("[DEBUG] start_notification error: %+v", err)
					c.emit("notification_error", err)
				})
			}, func(err interface{}) {
				dlog("[DEBUG] get_talk_statuses error: %+v", err)
			})
		}, func(err interface{}) {
			dlog("[DEBUG] get_talks error: %+v", err)
		})
	}, func(err interface{}) {
		dlog("[DEBUG] get_domains error: %+v", err)
	})
}

func countItems(v interface{}) int {
	if arr, ok := v.([]interface{}); ok {
		return len(arr)
	}
	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Close closes the WebSocket connection and stops all background goroutines.
// It is safe to call Close multiple times; subsequent calls are no-ops.
// After Close is called, the client cannot be reconnected; create a new client instead.
//
// Example:
//
//	defer client.Close()
func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.conn == nil || c.closed {
		return nil
	}

	c.closed = true
	close(c.Done)

	return c.conn.Close()
}

// On registers an event handler for the given event type.
// Multiple handlers can be registered for the same event, and they are called
// in the order they were registered.
//
// Event types are defined as constants (e.g., EventSessionCreated, EventNotifyCreateMessage).
// See events.go for the complete list of event constants.
// Handlers are called asynchronously in separate goroutines.
//
// Common events:
// - EventSessionCreated: Fired when the session is successfully created
// - EventNotifyCreateMessage: Fired when a new message is received
// - EventDataRecovered: Fired when initial data has been synced
// - EventError: Fired when a connection or protocol error occurs
//
// Example:
//
//	client.On(direct.EventSessionCreated, func(data interface{}) {
//		log.Println("Session created!")
//	})
func (c *Client) On(event string, handler EventHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers[event] = append(c.handlers[event], handler)
}

// OnMessage registers a callback for incoming messages.
// The handler is called for each message received from the Messages channel.
// The handler runs in a separate goroutine and receives ReceivedMessage objects
// with fields like Text, UserID, RoomID, and Type.
//
// Message objects may contain different content types based on the Message.Type:
// - MessageType.Text: Message.Text contains the message text
// - MessageType.Stamp: Message.Content contains stamp/emoji data
// - MessageType.File: Message.Content contains file attachment info
//
// Example:
//
//	client.OnMessage(func(msg direct.ReceivedMessage) {
//		log.Printf("Message from %s: %s", msg.UserID, msg.Text)
//	})
func (c *Client) OnMessage(handler func(ReceivedMessage)) {
	go func() {
		for msg := range c.Messages {
			handler(msg)
		}
	}()
}

// call sends an RPC request to the server.
func (c *Client) call(method string, params []interface{}, onSuccess func(interface{}), onError func(interface{})) {
	c.mu.Lock()

	if c.conn == nil {
		c.mu.Unlock()
		if onError != nil {
			onError(map[string]string{"message": "not connected"})
		}
		return
	}

	msgID := atomic.AddInt64(&c.msgID, 1)

	// Register response handler
	c.responseHandlers[msgID] = &ResponseHandler{
		Method:    method,
		OnSuccess: onSuccess,
		OnError:   onError,
	}

	c.mu.Unlock()

	// Build MessagePack RPC request: [type, msgId, method, params]
	request := []interface{}{RpcRequest, msgID, method, params}

	data, err := msgpack.Marshal(request)
	if err != nil {
		if onError != nil {
			onError(map[string]string{"message": err.Error()})
		}
		return
	}

	c.mu.Lock()
	err = c.conn.WriteMessage(websocket.BinaryMessage, data)
	c.mu.Unlock()

	if err != nil {
		if onError != nil {
			onError(map[string]string{"message": err.Error()})
		}
	}
}

// Call sends a synchronous RPC request to the direct API server.
// It blocks until a response is received or the 30-second timeout expires.
// Method names are defined as constants (e.g., MethodGetTalks, MethodCreateMessage).
// See events.go for the complete list of method constants.
//
// Parameters:
// - method: The RPC method name to call
// - params: An array of parameters to pass to the method
//
// Returns:
// - On success: The result value from the server and nil error
// - On server error: nil result and an error with "RPC error:" prefix
// - On timeout: nil result and an "RPC timeout" error
//
// Example:
//
//	talks, err := client.Call(direct.MethodGetTalks, []interface{}{})
//	if err != nil {
//		log.Printf("Failed to get talks: %v", err)
//		return
//	}
func (c *Client) Call(method string, params []interface{}) (interface{}, error) {
	resultCh := make(chan interface{}, 1)
	errCh := make(chan interface{}, 1)

	c.call(method, params, func(result interface{}) {
		resultCh <- result
	}, func(err interface{}) {
		errCh <- err
	})

	select {
	case result := <-resultCh:
		return result, nil
	case err := <-errCh:
		return nil, fmt.Errorf("RPC error: %v", err)
	case <-time.After(30 * time.Second):
		return nil, fmt.Errorf("RPC timeout")
	}
}

// Send sends a message with custom type and content to the specified room.
//
// Parameters:
// - roomID: The talk/room identifier (can be a string or numeric value)
// - msgType: One of the MessageType constants (MsgTypeText, MsgTypeStamp, etc.)
// - content: The message content, structure depends on msgType:
//   - MsgTypeText: string containing the message text
//   - MsgTypeStamp: map containing stamp/emoji data
//   - MsgTypeFile: map containing file attachment info
//   - MsgTypeSelect: map containing action stamp options
//
// Returns an error if the RPC call fails or times out.
//
// For most use cases, consider using SendTextWithContext instead.
//
// Example:
//
//	err := client.Send("123456", direct.MsgTypeText, "Hello World")
//	if err != nil {
//		log.Printf("Failed to send message: %v", err)
//	}
func (c *Client) Send(roomID interface{}, msgType int, content interface{}) error {
	_, err := c.Call("create_message", []interface{}{roomID, msgType, content})
	return err
}

// SendText sends a text message to the specified room.
// This is a convenience method that wraps Send with msgType=1 (text).
//
// Parameters:
// - roomID: The talk/room identifier as a string
// - text: The message text content
//
// Returns an error if the RPC call fails or times out.
//
// Deprecated: Use SendTextWithContext for better context support and control.
// SendText always uses message type 1 and internally converts roomID to the appropriate type.
func (c *Client) SendText(roomID string, text string) error {
	// For text messages, type is 1 and content is the text string
	// Convert roomID to uint64 for the API
	var talkID interface{} = roomID
	if id, err := strconv.ParseUint(roomID, 10, 64); err == nil {
		talkID = id
	}
	_, err := c.Call("create_message", []interface{}{talkID, 1, text})
	return err
}

// readLoop continuously reads messages from the WebSocket.
func (c *Client) readLoop() {
	defer close(c.Messages)

	for {
		c.mu.RLock()
		if c.closed {
			c.mu.RUnlock()
			return
		}
		conn := c.conn
		c.mu.RUnlock()

		msgType, data, err := conn.ReadMessage()
		if err != nil {
			if !c.closed {
				dlog("[DEBUG] ReadMessage error: %v", err)
				c.emit("error", map[string]string{"error": err.Error()})
			}
			return
		}

		dlog("[DEBUG] Raw WebSocket message: type=%d len=%d", msgType, len(data))

		c.handleMessage(data)
	}
}

// handleMessage processes an incoming WebSocket message.
func (c *Client) handleMessage(data []byte) {
	// Decode MessagePack
	var message []interface{}
	if err := msgpack.Unmarshal(data, &message); err != nil {
		dlog("[DEBUG] msgpack decode error: %v", err)
		c.emit("decode_error", map[string]string{"error": err.Error()})
		return
	}

	dlog("[DEBUG] Received message: len=%d type=%T", len(message), message)

	if len(message) < 4 {
		dlog("[DEBUG] Message too short: %v", message)
		return
	}

	// Get message type
	msgType, ok := toInt64(message[0])
	if !ok {
		dlog("[DEBUG] Could not get message type: %v", message[0])
		return
	}

	dlog("[DEBUG] Message type: %d", msgType)

	switch msgType {
	case RpcResponse:
		// Response: [1, msgId, error, result]
		c.handleResponse(message)

	case RpcRequest:
		// Request from server (notification): [0, msgId, method, params]
		c.handleNotification(message)
	}
}

// handleResponse processes an RPC response.
func (c *Client) handleResponse(message []interface{}) {
	msgID, ok := toInt64(message[1])
	if !ok {
		return
	}

	c.mu.Lock()
	handler := c.responseHandlers[msgID]
	delete(c.responseHandlers, msgID)
	c.mu.Unlock()

	if handler == nil {
		return
	}

	errVal := message[2]
	result := message[3]

	if errVal != nil {
		if handler.OnError != nil {
			handler.OnError(errVal)
		}
	} else {
		if handler.OnSuccess != nil {
			handler.OnSuccess(result)
		}
	}
}

// handleNotification processes a notification from the server.
func (c *Client) handleNotification(message []interface{}) {
	if len(message) < 4 {
		dlog("[DEBUG] Notification too short: %v", message)
		return
	}

	msgID, _ := toInt64(message[1])
	method, ok := message[2].(string)
	if !ok {
		dlog("[DEBUG] Method not a string: %v", message[2])
		return
	}

	dlog("[DEBUG] <<< SERVER NOTIFICATION: method=%s, msgID=%d", method, msgID)

	params, ok := message[3].([]interface{})
	if !ok || len(params) == 0 {
		dlog("[DEBUG] %s: params invalid or empty: %T %v", method, message[3], message[3])
		return
	}

	dlog("[DEBUG] Received notification: %s, params count: %d", method, len(params))

	// Emit the notification event
	c.emit(method, params[0])

	// Handle message notifications specially
	if method == "notify_create_message" || method == "create_message" {
		dlog("[DEBUG] Message notification received: %s", method)
		dlog("[DEBUG] Data: %+v", params[0])
		c.handleMessageNotification(params[0])
	}

	// Send acknowledgment response: [1, msgId, null, true]
	response := []interface{}{RpcResponse, msgID, nil, true}
	data, err := msgpack.Marshal(response)
	if err == nil {
		c.mu.Lock()
		if c.conn != nil {
			c.conn.WriteMessage(websocket.BinaryMessage, data)
		}
		c.mu.Unlock()
	}
}

// handleMessageNotification parses and queues a message notification.
func (c *Client) handleMessageNotification(data interface{}) {
	dlog("[DEBUG] handleMessageNotification: raw data: %+v", data)
	msg := parseMessage(data)

	// If DomainID is not in the message, look it up from cached talks
	if msg.DomainID == "" && msg.TalkID != "" {
		c.mu.RLock()
		dlog("[DEBUG] Looking up domain for talk_id=%s, cached talks count=%d", msg.TalkID, len(c.talkDomains))
		// Log all cached talk IDs for debugging
		for tid := range c.talkDomains {
			dlog("[DEBUG] Cached talk_id: %s", tid)
		}
		if domID, ok := c.talkDomains[msg.TalkID]; ok {
			msg.DomainID = domID
			dlog("[DEBUG] Resolved DomainID from talkDomains: %s", domID)
		} else {
			dlog("[DEBUG] talk_id %s not found in cache", msg.TalkID)
		}
		c.mu.RUnlock()
	}

	dlog("[DEBUG] handleMessageNotification: parsed msg: ID=%s UserID=%s TalkID=%s DomainID=%s Text=%s",
		msg.ID, msg.UserID, msg.TalkID, msg.DomainID, msg.Text)
	if msg.ID != "" {
		select {
		case c.Messages <- msg:
		default:
			// Channel full, drop message
		}
	}
}

// parseMessage converts a raw notification to a ReceivedMessage.
func parseMessage(data interface{}) ReceivedMessage {
	msg := ReceivedMessage{}

	m, ok := data.(map[string]interface{})
	if !ok {
		dlog("[DEBUG] parseMessage: data not a map, type=%T", data)
		return msg
	}

	dlog("[DEBUG] parseMessage: keys = %v", getMapKeys(m))

	if id, ok := m["message_id"]; ok {
		msg.ID = fmt.Sprintf("%v", id)
	} else if id, ok := m["id"]; ok {
		msg.ID = fmt.Sprintf("%v", id)
	}
	if talkId, ok := m["talk_id"]; ok {
		msg.TalkID = fmt.Sprintf("%v", talkId)
		msg.RoomID = msg.TalkID
	}
	if userId, ok := m["user_id"]; ok {
		msg.UserID = fmt.Sprintf("%v", userId)
	}
	if domainId, ok := m["domain_id"]; ok {
		msg.DomainID = fmt.Sprintf("%v", domainId)
	}
	if content, ok := m["content"]; ok {
		dlog("[DEBUG] content type=%T value=%v", content, content)
		msg.Content = content
		if text, ok := content.(string); ok {
			msg.Text = text
		} else if contentMap, ok := content.(map[string]interface{}); ok {
			if text, ok := contentMap["text"].(string); ok {
				msg.Text = text
			}
		}
	}
	if msgType, ok := m["type"]; ok {
		if t, ok := toInt64(msgType); ok {
			msg.Type = MessageType(t)
		}
	}

	dlog("[DEBUG] parsed: ID=%s TalkID=%s Text=%s", msg.ID, msg.TalkID, msg.Text)

	// Store raw data for custom parsing
	if rawData, err := json.Marshal(m); err == nil {
		msg.Raw = rawData
	}

	return msg
}

// getMapKeys returns the keys of a map for debugging
func getMapKeys(m map[string]interface{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// toInt64 converts various numeric types to int64.
func toInt64(v interface{}) (int64, bool) {
	switch n := v.(type) {
	case int:
		return int64(n), true
	case int8:
		return int64(n), true
	case int16:
		return int64(n), true
	case int32:
		return int64(n), true
	case int64:
		return n, true
	case uint:
		return int64(n), true
	case uint8:
		return int64(n), true
	case uint16:
		return int64(n), true
	case uint32:
		return int64(n), true
	case uint64:
		return int64(n), true
	case float32:
		return int64(n), true
	case float64:
		return int64(n), true
	default:
		return 0, false
	}
}

// emit dispatches an event to registered handlers.
func (c *Client) emit(event string, data interface{}) {
	c.mu.RLock()
	handlers := c.handlers[event]
	c.mu.RUnlock()

	for _, h := range handlers {
		go h(data)
	}
}

// GetTalksWithContext retrieves the list of talk rooms (conversations) with context support.
// Each Talk contains room metadata including participants, type (pair/group), and settings.
// This is the preferred method over the legacy GetTalks().
func (c *Client) GetTalksWithContext(ctx context.Context) ([]Talk, error) {
	result, err := c.Call(MethodGetTalks, []interface{}{})
	if err != nil {
		return nil, err
	}

	talks := []Talk{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if talkData, ok := item.(map[string]interface{}); ok {
				talk := parseTalk(talkData)
				talks = append(talks, *talk)
			}
		}
	}

	return talks, nil
}

// GetTalkStatusesWithContext retrieves the status of all talks with context support.
// Status includes unread count and latest message ID for each talk.
func (c *Client) GetTalkStatusesWithContext(ctx context.Context) ([]TalkStatus, error) {
	result, err := c.Call(MethodGetTalkStatuses, []interface{}{})
	if err != nil {
		return nil, err
	}

	statuses := []TalkStatus{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if statusData, ok := item.(map[string]interface{}); ok {
				status := TalkStatus{}
				if v, ok := statusData["talk_id"]; ok {
					status.TalkID = v
				}
				if v, ok := statusData["unread_count"].(int); ok {
					status.UnreadCount = v
				}
				if v, ok := statusData["latest_msg_id"]; ok {
					status.LatestMsgID = v
				}
				statuses = append(statuses, status)
			}
		}
	}

	return statuses, nil
}

// GetMeWithContext retrieves the current authenticated user's profile with context support.
// Returns user information including display name, email, status, and other profile details.
// This is the preferred method over the legacy GetMe().
func (c *Client) GetMeWithContext(ctx context.Context) (*UserInfo, error) {
	result, err := c.Call(MethodGetMe, []interface{}{})
	if err != nil {
		return nil, err
	}

	if userData, ok := result.(map[string]interface{}); ok {
		user := parseUserInfo(userData)
		return &user, nil
	}

	return nil, nil
}

// SendTextWithContext sends a text message to the specified room with context support.
//
// Parameters:
// - ctx: Context for cancellation and timeouts
// - roomID: The talk/room identifier as a string
// - text: The message text content
//
// Returns an error if the RPC call fails or times out.
//
// This is the preferred method for sending text messages. It supports context
// for timeout control and graceful cancellation.
//
// Example:
//
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//	err := client.SendTextWithContext(ctx, "room-123", "Hello!")
func (c *Client) SendTextWithContext(ctx context.Context, roomID string, text string) error {
	_, err := c.Call(MethodCreateMessage, []interface{}{roomID, 1, text})
	return err
}

// Legacy methods below - deprecated, use context-aware versions instead.
// These methods are maintained for backwards compatibility but should not be used
// in new code.

// GetTalks retrieves the list of talk rooms.
// Deprecated: Use GetTalksWithContext instead for better context support and type safety.
func (c *Client) GetTalks() (interface{}, error) {
	return c.Call("get_talks", []interface{}{})
}

// GetDomains retrieves the list of domains (organizations).
// Deprecated: Use GetDomainsWithContext instead for better context support and type safety.
func (c *Client) GetDomains() (interface{}, error) {
	return c.Call("get_domains", []interface{}{})
}

// GetDomainInvites retrieves pending domain invitations for the current user.
// Deprecated: Use GetDomainInvitesWithContext instead for better context support and type safety.
func (c *Client) GetDomainInvites() (interface{}, error) {
	return c.Call("get_domain_invites", []interface{}{})
}

// AcceptDomainInvite accepts a pending domain invitation.
// Parameters:
// - inviteID: The ID of the domain invitation to accept
// Deprecated: Use AcceptDomainInviteWithContext instead for better context support and type safety.
func (c *Client) AcceptDomainInvite(inviteID interface{}) (interface{}, error) {
	return c.Call("accept_domain_invite", []interface{}{inviteID})
}

// GetMe retrieves the profile information of the current authenticated user.
// Deprecated: Use GetMeWithContext instead for better context support and type safety.
func (c *Client) GetMe() (interface{}, error) {
	return c.Call("get_me", []interface{}{})
}
