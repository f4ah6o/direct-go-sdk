package direct

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"testing"
	"time"

	"go.opentelemetry.io/otel/trace"
	"pgregory.net/rapid"

	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
	"github.com/f4ah6o/direct-go-sdk/direct-go/testutil"
)

func TestClientConnect(t *testing.T) {
	// Create mock server
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Create client
	client := NewClient(Options{
		Endpoint:    mockServer.URL(),
		AccessToken: "test-token",
	})

	// Setup mock handlers for session creation
	mockServer.OnSimple("create_session", map[string]interface{}{
		"user_id": "test-user",
		"token":   "test-token",
	})
	mockServer.OnSimple("get_domains", []interface{}{})
	mockServer.OnSimple("get_talks", []interface{}{
		map[string]interface{}{
			"id":        "talk-1",
			"domain_id": "domain-1",
		},
	})
	mockServer.OnSimple("get_talk_statuses", []interface{}{})
	mockServer.OnSimple("start_notification", true)

	// Connect
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	// Wait until notification startup has reached the final initialization RPC.
	deadline := time.Now().Add(time.Second)
	for mockServer.GetCallCount("start_notification") == 0 && time.Now().Before(deadline) {
		time.Sleep(time.Millisecond)
	}
	if got := mockServer.GetCallCount("start_notification"); got == 0 {
		t.Fatal("start_notification was not called")
	}

	// Verify create_session was called
	messages := mockServer.GetReceivedMessages()
	if len(messages) == 0 {
		t.Fatal("No messages received by mock server")
	}

	foundCreateSession := false
	for i := 0; i < len(messages); i++ {
		method := mockServer.GetReceivedMethod(i)
		if method == "create_session" {
			foundCreateSession = true
			break
		}
	}

	if !foundCreateSession {
		t.Error("create_session was not called")
	}

	if got := mockServer.GetCallCount("create_message"); got != 0 {
		t.Fatalf("startup sent create_message %d times", got)
	}

	client.mu.RLock()
	domainID := client.talkDomains["talk-1"]
	client.mu.RUnlock()
	if domainID != "domain-1" {
		t.Fatalf("talk-domain cache = %q, want domain-1", domainID)
	}
}

func TestClientDiagnosticsRedactCredentialsAndMessages(t *testing.T) {
	const (
		accessToken = "access-token-secret"
		messageBody = "private message body"
	)

	for _, level := range []int{debuglog.LevelNormal, debuglog.LevelVerbose} {
		t.Run(fmt.Sprintf("level-%d", level), func(t *testing.T) {
			mockServer := testutil.NewMockServer()
			defer mockServer.Close()
			mockServer.OnSimple("create_session", map[string]interface{}{
				"user_id": "user-1",
				"token":   accessToken,
			})
			mockServer.OnSimple("get_domains", []interface{}{})
			mockServer.OnSimple("get_talks", []interface{}{})
			mockServer.OnSimple("get_talk_statuses", []interface{}{})
			mockServer.OnSimple("start_notification", true)

			var output synchronizedBuffer
			logger := debuglog.NewLogger(debuglog.LoggerOptions{Level: level, Writer: &output})
			client := NewClient(Options{
				Endpoint:    mockServer.URL(),
				AccessToken: accessToken,
				DebugLogger: logger,
			})
			if err := client.ConnectWithContext(context.Background()); err != nil {
				t.Fatalf("ConnectWithContext failed: %v", err)
			}
			defer client.Close()

			if err := mockServer.SendNotification("notify_create_message", map[string]interface{}{
				"message_id": "message-1",
				"talk_id":    "talk-1",
				"user_id":    "user-1",
				"content":    messageBody,
			}); err != nil {
				t.Fatalf("SendNotification failed: %v", err)
			}

			deadline := time.Now().Add(time.Second)
			for !strings.Contains(output.String(), "parsed msg") && time.Now().Before(deadline) {
				time.Sleep(time.Millisecond)
			}
			got := output.String()
			for _, unwanted := range []string{accessToken, messageBody} {
				if strings.Contains(got, unwanted) {
					t.Fatalf("client diagnostics leaked %q: %q", unwanted, got)
				}
			}
			if !strings.Contains(got, "access_token=configured") || !strings.Contains(got, "text=string(len=") {
				t.Fatalf("client diagnostics did not preserve safe metadata: %q", got)
			}
		})
	}
}

type synchronizedBuffer struct {
	mu  sync.Mutex
	buf bytes.Buffer
}

func (b *synchronizedBuffer) Write(p []byte) (int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.Write(p)
}

func (b *synchronizedBuffer) String() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.buf.String()
}

func configureSessionStartup(mockServer *testutil.MockServer) {
	mockServer.OnSimple("get_domains", []interface{}{})
	mockServer.OnSimple("get_talks", []interface{}{})
	mockServer.OnSimple("get_talk_statuses", []interface{}{})
	mockServer.OnSimple("start_notification", true)
}

