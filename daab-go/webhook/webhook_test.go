package webhook

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"pgregory.net/rapid"
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

// Property-Based Tests for WebhookResponse.Validate using Rapid

// TestValidate_ActionNone verifies "none" action is always valid
func TestValidate_ActionNone(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		resp := WebhookResponse{
			Action: "none",
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for 'none' action = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_Reply verifies valid reply actions pass validation
func TestValidate_Reply(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		text := rapid.StringMatching(`[a-zA-Z0-9\s,.!?]+`).Draw(t, "text")
		resp := WebhookResponse{
			Action: "reply",
			Text:   text,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid reply = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_ReplyMissingText verifies reply without text returns error
func TestValidate_ReplyMissingText(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		resp := WebhookResponse{
			Action: "reply",
			Text:   "",
		}
		result := resp.Validate()
		if result != ErrorCodeMissingText {
			t.Fatalf("Validate() for reply without text = %s, want %s", result, ErrorCodeMissingText)
		}
	})
}

// TestValidate_Send verifies valid send actions pass validation
func TestValidate_Send(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		text := rapid.StringMatching(`[a-zA-Z0-9\s]+`).Draw(t, "text")
		resp := WebhookResponse{
			Action: "send",
			RoomID: roomID,
			Text:   text,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid send = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_SendMissingRoomID verifies send without roomId returns error
func TestValidate_SendMissingRoomID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		text := rapid.StringMatching(`[a-zA-Z0-9\s]+`).Draw(t, "text")
		resp := WebhookResponse{
			Action: "send",
			RoomID: "",
			Text:   text,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingRoomID {
			t.Fatalf("Validate() for send without roomId = %s, want %s", result, ErrorCodeMissingRoomID)
		}
	})
}

// TestValidate_SendMissingText verifies send without text returns error
func TestValidate_SendMissingText(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		resp := WebhookResponse{
			Action: "send",
			RoomID: roomID,
			Text:   "",
		}
		result := resp.Validate()
		if result != ErrorCodeMissingText {
			t.Fatalf("Validate() for send without text = %s, want %s", result, ErrorCodeMissingText)
		}
	})
}

// TestValidate_SendSelect verifies valid send_select actions pass validation
func TestValidate_SendSelect(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		question := rapid.StringMatching(`[a-zA-Z0-9\s?]+`).Draw(t, "question")
		options := rapid.SliceOf(rapid.StringMatching(`[a-zA-Z0-9]+`)).Filter(func(s []string) bool {
			return len(s) > 0
		}).Draw(t, "options")
		resp := WebhookResponse{
			Action:   "send_select",
			RoomID:   roomID,
			Question: question,
			Options:  options,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid send_select = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_SendSelectMissingRoomID verifies send_select without roomId returns error
func TestValidate_SendSelectMissingRoomID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		question := rapid.StringMatching(`[a-zA-Z0-9\s?]+`).Draw(t, "question")
		options := rapid.SliceOf(rapid.StringMatching(`[a-zA-Z0-9]+`)).Filter(func(s []string) bool {
			return len(s) > 0
		}).Draw(t, "options")
		resp := WebhookResponse{
			Action:   "send_select",
			RoomID:   "",
			Question: question,
			Options:  options,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingRoomID {
			t.Fatalf("Validate() for send_select without roomId = %s, want %s", result, ErrorCodeMissingRoomID)
		}
	})
}

// TestValidate_SendSelectMissingQuestion verifies send_select without question returns error
func TestValidate_SendSelectMissingQuestion(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		options := rapid.SliceOf(rapid.StringMatching(`[a-zA-Z0-9]+`)).Filter(func(s []string) bool {
			return len(s) > 0
		}).Draw(t, "options")
		resp := WebhookResponse{
			Action:   "send_select",
			RoomID:   roomID,
			Question: "",
			Options:  options,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingQuestion {
			t.Fatalf("Validate() for send_select without question = %s, want %s", result, ErrorCodeMissingQuestion)
		}
	})
}

// TestValidate_SendSelectMissingOptions verifies send_select without options returns error
func TestValidate_SendSelectMissingOptions(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		question := rapid.StringMatching(`[a-zA-Z0-9\s?]+`).Draw(t, "question")
		resp := WebhookResponse{
			Action:   "send_select",
			RoomID:   roomID,
			Question: question,
			Options:  []string{},
		}
		result := resp.Validate()
		if result != ErrorCodeMissingOptions {
			t.Fatalf("Validate() for send_select without options = %s, want %s", result, ErrorCodeMissingOptions)
		}
	})
}

