// Package testutil provides testing utilities for direct-go.
package testutil

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/vmihailenco/msgpack/v5"
)

const (
	// MessagePack RPC message types
	RpcRequest  = 0
	RpcResponse = 1
)

// RPCHandler is a function that handles an RPC method call.
type RPCHandler func(params []interface{}) (interface{}, error)

// MockServer is a mock WebSocket server for testing.
type MockServer struct {
	server    *httptest.Server
	upgrader  websocket.Upgrader
	handlers  map[string]RPCHandler
	mu        sync.RWMutex
	conn      *websocket.Conn
	connMu    sync.Mutex
	messages  [][]interface{} // Stores received RPC requests for assertions
	messagesMu sync.Mutex
}

// NewMockServer creates a new mock WebSocket server.
func NewMockServer() *MockServer {
	ms := &MockServer{
		handlers: make(map[string]RPCHandler),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true // Allow all origins in tests
			},
		},
		messages: make([][]interface{}, 0),
	}

	ms.server = httptest.NewServer(http.HandlerFunc(ms.handleWebSocket))

	return ms
}

// URL returns the WebSocket URL of the mock server.
func (ms *MockServer) URL() string {
	return "ws" + ms.server.URL[4:] // Replace "http" with "ws"
}

// Close stops the mock server.
func (ms *MockServer) Close() {
	ms.server.Close()
	if ms.conn != nil {
		ms.conn.Close()
	}
}

// On registers a handler for the specified RPC method.
func (ms *MockServer) On(method string, handler RPCHandler) {
	ms.mu.Lock()
	defer ms.mu.Unlock()
	ms.handlers[method] = handler
}

// OnSimple registers a simple handler that returns a constant result.
func (ms *MockServer) OnSimple(method string, result interface{}) {
	ms.On(method, func(params []interface{}) (interface{}, error) {
		return result, nil
	})
}

// OnError registers a handler that returns an error.
func (ms *MockServer) OnError(method string, errMsg string) {
	ms.On(method, func(params []interface{}) (interface{}, error) {
		return nil, fmt.Errorf("%s", errMsg)
	})
}

// GetReceivedMessages returns all received RPC requests.
func (ms *MockServer) GetReceivedMessages() [][]interface{} {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	return ms.messages
}

// GetReceivedMethod returns the method name of a specific received message.
func (ms *MockServer) GetReceivedMethod(index int) string {
	messages := ms.GetReceivedMessages()
	if index < 0 || index >= len(messages) {
		return ""
	}
	if len(messages[index]) >= 3 {
		if method, ok := messages[index][2].(string); ok {
			return method
		}
	}
	return ""
}

// handleWebSocket handles the WebSocket connection.
func (ms *MockServer) handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := ms.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	ms.connMu.Lock()
	ms.conn = conn
	ms.connMu.Unlock()

	defer conn.Close()

	for {
		msgType, data, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if msgType != websocket.BinaryMessage {
			continue
		}

		// Decode MessagePack RPC message
		var message []interface{}
		if err := msgpack.Unmarshal(data, &message); err != nil {
			continue
		}

		// Store received message for assertions
		ms.messagesMu.Lock()
		ms.messages = append(ms.messages, message)
		ms.messagesMu.Unlock()

		// Handle RPC request: [type, msgId, method, params]
		if len(message) < 4 {
			continue
		}

		msgID, ok := message[1].(int64)
		if !ok {
			// Try other numeric types
			if id, ok := message[1].(int); ok {
				msgID = int64(id)
			} else {
				continue
			}
		}

		method, ok := message[2].(string)
		if !ok {
			continue
		}

		params, ok := message[3].([]interface{})
		if !ok {
			params = []interface{}{}
		}

		// Find and call handler
		ms.mu.RLock()
		handler := ms.handlers[method]
		ms.mu.RUnlock()

		var response []interface{}
		if handler != nil {
			result, err := handler(params)
			if err != nil {
				// Response with error: [1, msgId, error, nil]
				response = []interface{}{RpcResponse, msgID, map[string]string{"message": err.Error()}, nil}
			} else {
				// Response with result: [1, msgId, nil, result]
				response = []interface{}{RpcResponse, msgID, nil, result}
			}
		} else {
			// Method not found
			response = []interface{}{RpcResponse, msgID, map[string]string{"message": "method not found"}, nil}
		}

		// Send response
		responseData, err := msgpack.Marshal(response)
		if err != nil {
			continue
		}

		conn.WriteMessage(websocket.BinaryMessage, responseData)
	}
}

// SendNotification sends a notification to the client (simulates server push).
func (ms *MockServer) SendNotification(method string, params interface{}) error {
	ms.connMu.Lock()
	conn := ms.conn
	ms.connMu.Unlock()

	if conn == nil {
		return fmt.Errorf("no active connection")
	}

	// Notification: [0, 0, method, [params]]
	notification := []interface{}{RpcRequest, int64(0), method, []interface{}{params}}

	data, err := msgpack.Marshal(notification)
	if err != nil {
		return err
	}

	return conn.WriteMessage(websocket.BinaryMessage, data)
}

// Reset clears all received messages (useful for test isolation).
func (ms *MockServer) Reset() {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	ms.messages = make([][]interface{}, 0)
}

// GetCallCount returns the number of times a specific method was called.
func (ms *MockServer) GetCallCount(method string) int {
	ms.messagesMu.Lock()
	defer ms.messagesMu.Unlock()
	
	count := 0
	for _, msg := range ms.messages {
		if len(msg) >= 3 {
			if m, ok := msg[2].(string); ok && m == method {
				count++
			}
		}
	}
	return count
}

// OnDynamic registers a handler that can respond based on parameters.
func (ms *MockServer) OnDynamic(method string, handler RPCHandler) {
	ms.On(method, handler)
}

