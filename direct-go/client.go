// Package direct provides a Go client for the direct chat service.
//
// This package implements the MessagePack RPC protocol used by direct-js
// to communicate with the direct API server over WebSocket.
package direct

import (
	"context"
	"encoding/json"
	"errors"
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
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// EnableDebugServer enables sending logs to a debug WebSocket server.
// The debug server receives all WebSocket messages for debugging purposes.
//
// The url parameter should be a WebSocket server URL (e.g., "ws://localhost:8080").
// Use the direct-logserver tool to start a debug server.
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

// Protocol constants
const (
	// MessagePack RPC message types
	RpcRequest  = 0
	RpcResponse = 1

	// API version
	APIVersion = "1.128"

	// Default endpoint
	DefaultEndpoint = "wss://api.direct4b.com/albero-app-server/api"
)

// Options configures the direct client behavior.
// All fields are optional; unset fields will use their default values.
type Options struct {
	// Endpoint is the WebSocket API endpoint (e.g., "wss://api.direct4b.com/...").
	// If empty, DefaultEndpoint is used.
	Endpoint string

	// AccessToken is the authentication token for the API.
	// Can be retrieved via the Auth type or from the HUBOT_DIRECT_TOKEN environment variable.
	AccessToken string

	// ProxyURL is an optional HTTP proxy URL for WebSocket connections.
	// If set, all WebSocket traffic will be routed through this proxy.
	ProxyURL string

	// Host is the API host name.
	// If empty, it is derived from the Endpoint URL.
	Host string

	// Name is the bot name, used in log messages for debugging.
	Name string

	// MessageChannelSize controls the bounded Messages channel buffer.
	// If zero or negative, DefaultMessageChannelSize is used. Incoming message
	// delivery blocks when this buffer is full until the connection closes.
	MessageChannelSize int

	// WriteTimeout is the maximum duration allowed for a WebSocket write.
	// If zero or negative, DefaultWriteTimeout is used.
	WriteTimeout time.Duration

	// TracerProvider is an optional OpenTelemetry tracer provider used to create
	// spans for RPC calls. If nil, the global OpenTelemetry tracer provider is used.
	TracerProvider trace.TracerProvider
}

// ResponseHandler handles RPC responses for async requests.
// It is used internally by the call method for callback-based handling.
// For synchronous requests, use the Call method instead.
type ResponseHandler struct {
	// Method is the RPC method name that was called.
	Method string

	// OnSuccess is called when the RPC call succeeds.
	// The result parameter contains the unmarshaled response data.
	OnSuccess func(result interface{})

	// OnError is called when the RPC call fails.
	// The err parameter contains the error information from the server.
	OnError func(err interface{})
}

// Client is a direct API client.
//
// The lifecycle is reconnectable: after Close returns, or after a transport
// or session-authentication failure, Connect may be called again. Each
// successful Connect creates a new connection generation. Close clears the
// authenticated state and fails in-flight RPCs with ErrConnectionClosed.
// Messages remains usable across generations; Done is closed when the current
// generation ends and replaced by the next successful Connect.
type Client struct {
	options          Options
	conn             *websocket.Conn
	mu               sync.RWMutex
	handlers         map[string][]EventHandler
	responseHandlers map[int64]*ResponseHandler
	msgID            int64
	closed           bool
	connected        bool
	connDone         chan struct{}
	connReady        *connectionReadiness
	connWriteMu      *sync.Mutex
	connWriteTimeout time.Duration
	connecting       bool
	connectCancel    context.CancelFunc
	connectAttempt   uint64
	messageHandlers  []func(ReceivedMessage)

	// talkDomains maps talk_id to domain_id for user lookups
	talkDomains map[string]string

	// Channels for events. Messages remains usable across reconnects and is not
	// closed when one connection ends; Done belongs to the current connection
	// and is replaced on the next successful Connect.
	Messages chan ReceivedMessage
	Done     chan struct{}

	// Metrics records observability metrics. Use SetMetrics to set a custom implementation.
	metrics Metrics

	// tracer records OpenTelemetry spans for RPC calls.
	tracer trace.Tracer
}

// EventHandler is a callback for events.
type EventHandler func(data interface{})

// NewClient creates a new direct API client with the given options.
// The Options struct allows configuring the client's behavior.
//
// If no endpoint is provided in opts, DefaultEndpoint is used.
// If no host is provided, it is derived from the endpoint URL.
// The client must be connected via Connect() before calling any RPC methods.
//
// Example:
//
//	client := direct.NewClient(direct.Options{
//	    AccessToken: "your-token-here",
//	    ProxyURL:    "http://proxy.example.com:8080", // optional
//	})
//	err := client.Connect()
func NewClient(opts Options) *Client {
	if opts.Endpoint == "" {
		opts.Endpoint = DefaultEndpoint
	}
	if opts.Host == "" {
		if u, err := url.Parse(opts.Endpoint); err == nil {
			opts.Host = u.Host
		}
	}
	messageChannelSize := opts.MessageChannelSize
	if messageChannelSize <= 0 {
		messageChannelSize = DefaultMessageChannelSize
	}

	tracerProvider := opts.TracerProvider
	if tracerProvider == nil {
		tracerProvider = otel.GetTracerProvider()
	}

	return &Client{
		options:          opts,
		handlers:         make(map[string][]EventHandler),
		responseHandlers: make(map[int64]*ResponseHandler),
		talkDomains:      make(map[string]string),
		Messages:         make(chan ReceivedMessage, messageChannelSize),
		Done:             make(chan struct{}),
		metrics:          &NoopMetrics{},
		tracer:           tracerProvider.Tracer(instrumentationName),
	}
}

// Connect establishes a WebSocket connection to the direct API and returns
// after the WebSocket handshake completes. If an access token is provided,
// authentication and notification initialization continue asynchronously.
// Use ConnectWithContext or WaitReady when authenticated API calls must not
// begin until startup is complete.
// Returns ErrAlreadyConnected if the client is already connected.
// Returns a ConnectionError if the WebSocket connection fails.
// Connect may be called again after Close or a failed connection attempt.
func (c *Client) Connect() error {
	return c.connect(context.Background())
}

// ConnectWithContext establishes a WebSocket connection and waits until the
// client is authenticated and notification initialization is complete.
// When the context is canceled while dialing or waiting, the connection is
// closed and the context error is returned.
func (c *Client) ConnectWithContext(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return err
	}
	if err := c.connect(ctx); err != nil {
		if ctxErr := ctx.Err(); ctxErr != nil {
			return ctxErr
		}
		return err
	}
	if err := c.WaitReady(ctx); err != nil {
		_ = c.Close()
		return err
	}
	return nil
}

