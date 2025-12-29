package direct

import (
	"context"
	"time"
)

// Metrics defines the interface for recording observability metrics.
// Implementations can forward metrics to observability platforms like
// Prometheus, Datadog, or OpenTelemetry.
//
// The SDK provides a NoopMetrics implementation that discards all metrics.
// Users can set their own implementation via Client.SetMetrics().
type Metrics interface {
	// RecordRequest records the duration of an RPC request.
	RecordRequest(method string, duration time.Duration)

	// RecordError records an error that occurred during an RPC request.
	RecordError(method string, err error)

	// RecordConnectionState records a change in the WebSocket connection state.
	// Valid states are: "connecting", "connected", "disconnected", "reconnecting"
	RecordConnectionState(state string)
}

// NoopMetrics is a no-op implementation of Metrics that discards all metrics.
// Use this as the default or for testing.
type NoopMetrics struct{}

// RecordRequest does nothing.
func (m *NoopMetrics) RecordRequest(method string, duration time.Duration) {}

// RecordError does nothing.
func (m *NoopMetrics) RecordError(method string, err error) {}

// RecordConnectionState does nothing.
func (m *NoopMetrics) RecordConnectionState(state string) {}

// HealthChecker defines the interface for health status checks.
// Implementations can return detailed health status for liveness/readiness probes.
type HealthChecker interface {
	// Health checks if the client is healthy.
	// Returns nil if healthy, an error describing the issue otherwise.
	Health(ctx context.Context) error
}

// HealthStatus represents the current health status of the client.
type HealthStatus struct {
	// Connected is true if the WebSocket connection is active.
	Connected bool `json:"connected"`

	// Authenticated is true if a session has been created.
	Authenticated bool `json:"authenticated"`

	// Endpoint is the WebSocket API endpoint being used.
	Endpoint string `json:"endpoint"`

	// Error contains any error that caused the unhealthy status.
	Error string `json:"error,omitempty"`
}
