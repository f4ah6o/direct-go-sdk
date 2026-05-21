package teams

import (
	"strings"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

func TestTeamsThreadConversationID(t *testing.T) {
	got := teamsThreadConversationID("19:channel@thread.tacv2", "123")
	want := "19:channel@thread.tacv2;messageid=123"
	if got != want {
		t.Fatalf("teamsThreadConversationID() = %q, want %q", got, want)
	}
}

func TestTeamsThreadConversationIDAlreadyThread(t *testing.T) {
	id := "19:channel@thread.tacv2;messageid=123"
	if got := teamsThreadConversationID(id, "456"); got != id {
		t.Fatalf("teamsThreadConversationID() = %q, want %q", got, id)
	}
}

func TestFormatDirectRootAndReplyMessages(t *testing.T) {
	client := NewClient(config.BotConfig{}, "https://bridge.example.com")
	msg := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		UserID:    "1792959268018716672",
		Text:      "こんにちは",
		Attachments: []model.Attachment{{
			Name: "image.png",
			URL:  "https://api.direct4b.com/albero-app-server/files/file/token?message_id=1",
		}},
	}

	root := client.formatDirectRootMessage(msg)
	if !strings.HasPrefix(root, "[direct:bot-trial] room=1792967566075891712\nuser=1792959268018716672\n") {
		t.Fatalf("unexpected root message: %q", root)
	}
	if !strings.Contains(root, "[attachment: image.png](https://bridge.example.com/files/direct?account=bot-trial&url=") {
		t.Fatalf("expected proxied markdown link: %q", root)
	}

	reply := client.formatDirectReplyMessage(msg)
	if strings.Contains(reply, "[direct:") || strings.Contains(reply, "room=") {
		t.Fatalf("reply should not repeat thread title details: %q", reply)
	}
	if !strings.HasPrefix(reply, "user=1792959268018716672\nこんにちは") {
		t.Fatalf("unexpected reply message: %q", reply)
	}
}
