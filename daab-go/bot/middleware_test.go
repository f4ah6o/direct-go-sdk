package bot

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"
	"testing"
	"time"

	"pgregory.net/rapid"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

// TestMiddleware tests the basic middleware functionality.
func TestMiddleware(t *testing.T) {
	robot := New()

	var callOrder []string
	var mu sync.Mutex

	// Add middlewares that record their execution
	robot.Use(func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			callOrder = append(callOrder, "middleware1")
			mu.Unlock()
			next(ctx, res)
		}
	})

	robot.Use(func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			callOrder = append(callOrder, "middleware2")
			mu.Unlock()
			next(ctx, res)
		}
	})

	// Register a handler
	robot.Hear("test", func(ctx context.Context, res Response) {
		mu.Lock()
		callOrder = append(callOrder, "handler")
		mu.Unlock()
	})

	// Simulate a message
	msg := direct.ReceivedMessage{Text: "test"}
	robot.handleMessage(context.Background(), msg)
	time.Sleep(50 * time.Millisecond)

	mu.Lock()
	order := callOrder
	mu.Unlock()

	// Middleware should execute in order: middleware1 -> middleware2 -> handler
	if len(order) != 3 {
		t.Fatalf("Expected 3 calls, got %d: %v", len(order), order)
	}

	if order[0] != "middleware1" {
		t.Errorf("Expected first call to be middleware1, got %s", order[0])
	}
	if order[1] != "middleware2" {
		t.Errorf("Expected second call to be middleware2, got %s", order[1])
	}
	if order[2] != "handler" {
		t.Errorf("Expected third call to be handler, got %s", order[2])
	}
}

// TestRecoveryMiddleware tests that recovery middleware catches panics.
func TestRecoveryMiddleware(t *testing.T) {
	tw := &testWriter{}
	logger := log.New(tw, "", 0)

	robot := New()

	var handlerCalled bool
	var handlerMu sync.Mutex
	done := make(chan struct{})
	// Add wrapper to signal completion after everything (including panic recovery)
	// This must be added BEFORE RecoveryMiddleware so RecoveryMiddleware's defer runs first
	robot.Use(func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			defer close(done)
			next(ctx, res)
		}
	})
	robot.Use(RecoveryMiddleware(logger))

	robot.Hear("panic", func(ctx context.Context, res Response) {
		handlerMu.Lock()
		handlerCalled = true
		handlerMu.Unlock()
		panic("test panic")
	})

	// Simulate a message
	msg := direct.ReceivedMessage{Text: "panic"}
	robot.handleMessage(context.Background(), msg)
	<-done

	handlerMu.Lock()
	called := handlerCalled
	handlerMu.Unlock()
	if !called {
		t.Error("Handler was not called")
	}

	logOutput := tw.String()
	if !strings.Contains(logOutput, "RECOVER") {
		t.Errorf("Expected recovery log, got: %s", logOutput)
	}
	if !strings.Contains(logOutput, "test panic") {
		t.Errorf("Expected panic message in log, got: %s", logOutput)
	}
}

// TestLoggingMiddleware tests the logging middleware.
func TestLoggingMiddleware(t *testing.T) {
	// Use thread-safe buffer wrapper
	tw := &testWriter{}
	logger := log.New(tw, "", 0)

	robot := New()
	robot.Use(LoggingMiddleware(logger))

	// Add a wrapper middleware that closes done after everything finishes
	done := make(chan struct{})
	robot.Use(func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			defer close(done)
			next(ctx, res)
		}
	})

	robot.Hear("test", func(ctx context.Context, res Response) {
		// Handler does nothing
	})

	msg := direct.ReceivedMessage{
		Text:   "test message",
		UserID: "user123",
		TalkID: "room456",
	}
	robot.handleMessage(context.Background(), msg)

	// Wait for handler and all middleware defers to complete
	<-done

	logStr := tw.String()

	if !strings.Contains(logStr, "[START]") {
		t.Errorf("Expected [START] in log, got: %s", logStr)
	}
	if !strings.Contains(logStr, "[END]") {
		t.Errorf("Expected [END] in log, got: %s", logStr)
	}
	if !strings.Contains(logStr, "test message") {
		t.Errorf("Expected message text in log, got: %s", logStr)
	}
	if !strings.Contains(logStr, "user123") {
		t.Errorf("Expected user ID in log, got: %s", logStr)
	}
}