func (c *Client) connect(ctx context.Context) error {
	c.mu.Lock()
	if c.conn != nil || c.connecting {
		c.mu.Unlock()
		return ErrAlreadyConnected
	}
	connectCtx, cancel := context.WithCancel(ctx)
	c.connecting = true
	c.connectCancel = cancel
	c.connectAttempt++
	attempt := c.connectAttempt
	endpoint := c.options.Endpoint
	accessToken := c.options.AccessToken
	proxyURLString := c.options.ProxyURL
	writeTimeout := c.options.WriteTimeout
	if writeTimeout <= 0 {
		writeTimeout = DefaultWriteTimeout
	}
	c.mu.Unlock()

	dialer := websocket.Dialer{
		HandshakeTimeout: DefaultHandshakeTimeout,
	}

	// Set up proxy if configured
	if proxyURLString != "" {
		proxyURL, err := url.Parse(proxyURLString)
		if err != nil {
			c.mu.Lock()
			if c.connectAttempt == attempt {
				c.connecting = false
				c.connectCancel = nil
			}
			c.mu.Unlock()
			cancel()
			return fmt.Errorf("invalid proxy URL: %w", err)
		}
		dialer.Proxy = http.ProxyURL(proxyURL)
	}

	header := http.Header{}
	// Don't set Origin - let the library handle it
	// header.Set("Origin", "https://"+c.options.Host)

	// HTTP response from dialer is not needed after connection is established
	conn, _ /* *http.Response */, err := dialer.DialContext(connectCtx, endpoint, header)
	if err != nil {
		c.mu.Lock()
		if c.connectAttempt == attempt {
			c.connecting = false
			c.connectCancel = nil
		}
		c.mu.Unlock()
		cancel()
		return NewConnectionError(err)
	}
	cancel()

	c.mu.Lock()
	if c.connectAttempt != attempt || !c.connecting {
		c.mu.Unlock()
		_ = conn.Close()
		return ErrConnectionClosed
	}
	c.connecting = false
	c.connectCancel = nil
	c.conn = conn
	c.closed = false
	c.connected = false
	c.connDone = make(chan struct{})
	ready := newConnectionReadiness()
	c.connReady = ready
	c.connWriteMu = &sync.Mutex{}
	c.connWriteTimeout = writeTimeout
	c.Done = c.connDone
	done := c.connDone
	writeMu := c.connWriteMu
	metrics := c.metrics
	c.mu.Unlock()

	// Record connection
	metrics.RecordConnectionState("connected")

	// Set up pong handler
	conn.SetPongHandler(func(appData string) error {
		vlog("[DEBUG] Received pong: %s", appData)
		return nil
	})

	// Start message reader
	go c.readLoop(conn)

	// Start ping keepalive (every 45 seconds like direct-js)
	go c.pingLoop(conn, done, writeMu, writeTimeout)

	// Create session if access token is provided
	if accessToken != "" {
		go c.createSession(conn)
	} else {
		ready.complete(nil)
	}

	return nil
}

