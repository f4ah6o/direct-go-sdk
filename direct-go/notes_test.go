package direct

import (
	"context"
	"testing"
)

func TestGetNoteStatuses(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("get_note_statuses", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "domain1", "talk1", int8(60), "marker1")
		return map[string]interface{}{
			"marker":      "marker1",
			"next_marker": "marker2",
			"contents": []interface{}{
				map[string]interface{}{"note_id": "note1", "title": "Note 1"},
			},
		}, nil
	})

	statuses, err := client.GetNoteStatuses(context.Background(), "domain1", "talk1", 60, "marker1")
	if err != nil {
		t.Fatalf("GetNoteStatuses failed: %v", err)
	}
	if statuses.NextMarker != "marker2" || len(statuses.Contents) != 1 || statuses.Contents[0].NoteID != "note1" {
		t.Fatalf("unexpected statuses: %+v", statuses)
	}
}

func TestGetNote(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("get_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1")
		return map[string]interface{}{
			"note_id": "note1",
			"note_revision": map[string]interface{}{
				"rich_text_compressed": []byte("compressed"),
			},
		}, nil
	})

	note, err := client.GetNote(context.Background(), "note1")
	if err != nil {
		t.Fatalf("GetNote failed: %v", err)
	}
	if note.ID != "note1" || note.Raw["note_revision"] == nil {
		t.Fatalf("unexpected note: %+v", note)
	}
}

func TestUpdateNoteSetting(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	setting := map[string]interface{}{"editable": false}
	mockServer.OnDynamic("update_note_setting", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1", setting, int64(3))
		return map[string]interface{}{"note_id": "note1", "setting": setting}, nil
	})

	note, err := client.UpdateNoteSetting(context.Background(), "note1", setting, int64(3))
	if err != nil {
		t.Fatalf("UpdateNoteSetting failed: %v", err)
	}
	if note.ID != "note1" || note.Raw["setting"] == nil {
		t.Fatalf("unexpected note: %+v", note)
	}
}

func TestDeleteNote(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("delete_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1")
		return map[string]interface{}{"note_id": "note1", "deleted": true}, nil
	})

	deleted, err := client.DeleteNote(context.Background(), "note1")
	if err != nil {
		t.Fatalf("DeleteNote failed: %v", err)
	}
	if deleted.ID != "note1" || deleted.Raw["deleted"] != true {
		t.Fatalf("unexpected deleted note: %+v", deleted)
	}
}

func TestLockNote(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("lock_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1", int64(3))
		return true, nil
	})

	if err := client.LockNote(context.Background(), "note1", int64(3)); err != nil {
		t.Fatalf("LockNote failed: %v", err)
	}
}

func TestUnlockNote(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("unlock_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1", int64(3))
		return true, nil
	})

	if err := client.UnlockNote(context.Background(), "note1", int64(3)); err != nil {
		t.Fatalf("UnlockNote failed: %v", err)
	}
}

func TestNoteMethodsErrorPropagation(t *testing.T) {
	tests := []struct {
		name   string
		method string
		call   func(*Client) error
	}{
		{"statuses", "get_note_statuses", func(c *Client) error {
			_, err := c.GetNoteStatuses(context.Background(), "domain1", "talk1", 60, nil)
			return err
		}},
		{"get", "get_note", func(c *Client) error {
			_, err := c.GetNote(context.Background(), "note1")
			return err
		}},
		{"update setting", "update_note_setting", func(c *Client) error {
			_, err := c.UpdateNoteSetting(context.Background(), "note1", map[string]interface{}{}, int64(1))
			return err
		}},
		{"delete", "delete_note", func(c *Client) error {
			_, err := c.DeleteNote(context.Background(), "note1")
			return err
		}},
		{"lock", "lock_note", func(c *Client) error {
			return c.LockNote(context.Background(), "note1", int64(1))
		}},
		{"unlock", "unlock_note", func(c *Client) error {
			return c.UnlockNote(context.Background(), "note1", int64(1))
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer, client := newConnectedMockClient(t)
			mockServer.OnError(tt.method, "boom")
			if err := tt.call(client); err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}
