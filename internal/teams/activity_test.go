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