// pingLoop sends periodic pings to keep the connection alive
func (c *Client) pingLoop(conn *websocket.Conn, done <-chan struct{}, writeMu *sync.Mutex, writeTimeout time.Duration) {
	ticker := time.NewTicker(DefaultPingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if !c.isCurrentConnection(conn) {
				return
			}

			vlog("[DEBUG] Sending ping...")
			if err := writeMessage(conn, writeMu, writeTimeout, websocket.PingMessage, []byte("PING")); err != nil {
				vlog("[DEBUG] Ping error: %v", err)
				// Close this connection to trigger proper cleanup.
				_ = c.closeConnection(conn, newConnectionWriteError(err))
				return
			}
		case <-done:
			return
		}
	}
}

func writeMessage(conn *websocket.Conn, writeMu *sync.Mutex, timeout time.Duration, messageType int, data []byte) error {
	if conn == nil || writeMu == nil {
		return ErrConnectionClosed
	}
	if timeout <= 0 {
		timeout = DefaultWriteTimeout
	}

	writeMu.Lock()
	defer writeMu.Unlock()

	if err := conn.SetWriteDeadline(time.Now().Add(timeout)); err != nil {
		return err
	}
	err := conn.WriteMessage(messageType, data)
	clearErr := conn.SetWriteDeadline(time.Time{})
	if err != nil {
		return err
	}
	return clearErr
}

func writeCloseMessage(conn *websocket.Conn, writeMu *sync.Mutex, timeout time.Duration) error {
	if conn == nil || writeMu == nil {
		return ErrConnectionClosed
	}
	if timeout <= 0 {
		timeout = DefaultWriteTimeout
	}

	writeMu.Lock()
	defer writeMu.Unlock()

	deadline := time.Now().Add(timeout)
	if err := conn.SetWriteDeadline(deadline); err != nil {
		return err
	}
	err := conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), deadline)
	clearErr := conn.SetWriteDeadline(time.Time{})
	if err != nil {
		return err
	}
	return clearErr
}

type connectionWriteError struct {
	cause error
}

func (e *connectionWriteError) Error() string {
	return fmt.Sprintf("websocket write failed: %v", e.cause)
}

func (e *connectionWriteError) Unwrap() error {
	return errors.Join(ErrConnectionClosed, e.cause)
}

func newConnectionWriteError(err error) error {
	return &connectionWriteError{cause: err}
}

func isConnectionWriteError(err error) bool {
	var writeErr *connectionWriteError
	return errors.As(err, &writeErr)
}

type connectionReadiness struct {
	done chan struct{}
	err  error
	once sync.Once
}

func newConnectionReadiness() *connectionReadiness {
	return &connectionReadiness{done: make(chan struct{})}
}

func (r *connectionReadiness) complete(err error) {
	r.once.Do(func() {
		r.err = err
		close(r.done)
	})
}

// WaitReady waits until the current connection is authenticated and its
// notification initialization has completed. Clients without an access token
// are ready as soon as the WebSocket connection is established.
//
// If the connection closes before readiness, the close or startup error is
// returned. A canceled context returns its context error and leaves the
// connection open for other waiters.
func (c *Client) WaitReady(ctx context.Context) error {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return err
	}

	c.mu.RLock()
	ready := c.connReady
	connected := c.conn != nil && !c.closed
	closed := c.closed
	c.mu.RUnlock()
	if ready == nil || !connected {
		if closed {
			return ErrConnectionClosed
		}
		return ErrNotConnected
	}

	select {
	case <-ready.done:
		return ready.err
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (c *Client) completeReady(conn *websocket.Conn, err error) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != conn || c.closed || c.connReady == nil {
		return false
	}
	c.connReady.complete(err)
	return true
}

func (c *Client) failReady(conn *websocket.Conn, err error) {
	if err == nil {
		err = ErrConnectionClosed
	}
	if !c.completeReady(conn, err) {
		return
	}
	_ = c.closeConnection(conn, err)
}

