package teams

import "testing"

func TestParseBindAlias(t *testing.T) {
	a := Activity{
		Text:      `<at>Bridge</at> bind support`,
		Recipient: ChannelAccount{ID: "bot"},
		Entities:  []Entity{{Type: "mention", Text: "Bridge", Mentioned: ChannelAccount{ID: "bot"}}},
	}
	alias, ok := ParseBindAlias(a)
	if !ok || alias != "support" {
		t.Fatalf("ParseBindAlias() = %q, %v", alias, ok)
	}
}

func TestStripRecipientMention(t *testing.T) {
	a := Activity{
		Text:      `<at>Bridge</at> hello<br>world`,
		Recipient: ChannelAccount{ID: "bot"},
		Entities:  []Entity{{Type: "mention", Text: "Bridge", Mentioned: ChannelAccount{ID: "bot"}}},
	}
	if got := StripRecipientMention(a); got != "hello\nworld" {
		t.Fatalf("StripRecipientMention() = %q", got)
	}
}

func TestParseCommand(t *testing.T) {
	a := Activity{
		Text:      `<at>direct</at> Hi`,
		Recipient: ChannelAccount{ID: "bot"},
		Entities:  []Entity{{Type: "mention", Text: "direct", Mentioned: ChannelAccount{ID: "bot"}}},
	}
	if got := ParseCommand(a); got != "hi" {
		t.Fatalf("ParseCommand() = %q", got)
	}
}

func TestBotWasAdded(t *testing.T) {
	a := Activity{
		Recipient:    ChannelAccount{ID: "bot"},
		MembersAdded: []ChannelAccount{{ID: "user"}, {ID: "bot"}},
	}
	if !BotWasAdded(a) {
		t.Fatalf("expected bot to be detected in membersAdded")
	}
}
