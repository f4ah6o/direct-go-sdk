package bot

import (
	"context"
	"errors"
	"sync"
	"testing"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-go/testutil"
)

func TestNew(t *testing.T) {
	robot := New()
	if robot == nil {
		t.Fatal("Expected robot to be created")
	}
	if robot.Name != "daabgo" {
		t.Errorf("Expected default name 'daabgo', got %s", robot.Name)
	}
	if robot.listeners == nil {
		t.Error("Expected listeners to be initialized")
	}
	if robot.auth == nil {
		t.Error("Expected auth to be initialized")
	}
}

func TestNewWithOptions(t *testing.T) {
	robot := New(
		WithName("testbot"),
		WithToken("test-token-123"),
		WithEndpoint("wss://test.example.com"),
		WithProxy("http://proxy.example.com"),
	)

	if robot.Name != "testbot" {
		t.Errorf("Expected name 'testbot', got %s", robot.Name)
	}
	if robot.Token != "test-token-123" {
		t.Errorf("Expected token 'test-token-123', got %s", robot.Token)
	}
	if robot.endpoint != "wss://test.example.com" {
		t.Errorf("Expected endpoint 'wss://test.example.com', got %s", robot.endpoint)
	}
	if robot.proxyURL != "http://proxy.example.com" {
		t.Errorf("Expected proxy 'http://proxy.example.com', got %s", robot.proxyURL)
	}
}

func TestHear(t *testing.T) {
	robot := New()
	var called bool
	var mu sync.Mutex
	robot.Hear("hello", func(ctx context.Context, res Response) {
		mu.Lock()
		called = true
		mu.Unlock()
	})

	if len(robot.listeners) != 1 {
		t.Errorf("Expected 1 listener, got %d", len(robot.listeners))
	}

	listener := robot.listeners[0]
	if listener.IsDirect {
		t.Error("Expected Hear listener to not be direct")
	}

	// Test pattern matching
	matches := listener.Pattern.FindStringSubmatch("say hello world")
	if matches == nil {
		t.Error("Expected pattern to match 'say hello world'")
	}

	// Simulate handler call
	msg := direct.ReceivedMessage{Text: "hello"}
	robot.handleMessage(context.Background(), msg)
	time.Sleep(10 * time.Millisecond) // Give goroutine time to execute

	mu.Lock()
	wasCalled := called
	mu.Unlock()
	if !wasCalled {
		t.Error("Expected handler to be called")
	}
}

func TestRespond(t *testing.T) {
	robot := New(WithName("testbot"))
	var called bool
	var mu sync.Mutex
	robot.Respond("ping", func(ctx context.Context, res Response) {
		mu.Lock()
		called = true
		mu.Unlock()
	})

	if len(robot.listeners) != 1 {
		t.Errorf("Expected 1 listener, got %d", len(robot.listeners))
	}

	listener := robot.listeners[0]
	if !listener.IsDirect {
		t.Error("Expected Respond listener to be direct")
	}

	// Test pattern matching for direct address
	matches := listener.Pattern.FindStringSubmatch("@testbot ping")
	if matches == nil {
		t.Error("Expected pattern to match '@testbot ping'")
	}

	matches = listener.Pattern.FindStringSubmatch("testbot: ping")
	if matches == nil {
		t.Error("Expected pattern to match 'testbot: ping'")
	}

	// Should not match without bot name
	matches = listener.Pattern.FindStringSubmatch("ping")
	if matches != nil {
		t.Error("Expected pattern to not match 'ping' without bot name")
	}

	// Simulate handler call
	msg := direct.ReceivedMessage{Text: "@testbot ping"}
	robot.handleMessage(context.Background(), msg)
	time.Sleep(10 * time.Millisecond) // Give goroutine time to execute

	mu.Lock()
	wasCalled := called
	mu.Unlock()
	if !wasCalled {
		t.Error("Expected handler to be called")
	}
}

func TestOnEvent(t *testing.T) {
	robot := New()
	var connectedCalled, readyCalled bool
	var mu sync.Mutex

	robot.On(EventConnected, func() {
		mu.Lock()
		connectedCalled = true
		mu.Unlock()
	})

	robot.On(EventReady, func() {
		mu.Lock()
		readyCalled = true
		mu.Unlock()
	})

	// Emit events
	robot.emit(EventConnected)
	time.Sleep(10 * time.Millisecond)
	mu.Lock()
	wasConnectedCalled := connectedCalled
	mu.Unlock()
	if !wasConnectedCalled {
		t.Error("Expected connected event handler to be called")
	}

	robot.emit(EventReady)
	time.Sleep(10 * time.Millisecond)
	mu.Lock()
	wasReadyCalled := readyCalled
	mu.Unlock()
	if !wasReadyCalled {
		t.Error("Expected ready event handler to be called")
	}
}

