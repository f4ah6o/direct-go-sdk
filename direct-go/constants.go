package direct

import "time"

// WebSocket configuration constants
const (
	// DefaultPingInterval is the interval between WebSocket ping messages
	DefaultPingInterval = 45 * time.Second

	// DefaultRequestTimeout is the default timeout for RPC requests
	DefaultRequestTimeout = 30 * time.Second

	// DefaultHandshakeTimeout is the timeout for WebSocket handshake
	DefaultHandshakeTimeout = 10 * time.Second
)

// Buffer size constants
const (
	// DefaultMessageChannelSize is the buffer size for the Messages channel
	DefaultMessageChannelSize = 100

	// DefaultResultChannelSize is the buffer size for RPC result channels
	DefaultResultChannelSize = 1
)
