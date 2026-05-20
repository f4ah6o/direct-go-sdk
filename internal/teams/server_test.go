package teams

import "testing"

func TestParseReplyResource(t *testing.T) {
	teamID, channelID, rootID, replyID, ok := parseReplyResource("teams/team-1/channels/channel-1/messages/root-1/replies/reply-1")
	if !ok {
		t.Fatalf("expected parse success")
	}
	if teamID != "team-1" || channelID != "channel-1" || rootID != "root-1" || replyID != "reply-1" {
		t.Fatalf("unexpected parse: %s %s %s %s", teamID, channelID, rootID, replyID)
	}
}

func TestMentionsAndStripMentions(t *testing.T) {
	msg := &ChatMessage{
		Body: ItemBody{Content: `<at id="0">Bridge</at> hello<br>world`},
		Mentions: []Mention{{
			MentionText: "Bridge",
			Mentioned:   Mentioned{User: &UserIdentity{ID: "user-1"}},
		}},
	}
	if !MentionsUser(msg, "user-1") {
		t.Fatalf("expected mention")
	}
	if got := StripMentions(msg); got != "hello\nworld" {
		t.Fatalf("unexpected stripped text: %q", got)
	}
}