// TestValidate_SendYesno verifies valid send_yesno actions pass validation
func TestValidate_SendYesno(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		question := rapid.StringMatching(`[a-zA-Z0-9\s?]+`).Draw(t, "question")
		resp := WebhookResponse{
			Action:   "send_yesno",
			RoomID:   roomID,
			Question: question,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid send_yesno = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_SendYesnoMissingRoomID verifies send_yesno without roomId returns error
func TestValidate_SendYesnoMissingRoomID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		question := rapid.StringMatching(`[a-zA-Z0-9\s?]+`).Draw(t, "question")
		resp := WebhookResponse{
			Action:   "send_yesno",
			RoomID:   "",
			Question: question,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingRoomID {
			t.Fatalf("Validate() for send_yesno without roomId = %s, want %s", result, ErrorCodeMissingRoomID)
		}
	})
}

// TestValidate_SendYesnoMissingQuestion verifies send_yesno without question returns error
func TestValidate_SendYesnoMissingQuestion(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		resp := WebhookResponse{
			Action:   "send_yesno",
			RoomID:   roomID,
			Question: "",
		}
		result := resp.Validate()
		if result != ErrorCodeMissingQuestion {
			t.Fatalf("Validate() for send_yesno without question = %s, want %s", result, ErrorCodeMissingQuestion)
		}
	})
}

// TestValidate_SendTask verifies valid send_task actions pass validation
func TestValidate_SendTask(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		title := rapid.StringMatching(`[a-zA-Z0-9\s]+`).Draw(t, "title")
		resp := WebhookResponse{
			Action: "send_task",
			RoomID: roomID,
			Title:  title,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid send_task = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_SendTaskMissingRoomID verifies send_task without roomId returns error
func TestValidate_SendTaskMissingRoomID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		title := rapid.StringMatching(`[a-zA-Z0-9\s]+`).Draw(t, "title")
		resp := WebhookResponse{
			Action: "send_task",
			RoomID: "",
			Title:  title,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingRoomID {
			t.Fatalf("Validate() for send_task without roomId = %s, want %s", result, ErrorCodeMissingRoomID)
		}
	})
}

// TestValidate_SendTaskMissingTitle verifies send_task without title returns error
func TestValidate_SendTaskMissingTitle(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		roomID := rapid.StringMatching(`[0-9]+`).Draw(t, "roomID")
		resp := WebhookResponse{
			Action: "send_task",
			RoomID: roomID,
			Title:  "",
		}
		result := resp.Validate()
		if result != ErrorCodeMissingTitle {
			t.Fatalf("Validate() for send_task without title = %s, want %s", result, ErrorCodeMissingTitle)
		}
	})
}

// TestValidate_ReplySelect verifies valid reply_select actions pass validation
func TestValidate_ReplySelect(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		inReplyTo := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "inReplyTo")
		response := rapid.Int().Draw(t, "response")
		resp := WebhookResponse{
			Action:    "reply_select",
			InReplyTo: inReplyTo,
			Response:  &response,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid reply_select = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_ReplySelectMissingInReplyTo verifies reply_select without inReplyTo returns error
func TestValidate_ReplySelectMissingInReplyTo(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		response := rapid.Int().Draw(t, "response")
		resp := WebhookResponse{
			Action:    "reply_select",
			InReplyTo: "",
			Response:  &response,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingInReplyTo {
			t.Fatalf("Validate() for reply_select without inReplyTo = %s, want %s", result, ErrorCodeMissingInReplyTo)
		}
	})
}

// TestValidate_ReplySelectMissingResponse verifies reply_select without response returns error
func TestValidate_ReplySelectMissingResponse(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		inReplyTo := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "inReplyTo")
		resp := WebhookResponse{
			Action:    "reply_select",
			InReplyTo: inReplyTo,
			Response:  nil,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingResponse {
			t.Fatalf("Validate() for reply_select without response = %s, want %s", result, ErrorCodeMissingResponse)
		}
	})
}

// TestValidate_ReplyYesno verifies valid reply_yesno actions pass validation
func TestValidate_ReplyYesno(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		inReplyTo := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "inReplyTo")
		response := rapid.Bool().Draw(t, "response")
		resp := WebhookResponse{
			Action:       "reply_yesno",
			InReplyTo:    inReplyTo,
			ResponseBool: &response,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid reply_yesno = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_ReplyYesnoMissingInReplyTo verifies reply_yesno without inReplyTo returns error
func TestValidate_ReplyYesnoMissingInReplyTo(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		response := rapid.Bool().Draw(t, "response")
		resp := WebhookResponse{
			Action:       "reply_yesno",
			InReplyTo:    "",
			ResponseBool: &response,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingInReplyTo {
			t.Fatalf("Validate() for reply_yesno without inReplyTo = %s, want %s", result, ErrorCodeMissingInReplyTo)
		}
	})
}