func (c *Client) isCurrentConnection(conn *websocket.Conn) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.conn == conn && !c.closed
}

// createSession authenticates with the server.
func (c *Client) createSession(conn *websocket.Conn) {
	dlog("[DEBUG] Creating session with token: ***...")
	osString := DefaultBotOS
	params := []interface{}{c.options.AccessToken, APIVersion, osString}

	c.callOnConnection(conn, "create_session", params, func(result interface{}) {
		if !c.isCurrentConnection(conn) {
			return
		}
		dlog("[DEBUG] Session created successfully: %+v", result)
		c.mu.Lock()
		if c.conn != conn || c.closed {
			c.mu.Unlock()
			return
		}
		c.connected = true
		c.mu.Unlock()

		c.emit("session_created", result)

		// Start notification after session is created
		c.startNotification(conn)
	}, func(err interface{}) {
		if !c.isCurrentConnection(conn) {
			return
		}
		dlog("[DEBUG] Session error: %+v", err)
		c.emit("session_error", err)
		c.failReady(conn, fmt.Errorf("direct: session authentication failed: %v", err))
	})
}

// startNotification tells the server to start sending notifications.
func (c *Client) startNotification(conn *websocket.Conn) {
	dlog("[DEBUG] Starting notification...")

	// First, get domains to initialize data
	c.callOnConnection(conn, "get_domains", []interface{}{}, func(result interface{}) {
		if !c.isCurrentConnection(conn) {
			return
		}
		dlog("[DEBUG] get_domains success: %d domains", countItems(result))

		// Then get talks
		c.callOnConnection(conn, "get_talks", []interface{}{}, func(result interface{}) {
			if !c.isCurrentConnection(conn) {
				return
			}
			dlog("[DEBUG] get_talks success: %d talks", countItems(result))

			// Cache talk->domain mappings needed by notification parsing.
			if talks, ok := result.([]interface{}); ok {
				for _, talk := range talks {
					if talkMap, ok := talk.(map[string]interface{}); ok {
						// Cache talk_id -> domain_id mapping
						var talkID, domainID string
						if id, ok := talkMap["talk_id"]; ok {
							talkID = fmt.Sprintf("%v", id)
						} else if id, ok := talkMap["id"]; ok {
							talkID = fmt.Sprintf("%v", id)
						}
						if domID, ok := talkMap["domain_id"]; ok {
							domainID = fmt.Sprintf("%v", domID)
						}
						if talkID != "" && domainID != "" {
							c.mu.Lock()
							if c.conn == conn && !c.closed {
								c.talkDomains[talkID] = domainID
							}
							c.mu.Unlock()
							dlog("[DEBUG] Cached talk->domain: %s -> %s", talkID, domainID)
						}
					}
				}
			} else {
				dlog("[DEBUG] get_talks result is not []interface{}, type=%T", result)
			}

			// Then get talk statuses
			c.callOnConnection(conn, "get_talk_statuses", []interface{}{}, func(result interface{}) {
				if !c.isCurrentConnection(conn) {
					return
				}
				dlog("[DEBUG] get_talk_statuses success")

				// Try start_notification first
				c.callOnConnection(conn, "start_notification", []interface{}{}, func(result interface{}) {
					if !c.isCurrentConnection(conn) {
						return
					}
					dlog("[DEBUG] start_notification result: %+v", result)

					// If false, try reset_notification and then start_notification again
					if result == false {
						dlog("[DEBUG] start_notification returned false, trying reset_notification...")
						c.callOnConnection(conn, "reset_notification", []interface{}{}, func(result interface{}) {
							if !c.isCurrentConnection(conn) {
								return
							}
							dlog("[DEBUG] reset_notification result: %+v", result)

							// After reset, call start_notification again
							c.callOnConnection(conn, "start_notification", []interface{}{}, func(result interface{}) {
								if !c.isCurrentConnection(conn) {
									return
								}
								dlog("[DEBUG] start_notification (after reset) result: %+v", result)

								// Call update_last_used_at to mark session as active
								c.callOnConnection(conn, "update_last_used_at", []interface{}{}, func(result interface{}) {
									if !c.isCurrentConnection(conn) {
										return
									}
									dlog("[DEBUG] update_last_used_at result: %+v", result)
									c.completeReady(conn, nil)
									c.emit("data_recovered", result)
								}, func(err interface{}) {
									if !c.isCurrentConnection(conn) {
										return
									}
									dlog("[DEBUG] update_last_used_at error: %+v", err)
									c.completeReady(conn, nil)
									c.emit("data_recovered", nil)
								})
							}, func(err interface{}) {
								if !c.isCurrentConnection(conn) {
									return
								}
								dlog("[DEBUG] start_notification (after reset) error: %+v", err)
								c.emit("notification_error", err)
								c.failReady(conn, fmt.Errorf("direct: notification initialization failed: %v", err))
							})
						}, func(err interface{}) {
							if !c.isCurrentConnection(conn) {
								return
							}
							dlog("[DEBUG] reset_notification error: %+v", err)
							c.emit("notification_error", err)
							c.failReady(conn, fmt.Errorf("direct: notification initialization failed: %v", err))
						})
					} else {
						c.completeReady(conn, nil)
						c.emit("data_recovered", result)
					}
				}, func(err interface{}) {
					if !c.isCurrentConnection(conn) {
						return
					}
					dlog("[DEBUG] start_notification error: %+v", err)
					c.emit("notification_error", err)
					c.failReady(conn, fmt.Errorf("direct: notification initialization failed: %v", err))
				})
			}, func(err interface{}) {
				if !c.isCurrentConnection(conn) {
					return
				}
				dlog("[DEBUG] get_talk_statuses error: %+v", err)
				c.emit("notification_error", err)
				c.failReady(conn, fmt.Errorf("direct: notification initialization failed: %v", err))
			})
		}, func(err interface{}) {
			if !c.isCurrentConnection(conn) {
				return
			}
			dlog("[DEBUG] get_talks error: %+v", err)
			c.emit("notification_error", err)
			c.failReady(conn, fmt.Errorf("direct: notification initialization failed: %v", err))
		})
	}, func(err interface{}) {
		if !c.isCurrentConnection(conn) {
			return
		}
		dlog("[DEBUG] get_domains error: %+v", err)
		c.emit("notification_error", err)
		c.failReady(conn, fmt.Errorf("direct: notification initialization failed: %v", err))
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
// It is safe to call Close multiple times.
func (c *Client) Close() error {
	return c.closeConnection(nil, ErrConnectionClosed)
}

// closeConnection closes the specified active connection, or the current
// connection when conn is nil. A stale read/ping loop cannot close a newer
// connection because the pointer must still match the active connection.
func (c *Client) closeConnection(conn *websocket.Conn, cause error) error {
	if cause == nil {
		cause = ErrConnectionClosed
	}

	c.mu.Lock()
	if conn != nil && c.conn != conn {
		c.mu.Unlock()
		return nil
	}
	activeConn := c.conn
	if activeConn == nil && !c.connecting && c.closed {
		c.mu.Unlock()
		return nil
	}
	connectCancel := c.connectCancel
	wasActive := activeConn != nil || c.connecting
	writeMu := c.connWriteMu
	writeTimeout := c.connWriteTimeout
	ready := c.connReady
	c.connectAttempt++
	c.connecting = false
	c.connectCancel = nil
	c.conn = nil
	c.connDone = nil
	c.connReady = nil
	c.connWriteMu = nil
	c.connWriteTimeout = 0
	c.closed = true
	c.connected = false
	done := c.Done
	if done != nil {
		close(done)
	}
	pending := make([]*ResponseHandler, 0, len(c.responseHandlers))
	for _, handler := range c.responseHandlers {
		pending = append(pending, handler)
	}
	c.responseHandlers = make(map[int64]*ResponseHandler)
	c.talkDomains = make(map[string]string)
	if ready != nil {
		ready.complete(cause)
	}
	metrics := c.metrics
	c.mu.Unlock()

	if connectCancel != nil {
		connectCancel()
	}

	for _, handler := range pending {
		if handler != nil && handler.OnError != nil {
			handler.OnError(cause)
		}
	}

	if wasActive {
		metrics.RecordConnectionState("disconnected")
	}

	if activeConn != nil {
		var closeErr error
		if writeMu != nil {
			closeErr = writeCloseMessage(activeConn, writeMu, writeTimeout)
		}
		if err := activeConn.Close(); closeErr == nil {
			closeErr = err
		}
		if causeWrite := isConnectionWriteError(cause); causeWrite {
			c.emit("error", map[string]string{"error": cause.Error()})
		}
		return closeErr
	}
	return nil
}

// SetMetrics sets the metrics implementation for this client.
// Pass a custom Metrics implementation to record observability metrics.
// To disable metrics, pass &NoopMetrics{}.
func (c *Client) SetMetrics(m Metrics) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if m == nil {
		m = &NoopMetrics{}
	}
	c.metrics = m
}

// SetTracerProvider sets the OpenTelemetry tracer provider for this client.
// Passing nil resets the client to the global OpenTelemetry tracer provider.
func (c *Client) SetTracerProvider(provider trace.TracerProvider) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if provider == nil {
		provider = otel.GetTracerProvider()
	}
	c.tracer = provider.Tracer(instrumentationName)
}

// Health checks the health status of the client connection.
// Returns nil if the client is connected and healthy, an error otherwise.
//
// The returned HealthStatus contains detailed information about the client state:
//   - Connected: true if WebSocket is active
//   - Authenticated: true if a session has been created
//   - Endpoint: the API endpoint being used
//   - Error: any error that caused the unhealthy status
func (c *Client) Health() *HealthStatus {
	c.mu.RLock()
	defer c.mu.RUnlock()

	status := &HealthStatus{
		Connected:     c.conn != nil && !c.closed,
		Authenticated: c.connected,
		Endpoint:      c.options.Endpoint,
	}

	if !status.Connected {
		status.Error = "not connected"
	}

	return status
}

// On registers an event handler for the given event type.
// Multiple handlers can be registered for the same event.
// Event types are defined as constants (e.g., EventSessionCreated, EventNotifyCreateMessage).
// Handlers are called asynchronously in separate goroutines.
func (c *Client) On(event string, handler EventHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.handlers[event] = append(c.handlers[event], handler)
}

// OnMessage registers a callback for incoming messages.
// The handler runs in a separate goroutine for each message and remains
// registered across reconnects. Callback panics are recovered and written to
// the debug log; a panicking callback does not stop message delivery.
func (c *Client) OnMessage(handler func(ReceivedMessage)) {
	c.mu.Lock()
	c.messageHandlers = append(c.messageHandlers, handler)
	c.mu.Unlock()
}

// call sends an RPC request to the current connection.
func (c *Client) call(method string, params []interface{}, onSuccess func(interface{}), onError func(interface{})) int64 {
	return c.callOnConnection(nil, method, params, onSuccess, onError)
}

// callOnConnection sends an RPC request only if expected is still the active
// connection. This prevents callbacks from an older connection from issuing
// requests on a newer one after reconnect.
func (c *Client) callOnConnection(expected *websocket.Conn, method string, params []interface{}, onSuccess func(interface{}), onError func(interface{})) int64 {
	c.mu.Lock()

	conn := c.conn
	if conn == nil {
		c.mu.Unlock()
		if onError != nil {
			onError(map[string]string{"message": "not connected"})
		}
		return 0
	}
	if c.closed || (expected != nil && conn != expected) {
		c.mu.Unlock()
		if onError != nil {
			onError(ErrConnectionClosed)
		}
		return 0
	}
	writeMu := c.connWriteMu
	writeTimeout := c.connWriteTimeout

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
		c.mu.Lock()
		delete(c.responseHandlers, msgID)
		c.mu.Unlock()
		if onError != nil {
			onError(map[string]string{"message": err.Error()})
		}
		return 0
	}

	c.mu.Lock()
	validConnection := c.conn == conn && !c.closed && (expected == nil || c.conn == expected)
	c.mu.Unlock()
	if !validConnection {
		c.mu.Lock()
		delete(c.responseHandlers, msgID)
		c.mu.Unlock()
		if onError != nil {
			onError(ErrConnectionClosed)
		}
		return 0
	}
	err = writeMessage(conn, writeMu, writeTimeout, websocket.BinaryMessage, data)

	if err != nil {
		_ = c.closeConnection(conn, newConnectionWriteError(err))
	}

	return msgID
}

