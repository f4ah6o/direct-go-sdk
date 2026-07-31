package direct

import (
	"context"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
)

const instrumentationName = "github.com/f4ah6o/direct-go-sdk/direct-go"

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

// MessageMetrics is an optional extension to Metrics for message-delivery
// observability. It is intentionally separate from Metrics so existing custom
// Metrics implementations remain source-compatible.
type MessageMetrics interface {
	// RecordMessageDrop records a message that could not be delivered because
	// the connection was shut down while applying channel backpressure.
	RecordMessageDrop(reason string)
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

// RecordMessageDrop does nothing.
func (m *NoopMetrics) RecordMessageDrop(reason string) {}

// OpenTelemetryMetrics records Direct RPC and connection metrics using
// OpenTelemetry. It uses the global MeterProvider by default, so applications
// only need to configure their OpenTelemetry exporter/provider once.
type OpenTelemetryMetrics struct {
	requestDuration       metric.Float64Histogram
	requestCount          metric.Int64Counter
	requestErrorCount     metric.Int64Counter
	connectionStateChange metric.Int64Counter
	messageDropCount      metric.Int64Counter
}

// OpenTelemetryMetricsOption configures OpenTelemetryMetrics.
type OpenTelemetryMetricsOption func(*openTelemetryMetricsConfig)

type openTelemetryMetricsConfig struct {
	meterProvider metric.MeterProvider
}

// WithMeterProvider configures the OpenTelemetry MeterProvider used by
// OpenTelemetryMetrics. If unset, the global OpenTelemetry MeterProvider is used.
func WithMeterProvider(provider metric.MeterProvider) OpenTelemetryMetricsOption {
	return func(cfg *openTelemetryMetricsConfig) {
		cfg.meterProvider = provider
	}
}

// NewOpenTelemetryMetrics creates a Metrics implementation backed by
// OpenTelemetry instruments.
func NewOpenTelemetryMetrics(opts ...OpenTelemetryMetricsOption) (*OpenTelemetryMetrics, error) {
	cfg := openTelemetryMetricsConfig{meterProvider: otel.GetMeterProvider()}
	for _, opt := range opts {
		opt(&cfg)
	}
	if cfg.meterProvider == nil {
		cfg.meterProvider = otel.GetMeterProvider()
	}

	meter := cfg.meterProvider.Meter(instrumentationName)

	requestDuration, err := meter.Float64Histogram(
		"direct.rpc.request.duration",
		metric.WithDescription("Duration of Direct RPC requests"),
		metric.WithUnit("ms"),
	)
	if err != nil {
		return nil, err
	}
	requestCount, err := meter.Int64Counter(
		"direct.rpc.request.count",
		metric.WithDescription("Number of completed Direct RPC requests"),
		metric.WithUnit("{request}"),
	)
	if err != nil {
		return nil, err
	}
	requestErrorCount, err := meter.Int64Counter(
		"direct.rpc.request.error.count",
		metric.WithDescription("Number of failed Direct RPC requests"),
		metric.WithUnit("{error}"),
	)
	if err != nil {
		return nil, err
	}
	connectionStateChange, err := meter.Int64Counter(
		"direct.websocket.connection.state_change.count",
		metric.WithDescription("Number of Direct WebSocket connection state changes"),
		metric.WithUnit("{change}"),
	)
	if err != nil {
		return nil, err
	}
	messageDropCount, err := meter.Int64Counter(
		"direct.message.drop.count",
		metric.WithDescription("Number of incoming messages canceled during delivery"),
		metric.WithUnit("{message}"),
	)
	if err != nil {
		return nil, err
	}

	return &OpenTelemetryMetrics{
		requestDuration:       requestDuration,
		requestCount:          requestCount,
		requestErrorCount:     requestErrorCount,
		connectionStateChange: connectionStateChange,
		messageDropCount:      messageDropCount,
	}, nil
}

// RecordRequest records the duration and count of a successful RPC request.
func (m *OpenTelemetryMetrics) RecordRequest(method string, duration time.Duration) {
	attrs := metric.WithAttributes(attribute.String("rpc.method", method))
	ctx := context.Background()
	m.requestDuration.Record(ctx, float64(duration)/float64(time.Millisecond), attrs)
	m.requestCount.Add(ctx, 1, attrs)
}

// RecordError records an RPC error.
func (m *OpenTelemetryMetrics) RecordError(method string, err error) {
	attrs := []attribute.KeyValue{attribute.String("rpc.method", method)}
	if err != nil {
		attrs = append(attrs, attribute.String("error.type", err.Error()))
	}
	m.requestErrorCount.Add(context.Background(), 1, metric.WithAttributes(attrs...))
}

// RecordConnectionState records a WebSocket connection state change.
func (m *OpenTelemetryMetrics) RecordConnectionState(state string) {
	m.connectionStateChange.Add(context.Background(), 1, metric.WithAttributes(attribute.String("connection.state", state)))
}

// RecordMessageDrop records an incoming message canceled during connection shutdown.
func (m *OpenTelemetryMetrics) RecordMessageDrop(reason string) {
	m.messageDropCount.Add(context.Background(), 1, metric.WithAttributes(attribute.String("reason", reason)))
}

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