// testWriter is an io.Writer that stores output in an atomic.Value.
type testWriter struct {
	mu     sync.Mutex
	output strings.Builder
}

func (w *testWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.output.Write(p)
}

func (w *testWriter) String() string {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.output.String()
}

// TestFilterMiddleware tests the filter middleware.
func TestFilterMiddleware(t *testing.T) {
	robot := New()

	// Filter to only allow messages from specific user
	robot.Use(FilterMiddleware(func(ctx context.Context, res Response) bool {
		return res.Message.UserID == "allowed_user"
	}))

	var handlerCalled bool
	var handlerMu sync.Mutex
	done := make(chan struct{})
	robot.Hear("test", func(ctx context.Context, res Response) {
		handlerMu.Lock()
		handlerCalled = true
		handlerMu.Unlock()
		close(done)
	})

	// Test with allowed user
	msg1 := direct.ReceivedMessage{
		Text:   "test",
		UserID: "allowed_user",
	}
	robot.handleMessage(context.Background(), msg1)
	<-done
	time.Sleep(10 * time.Millisecond)

	handlerMu.Lock()
	called := handlerCalled
	handlerMu.Unlock()
	if !called {
		t.Error("Handler should be called for allowed user")
	}

	// Test with blocked user
	handlerCalled = false
	done = make(chan struct{})
	robot.Hear("test2", func(ctx context.Context, res Response) {
		handlerMu.Lock()
		handlerCalled = true
		handlerMu.Unlock()
		close(done)
	})

	msg2 := direct.ReceivedMessage{
		Text:   "test",
		UserID: "blocked_user",
	}
	robot.handleMessage(context.Background(), msg2)
	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
	}

	handlerMu.Lock()
	called = handlerCalled
	handlerMu.Unlock()
	if called {
		t.Error("Handler should not be called for blocked user")
	}
}

// TestRateLimitMiddleware tests the rate limit middleware.
func TestRateLimitMiddleware(t *testing.T) {
	robot := New()
	cooldown := 100 * time.Millisecond
	robot.Use(RateLimitMiddleware(cooldown))

	var callCount int
	var mu sync.Mutex
	robot.Hear("test", func(ctx context.Context, res Response) {
		mu.Lock()
		callCount++
		mu.Unlock()
	})

	msg := direct.ReceivedMessage{
		Text:   "test",
		UserID: "user123",
	}

	// First call should succeed
	robot.handleMessage(context.Background(), msg)
	time.Sleep(20 * time.Millisecond)

	mu.Lock()
	count := callCount
	mu.Unlock()
	if count != 1 {
		t.Errorf("Expected 1 call, got %d", count)
	}

	// Immediate second call should be rate limited
	robot.handleMessage(context.Background(), msg)
	time.Sleep(20 * time.Millisecond)

	mu.Lock()
	count = callCount
	mu.Unlock()
	if count != 1 {
		t.Errorf("Expected 1 call (rate limited), got %d", count)
	}

	// Wait for cooldown to expire
	time.Sleep(100 * time.Millisecond)

	// Third call should succeed after cooldown
	robot.handleMessage(context.Background(), msg)
	time.Sleep(20 * time.Millisecond)

	mu.Lock()
	count = callCount
	mu.Unlock()
	if count != 2 {
		t.Errorf("Expected 2 calls (after cooldown), got %d", count)
	}
}

// TestContextMiddleware tests the context middleware.
func TestContextMiddleware(t *testing.T) {
	robot := New()

	robot.Use(ContextMiddleware("key1", "value1", "key2", 42))

	var capturedValue interface{}
	done := make(chan struct{})
	robot.Hear("test", func(ctx context.Context, res Response) {
		capturedValue = GetContextValue(ctx, "key1")
		close(done)
	})

	msg := direct.ReceivedMessage{Text: "test"}
	robot.handleMessage(context.Background(), msg)
	<-done

	if capturedValue != "value1" {
		t.Errorf("Expected 'value1', got %v", capturedValue)
	}
}