// Call sends a synchronous RPC request to the direct API server.
// It blocks until a response is received or the 30-second timeout expires.
// If an access token is configured, it also waits for authenticated readiness.
// Method names are defined as constants (e.g., MethodGetTalks, MethodCreateMessage).
// Returns the result on success, or an error on failure or timeout.
// Returns ErrNotConnected if the client is not connected.
// Returns ErrTimeout if the request times out.
//
// Deprecated: Use CallWithContext to control cancellation and deadlines.
func (c *Client) Call(method string, params []interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultRequestTimeout)
	defer cancel()

	result, err := c.CallWithContext(ctx, method, params)
	if errors.Is(err, context.DeadlineExceeded) {
		return nil, fmt.Errorf("RPC error: %w", ErrTimeout)
	}
	return result, err
}

// CallWithContext sends a synchronous RPC request using the supplied context.
// Cancellation and deadlines stop waiting for the response and remove the
// request handler so a late response cannot be delivered to a canceled call.
// If an access token is configured, it waits for authenticated readiness before
// sending the request.
// The span created for the RPC inherits the caller's context.
func (c *Client) CallWithContext(ctx context.Context, method string, params []interface{}) (interface{}, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	if err := c.waitForAuthenticatedReady(ctx); err != nil {
		return nil, err
	}

	start := time.Now()
	c.mu.RLock()
	tracer := c.tracer
	metrics := c.metrics
	c.mu.RUnlock()
	_, span := tracer.Start(ctx, "direct.rpc", trace.WithAttributes(attribute.String("rpc.system", "direct"), attribute.String("rpc.method", method)))
	defer span.End()
	resultCh := make(chan interface{}, DefaultResultChannelSize)
	errCh := make(chan interface{}, DefaultResultChannelSize)

	msgID := c.call(method, params, func(result interface{}) {
		resultCh <- result
	}, func(err interface{}) {
		errCh <- err
	})

	select {
	case result := <-resultCh:
		metrics.RecordRequest(method, time.Since(start))
		span.SetStatus(codes.Ok, "")
		return result, nil
	case err := <-errCh:
		var rpcErr error
		if errValue, ok := err.(error); ok {
			rpcErr = fmt.Errorf("RPC error: %w", errValue)
		} else {
			rpcErr = fmt.Errorf("RPC error: %v", err)
		}
		metrics.RecordError(method, rpcErr)
		span.RecordError(rpcErr)
		span.SetStatus(codes.Error, rpcErr.Error())
		return nil, rpcErr
	case <-ctx.Done():
		err := ctx.Err()
		if msgID != 0 {
			c.mu.Lock()
			delete(c.responseHandlers, msgID)
			c.mu.Unlock()
		}
		metrics.RecordError(method, err)
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
		return nil, err
	}
}

