package slackcompat

import (
	"testing"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

func TestConvertMessageEvent(t *testing.T) {
	mapper := NewMapper()
	msg := direct.ReceivedMessage{
		ID:        "msg-1",
		TalkID:    "talk-1",
		UserID:    "user-1",
		Text:      "hello",
		Timestamp: time.Unix(1710000000, 0).UTC(),
	}
	event := ConvertMessageEvent(mapper, "Tdirect", "account-a", msg)
	if event.Type != "event_callback" {
		t.Fatalf("type = %q", event.Type)
	}
	if event.TeamID != "Tdirect" {
		t.Fatalf("team_id = %q", event.TeamID)
	}
	if event.Event.Type != "message" || event.Event.Text != "hello" {
		t.Fatalf("event = %+v", event.Event)
	}
	if event.Event.Channel != mapper.ChannelID("account-a", "talk-1") {
		t.Fatalf("channel = %q", event.Event.Channel)
	}
	if event.Event.User != mapper.UserID("user-1") {
		t.Fatalf("user = %q", event.Event.User)
	}
	if event.Event.TS != "1710000000.000000" {
		t.Fatalf("ts = %q", event.Event.TS)
	}
}