func TestClientConnectWithContextWaitsForReady(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	createReceived := make(chan struct{})
	release := make(chan struct{})
	var releaseOnce sync.Once
	defer releaseOnce.Do(func() { close(release) })
	mockServer.On("create_session", func(params []interface{}) (interface{}, error) {
		close(createReceived)
		<-release
		return map[string]interface{}{"user_id": "user123"}, nil
	})
	mockServer.OnSimple("get_me", map[string]interface{}{"user_id": "user123"})
	configureSessionStartup(mockServer)

	client := NewClient(Options{Endpoint: mockServer.URL(), AccessToken: "token"})
	defer client.Close()
	readyErr := make(chan error, 1)
	go func() {
		readyErr <- client.ConnectWithContext(context.Background())
	}()
	select {
	case <-createReceived:
	case <-time.After(time.Second):
		t.Fatal("create_session was not called")
	}
	select {
	case err := <-readyErr:
		t.Fatalf("ConnectWithContext returned before authentication: %v", err)
	default:
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	_, err := client.GetMeWithContext(ctx)
	cancel()
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("API call before readiness error = %v, want context deadline", err)
	}
	if got := mockServer.GetCallCount("get_me"); got != 0 {
		t.Fatalf("get_me was sent before readiness: %d calls", got)
	}
	releaseOnce.Do(func() { close(release) })

	select {
	case err := <-readyErr:
		if err != nil {
			t.Fatalf("ConnectWithContext failed: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("ConnectWithContext did not wait for readiness")
	}
	if health := client.Health(); !health.Connected || !health.Authenticated {
		t.Fatalf("health after readiness = %+v", health)
	}
}

func TestClientWaitReadyUnblocksOnCancellationAndClose(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	createReceived := make(chan struct{})
	release := make(chan struct{})
	var releaseOnce sync.Once
	defer releaseOnce.Do(func() { close(release) })
	mockServer.On("create_session", func(params []interface{}) (interface{}, error) {
		close(createReceived)
		<-release
		return map[string]interface{}{"user_id": "user123"}, nil
	})
	configureSessionStartup(mockServer)

	client := NewClient(Options{Endpoint: mockServer.URL(), AccessToken: "token"})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()
	select {
	case <-createReceived:
	case <-time.After(time.Second):
		t.Fatal("create_session was not called")
	}

	ctx, cancel := context.WithCancel(context.Background())
	waitErr := make(chan error, 1)
	go func() { waitErr <- client.WaitReady(ctx) }()
	cancel()
	select {
	case err := <-waitErr:
		if !errors.Is(err, context.Canceled) {
			t.Fatalf("WaitReady cancellation error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("WaitReady did not unblock on cancellation")
	}

	closeErr := make(chan error, 1)
	go func() { closeErr <- client.WaitReady(context.Background()) }()
	if err := client.Close(); err != nil {
		t.Fatalf("Close failed: %v", err)
	}
	select {
	case err := <-closeErr:
		if !errors.Is(err, ErrConnectionClosed) {
			t.Fatalf("WaitReady close error = %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("WaitReady did not unblock on close")
	}
	releaseOnce.Do(func() { close(release) })
}

func TestClientConnectWithContextReturnsSessionError(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnError("create_session", "invalid token")

	client := NewClient(Options{Endpoint: mockServer.URL(), AccessToken: "bad-token"})
	err := client.ConnectWithContext(context.Background())
	if err == nil {
		t.Fatal("ConnectWithContext unexpectedly succeeded")
	}
	if !strings.Contains(err.Error(), "invalid token") {
		t.Fatalf("session error = %v, want invalid token", err)
	}
	if health := client.Health(); health.Connected || health.Authenticated {
		t.Fatalf("health after session failure = %+v", health)
	}
}

func TestClientCallRPC(t *testing.T) {
	// Create mock server
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Setup mock handler
	mockServer.OnSimple("get_me", map[string]interface{}{
		"user_id": "123",
		"name":    "Test User",
		"email":   "test@example.com",
	})

	// Create client without access token (to skip auto session creation)
	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	// Call RPC method
	result, err := client.Call("get_me", []interface{}{})
	if err != nil {
		t.Fatalf("Call failed: %v", err)
	}

	// Verify result
	if result == nil {
		t.Fatal("Result is nil")
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	if resultMap["user_id"] != "123" {
		t.Errorf("Expected user_id=123, got %v", resultMap["user_id"])
	}
}

func TestCreateAccessTokenUsesDeviceIDAndBotOS(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	var got []interface{}
	mockServer.On(MethodCreateAccessToken, func(params []interface{}) (interface{}, error) {
		got = params
		return "token-123", nil
	})
	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	token, err := client.CreateAccessToken("user@example.com", "password", "device-id", DefaultBotOS)
	if err != nil {
		t.Fatalf("CreateAccessToken failed: %v", err)
	}
	if token != "token-123" {
		t.Fatalf("token = %q", token)
	}
	want := []interface{}{"user@example.com", "password", "device-id", "bot", ""}
	if fmt.Sprint(got) != fmt.Sprint(want) {
		t.Fatalf("params = %#v, want %#v", got, want)
	}
}

func TestClientRPCError(t *testing.T) {
	// Create mock server
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Setup error handler
	mockServer.OnError("invalid_method", "method not implemented")

	// Create client
	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	// Call RPC method that returns error
	_, err = client.Call("invalid_method", []interface{}{})
	if err == nil {
		t.Fatal("Expected error, got nil")
	}
}

func TestCallWithContextPreCanceledDoesNotSendRequest(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("ping", true)

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()
	if _, err := client.Call("ping", []interface{}{}); err != nil {
		t.Fatalf("connection synchronization call failed: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	started := time.Now()
	_, err := client.GetMeWithContext(ctx)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error = %v, want context.Canceled", err)
	}
	if elapsed := time.Since(started); elapsed > 100*time.Millisecond {
		t.Fatalf("pre-canceled call took %v", elapsed)
	}
	if got := mockServer.GetCallCount("get_me"); got != 0 {
		t.Fatalf("pre-canceled call sent get_me %d times", got)
	}

	client.mu.RLock()
	defer client.mu.RUnlock()
	if got := len(client.responseHandlers); got != 0 {
		t.Fatalf("response handlers after pre-canceled call = %d, want 0", got)
	}
}

func TestCallWithContextDeadlineCleansUpHandler(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	started := make(chan struct{})
	release := make(chan struct{})
	finished := make(chan struct{})
	mockServer.On("get_me", func(params []interface{}) (interface{}, error) {
		close(started)
		<-release
		close(finished)
		return map[string]interface{}{"id": "user123"}, nil
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()
	errCh := make(chan error, 1)
	go func() {
		_, err := client.CallWithContext(ctx, "get_me", []interface{}{})
		errCh <- err
	}()

	select {
	case <-started:
	case <-time.After(time.Second):
		t.Fatal("mock server did not receive get_me")
	}

	err := <-errCh
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error = %v, want context.DeadlineExceeded", err)
	}
	client.mu.RLock()
	handlerCount := len(client.responseHandlers)
	client.mu.RUnlock()
	if handlerCount != 0 {
		t.Fatalf("response handlers after deadline = %d, want 0", handlerCount)
	}

	close(release)
	select {
	case <-finished:
	case <-time.After(time.Second):
		t.Fatal("mock server handler did not finish")
	}
}

func TestCallWithContextPropagatesTraceParent(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("get_me", map[string]interface{}{"id": "user123"})

	provider := &recordingTracerProvider{}
	client := NewClient(Options{Endpoint: mockServer.URL(), TracerProvider: provider})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	parent := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    trace.TraceID{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
		SpanID:     trace.SpanID{1, 2, 3, 4, 5, 6, 7, 8},
		TraceFlags: trace.FlagsSampled,
	})
	ctx := trace.ContextWithSpanContext(context.Background(), parent)
	if _, err := client.GetMeWithContext(ctx); err != nil {
		t.Fatalf("GetMeWithContext failed: %v", err)
	}
	got := provider.parent
	if got.TraceID() != parent.TraceID() || got.SpanID() != parent.SpanID() || got.TraceFlags() != parent.TraceFlags() {
		t.Fatalf("trace parent = %v, want %v", got, parent)
	}
}

func TestClientReconnectAfterClose(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("get_me", map[string]interface{}{"id": "user123"})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("first Connect failed: %v", err)
	}
	firstDone := client.Done
	if _, err := client.Call("get_me", []interface{}{}); err != nil {
		t.Fatalf("first RPC failed: %v", err)
	}
	if health := client.Health(); !health.Connected || health.Authenticated {
		t.Fatalf("health after first connect = %+v", health)
	}

	if err := client.Close(); err != nil {
		t.Fatalf("first Close failed: %v", err)
	}
	select {
	case <-firstDone:
	case <-time.After(time.Second):
		t.Fatal("first Done channel was not closed")
	}
	if health := client.Health(); health.Connected || health.Authenticated {
		t.Fatalf("health after close = %+v", health)
	}
	select {
	case _, ok := <-client.Messages:
		if !ok {
			t.Fatal("Messages was closed during reconnectable Close")
		}
	default:
	}
	if err := client.Close(); err != nil {
		t.Fatalf("second Close failed: %v", err)
	}

	if err := client.Connect(); err != nil {
		t.Fatalf("reconnect failed: %v", err)
	}
	if client.Done == firstDone {
		t.Fatal("reconnect reused the closed Done channel")
	}
	if _, err := client.Call("get_me", []interface{}{}); err != nil {
		t.Fatalf("RPC after reconnect failed: %v", err)
	}
	if health := client.Health(); !health.Connected || health.Authenticated {
		t.Fatalf("health after reconnect = %+v", health)
	}
	if err := client.Close(); err != nil {
		t.Fatalf("final Close failed: %v", err)
	}
}

func TestClientCloseReleasesPendingCall(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	started := make(chan struct{})
	release := make(chan struct{})
	mockServer.On("get_me", func(params []interface{}) (interface{}, error) {
		close(started)
		<-release
		return map[string]interface{}{"id": "late"}, nil
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}

	errCh := make(chan error, 1)
	go func() {
		_, err := client.CallWithContext(context.Background(), "get_me", []interface{}{})
		errCh <- err
	}()
	select {
	case <-started:
	case <-time.After(time.Second):
		client.Close()
		close(release)
		t.Fatal("mock server did not receive get_me")
	}

	if err := client.Close(); err != nil {
		t.Fatalf("Close failed: %v", err)
	}
	select {
	case err := <-errCh:
		if !errors.Is(err, ErrConnectionClosed) {
			t.Fatalf("pending call error = %v, want ErrConnectionClosed", err)
		}
	case <-time.After(200 * time.Millisecond):
		t.Fatal("pending call was not released by Close")
	}

	client.mu.RLock()
	handlerCount := len(client.responseHandlers)
	client.mu.RUnlock()
	if handlerCount != 0 {
		t.Fatalf("response handlers after close = %d, want 0", handlerCount)
	}
	close(release)
}

func TestClientFailedDialCanRetry(t *testing.T) {
	failedServer := testutil.NewMockServer()
	failedURL := failedServer.URL()
	failedServer.Close()

	workingServer := testutil.NewMockServer()
	defer workingServer.Close()
	workingServer.OnSimple("get_me", map[string]interface{}{"id": "user123"})

	client := NewClient(Options{Endpoint: failedURL})
	err := client.Connect()
	if err == nil {
		t.Fatal("Connect unexpectedly succeeded against a closed server")
	}
	if health := client.Health(); health.Connected || health.Authenticated {
		t.Fatalf("health after failed dial = %+v", health)
	}

	client.mu.Lock()
	client.options.Endpoint = workingServer.URL()
	client.mu.Unlock()
	if err := client.Connect(); err != nil {
		t.Fatalf("retry Connect failed: %v", err)
	}
	defer client.Close()
	if _, err := client.Call("get_me", []interface{}{}); err != nil {
		t.Fatalf("RPC after dial retry failed: %v", err)
	}
}

func TestClientFailedAuthenticationCanRetry(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.On("create_session", func(params []interface{}) (interface{}, error) {
		if mockServer.GetCallCount("create_session") == 1 {
			return nil, fmt.Errorf("invalid token")
		}
		return map[string]interface{}{"user_id": "user123"}, nil
	})
	mockServer.OnSimple("get_domains", []interface{}{})
	mockServer.OnSimple("get_talks", []interface{}{})
	mockServer.OnSimple("get_talk_statuses", []interface{}{})
	mockServer.OnSimple("start_notification", true)

	client := NewClient(Options{Endpoint: mockServer.URL(), AccessToken: "token"})
	if err := client.Connect(); err != nil {
		t.Fatalf("first Connect failed: %v", err)
	}
	firstDone := client.Done
	select {
	case <-firstDone:
	case <-time.After(time.Second):
		t.Fatal("failed authentication did not close Done")
	}
	if health := client.Health(); health.Connected || health.Authenticated {
		t.Fatalf("health after failed authentication = %+v", health)
	}

	sessionCreated := make(chan struct{}, 1)
	client.On("session_created", func(data interface{}) {
		sessionCreated <- struct{}{}
	})
	if err := client.Connect(); err != nil {
		t.Fatalf("retry after failed authentication failed: %v", err)
	}
	defer client.Close()
	select {
	case <-sessionCreated:
	case <-time.After(time.Second):
		t.Fatal("retry did not authenticate")
	}
	if health := client.Health(); !health.Connected || !health.Authenticated {
		t.Fatalf("health after authentication retry = %+v", health)
	}
}

func TestMessageDeliveryFanoutBackpressureAndShutdown(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("ping", true)

	drops := make(chan string, 1)
	metrics := &messageDropMetrics{drops: drops}
	client := NewClient(Options{
		Endpoint:           mockServer.URL(),
		MessageChannelSize: 1,
	})
	client.SetMetrics(metrics)
	if cap(client.Messages) != 1 {
		t.Fatalf("message channel capacity = %d, want 1", cap(client.Messages))
	}

	firstHandler := make(chan string, 2)
	secondHandler := make(chan string, 2)
	client.OnMessage(func(msg ReceivedMessage) {
		firstHandler <- msg.ID
	})
	client.OnMessage(func(msg ReceivedMessage) {
		secondHandler <- msg.ID
	})
	client.OnMessage(func(msg ReceivedMessage) {
		panic("handler failure")
	})

	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()
	if _, err := client.Call("ping", []interface{}{}); err != nil {
		t.Fatalf("connection synchronization call failed: %v", err)
	}

	client.mu.RLock()
	conn := client.conn
	client.mu.RUnlock()
	if conn == nil {
		t.Fatal("client connection is nil")
	}

	client.handleMessageNotification(conn, map[string]interface{}{
		"message_id": "first",
		"talk_id":    "talk-1",
		"content":    "one",
	})
	if got := <-firstHandler; got != "first" {
		t.Fatalf("first handler message = %q, want first", got)
	}
	if got := <-secondHandler; got != "first" {
		t.Fatalf("second handler message = %q, want first", got)
	}
	if got := len(client.Messages); got != 1 {
		t.Fatalf("channel length after first message = %d, want 1", got)
	}

	deliveryDone := make(chan struct{})
	go func() {
		client.handleMessageNotification(conn, map[string]interface{}{
			"message_id": "second",
			"talk_id":    "talk-1",
			"content":    "two",
		})
		close(deliveryDone)
	}()
	if got := <-firstHandler; got != "second" {
		t.Fatalf("first handler message = %q, want second", got)
	}
	if got := <-secondHandler; got != "second" {
		t.Fatalf("second handler message = %q, want second", got)
	}

	closeDone := make(chan error, 1)
	go func() {
		closeDone <- client.Close()
	}()
	select {
	case err := <-closeDone:
		if err != nil {
			t.Fatalf("Close failed: %v", err)
		}
	case <-time.After(time.Second):
		t.Fatal("Close blocked on a full message channel")
	}
	select {
	case <-deliveryDone:
	case <-time.After(time.Second):
		t.Fatal("message delivery did not stop after connection shutdown")
	}
	select {
	case reason := <-drops:
		if reason != "connection_closed" {
			t.Fatalf("message drop reason = %q, want connection_closed", reason)
		}
	case <-time.After(time.Second):
		t.Fatal("canceled message delivery was not observable")
	}

	if got := (<-client.Messages).ID; got != "first" {
		t.Fatalf("channel message = %q, want first", got)
	}
	select {
	case msg := <-client.Messages:
		t.Fatalf("second message unexpectedly reached channel: %q", msg.ID)
	default:
	}
}

type messageDropMetrics struct {
	NoopMetrics
	drops chan<- string
}

func (m *messageDropMetrics) RecordMessageDrop(reason string) {
	m.drops <- reason
}

func TestClientConcurrentRPCAndNotificationWrites(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("ping", true)
	mockServer.On("echo", func(params []interface{}) (interface{}, error) {
		if len(params) != 1 {
			return nil, fmt.Errorf("got %d params", len(params))
		}
		return params[0], nil
	})

	client := NewClient(Options{Endpoint: mockServer.URL(), WriteTimeout: time.Second})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()
	if _, err := client.Call("ping", []interface{}{}); err != nil {
		t.Fatalf("connection synchronization call failed: %v", err)
	}

	client.mu.RLock()
	conn := client.conn
	client.mu.RUnlock()
	if conn == nil {
		t.Fatal("client connection is nil")
	}

	const writers = 32
	errs := make(chan error, writers)
	var wg sync.WaitGroup
	wg.Add(writers * 2)
	for i := 0; i < writers; i++ {
		i := i
		go func() {
			defer wg.Done()
			result, err := client.Call("echo", []interface{}{i})
			if err != nil {
				errs <- err
				return
			}
			got, ok := toInt64(result)
			if !ok || got != int64(i) {
				errs <- fmt.Errorf("echo result = %v, want %d", result, i)
			}
		}()
		go func() {
			defer wg.Done()
			client.handleNotification(conn, []interface{}{
				RpcRequest,
				int64(i),
				"notify_test",
				[]interface{}{map[string]interface{}{"id": i}},
			})
		}()
	}
	wg.Wait()
	close(errs)
	for err := range errs {
		t.Error(err)
	}
	if health := client.Health(); !health.Connected {
		t.Fatalf("health after concurrent writes = %+v", health)
	}
}

type recordingTracerProvider struct {
	trace.TracerProvider
	parent trace.SpanContext
}

func (p *recordingTracerProvider) Tracer(name string, options ...trace.TracerOption) trace.Tracer {
	return &recordingTracer{
		Tracer: trace.NewNoopTracerProvider().Tracer(name, options...),
		parent: &p.parent,
	}
}

type recordingTracer struct {
	trace.Tracer
	parent *trace.SpanContext
}

func (t *recordingTracer) Start(ctx context.Context, name string, options ...trace.SpanStartOption) (context.Context, trace.Span) {
	*t.parent = trace.SpanFromContext(ctx).SpanContext()
	return t.Tracer.Start(ctx, name, options...)
}

func TestGetMeWithContext(t *testing.T) {
	// Create mock server
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Setup mock response
	mockServer.OnSimple("get_me", map[string]interface{}{
		"id":           "user123",
		"display_name": "Test User",
		"email":        "test@example.com",
	})

	// Create client
	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	// Call GetMeWithContext
	ctx := context.Background()
	user, err := client.GetMeWithContext(ctx)
	if err != nil {
		t.Fatalf("GetMeWithContext failed: %v", err)
	}

	if user == nil {
		t.Fatal("User is nil")
	}

	// UserInfo.ID is interface{}, need to convert
	if user.ID != "user123" {
		t.Errorf("Expected ID=user123, got %v", user.ID)
	}
}

func TestSendTextWithContext(t *testing.T) {
	// Create mock server
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Setup mock response for create_message
	mockServer.OnSimple("create_message", map[string]interface{}{
		"id":      "msg123",
		"talk_id": "talk456",
		"content": "Hello",
	})

	// Create client
	client := NewClient(Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	// Send text message
	ctx := context.Background()
	err = client.SendTextWithContext(ctx, "talk456", "Hello")
	if err != nil {
		t.Fatalf("SendTextWithContext failed: %v", err)
	}

	// Verify create_message was called with correct params
	found := false
	messages := mockServer.GetReceivedMessages()
	for _, msg := range messages {
		if len(msg) >= 4 && msg[2] == "create_message" {
			params := msg[3].([]interface{})
			t.Logf("create_message params: %v (types: %T, %T, %T)", params, params[0], params[1], params[2])

			if len(params) == 3 {
				// Params: [roomID, msgType, content]
				// msgType can be various integer types depending on msgpack encoding
				msgType := int64(0)
				switch v := params[1].(type) {
				case int:
					msgType = int64(v)
				case int8:
					msgType = int64(v)
				case int64:
					msgType = v
				case uint8:
					msgType = int64(v)
				}

				if params[0] == "talk456" && msgType == 1 && params[2] == "Hello" {
					found = true
					break
				}
			}
		}
	}

	if !found {
		t.Error("create_message was not called with expected params")
	}
}

func TestSendTextWithContextConvertsNumericTalkID(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("create_message", map[string]interface{}{
		"id": "msg123",
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	if err := client.SendTextWithContext(context.Background(), "1792967566075891712", "Hello"); err != nil {
		t.Fatalf("SendTextWithContext failed: %v", err)
	}

	for _, msg := range mockServer.GetReceivedMessages() {
		if len(msg) >= 4 && msg[2] == "create_message" {
			params := msg[3].([]interface{})
			if _, ok := params[0].(uint64); !ok {
				t.Fatalf("talk id type = %T, want uint64", params[0])
			}
			return
		}
	}
	t.Fatal("create_message was not called")
}

func TestCreateTextMessageWithContextReturnsMessageID(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()
	mockServer.OnSimple("create_message", map[string]interface{}{
		"id": "msg123",
	})

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	id, err := client.CreateTextMessageWithContext(context.Background(), "1792967566075891712", "Hello")
	if err != nil {
		t.Fatalf("CreateTextMessageWithContext failed: %v", err)
	}
	if id != "msg123" {
		t.Fatalf("message id = %q, want msg123", id)
	}
}

// Property-Based Tests for toInt64 using Rapid

// TestToInt64_Int verifies int values are correctly converted to int64
func TestToInt64_Int(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Int().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Int8 verifies int8 values are correctly converted to int64
func TestToInt64_Int8(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Int8().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Int16 verifies int16 values are correctly converted to int64
func TestToInt64_Int16(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Int16().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Int32 verifies int32 values are correctly converted to int64
func TestToInt64_Int32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Int32().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Int64 verifies int64 values are returned as-is
func TestToInt64_Int64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Int64().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != val {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, val)
		}
	})
}

// TestToInt64_Uint verifies uint values are correctly converted to int64
func TestToInt64_Uint(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Uint().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Uint8 verifies uint8 values are correctly converted to int64
func TestToInt64_Uint8(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Uint8().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Uint16 verifies uint16 values are correctly converted to int64
func TestToInt64_Uint16(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Uint16().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Uint32 verifies uint32 values are correctly converted to int64
func TestToInt64_Uint32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Uint32().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Uint64 verifies uint64 values (within int64 range) are correctly converted
func TestToInt64_Uint64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Only test values within int64 range to avoid overflow
		val := rapid.Uint64Range(0, 1<<63-1).Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%d) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%d) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Float32 verifies float32 values are correctly converted to int64
func TestToInt64_Float32(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Float32().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%f) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%f) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_Float64 verifies float64 values are correctly converted to int64
func TestToInt64_Float64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.Float64().Draw(t, "val")
		result, ok := toInt64(val)
		if !ok {
			t.Fatalf("toInt64(%f) returned false", val)
		}
		if result != int64(val) {
			t.Fatalf("toInt64(%f) = %d, want %d", val, result, int64(val))
		}
	})
}

// TestToInt64_String verifies non-numeric types return false
func TestToInt64_String(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.String().Draw(t, "val")
		_, ok := toInt64(val)
		if ok {
			t.Fatalf("toInt64(%q) should return false for string", val)
		}
	})
}

// TestToInt64_Slice verifies non-numeric types return false
func TestToInt64_Slice(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.SliceOf(rapid.Int()).Draw(t, "val")
		_, ok := toInt64(val)
		if ok {
			t.Fatalf("toInt64(%v) should return false for slice", val)
		}
	})
}

// TestToInt64_Map verifies non-numeric types return false
func TestToInt64_Map(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		val := rapid.MapOf(rapid.String(), rapid.Int()).Draw(t, "val")
		_, ok := toInt64(val)
		if ok {
			t.Fatalf("toInt64(%v) should return false for map", val)
		}
	})
}

// TestToInt64_Nil verifies nil returns false
func TestToInt64_Nil(t *testing.T) {
	_, ok := toInt64(nil)
	if ok {
		t.Fatal("toInt64(nil) should return false")
	}
}

// Property-Based Tests for parseMessage using Rapid

// TestParseMessage_NonMap returns empty message for non-map input
func TestParseMessage_NonMap(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate non-map values - test each type separately
		kind := rapid.IntRange(0, 3).Draw(t, "kind")
		var val interface{}
		switch kind {
		case 0:
			val = rapid.String().Draw(t, "str")
		case 1:
			val = rapid.Int().Draw(t, "int")
		case 2:
			val = rapid.Bool().Draw(t, "bool")
		default:
			val = rapid.SliceOf(rapid.Int()).Draw(t, "slice")
		}

		result := parseMessage(val)

		// Non-map input should return empty message
		if result.ID != "" {
			t.Fatalf("parseMessage(%v) should have empty ID, got %q", val, result.ID)
		}
		if result.TalkID != "" {
			t.Fatalf("parseMessage(%v) should have empty TalkID, got %q", val, result.TalkID)
		}
	})
}