func (c *Client) waitForAuthenticatedReady(ctx context.Context) error {
	c.mu.RLock()
	needsReady := c.conn != nil && !c.closed && c.options.AccessToken != ""
	c.mu.RUnlock()
	if !needsReady {
		return nil
	}
	return c.WaitReady(ctx)
}

// Send sends a message with custom type and content to the specified room.
// roomID can be a string or numeric room/talk identifier.
// msgType should be one of the MessageType constants (e.g., MsgTypeText, MsgTypeStamp).
// content structure depends on the message type.
func (c *Client) Send(roomID interface{}, msgType int, content interface{}) error {
	_, err := c.Call("create_message", []interface{}{roomID, msgType, content})
	return err
}

// SendText sends a text message to the specified room.
// This is a convenience method that wraps Send with msgType=1 (text).
// Deprecated: Use SendTextWithContext for better context support.
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
func (c *Client) readLoop(conn *websocket.Conn) {
	for {
		if !c.isCurrentConnection(conn) {
			return
		}

		msgType, data, err := conn.ReadMessage()
		if err != nil {
			if c.isCurrentConnection(conn) {
				dlog("[DEBUG] ReadMessage error: %v", err)
				c.emit("error", map[string]string{"error": err.Error()})
				// Close only this connection. A newer connection may already be active.
				_ = c.closeConnection(conn, fmt.Errorf("%w: %v", ErrConnectionClosed, err))
			}
			return
		}

		dlog("[DEBUG] Raw WebSocket message: type=%d len=%d", msgType, len(data))

		c.handleMessage(conn, data)
	}
}

