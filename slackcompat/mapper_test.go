package slackcompat

import (
	"testing"
	"time"
)

func TestMapperRoundTrip(t *testing.T) {
	mapper := NewMapper()
	channel := mapper.ChannelID("account-a", "talk-123")
	accountID, talkID, ok := mapper.TalkID(channel)
	if !ok {
		t.Fatalf("channel did not decode: %s", channel)
	}
	if accountID != "account-a" || talkID != "talk-123" {
		t.Fatalf("decoded %q %q, want account-a talk-123", accountID, talkID)
	}

	user := mapper.UserID("user-456")
	directUser, ok := mapper.DirectUserID(user)
	if !ok || directUser != "user-456" {
		t.Fatalf("decoded user %q ok=%v, want user-456 true", directUser, ok)
	}
}

func TestSlackTimestampFormat(t *testing.T) {
	got := formatSlackTS(time.Unix(1710000000, 123456000).UTC())
	if got != "1710000000.123456" {
		t.Fatalf("ts = %q", got)
	}
	parsed, ok := parseSlackTS(got)
	if !ok {
		t.Fatalf("failed to parse %q", got)
	}
	if !parsed.Equal(time.Unix(1710000000, 123456000).UTC()) {
		t.Fatalf("parsed = %s", parsed)
	}
}
