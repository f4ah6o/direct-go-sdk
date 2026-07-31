package direct

import (
	"errors"
	"fmt"
)

// Sentinel errors that can be checked with errors.Is().
var (
	// ErrNotConnected is returned when the client is not connected.
	ErrNotConnected = errors.New("direct: not connected")

	// ErrTimeout is returned when a request times out.
	ErrTimeout = errors.New("direct: request timeout")

	// ErrInvalidResponse is returned when the server returns an invalid response.
	ErrInvalidResponse = errors.New("direct: invalid response")

	// ErrAlreadyConnected is returned when trying to connect an already connected client.
	ErrAlreadyConnected = errors.New("direct: already connected")

	// ErrConnectionClosed is returned when an active RPC is interrupted by connection closure.
	ErrConnectionClosed = errors.New("direct: connection closed")

	// ErrRoomNotFound is returned when a room/talk cannot be found.
	ErrRoomNotFound = errors.New("direct: room not found")

	// ErrUserNotFound is returned when a user cannot be found.
	ErrUserNotFound = errors.New("direct: user not found")
)

// RPCError represents an error returned by the RPC server.
type RPCError struct {
	Code    int
	Message string
}

// Error returns the error message.
func (e *RPCError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("RPC error %d: %s", e.Code, e.Message)
	}
	return fmt.Sprintf("RPC error %d", e.Code)
}

// RPCError creates a new RPCError with the given code and message.
func NewRPCError(code int, message string) *RPCError {
	return &RPCError{Code: code, Message: message}
}

// ConnectionError represents a WebSocket connection error.
type ConnectionError struct {
	Cause error
}

// Error returns the error message.
func (e *ConnectionError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("connection failed: %v", e.Cause)
	}
	return "connection failed"
}

// Unwrap returns the underlying cause.
func (e *ConnectionError) Unwrap() error {
	return e.Cause
}

// NewConnectionError creates a new ConnectionError wrapping the given cause.
func NewConnectionError(cause error) *ConnectionError {
	return &ConnectionError{Cause: cause}
}