// handleMessage processes an incoming WebSocket message.
func (c *Client) handleMessage(conn *websocket.Conn, data []byte) {
	if !c.isCurrentConnection(conn) {
		return
	}

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
		c.handleNotification(conn, message)
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
func (c *Client) handleNotification(conn *websocket.Conn, message []interface{}) {
	if !c.isCurrentConnection(conn) {
		return
	}

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
	if !c.isCurrentConnection(conn) {
		return
	}

	// Emit the notification event
	c.emit(method, params[0])

	// Handle message notifications specially
	if method == "notify_create_message" || method == "create_message" {
		dlog("[DEBUG] Message notification received: %s", method)
		dlog("[DEBUG] Data: %+v", params[0])
		c.handleMessageNotification(conn, params[0])
	}

	// Send acknowledgment response: [1, msgId, null, true]
	response := []interface{}{RpcResponse, msgID, nil, true}
	data, err := msgpack.Marshal(response)
	if err != nil {
		dlog("[DEBUG] Could not encode notification acknowledgment: %v", err)
		return
	}
	writeMu, writeTimeout := c.connectionWriter(conn)
	if writeMu == nil {
		return
	}
	if err := writeMessage(conn, writeMu, writeTimeout, websocket.BinaryMessage, data); err != nil {
		dlog("[DEBUG] Notification acknowledgment write failed: %v", err)
		_ = c.closeConnection(conn, newConnectionWriteError(err))
	}
}

// handleMessageNotification parses and queues a message notification.
func (c *Client) handleMessageNotification(conn *websocket.Conn, data interface{}) {
	if !c.isCurrentConnection(conn) {
		return
	}

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
		c.mu.RLock()
		handlers := append([]func(ReceivedMessage){}, c.messageHandlers...)
		c.mu.RUnlock()
		for _, handler := range handlers {
			go c.runMessageHandler(handler, msg)
		}

		done := c.connectionDone(conn)
		if done == nil {
			c.recordMessageDrop("connection_closed")
			dlog("[DEBUG] Message delivery canceled: connection is no longer active")
			return
		}
		select {
		case c.Messages <- msg:
		case <-done:
			c.recordMessageDrop("connection_closed")
			dlog("[DEBUG] Message delivery canceled during connection shutdown")
		}
	}
}