// TestChain tests chaining multiple middlewares.
func TestChain(t *testing.T) {
	var callOrder []string
	var mu sync.Mutex

	m1 := func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			callOrder = append(callOrder, "m1")
			mu.Unlock()
			next(ctx, res)
		}
	}

	m2 := func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			callOrder = append(callOrder, "m2")
			mu.Unlock()
			next(ctx, res)
		}
	}

	m3 := func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			callOrder = append(callOrder, "m3")
			mu.Unlock()
			next(ctx, res)
		}
	}

	chain := Chain(m1, m2, m3)

	handler := chain(func(ctx context.Context, res Response) {
		mu.Lock()
		callOrder = append(callOrder, "handler")
		mu.Unlock()
	})

	handler(context.Background(), Response{})

	mu.Lock()
	order := callOrder
	mu.Unlock()

	// Order should be m1 -> m2 -> m3 -> handler
	if len(order) != 4 {
		t.Fatalf("Expected 4 calls, got %d: %v", len(order), order)
	}

	if order[0] != "m1" || order[1] != "m2" || order[2] != "m3" || order[3] != "handler" {
		t.Errorf("Wrong order: %v", order)
	}
}

// TestMiddlewareNilSlice tests that middleware works with nil middlewares.
func TestMiddlewareNilSlice(t *testing.T) {
	robot := New()

	// Don't add any middlewares
	var handlerCalled bool
	var handlerMu sync.Mutex
	done := make(chan struct{})
	robot.Hear("test", func(ctx context.Context, res Response) {
		handlerMu.Lock()
		handlerCalled = true
		handlerMu.Unlock()
		close(done)
	})

	msg := direct.ReceivedMessage{Text: "test"}
	robot.handleMessage(context.Background(), msg)
	<-done

	handlerMu.Lock()
	called := handlerCalled
	handlerMu.Unlock()
	if !called {
		t.Error("Handler should be called even without middlewares")
	}
}

// TestMiddlewareStopChain tests that middleware can stop the chain.
func TestMiddlewareStopChain(t *testing.T) {
	robot := New()

	// Add middleware that stops the chain
	robot.Use(func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			// Don't call next(), stopping the chain
			return
		}
	})

	var handlerCalled bool
	var handlerMu sync.Mutex
	done := make(chan struct{})
	robot.Hear("test", func(ctx context.Context, res Response) {
		handlerMu.Lock()
		handlerCalled = true
		handlerMu.Unlock()
		close(done)
	})

	msg := direct.ReceivedMessage{Text: "test"}
	robot.handleMessage(context.Background(), msg)

	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
	}

	handlerMu.Lock()
	called := handlerCalled
	handlerMu.Unlock()
	if called {
		t.Error("Handler should not be called when middleware stops chain")
	}
}

// TestMultipleListenersWithMiddleware tests that middleware applies to all listeners.
func TestMultipleListenersWithMiddleware(t *testing.T) {
	robot := New()

	var mu sync.Mutex
	var calls []string

	robot.Use(func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			calls = append(calls, "middleware")
			mu.Unlock()
			next(ctx, res)
		}
	})

	robot.Hear("hello", func(ctx context.Context, res Response) {
		mu.Lock()
		calls = append(calls, "hello")
		mu.Unlock()
	})

	robot.Hear("bye", func(ctx context.Context, res Response) {
		mu.Lock()
		calls = append(calls, "bye")
		mu.Unlock()
	})

	// Send two messages
	msg1 := direct.ReceivedMessage{Text: "hello"}
	robot.handleMessage(context.Background(), msg1)
	time.Sleep(50 * time.Millisecond)

	msg2 := direct.ReceivedMessage{Text: "bye"}
	robot.handleMessage(context.Background(), msg2)
	time.Sleep(50 * time.Millisecond)

	mu.Lock()
	callsList := calls
	mu.Unlock()

	// Should have: middleware, hello, middleware, bye
	if len(callsList) != 4 {
		t.Fatalf("Expected 4 calls, got %d: %v", len(callsList), callsList)
	}
}

// Property-Based Tests for Middleware

