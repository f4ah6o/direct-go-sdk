package direct

import (
	"context"
	"errors"
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

func TestGetNoteDecompressed(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	compressed, err := CompressNoteRichText("<doc>こんにちは</doc>")
	if err != nil {
		t.Fatalf("CompressNoteRichText failed: %v", err)
	}
	mockServer.OnDynamic("get_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1")
		return map[string]interface{}{
			"note_id": "note1",
			"note_revision": map[string]interface{}{
				"content": map[string]interface{}{
					"files":                []interface{}{},
					"rich_text_compressed": compressed,
				},
			},
		}, nil
	})

	note, err := client.GetNoteDecompressed(context.Background(), "note1")
	if err != nil {
		t.Fatalf("GetNoteDecompressed failed: %v", err)
	}
	revision := note.Raw["note_revision"].(map[string]interface{})
	content := revision["content"].(map[string]interface{})
	if content["rich_text"] != "<doc>こんにちは</doc>" {
		t.Fatalf("unexpected rich_text: %+v", content)
	}
	if _, ok := content["rich_text_compressed"]; ok {
		t.Fatalf("rich_text_compressed should be removed: %+v", content)
	}
}

func TestNoteRichTextCompressionRoundTrip(t *testing.T) {
	text := "<doc><p>hello</p><p>こんにちは</p></doc>"
	compressed, err := CompressNoteRichText(text)
	if err != nil {
		t.Fatalf("CompressNoteRichText failed: %v", err)
	}
	if len(compressed) == 0 {
		t.Fatal("expected compressed bytes")
	}
	decompressed, err := DecompressNoteRichText(compressed)
	if err != nil {
		t.Fatalf("DecompressNoteRichText failed: %v", err)
	}
	if decompressed != text {
		t.Fatalf("round trip mismatch: got %q want %q", decompressed, text)
	}
}

func TestDecompressNoteRichTextInvalidGzip(t *testing.T) {
	if _, err := DecompressNoteRichText([]byte("not gzip")); err == nil {
		t.Fatal("expected invalid gzip error")
	}
}

func TestCreateNoteText(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("create_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "talk1", "Title", int8(NoteContentTypeText), "body", true)
		return map[string]interface{}{"note_id": "note1"}, nil
	})

	note, err := client.CreateNote(context.Background(), "talk1", CreateNoteInput{
		Title:       "Title",
		Content:     "body",
		ContentType: NoteContentTypeText,
	})
	if err != nil {
		t.Fatalf("CreateNote failed: %v", err)
	}
	if note.ID != "note1" {
		t.Fatalf("unexpected note: %+v", note)
	}
}

func TestCreateNoteRichText(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	createMessage := false
	mockServer.OnDynamic("create_note", func(params []interface{}) (interface{}, error) {
		if len(params) != 5 {
			t.Fatalf("expected 5 params, got %d: %#v", len(params), params)
		}
		if params[0] != "talk1" || params[1] != "Title" || params[2] != int8(NoteContentTypeRichText) || params[4] != false {
			t.Fatalf("unexpected params: %#v", params)
		}
		content, ok := params[3].(map[string]interface{})
		if !ok {
			t.Fatalf("expected content map, got %T", params[3])
		}
		if _, ok := content["rich_text"]; ok {
			t.Fatalf("rich_text should not be sent: %#v", content)
		}
		compressed, ok := content["rich_text_compressed"].([]byte)
		if !ok {
			t.Fatalf("expected compressed bytes: %#v", content["rich_text_compressed"])
		}
		decompressed, err := DecompressNoteRichText(compressed)
		if err != nil {
			t.Fatalf("DecompressNoteRichText failed: %v", err)
		}
		if decompressed != "<doc>rich</doc>" {
			t.Fatalf("unexpected decompressed text: %q", decompressed)
		}
		return map[string]interface{}{
			"note_id": "note1",
			"note_revision": map[string]interface{}{
				"content": map[string]interface{}{
					"rich_text_compressed": compressed,
				},
			},
		}, nil
	})

	note, err := client.CreateNote(context.Background(), "talk1", CreateNoteInput{
		Title:         "Title",
		Content:       "<doc>rich</doc>",
		ContentType:   NoteContentTypeRichText,
		CreateMessage: &createMessage,
	})
	if err != nil {
		t.Fatalf("CreateNote failed: %v", err)
	}
	content := note.Raw["note_revision"].(map[string]interface{})["content"].(map[string]interface{})
	if content["rich_text"] != "<doc>rich</doc>" {
		t.Fatalf("unexpected note content: %+v", content)
	}
}