func (c *Client) connectionDone(conn *websocket.Conn) <-chan struct{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.conn != conn || c.closed {
		return nil
	}
	return c.connDone
}

func (c *Client) connectionWriter(conn *websocket.Conn) (*sync.Mutex, time.Duration) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if c.conn != conn || c.closed {
		return nil, 0
	}
	return c.connWriteMu, c.connWriteTimeout
}

func (c *Client) recordMessageDrop(reason string) {
	c.mu.RLock()
	metrics := c.metrics
	c.mu.RUnlock()
	if messageMetrics, ok := metrics.(MessageMetrics); ok {
		messageMetrics.RecordMessageDrop(reason)
	}
}

func (c *Client) runMessageHandler(handler func(ReceivedMessage), msg ReceivedMessage) {
	defer func() {
		if recovered := recover(); recovered != nil {
			dlog("[DEBUG] OnMessage handler panicked: %v", recovered)
		}
	}()
	handler(msg)
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
	result, err := c.CallWithContext(ctx, MethodGetTalks, []interface{}{})
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
	result, err := c.CallWithContext(ctx, MethodGetTalkStatuses, []interface{}{})
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
	result, err := c.CallWithContext(ctx, MethodGetMe, []interface{}{})
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
// roomID is the talk/room identifier, and text is the message content.
// This is the preferred method over the legacy SendText().
func (c *Client) SendTextWithContext(ctx context.Context, roomID string, text string) error {
	_, err := c.CreateTextMessageWithContext(ctx, roomID, text)
	return err
}

// CreateTextMessageWithContext sends a text message and returns the created
// message ID when the API response includes it.
func (c *Client) CreateTextMessageWithContext(ctx context.Context, roomID string, text string) (string, error) {
	var talkID interface{} = roomID
	if id, err := strconv.ParseUint(roomID, 10, 64); err == nil {
		talkID = id
	}
	result, err := c.CallWithContext(ctx, MethodCreateMessage, []interface{}{talkID, 1, text})
	if err != nil {
		return "", err
	}
	if m, ok := result.(map[string]interface{}); ok {
		if id, ok := m["id"].(string); ok {
			return id, nil
		}
		if id, ok := m["id"]; ok {
			return fmt.Sprint(id), nil
		}
	}
	return "", nil
}

// Legacy methods below - deprecated, use context-aware versions instead

// GetTalks retrieves the list of talk rooms.
// Deprecated: Use GetTalksWithContext instead.
func (c *Client) GetTalks() (interface{}, error) {
	return c.Call("get_talks", []interface{}{})
}

// GetDomains retrieves the list of domains.
// Deprecated: Use GetDomainsWithContext instead.
func (c *Client) GetDomains() (interface{}, error) {
	return c.Call("get_domains", []interface{}{})
}

// GetDomainInvites retrieves pending domain invitations.
// Deprecated: Use GetDomainInvitesWithContext instead.
func (c *Client) GetDomainInvites() (interface{}, error) {
	return c.Call("get_domain_invites", []interface{}{})
}

// AcceptDomainInvite accepts a domain invitation.
// Deprecated: Use AcceptDomainInviteWithContext instead.
func (c *Client) AcceptDomainInvite(inviteID interface{}) (interface{}, error) {
	return c.Call("accept_domain_invite", []interface{}{inviteID})
}

// GetMe retrieves the current user's profile.
// Deprecated: Use GetMeWithContext instead.
func (c *Client) GetMe() (interface{}, error) {
	return c.Call("get_me", []interface{}{})
}