// TestMiddleware_ChainPreservesHandler verifies that chaining middlewares preserves handler execution.
func TestMiddleware_ChainPreservesHandler(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		robot := New()

		// Generate random number of middlewares
		numMiddleware := rapid.IntRange(0, 10).Draw(t, "numMiddleware")

		var callCount int
		var mu sync.Mutex

		// Add middlewares that just pass through
		for i := 0; i < numMiddleware; i++ {
			robot.Use(func(next Handler) Handler {
				return func(ctx context.Context, res Response) {
					next(ctx, res)
				}
			})
		}

		robot.Hear("test", func(ctx context.Context, res Response) {
			mu.Lock()
			callCount++
			mu.Unlock()
		})

		msg := direct.ReceivedMessage{Text: "test"}
		robot.handleMessage(context.Background(), msg)
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		count := callCount
		mu.Unlock()

		if count != 1 {
			t.Fatalf("With %d middlewares, expected handler to be called once, got %d", numMiddleware, count)
		}
	})
}

// TestMiddleware_OrderIsPreserved verifies that middleware execution order is preserved.
func TestMiddleware_OrderIsPreserved(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		robot := New()

		// Generate random number of middlewares
		numMiddleware := rapid.IntRange(1, 5).Draw(t, "numMiddleware")

		var order []int
		var mu sync.Mutex

		// Add middlewares that record their index
		for i := 0; i < numMiddleware; i++ {
			idx := i
			robot.Use(func(next Handler) Handler {
				return func(ctx context.Context, res Response) {
					mu.Lock()
					order = append(order, idx)
					mu.Unlock()
					next(ctx, res)
				}
			})
		}

		robot.Hear("test", func(ctx context.Context, res Response) {
			mu.Lock()
			order = append(order, -1) // -1 marks the handler
			mu.Unlock()
		})

		msg := direct.ReceivedMessage{Text: "test"}
		robot.handleMessage(context.Background(), msg)
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		executedOrder := order
		mu.Unlock()

		// Check that order is sequential: 0, 1, 2, ..., n-1, -1
		if len(executedOrder) != numMiddleware+1 {
			t.Fatalf("Expected %d calls, got %d", numMiddleware+1, len(executedOrder))
		}

		for i, idx := range executedOrder {
			expected := i
			if i == numMiddleware {
				expected = -1 // Handler marker
			}
			if idx != expected {
				t.Fatalf("Order mismatch at position %d: expected %d, got %v", i, expected, executedOrder)
			}
		}
	})
}

// TestRecoveryMiddleware_DoesNotCrash verifies recovery middleware catches various panic types.
func TestRecoveryMiddleware_DoesNotCrash(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		tw := &testWriter{}
		logger := log.New(tw, "", 0)

		robot := New()
		robot.Use(RecoveryMiddleware(logger))

		// Generate different panic types (0-3)
		panicType := rapid.IntRange(0, 3).Draw(t, "panicType")

		done := make(chan struct{})
		robot.Hear("panic", func(ctx context.Context, res Response) {
			defer close(done)
			switch panicType {
			case 0:
				panic("string panic")
			case 1:
				panic(42)
			case 2:
				panic(true)
			case 3:
				panic(struct{ name string }{"test"})
			}
		})

		msg := direct.ReceivedMessage{Text: "panic"}
		robot.handleMessage(context.Background(), msg)
		<-done
		time.Sleep(10 * time.Millisecond)

		// Bot should still be alive (no actual Go panic crashed the test)
		logOutput := tw.String()
		if !strings.Contains(logOutput, "[RECOVER]") {
			t.Fatalf("Expected recovery log for panicType %d", panicType)
		}
	})
}

// TestFilterMiddleware_SelectiveExecution verifies filter middleware selectively executes.
func TestFilterMiddleware_SelectiveExecution(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		robot := New()

		// Generate a random threshold
		threshold := rapid.IntRange(0, 100).Draw(t, "threshold")

		robot.Use(FilterMiddleware(func(ctx context.Context, res Response) bool {
			// Filter based on message length
			return len(res.Message.Text) > threshold
		}))

		var handlerCalled bool
		var mu sync.Mutex
		// Use ".*" pattern to match any message text including empty
		robot.Hear(".*", func(ctx context.Context, res Response) {
			mu.Lock()
			handlerCalled = true
			mu.Unlock()
		})

		// Generate a message with random length
		msgLen := rapid.IntRange(0, 200).Draw(t, "msgLen")
		msgText := strings.Repeat("a", msgLen)

		msg := direct.ReceivedMessage{Text: msgText}
		robot.handleMessage(context.Background(), msg)
		time.Sleep(100 * time.Millisecond)

		mu.Lock()
		called := handlerCalled
		mu.Unlock()

		shouldCall := len(msgText) > threshold
		if called != shouldCall {
			t.Fatalf("Filter: threshold=%d, msgLen=%d, expected handlerCalled=%v, got %v",
				threshold, msgLen, shouldCall, called)
		}
	})
}