func TestResponseMethods(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	// Setup mock handlers
	mockServer.OnSimple("create_message", map[string]interface{}{
		"id":      "msg123",
		"talk_id": "talk456",
	})

	client := direct.NewClient(direct.Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	robot := New()
	robot.client = client

	msg := direct.ReceivedMessage{
		ID:     "123",
		TalkID: "talk456",
		UserID: "user789",
		Text:   "test message",
	}

	response := Response{
		Message: msg,
		Match:   []string{"test message"},
		Robot:   robot,
	}

	// Test Text()
	if response.Text() != "test message" {
		t.Errorf("Expected text 'test message', got %s", response.Text())
	}

	// Test RoomID()
	if response.RoomID() != "talk456" {
		t.Errorf("Expected roomID 'talk456', got %s", response.RoomID())
	}

	// Test UserID()
	if response.UserID() != "user789" {
		t.Errorf("Expected userID 'user789', got %s", response.UserID())
	}

	// Test Send()
	err = response.Send("reply text")
	if err != nil {
		t.Errorf("Send failed: %v", err)
	}
}

func TestSendText(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("create_message", map[string]interface{}{
		"id":      "msg123",
		"talk_id": "room456",
	})

	client := direct.NewClient(direct.Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	robot := New()
	robot.client = client

	err = robot.SendText("room456", "Hello world")
	if err != nil {
		t.Errorf("SendText failed: %v", err)
	}
}

func TestSendTextNotConnected(t *testing.T) {
	robot := New()
	err := robot.SendText("room456", "Hello")
	if !errors.Is(err, ErrNotConnected) {
		t.Errorf("Expected ErrNotConnected, got %v", err)
	}
}

func TestCallMethod(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("test_method", map[string]interface{}{
		"result": "success",
	})

	client := direct.NewClient(direct.Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	robot := New()
	robot.client = client

	result, err := robot.Call("test_method", []interface{}{"param1"})
	if err != nil {
		t.Errorf("Call failed: %v", err)
	}

	if result == nil {
		t.Fatal("Result is nil")
	}

	resultMap, ok := result.(map[string]interface{})
	if !ok {
		t.Fatalf("Result is not a map: %T", result)
	}

	if resultMap["result"] != "success" {
		t.Errorf("Expected result=success, got %v", resultMap["result"])
	}
}

func TestCallNotConnected(t *testing.T) {
	robot := New()
	_, err := robot.Call("test_method", []interface{}{})
	if !errors.Is(err, ErrNotConnected) {
		t.Errorf("Expected ErrNotConnected, got %v", err)
	}
}

func TestNormalizeRoomID(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"12345", uint64(12345)},
		{"room-abc", "room-abc"},
		{"0", uint64(0)},
	}

	for _, tt := range tests {
		result := normalizeRoomID(tt.input)
		if result != tt.expected {
			t.Errorf("normalizeRoomID(%s) = %v (type %T), want %v (type %T)",
				tt.input, result, result, tt.expected, tt.expected)
		}
	}
}

func TestExtractMessageID(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name: "map with message_id",
			input: map[string]interface{}{
				"message_id": "msg123",
			},
			expected: "msg123",
		},
		{
			name: "map with id",
			input: map[string]interface{}{
				"id": "msg456",
			},
			expected: "msg456",
		},
		{
			name:     "string",
			input:    "msg789",
			expected: "msg789",
		},
		{
			name:     "integer",
			input:    12345,
			expected: "12345",
		},
		{
			name:     "nil",
			input:    nil,
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractMessageID(tt.input)
			if result != tt.expected {
				t.Errorf("extractMessageID() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestSendSelect(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("create_message", map[string]interface{}{
		"message_id": "msg123",
	})

	client := direct.NewClient(direct.Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	robot := New()
	robot.client = client

	msg := direct.ReceivedMessage{
		TalkID: "talk456",
	}

	response := Response{
		Message: msg,
		Robot:   robot,
	}

	messageID, err := response.SendSelect("Choose one:", []string{"Option A", "Option B", "Option C"})
	if err != nil {
		t.Errorf("SendSelect failed: %v", err)
	}

	if messageID != "msg123" {
		t.Errorf("Expected message ID 'msg123', got %s", messageID)
	}
}

func TestReply(t *testing.T) {
	mockServer := testutil.NewMockServer()
	defer mockServer.Close()

	mockServer.OnSimple("create_message", map[string]interface{}{
		"id": "msg123",
	})

	client := direct.NewClient(direct.Options{
		Endpoint: mockServer.URL(),
	})

	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	robot := New()
	robot.client = client

	msg := direct.ReceivedMessage{
		TalkID: "talk456",
		UserID: "user789",
	}

	response := Response{
		Message: msg,
		Robot:   robot,
	}

	err = response.Reply("Hello!")
	if err != nil {
		t.Errorf("Reply failed: %v", err)
	}

	// Verify that the message contains the mention
	time.Sleep(10 * time.Millisecond)
	messages := mockServer.GetReceivedMessages()
	found := false
	for _, msg := range messages {
		if len(msg) >= 4 && msg[2] == "create_message" {
			params := msg[3].([]interface{})
			if len(params) >= 3 {
				text, ok := params[2].(string)
				if ok && text == "@user789 Hello!" {
					found = true
					break
				}
			}
		}
	}

	if !found {
		t.Error("Expected reply message to contain mention")
	}
}