// TestValidate_ReplyYesnoMissingResponse verifies reply_yesno without response returns error
func TestValidate_ReplyYesnoMissingResponse(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		inReplyTo := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "inReplyTo")
		resp := WebhookResponse{
			Action:       "reply_yesno",
			InReplyTo:    inReplyTo,
			ResponseBool: nil,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingResponse {
			t.Fatalf("Validate() for reply_yesno without response = %s, want %s", result, ErrorCodeMissingResponse)
		}
	})
}

// TestValidate_ReplyTask verifies valid reply_task actions pass validation
func TestValidate_ReplyTask(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		inReplyTo := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "inReplyTo")
		done := rapid.Bool().Draw(t, "done")
		resp := WebhookResponse{
			Action:    "reply_task",
			InReplyTo: inReplyTo,
			Done:      &done,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid reply_task = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_ReplyTaskMissingInReplyTo verifies reply_task without inReplyTo returns error
func TestValidate_ReplyTaskMissingInReplyTo(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		done := rapid.Bool().Draw(t, "done")
		resp := WebhookResponse{
			Action:    "reply_task",
			InReplyTo: "",
			Done:      &done,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingInReplyTo {
			t.Fatalf("Validate() for reply_task without inReplyTo = %s, want %s", result, ErrorCodeMissingInReplyTo)
		}
	})
}

// TestValidate_ReplyTaskMissingDone verifies reply_task without done returns error
func TestValidate_ReplyTaskMissingDone(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		inReplyTo := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "inReplyTo")
		resp := WebhookResponse{
			Action:    "reply_task",
			InReplyTo: inReplyTo,
			Done:      nil,
		}
		result := resp.Validate()
		if result != ErrorCodeMissingResponse {
			t.Fatalf("Validate() for reply_task without done = %s, want %s", result, ErrorCodeMissingResponse)
		}
	})
}

// TestValidate_CloseSelect verifies valid close_select actions pass validation
func TestValidate_CloseSelect(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		messageID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "messageID")
		resp := WebhookResponse{
			Action:    "close_select",
			MessageID: messageID,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid close_select = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_CloseSelectMissingMessageID verifies close_select without messageId returns error
func TestValidate_CloseSelectMissingMessageID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		resp := WebhookResponse{
			Action:    "close_select",
			MessageID: "",
		}
		result := resp.Validate()
		if result != ErrorCodeMissingMessageID {
			t.Fatalf("Validate() for close_select without messageId = %s, want %s", result, ErrorCodeMissingMessageID)
		}
	})
}

// TestValidate_CloseYesno verifies valid close_yesno actions pass validation
func TestValidate_CloseYesno(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		messageID := rapid.StringMatching(`[a-zA-Z0-9_-]+`).Draw(t, "messageID")
		resp := WebhookResponse{
			Action:    "close_yesno",
			MessageID: messageID,
		}
		result := resp.Validate()
		if result != ErrorCodeOK {
			t.Fatalf("Validate() for valid close_yesno = %s, want %s", result, ErrorCodeOK)
		}
	})
}

// TestValidate_CloseYesnoMissingMessageID verifies close_yesno without messageId returns error
func TestValidate_CloseYesnoMissingMessageID(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		resp := WebhookResponse{
			Action:    "close_yesno",
			MessageID: "",
		}
		result := resp.Validate()
		if result != ErrorCodeMissingMessageID {
			t.Fatalf("Validate() for close_yesno without messageId = %s, want %s", result, ErrorCodeMissingMessageID)
		}
	})
}

// TestValidate_InvalidAction verifies unknown actions return invalid action error
func TestValidate_InvalidAction(t *testing.T) {
	rapid.Check(t, func(t *rapid.T) {
		// Generate a string that's not a valid action
		action := rapid.StringMatching(`[a-z_]+`).Filter(func(s string) bool {
			validActions := map[string]bool{
				"none":        true,
				"reply":       true,
				"send":        true,
				"send_select": true,
				"send_yesno":  true,
				"send_task":   true,
				"reply_select": true,
				"reply_yesno":  true,
				"reply_task":   true,
				"close_select": true,
				"close_yesno":  true,
			}
			return !validActions[s]
		}).Draw(t, "action")
		resp := WebhookResponse{
			Action: action,
		}
		result := resp.Validate()
		if result != ErrorCodeInvalidAction {
			t.Fatalf("Validate() for invalid action %q = %s, want %s", action, result, ErrorCodeInvalidAction)
		}
	})
}

// TestValidate_EmptyAction verifies empty action returns missing action error
func TestValidate_EmptyAction(t *testing.T) {
	resp := WebhookResponse{
		Action: "",
	}
	result := resp.Validate()
	if result != ErrorCodeMissingAction {
		t.Fatalf("Validate() for empty action = %s, want %s", result, ErrorCodeMissingAction)
	}
}
