package webhook

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewPayload(t *testing.T) {
	msg := MessageData{
		ID:       "123",
		TalkID:   "456",
		UserID:   "789",
		Type:     1,
		TypeName: "text",
		Text:     "hello",
		Created:  1702345678,
	}

	payload := NewPayload("message_created", "testbot", msg)

	if payload.Version != "1.0" {
		t.Errorf("Expected version 1.0, got %s", payload.Version)
	}
	if payload.EventType != "message_created" {
		t.Errorf("Expected eventType message_created, got %s", payload.EventType)
	}
	if payload.Bot.Name != "testbot" {
		t.Errorf("Expected bot name testbot, got %s", payload.Bot.Name)
	}
	if payload.Message.ID != "123" {
		t.Errorf("Expected message ID 123, got %s", payload.Message.ID)
	}
}

func TestMessageTypeToName(t *testing.T) {
	tests := []struct {
		msgType  int
		expected string
	}{
		{1, "text"},
		{2, "stamp"},
		{13, "yesno"},
		{15, "select"},
		{17, "task"},
		{999, "unknown"},
	}

	for _, tt := range tests {
		result := MessageTypeToName(tt.msgType)
		if result != tt.expected {
			t.Errorf("MessageTypeToName(%d) = %s, want %s", tt.msgType, result, tt.expected)
		}
	}
}

func TestWebhookResponseValidate(t *testing.T) {
	tests := []struct {
		name     string
		response WebhookResponse
		expected ErrorCode
	}{
		{
			name:     "none action is valid",
			response: WebhookResponse{Action: "none"},
			expected: ErrorCodeOK,
		},
		{
			name:     "missing action",
			response: WebhookResponse{},
			expected: ErrorCodeMissingAction,
		},
		{
			name:     "reply without text",
			response: WebhookResponse{Action: "reply"},
			expected: ErrorCodeMissingText,
		},
		{
			name:     "reply with text",
			response: WebhookResponse{Action: "reply", Text: "hello"},
			expected: ErrorCodeOK,
		},
		{
			name:     "send without roomId",
			response: WebhookResponse{Action: "send", Text: "hello"},
			expected: ErrorCodeMissingRoomID,
		},
		{
			name:     "send without text",
			response: WebhookResponse{Action: "send", RoomID: "123"},
			expected: ErrorCodeMissingText,
		},
		{
			name:     "send with all fields",
			response: WebhookResponse{Action: "send", RoomID: "123", Text: "hello"},
			expected: ErrorCodeOK,
		},
		{
			name:     "send_select without question",
			response: WebhookResponse{Action: "send_select", RoomID: "123", Options: []string{"A", "B"}},
			expected: ErrorCodeMissingQuestion,
		},
		{
			name:     "send_select without options",
			response: WebhookResponse{Action: "send_select", RoomID: "123", Question: "Q?"},
			expected: ErrorCodeMissingOptions,
		},
		{
			name:     "send_select valid",
			response: WebhookResponse{Action: "send_select", RoomID: "123", Question: "Q?", Options: []string{"A", "B"}},
			expected: ErrorCodeOK,
		},
		{
			name:     "invalid action",
			response: WebhookResponse{Action: "unknown_action"},
			expected: ErrorCodeInvalidAction,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.response.Validate()
			if result != tt.expected {
				t.Errorf("Validate() = %s, want %s", result, tt.expected)
			}
		})
	}
}

func TestClientSend(t *testing.T) {
	// Create mock server
	mockResp := WebhookResponse{
		Action: "reply",
		Text:   "mock response",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("Expected POST, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Verify payload
		var payload WebhookPayload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			t.Errorf("Failed to decode payload: %v", err)
		}

		// Send response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mockResp)
	}))
	defer server.Close()

	// Create client
	client := NewClient(server.URL, "testbot")

	// Send payload
	payload := NewPayload("message_created", "testbot", MessageData{
		ID:     "123",
		TalkID: "456",
		Type:   1,
	})

	resp, err := client.Send(payload)
	if err != nil {
		t.Fatalf("Send failed: %v", err)
	}

	if resp.Action != "reply" {
		t.Errorf("Expected action reply, got %s", resp.Action)
	}
	if resp.Text != "mock response" {
		t.Errorf("Expected text 'mock response', got %s", resp.Text)
	}
}

func TestClientSendError(t *testing.T) {
	// Create server that returns error
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	client := NewClient(server.URL, "testbot")
	payload := NewPayload("message_created", "testbot", MessageData{ID: "123"})

	_, err := client.Send(payload)
	if err == nil {
		t.Error("Expected error for 500 status, got nil")
	}
}