// TestParseMessage_WithMessageID extracts message_id field
func TestParseMessage_WithMessageID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		id := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "id")
		data := map[string]interface{}{
			"message_id": id,
		}

		result := parseMessage(data)

		expectedID := fmt.Sprintf("%v", id)
		if result.ID != expectedID {
			t.Fatalf("parseMessage() ID = %q, want %q", result.ID, expectedID)
		}
	})
}

// TestParseMessage_WithID extracts id field when message_id is absent
func TestParseMessage_WithID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		id := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "id")
		data := map[string]interface{}{
			"id": id,
		}

		result := parseMessage(data)

		expectedID := fmt.Sprintf("%v", id)
		if result.ID != expectedID {
			t.Fatalf("parseMessage() ID = %q, want %q", result.ID, expectedID)
		}
	})
}

// TestParseMessage_MessageIDTakesPrecedence message_id takes precedence over id
func TestParseMessage_MessageIDTakesPrecedence(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		messageID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "messageID")
		regularID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "regularID")
		data := map[string]interface{}{
			"message_id": messageID,
			"id":         regularID,
		}

		result := parseMessage(data)

		expectedID := fmt.Sprintf("%v", messageID)
		if result.ID != expectedID {
			t.Fatalf("parseMessage() ID = %q, want message_id %q", result.ID, expectedID)
		}
	})
}

