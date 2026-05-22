package bridge

import (
	"log"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

func TestConsumePendingDirectMessageSuppressesTeamsToDirectEcho(t *testing.T) {
	s := NewService(nil, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	s.markPendingDirectMessage(outbound)

	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		UserID:    "1792959268018716672",
		Text:      "了解です（Taro Yamada）",
		MessageID: "direct-message-id",
	}
	if !s.consumePendingDirectMessage(inbound) {
		t.Fatalf("expected pending direct message to be consumed")
	}
	if s.consumePendingDirectMessage(inbound) {
		t.Fatalf("pending direct message should only be consumed once")
	}
}

func TestClearPendingDirectMessageAllowsLaterUserMessage(t *testing.T) {
	s := NewService(nil, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	s.markPendingDirectMessage(outbound)
	s.clearPendingDirectMessage(outbound)

	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	if s.consumePendingDirectMessage(inbound) {
		t.Fatalf("cleared pending direct message should not be consumed")
	}
}

func TestSuccessfulDirectSentKeepsPendingUntilDirectNotification(t *testing.T) {
	s := NewService(nil, nil, nil, nil, nil, nil, nil, log.Default())
	outbound := model.DirectOutbound{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "了解です（Taro Yamada）",
	}
	s.markPendingDirectMessage(outbound)

	inbound := model.DirectMessage{
		AccountID: "bot-trial",
		TalkID:    "1792967566075891712",
		Text:      "  了解です（Taro Yamada）\r\n",
		MessageID: "direct-message-id",
	}
	if !s.consumePendingDirectMessage(inbound) {
		t.Fatalf("successful send pending marker should remain until direct notification is consumed")
	}
}