func TestCreateNoteRawPreservesCompressedResponse(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("create_note", func(params []interface{}) (interface{}, error) {
		return map[string]interface{}{
			"note_id": "note1",
			"note_revision": map[string]interface{}{
				"content": map[string]interface{}{
					"rich_text_compressed": []byte("not gzip"),
				},
			},
		}, nil
	})

	note, err := client.CreateNoteRaw(context.Background(), "talk1", CreateNoteInput{
		Title:       "Title",
		Content:     "body",
		ContentType: NoteContentTypeText,
	})
	if err != nil {
		t.Fatalf("CreateNoteRaw failed: %v", err)
	}
	content := note.Raw["note_revision"].(map[string]interface{})["content"].(map[string]interface{})
	if string(content["rich_text_compressed"].([]byte)) != "not gzip" {
		t.Fatalf("unexpected raw content: %+v", content)
	}
}

func TestUpdateNoteText(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	notify := false
	mockServer.OnDynamic("update_note", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "note1", int64(3), "Title", int8(NoteContentTypeText), "body", false)
		return map[string]interface{}{"note_id": "note1"}, nil
	})

	note, err := client.UpdateNote(context.Background(), "note1", int64(3), UpdateNoteInput{
		Title:                     "Title",
		Content:                   "body",
		ContentType:               NoteContentTypeText,
		NotifyTalkMembersOfUpdate: &notify,
	})
	if err != nil {
		t.Fatalf("UpdateNote failed: %v", err)
	}
	if note.ID != "note1" {
		t.Fatalf("unexpected note: %+v", note)
	}
}

func TestUpdateNoteRichText(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("update_note", func(params []interface{}) (interface{}, error) {
		if len(params) != 6 {
			t.Fatalf("expected 6 params, got %d: %#v", len(params), params)
		}
		if params[0] != "note1" || params[1] != int64(3) || params[2] != "Title" || params[3] != int8(NoteContentTypeRichText) || params[5] != true {
			t.Fatalf("unexpected params: %#v", params)
		}
		content := params[4].(map[string]interface{})
		compressed := content["rich_text_compressed"].([]byte)
		decompressed, err := DecompressNoteRichText(compressed)
		if err != nil {
			t.Fatalf("DecompressNoteRichText failed: %v", err)
		}
		if decompressed != "<doc>updated</doc>" {
			t.Fatalf("unexpected decompressed text: %q", decompressed)
		}
		return map[string]interface{}{"note_id": "note1"}, nil
	})

	if _, err := client.UpdateNote(context.Background(), "note1", int64(3), UpdateNoteInput{
		Title:       "Title",
		Content:     "<doc>updated</doc>",
		ContentType: NoteContentTypeRichText,
	}); err != nil {
		t.Fatalf("UpdateNote failed: %v", err)
	}
}

func TestNoteUnsupportedContentType(t *testing.T) {
	client := &Client{}
	_, err := client.CreateNote(context.Background(), "talk1", CreateNoteInput{ContentType: 999})
	if !errors.Is(err, ErrUnsupportedNoteContentType) {
		t.Fatalf("expected ErrUnsupportedNoteContentType, got %v", err)
	}
}

func TestGetNoteDecompressedPropagatesGzipError(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("get_note", func(params []interface{}) (interface{}, error) {
		return map[string]interface{}{
			"note_id": "note1",
			"note_revision": map[string]interface{}{
				"content": map[string]interface{}{
					"rich_text_compressed": []byte("not gzip"),
				},
			},
		}, nil
	})

	if _, err := client.GetNoteDecompressed(context.Background(), "note1"); err == nil {
		t.Fatal("expected decompression error")
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
		{"create", "create_note", func(c *Client) error {
			_, err := c.CreateNote(context.Background(), "talk1", CreateNoteInput{ContentType: NoteContentTypeText})
			return err
		}},
		{"update", "update_note", func(c *Client) error {
			_, err := c.UpdateNote(context.Background(), "note1", int64(1), UpdateNoteInput{ContentType: NoteContentTypeText})
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
