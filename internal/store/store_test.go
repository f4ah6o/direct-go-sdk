package store

import (
	"path/filepath"
	"testing"
)

func TestStorePersistsMappingsAndIndexesThread(t *testing.T) {
	path := filepath.Join(t.TempDir(), "state.json")
	st, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	m := ThreadMapping{AccountID: "a", TalkID: "t", TeamID: "team", ChannelID: "channel", RootID: "root"}
	if err := st.PutMapping(m); err != nil {
		t.Fatal(err)
	}
	reopened, err := Open(path)
	if err != nil {
		t.Fatal(err)
	}
	got, ok := reopened.GetByThread("team", "channel", "root")
	if !ok {
		t.Fatalf("thread index missing")
	}
	if got.AccountID != "a" || got.TalkID != "t" {
		t.Fatalf("unexpected mapping: %+v", got)
	}
	if err := reopened.Forget("a", "t"); err != nil {
		t.Fatal(err)
	}
	if _, ok := reopened.GetByThread("team", "channel", "root"); ok {
		t.Fatalf("thread index remained after forget")
	}
}