// TestParseMessage_WithTalkID extracts talk_id field
func TestParseMessage_WithTalkID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		talkID := rapid.StringMatching(`[0-9]+`).Draw(t, "talkID")
		data := map[string]interface{}{
			"talk_id": talkID,
		}

		result := parseMessage(data)

		expectedTalkID := fmt.Sprintf("%v", talkID)
		if result.TalkID != expectedTalkID {
			t.Fatalf("parseMessage() TalkID = %q, want %q", result.TalkID, expectedTalkID)
		}
		if result.RoomID != expectedTalkID {
			t.Fatalf("parseMessage() RoomID = %q, want %q", result.RoomID, expectedTalkID)
		}
	})
}

// TestParseMessage_WithUserID extracts user_id field
func TestParseMessage_WithUserID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		userID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "userID")
		data := map[string]interface{}{
			"user_id": userID,
		}

		result := parseMessage(data)

		expectedUserID := fmt.Sprintf("%v", userID)
		if result.UserID != expectedUserID {
			t.Fatalf("parseMessage() UserID = %q, want %q", result.UserID, expectedUserID)
		}
	})
}

// TestParseMessage_WithDomainID extracts domain_id field
func TestParseMessage_WithDomainID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		domainID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "domainID")
		data := map[string]interface{}{
			"domain_id": domainID,
		}

		result := parseMessage(data)

		expectedDomainID := fmt.Sprintf("%v", domainID)
		if result.DomainID != expectedDomainID {
			t.Fatalf("parseMessage() DomainID = %q, want %q", result.DomainID, expectedDomainID)
		}
	})
}

