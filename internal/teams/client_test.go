package teams

import "testing"

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
