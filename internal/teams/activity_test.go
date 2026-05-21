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

func TestChannelConversationIDPrefersChannelID(t *testing.T) {
	a := Activity{
		Conversation: ConversationAccount{ID: "19:channel@thread.tacv2;messageid=123"},
		ChannelData:  ChannelData{Channel: ChannelInfo{ID: "19:channel@thread.tacv2"}},
	}
	if got := channelConversationID(a); got != "19:channel@thread.tacv2" {
		t.Fatalf("channelConversationID() = %q", got)
	}
}

func TestThreadReferenceUsesConversationMessageID(t *testing.T) {
	a := Activity{
		Conversation: ConversationAccount{ID: "19:channel@thread.tacv2;messageid=123"},
		ChannelData:  ChannelData{Channel: ChannelInfo{ID: "19:channel@thread.tacv2"}},
	}
	conversationID, rootID := threadReference(a)
	if conversationID != "19:channel@thread.tacv2" || rootID != "123" {
		t.Fatalf("threadReference() = %q, %q", conversationID, rootID)
	}
}

func TestAttachmentDownloadURL(t *testing.T) {
	a := Attachment{Content: map[string]interface{}{"downloadUrl": "https://example.com/file.png"}}
	if got := a.DownloadURL(); got != "https://example.com/file.png" {
		t.Fatalf("DownloadURL() = %q", got)
	}
}

func TestAppendTeamsSenderName(t *testing.T) {
	if got := appendTeamsSenderName("了解です", "Taro Yamada"); got != "了解です（Taro Yamada）" {
		t.Fatalf("appendTeamsSenderName() = %q", got)
	}
}

func TestAppendTeamsSenderNameJapaneseName(t *testing.T) {
	if got := appendTeamsSenderName("了解です", "山田 太郎"); got != "了解です（山田 太郎）" {
		t.Fatalf("appendTeamsSenderName() = %q", got)
	}
}