// TestParseMessage_WithContentString extracts text from string content
func TestParseMessage_WithContentString(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		text := rapid.StringMatching(`[a-zA-Z0-9\s,.!?]+`).Draw(t, "text")
		data := map[string]interface{}{
			"content": text,
		}

		result := parseMessage(data)

		if result.Text != text {
			t.Fatalf("parseMessage() Text = %q, want %q", result.Text, text)
		}
		if result.Content == nil {
			t.Fatal("parseMessage() Content should not be nil")
		}
	})
}

// TestParseMessage_WithContentMap extracts text from map content
func TestParseMessage_WithContentMap(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		text := rapid.StringMatching(`[a-zA-Z0-9\s]+`).Draw(t, "text")
		data := map[string]interface{}{
			"content": map[string]interface{}{
				"text": text,
			},
		}

		result := parseMessage(data)

		if result.Text != text {
			t.Fatalf("parseMessage() Text = %q, want %q", result.Text, text)
		}
	})
}

// TestParseMessage_WithType extracts and converts type field
func TestParseMessage_WithType(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		msgType := rapid.IntRange(0, 25).Draw(t, "msgType")
		data := map[string]interface{}{
			"type": msgType,
		}

		result := parseMessage(data)

		expectedType := MessageType(msgType)
		if result.Type != expectedType {
			t.Fatalf("parseMessage() Type = %d, want %d", result.Type, expectedType)
		}
	})
}

