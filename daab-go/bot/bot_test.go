package bot

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"

	"pgregory.net/rapid"

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

// Property-Based Tests for normalizeRoomID using Rapid

// TestNormalizeRoomID_NumericString verifies numeric strings are converted to uint64
func TestNormalizeRoomID_NumericString(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		num := rapid.Uint64().Draw(t, "num")
		roomID := strconv.FormatUint(num, 10)

		result := normalizeRoomID(roomID)

		resultUint, ok := result.(uint64)
		if !ok {
			t.Fatalf("normalizeRoomID(%q) returned %T, want uint64", roomID, result)
		}
		if resultUint != num {
			t.Fatalf("normalizeRoomID(%q) = %d, want %d", roomID, resultUint, num)
		}
	})
}

// TestNormalizeRoomID_NonNumericString verifies non-numeric strings are returned unchanged
func TestNormalizeRoomID_NonNumericString(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate string that starts with a letter (guaranteed non-numeric)
		firstChar := rapid.StringMatching(`[a-zA-Z]`).Draw(t, "firstChar")
		rest := rapid.StringMatching(`[a-zA-Z0-9_-]*`).Draw(t, "rest")
		roomID := firstChar + rest

		result := normalizeRoomID(roomID)

		resultStr, ok := result.(string)
		if !ok {
			t.Fatalf("normalizeRoomID(%q) returned %T, want string", roomID, result)
		}
		if resultStr != roomID {
			t.Fatalf("normalizeRoomID(%q) = %q, want %q", roomID, resultStr, roomID)
		}
	})
}

// TestNormalizeRoomID_EmptyString verifies empty string is returned as-is
func TestNormalizeRoomID_EmptyString(t *testing.T) {
	result := normalizeRoomID("")
	if result != "" {
		t.Fatalf("normalizeRoomID('') = %v, want ''", result)
	}
}

// TestNormalizeRoomID_LeadingZeros verifies numbers with leading zeros work correctly
func TestNormalizeRoomID_LeadingZeros(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate a small number and add leading zeros
		num := rapid.Uint64Range(0, 99999).Draw(t, "num")
		leadingZeros := rapid.IntRange(0, 10).Draw(t, "zeros")
		roomID := ""
		for i := 0; i < leadingZeros; i++ {
			roomID += "0"
		}
		roomID += strconv.FormatUint(num, 10)

		result := normalizeRoomID(roomID)

		resultUint, ok := result.(uint64)
		if !ok {
			t.Fatalf("normalizeRoomID(%q) returned %T, want uint64", roomID, result)
		}
		if resultUint != num {
			t.Fatalf("normalizeRoomID(%q) = %d, want %d", roomID, resultUint, num)
		}
	})
}

// TestNormalizeRoomID_SpecialChars verifies strings with special characters are returned unchanged
func TestNormalizeRoomID_SpecialChars(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate string with special characters
		special := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "base")
		suffix := rapid.StringMatching(`[-_@!#$%&*+]+`).Draw(t, "suffix")
		roomID := special + suffix

		result := normalizeRoomID(roomID)

		resultStr, ok := result.(string)
		if !ok {
			t.Fatalf("normalizeRoomID(%q) returned %T, want string", roomID, result)
		}
		if resultStr != roomID {
			t.Fatalf("normalizeRoomID(%q) = %q, want %q", roomID, resultStr, roomID)
		}
	})
}

// TestNormalizeRoomID_WhiteSpace verifies whitespace is preserved (returns original)
func TestNormalizeRoomID_WhiteSpace(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate string with spaces
		base := rapid.StringMatching(`[a-zA-Z0-9]+`).Draw(t, "base")
		spaces := rapid.IntRange(1, 5).Draw(t, "spaces")
		roomID := base
		for i := 0; i < spaces; i++ {
			roomID += " "
		}

		result := normalizeRoomID(roomID)

		// Strings with spaces should not be parsed as numbers
		resultStr, ok := result.(string)
		if !ok {
			t.Fatalf("normalizeRoomID(%q) returned %T, want string", roomID, result)
		}
		if resultStr != roomID {
			t.Fatalf("normalizeRoomID(%q) = %q, want %q", roomID, resultStr, roomID)
		}
	})
}

// Property-Based Tests for extractMessageID using Rapid

// TestExtractMessageID_MapWithMessageID verifies maps with message_id field return the ID
func TestExtractMessageID_MapWithMessageID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		id := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "id")
		result := map[string]interface{}{
			"message_id": id,
			"other":      "ignored",
		}

		extracted := extractMessageID(result)
		if extracted != fmt.Sprintf("%v", id) {
			t.Fatalf("extractMessageID(%v) = %q, want %q", result, extracted, id)
		}
	})
}

