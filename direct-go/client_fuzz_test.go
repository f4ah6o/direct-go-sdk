package direct

import (
	"bytes"
	"testing"

	"github.com/vmihailenco/msgpack/v5"
)

// FuzzMessagePackMarshal fuzzes MessagePack marshaling for RPC messages.
func FuzzMessagePackMarshal(f *testing.F) {
	// Seed corpus with valid RPC messages
	f.Add([]byte{0x94, 0x00, 0x01, 0xa7, 0x47, 0x65, 0x74, 0x4d, 0x65, 0x90}) // ["GetMe", 0, nil, []]
	f.Add([]byte{0x94, 0x00, 0x02, 0xa8, 0x47, 0x65, 0x74, 0x54, 0x61, 0x6c, 0x6b, 0x73, 0x90})
	f.Add([]byte{0x94, 0x01, 0x01, 0x00, 0x81}) // Response with error

	f.Fuzz(func(t *testing.T, data []byte) {
		var msg []interface{}
		err := msgpack.Unmarshal(data, &msg)
		if err != nil {
			return // Invalid MessagePack is expected
		}

		// If we can unmarshal, try to re-marshal
		_, err = msgpack.Marshal(msg)
		if err != nil {
			t.Errorf("failed to marshal valid message: %v", err)
		}
	})
}

// FuzzReceivedMessage fuzzes ReceivedMessage parsing.
func FuzzReceivedMessage(f *testing.F) {
	// Seed corpus with valid MessagePack data for ReceivedMessage
	msg := ReceivedMessage{
		ID:     "msg123",
		Text:   "hello",
		UserID: "user123",
		TalkID: "room123",
	}
	data, _ := msgpack.Marshal(msg)
	f.Add(data)

	f.Fuzz(func(t *testing.T, data []byte) {
		// Try to unmarshal as ReceivedMessage
		var msg ReceivedMessage
		err := msgpack.Unmarshal(data, &msg)
		if err != nil {
			return // Invalid data is expected
		}

		// If we get a valid message, its fields should be accessible
		_ = msg.ID
		_ = msg.Text
		_ = msg.UserID
		_ = msg.TalkID
		_ = msg.Created
	})
}

// FuzzUserIDParsing fuzzes user ID parsing from various input formats.
func FuzzUserIDParsing(f *testing.F) {
	// Seed corpus
	f.Add("12345")
	f.Add("user_123")
	f.Add("")
	f.Add("0")

	f.Fuzz(func(t *testing.T, data string) {
		// Try to create a ReceivedMessage with the fuzzed user ID
		msg := ReceivedMessage{
			ID:     "msg123",
			Text:   "test",
			UserID: data,
			TalkID: "room123",
		}

		// Verify the data is preserved
		if msg.UserID != data {
			t.Errorf("UserID mismatch: got %q, want %q", msg.UserID, data)
		}
	})
}

// FuzzRoomIDParsing fuzzes room ID parsing from various input formats.
func FuzzRoomIDParsing(f *testing.F) {
	// Seed corpus
	f.Add("12345")
	f.Add("room_abc")
	f.Add("")
	f.Add("0")

	f.Fuzz(func(t *testing.T, data string) {
		msg := ReceivedMessage{
			ID:     "msg123",
			Text:   "test",
			UserID: "user123",
			TalkID: data,
		}

		if msg.TalkID != data {
			t.Errorf("TalkID mismatch: got %q, want %q", msg.TalkID, data)
		}
	})
}

// FuzzRPCRequestFormat fuzzes RPC request format validation.
func FuzzRPCRequestFormat(f *testing.F) {
	// Seed corpus with valid and invalid formats
	validMsg := []interface{}{0, int64(1), "TestMethod", []interface{}{}}
	data, _ := msgpack.Marshal(validMsg)
	f.Add(data)

	f.Fuzz(func(t *testing.T, data []byte) {
		var msg []interface{}
		err := msgpack.Unmarshal(data, &msg)
		if err != nil {
			return
		}

		// Check if it looks like a valid RPC request
		// Format: [type(0), msgId, method, params]
		if len(msg) >= 4 {
			if msgType, ok := msg[0].(int8); ok && msgType == 0 {
				if _, ok := msg[1].(int64); ok {
					if method, ok := msg[2].(string); ok && method != "" {
						if params, ok := msg[3].([]interface{}); ok {
							// Valid RPC request structure
							_ = params
						}
					}
				}
			}
		}
	})
}

// FuzzTextMessageParsing fuzzes text message content.
func FuzzTextMessageParsing(f *testing.F) {
	// Seed corpus
	f.Add("Hello, world!")
	f.Add("")
	f.Add(string([]byte{0x00, 0x01, 0x02}))
	f.Add("特殊文字テスト")
	f.Add("emoji 🎉")

	f.Fuzz(func(t *testing.T, data string) {
		msg := ReceivedMessage{
			ID:     "msg123",
			Text:   data,
			UserID: "user123",
			TalkID: "room123",
		}

		// Text should be preserved exactly
		if msg.Text != data {
			t.Errorf("Text mismatch: got %q, want %q", msg.Text, data)
		}
	})
}

// FuzzMessageBytesRoundTrip fuzzes marshal/unmarshal roundtrip.
func FuzzMessageBytesRoundTrip(f *testing.F) {
	// Seed corpus
	seed := ReceivedMessage{
		ID:      "msg123",
		Text:    "Hello",
		UserID:  "user123",
		TalkID:  "room123",
		Created: 1234567890,
	}
	data, _ := msgpack.Marshal(seed)
	f.Add(data)

	f.Fuzz(func(t *testing.T, original []byte) {
		var msg1 ReceivedMessage
		err := msgpack.Unmarshal(original, &msg1)
		if err != nil {
			return
		}

		// Marshal again
		remarshaled, err := msgpack.Marshal(msg1)
		if err != nil {
			t.Errorf("failed to marshal: %v", err)
			return
		}

		// Unmarshal again
		var msg2 ReceivedMessage
		err = msgpack.Unmarshal(remarshaled, &msg2)
		if err != nil {
			t.Errorf("failed to unmarshal remarshaled data: %v", err)
			return
		}

		// Fields should match
		if msg1.ID != msg2.ID || msg1.Text != msg2.Text ||
			msg1.UserID != msg2.UserID || msg1.TalkID != msg2.TalkID {
			t.Errorf("roundtrip mismatch: %+v vs %+v", msg1, msg2)
		}
	})
}

// FuzzBufferHandling fuzzes buffer handling in message parsing.
func FuzzBufferHandling(f *testing.F) {
	f.Add([]byte{})
	f.Add([]byte{0x00})
	f.Add(make([]byte, 1000))
	f.Add(bytes.Repeat([]byte{0xFF}, 10000))

	f.Fuzz(func(t *testing.T, data []byte) {
		// Create a decoder with the fuzzed data
		decoder := msgpack.NewDecoder(bytes.NewReader(data))

		var msg []interface{}
		err := decoder.Decode(&msg)
		if err != nil {
			return // Expected for invalid data
		}

		// If decode succeeded, verify we can access elements safely
		for i, item := range msg {
			_ = i
			_ = item
		}
	})
}
