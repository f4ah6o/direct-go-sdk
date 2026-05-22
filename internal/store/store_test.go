package store

import (
	"path/filepath"
	"regexp"
	"testing"
)

func TestStorePersistsMappingsAndIndexesThread(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	st, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	m := ThreadMapping{AccountID: "a", TalkID: "t", ConversationID: "conversation", RootID: "root"}
	if err := st.PutMapping(m); err != nil {
		t.Fatal(err)
	}
	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := reopened.GetByThread("conversation", "root")
	if !ok {
		t.Fatalf("thread index missing")
	}
	if got.AccountID != "a" || got.TalkID != "t" {
		t.Fatalf("unexpected mapping: %+v", got)
	}
	if err := reopened.Forget("a", "t"); err != nil {
		t.Fatal(err)
	}
	if _, ok := reopened.GetByThread("conversation", "root"); ok {
		t.Fatalf("thread index remained after forget")
	}
}

func TestStorePersistsChannelBindings(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	st, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutChannelBinding(TeamsChannelBinding{Alias: "support", ConversationID: "conv", ServiceURL: "https://service"}); err != nil {
		t.Fatal(err)
	}
	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := reopened.GetChannelBinding("support")
	if !ok || got.ConversationID != "conv" {
		t.Fatalf("unexpected binding: %+v %v", got, ok)
	}
}

func TestStoreDirectDevicePersistsAndResets(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	st, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	first, err := st.EnsureDirectDevice("bot-trial")
	if err != nil {
		t.Fatal(err)
	}
	second, err := st.EnsureDirectDevice("bot-trial")
	if err != nil {
		t.Fatal(err)
	}
	if first == "" || second != first {
		t.Fatalf("device id was not stable: first=%q second=%q", first, second)
	}
	reset, err := st.ResetDirectDevice("bot-trial")
	if err != nil {
		t.Fatal(err)
	}
	if reset == "" || reset == first {
		t.Fatalf("device id was not reset: first=%q reset=%q", first, reset)
	}
	if !regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$`).MatchString(reset) {
		t.Fatalf("device id is not UUID-like: %q", reset)
	}
}

func TestStoreUnbindChannelRemovesBindingAndMatchingMappings(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	st, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutChannelBinding(TeamsChannelBinding{Alias: "trial", ConversationID: "conv-a", ServiceURL: "https://service"}); err != nil {
		t.Fatal(err)
	}
	if err := st.PutMapping(ThreadMapping{AccountID: "a", TalkID: "1", ChannelAlias: "trial", ConversationID: "conv-a", RootID: "root-a"}); err != nil {
		t.Fatal(err)
	}
	if err := st.PutMapping(ThreadMapping{AccountID: "b", TalkID: "2", ChannelAlias: "trial", ConversationID: "conv-b", RootID: "root-b"}); err != nil {
		t.Fatal(err)
	}
	removed, err := st.UnbindChannel("trial", "conv-a")
	if err != nil {
		t.Fatal(err)
	}
	if removed != 1 {
		t.Fatalf("removed = %d, want 1", removed)
	}
	if _, ok := st.GetChannelBinding("trial"); ok {
		t.Fatalf("binding remained after unbind")
	}
	if _, ok := st.GetByThread("conv-a", "root-a"); ok {
		t.Fatalf("matching mapping remained after unbind")
	}
	if _, ok := st.GetByThread("conv-b", "root-b"); !ok {
		t.Fatalf("mapping for another conversation was removed")
	}
}
