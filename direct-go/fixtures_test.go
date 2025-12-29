package direct

import (
	"testing"
	"time"
)

// NewTestMessage creates a test ReceivedMessage with common fields.
func NewTestMessage(id, text, userID, talkID string) ReceivedMessage {
	return ReceivedMessage{
		ID:     id,
		Text:   text,
		UserID: userID,
		TalkID: talkID,
		Created: time.Now().Unix(),
	}
}

// NewTestUser creates a test User.
func NewTestUser(id, name string) *User {
	return &User{
		ID:   id,
		Name: name,
	}
}

// NewTestRoom creates a test Room.
func NewTestRoom(id interface{}, name string) *Room {
	return &Room{
		ID:   id,
		Name: name,
	}
}

// WaitForMessage waits for a message on a channel with timeout.
// Returns the message or nil if timeout occurs.
func WaitForMessage[T any](t *testing.T, ch <-chan T, timeout time.Duration) *T {
	t.Helper()
	timeoutCh := time.After(timeout)
	select {
	case msg := <-ch:
		return &msg
	case <-timeoutCh:
		t.Fatal("timeout waiting for message")
		return nil
	}
}

// AssertNoError fails the test if err is not nil.
func AssertNoError(t *testing.T, err error, msg ...string) {
	t.Helper()
	if err != nil {
		if len(msg) > 0 {
			t.Fatalf("%s: %v", msg[0], err)
		} else {
			t.Fatalf("unexpected error: %v", err)
		}
	}
}

// AssertError fails the test if err is nil.
func AssertError(t *testing.T, err error, msg ...string) {
	t.Helper()
	if err == nil {
		if len(msg) > 0 {
			t.Fatalf("%s: expected error, got nil", msg[0])
		} else {
			t.Fatal("expected error, got nil")
		}
	}
}

// AssertEqual fails the test if want and got are not equal.
func AssertEqual[T comparable](t *testing.T, want, got T, msg ...string) {
	t.Helper()
	if want != got {
		if len(msg) > 0 {
			t.Fatalf("%s: want %v, got %v", msg[0], want, got)
		} else {
			t.Fatalf("want %v, got %v", want, got)
		}
	}
}