// TestParseMessage_CompleteMessage parses all fields correctly
func TestParseMessage_CompleteMessage(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		messageID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "messageID")
		talkID := rapid.StringMatching(`[0-9]+`).Draw(t, "talkID")
		userID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "userID")
		domainID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "domainID")
		text := rapid.StringMatching(`[a-zA-Z0-9\s]+`).Draw(t, "text")
		msgType := rapid.IntRange(0, 25).Draw(t, "msgType")

		data := map[string]interface{}{
			"message_id": messageID,
			"talk_id":    talkID,
			"user_id":    userID,
			"domain_id":  domainID,
			"content":    text,
			"type":       msgType,
		}

		result := parseMessage(data)

		if result.ID != fmt.Sprintf("%v", messageID) {
			t.Fatalf("ID mismatch: got %q, want %q", result.ID, messageID)
		}
		if result.TalkID != fmt.Sprintf("%v", talkID) {
			t.Fatalf("TalkID mismatch: got %q, want %q", result.TalkID, talkID)
		}
		if result.RoomID != fmt.Sprintf("%v", talkID) {
			t.Fatalf("RoomID should equal TalkID")
		}
		if result.UserID != fmt.Sprintf("%v", userID) {
			t.Fatalf("UserID mismatch: got %q, want %q", result.UserID, userID)
		}
		if result.DomainID != fmt.Sprintf("%v", domainID) {
			t.Fatalf("DomainID mismatch: got %q, want %q", result.DomainID, domainID)
		}
		if result.Text != text {
			t.Fatalf("Text mismatch: got %q, want %q", result.Text, text)
		}
		if result.Type != MessageType(msgType) {
			t.Fatalf("Type mismatch: got %d, want %d", result.Type, msgType)
		}
	})
}

