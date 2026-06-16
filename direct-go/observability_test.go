package direct

import (
	"errors"
	"testing"
	"time"
)

func TestNewOpenTelemetryMetrics(t *testing.T) {
	metrics, err := NewOpenTelemetryMetrics()
	if err != nil {
		t.Fatalf("NewOpenTelemetryMetrics returned error: %v", err)
	}

	metrics.RecordRequest(MethodGetMe, time.Millisecond)
	metrics.RecordError(MethodGetMe, errors.New("boom"))
	metrics.RecordConnectionState("connected")
}

func TestSetMetricsNilUsesNoop(t *testing.T) {
	client := NewClient(Options{})
	client.SetMetrics(nil)

	if _, ok := client.metrics.(*NoopMetrics); !ok {
		t.Fatalf("expected nil metrics to reset to *NoopMetrics, got %T", client.metrics)
	}
}