// TestContextMiddleware_ValuesAreSet verifies context middleware sets values correctly.
func TestContextMiddleware_ValuesAreSet(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Use unique pattern for each iteration to avoid listener accumulation
		pattern := fmt.Sprintf("test_rapid_%d", rapid.Int64().Draw(t, "seed"))

		robot := New()

		// Generate random key-value pairs with unique keys
		numPairs := rapid.IntRange(1, 5).Draw(t, "numPairs")
		type kv struct{ key, value string }
		kvPairs := make([]kv, 0, numPairs)
		seenKeys := make(map[string]bool)

		for i := 0; i < numPairs*2; i++ {
			key := rapid.StringMatching(`[a-z]+`).Draw(t, "key")
			if !seenKeys[key] {
				seenKeys[key] = true
				value := rapid.String().Draw(t, "value")
				kvPairs = append(kvPairs, kv{key, value})
				if len(kvPairs) >= numPairs {
					break
				}
			}
		}

		// Build key-value args
		args := make([]interface{}, 0, len(kvPairs)*2)
		for _, pair := range kvPairs {
			args = append(args, pair.key, pair.value)
		}

		robot.Use(ContextMiddleware(args...))

		// Check that values are retrievable
		capturedValues := make(map[string]interface{})
		var capturedValuesMu sync.Mutex
		done := make(chan struct{})
		robot.Hear(pattern, func(ctx context.Context, res Response) {
			capturedValuesMu.Lock()
			defer capturedValuesMu.Unlock()
			for _, pair := range kvPairs {
				capturedValues[pair.key] = GetContextValue(ctx, pair.key)
			}
			close(done)
		})

		msg := direct.ReceivedMessage{Text: pattern}
		robot.handleMessage(context.Background(), msg)
		<-done // Wait for handler to complete

		// Verify all values were set correctly
		capturedValuesMu.Lock()
		defer capturedValuesMu.Unlock()
		for _, pair := range kvPairs {
			if capturedValues[pair.key] != pair.value {
				t.Fatalf("Context key %s: expected %v, got %v", pair.key, pair.value, capturedValues[pair.key])
			}
		}
	})
}

// TestChain_EmptySlice verifies empty middleware chain works.
func TestChain_EmptySlice(t *testing.T) {
	var handlerCalled bool
	var handlerMu sync.Mutex
	done := make(chan struct{})
	handler := func(ctx context.Context, res Response) {
		handlerMu.Lock()
		handlerCalled = true
		handlerMu.Unlock()
		close(done)
	}

	// Empty chain
	chain := Chain()
	wrapped := chain(handler)

	wrapped(context.Background(), Response{})
	<-done

	handlerMu.Lock()
	called := handlerCalled
	handlerMu.Unlock()
	if !called {
		t.Error("Handler should be called with empty middleware chain")
	}
}

// TestChain_SingleMiddleware verifies single middleware in chain works.
func TestChain_SingleMiddleware(t *testing.T) {
	var middlewareCalled bool
	var handlerCalled bool
	var mu sync.Mutex
	done := make(chan struct{})

	middleware := func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			mu.Lock()
			middlewareCalled = true
			mu.Unlock()
			next(ctx, res)
		}
	}

	handler := func(ctx context.Context, res Response) {
		mu.Lock()
		handlerCalled = true
		mu.Unlock()
		close(done)
	}

	chain := Chain(middleware)
	wrapped := chain(handler)

	wrapped(context.Background(), Response{})
	<-done

	mu.Lock()
	midCalled := middlewareCalled
	handCalled := handlerCalled
	mu.Unlock()

	if !midCalled {
		t.Error("Middleware should be called")
	}
	if !handCalled {
		t.Error("Handler should be called")
	}
}
