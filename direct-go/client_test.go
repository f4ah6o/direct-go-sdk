package direct

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"pgregory.net/rapid"

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
	mockServer.OnSimple("get_talks", []interface{}{})
	mockServer.OnSimple("get_talk_statuses", []interface{}{})
	mockServer.OnSimple("start_notification", true)

	// Connect
	err := client.Connect()
	if err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	defer client.Close()

	// Give some time for session creation
	time.Sleep(100 * time.Millisecond)

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