// TestParseMessage_RawIsSet verifies Raw field is populated with JSON
func TestParseMessage_RawIsSet(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		messageID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "messageID")
		text := rapid.StringMatching(`[a-zA-Z0-9]+`).Draw(t, "text")

		data := map[string]interface{}{
			"message_id": messageID,
			"text":       text,
		}

		result := parseMessage(data)

		if len(result.Raw) == 0 {
			t.Fatal("parseMessage() Raw should be populated with JSON")
		}
		// Verify it's valid JSON
		var parsed map[string]interface{}
		if err := json.Unmarshal(result.Raw, &parsed); err != nil {
			t.Fatalf("parseMessage() Raw should be valid JSON: %v", err)
		}
	})
}

// TestParseMessage_EmptyMap returns empty message
func TestParseMessage_EmptyMap(t *testing.T) {
	data := map[string]interface{}{}

	result := parseMessage(data)

	if result.ID != "" {
		t.Fatalf("parseMessage(empty) ID should be empty, got %q", result.ID)
	}
	if result.TalkID != "" {
		t.Fatalf("parseMessage(empty) TalkID should be empty, got %q", result.TalkID)
	}
	if result.UserID != "" {
		t.Fatalf("parseMessage(empty) UserID should be empty, got %q", result.UserID)
	}
}

// TestParseMessage_NilInput returns empty message
func TestParseMessage_NilInput(t *testing.T) {
	result := parseMessage(nil)

	if result.ID != "" {
		t.Fatalf("parseMessage(nil) ID should be empty, got %q", result.ID)
	}
}