// TestExtractMessageID_MapWithID verifies maps with id field (no message_id) return the ID
func TestExtractMessageID_MapWithID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		id := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "id")
		result := map[string]interface{}{
			"id": id,
		}

		extracted := extractMessageID(result)
		if extracted != fmt.Sprintf("%v", id) {
			t.Fatalf("extractMessageID(%v) = %q, want %q", result, extracted, id)
		}
	})
}

// TestExtractMessageID_MapWithBoth verifies message_id takes precedence over id
func TestExtractMessageID_MapWithBoth(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		messageID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "messageID")
		regularID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "regularID")
		result := map[string]interface{}{
			"message_id": messageID,
			"id":         regularID,
		}

		extracted := extractMessageID(result)
		if extracted != fmt.Sprintf("%v", messageID) {
			t.Fatalf("extractMessageID(%v) = %q, want message_id %q", result, extracted, messageID)
		}
	})
}

// TestExtractMessageID_String verifies strings are returned as-is
func TestExtractMessageID_String(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		str := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "str")

		extracted := extractMessageID(str)
		if extracted != str {
			t.Fatalf("extractMessageID(%q) = %q, want %q", str, extracted, str)
		}
	})
}

// TestExtractMessageID_Int verifies integers are converted to strings
func TestExtractMessageID_Int(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		num := rapid.Int().Draw(t, "num")

		extracted := extractMessageID(num)
		expected := fmt.Sprintf("%v", num)
		if extracted != expected {
			t.Fatalf("extractMessageID(%d) = %q, want %q", num, extracted, expected)
		}
	})
}

// TestExtractMessageID_Int64 verifies int64 values are converted to strings
func TestExtractMessageID_Int64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		num := rapid.Int64().Draw(t, "num")

		extracted := extractMessageID(num)
		expected := fmt.Sprintf("%v", num)
		if extracted != expected {
			t.Fatalf("extractMessageID(%d) = %q, want %q", num, extracted, expected)
		}
	})
}

// TestExtractMessageID_Uint64 verifies uint64 values are converted to strings
func TestExtractMessageID_Uint64(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		num := rapid.Uint64().Draw(t, "num")

		extracted := extractMessageID(num)
		expected := fmt.Sprintf("%v", num)
		if extracted != expected {
			t.Fatalf("extractMessageID(%d) = %q, want %q", num, extracted, expected)
		}
	})
}

// TestExtractMessageID_Float verifies floats are converted to strings
func TestExtractMessageID_Float(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		num := rapid.Float64().Draw(t, "num")

		extracted := extractMessageID(num)
		expected := fmt.Sprintf("%v", num)
		if extracted != expected {
			t.Fatalf("extractMessageID(%f) = %q, want %q", num, extracted, expected)
		}
	})
}

// TestExtractMessageID_Nil verifies nil returns empty string
func TestExtractMessageID_Nil(t *testing.T) {
	extracted := extractMessageID(nil)
	if extracted != "" {
		t.Fatalf("extractMessageID(nil) = %q, want empty string", extracted)
	}
}

// TestExtractMessageID_EmptyMap verifies empty map returns empty string
func TestExtractMessageID_EmptyMap(t *testing.T) {
	extracted := extractMessageID(map[string]interface{}{})
	if extracted != "" {
		t.Fatalf("extractMessageID(empty map) = %q, want empty string", extracted)
	}
}

// TestExtractMessageID_MapWithOtherKeys verifies map without id fields returns empty string
func TestExtractMessageID_MapWithOtherKeys(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		key := rapid.StringMatching(`[a-z]+`).Filter(func(s string) bool {
			return s != "message_id" && s != "id"
		}).Draw(t, "key")
		value := rapid.String().Draw(t, "value")
		result := map[string]interface{}{
			key: value,
		}

		extracted := extractMessageID(result)
		if extracted != "" {
			t.Fatalf("extractMessageID(%v) = %q, want empty string", result, extracted)
		}
	})
}

// TestExtractMessageID_Bool verifies booleans are converted to strings
func TestExtractMessageID_Bool(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		b := rapid.Bool().Draw(t, "b")

		extracted := extractMessageID(b)
		expected := fmt.Sprintf("%v", b)
		if extracted != expected {
			t.Fatalf("extractMessageID(%t) = %q, want %q", b, extracted, expected)
		}
	})
}

// TestExtractMessageID_Slice verifies slices are converted to strings
func TestExtractMessageID_Slice(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		slice := rapid.SliceOf(rapid.StringMatching(`[a-z]+`)).Draw(t, "slice")

		extracted := extractMessageID(slice)
		expected := fmt.Sprintf("%v", slice)
		if extracted != expected {
			t.Fatalf("extractMessageID(%v) = %q, want %q", slice, extracted, expected)
		}
	})
}
