package direct

import "context"

// Note contains a raw note payload returned by the Direct API.
type Note struct {
	ID  interface{}
	Raw map[string]interface{}
}

// NoteStatus contains a raw note status payload.
type NoteStatus struct {
	NoteID interface{}
	Raw    map[string]interface{}
}

// NoteStatusesResult contains note status list metadata and payloads.
type NoteStatusesResult struct {
	Marker     interface{}
	NextMarker interface{}
	Contents   []NoteStatus
	Raw        map[string]interface{}
}

// DeletedNote contains the raw delete-note result payload.
type DeletedNote struct {
	ID  interface{}
	Raw map[string]interface{}
}

// GetNoteStatuses retrieves note statuses for a talk.
func (c *Client) GetNoteStatuses(ctx context.Context, domainID, talkID interface{}, limit int, marker interface{}) (*NoteStatusesResult, error) {
	result, err := c.Call(MethodGetNoteStatuses, []interface{}{domainID, talkID, limit, marker})
	if err != nil {
		return nil, err
	}
	return parseNoteStatusesResult(result), nil
}

// GetNote retrieves a note by ID. Compressed content is preserved as returned by the API.
func (c *Client) GetNote(ctx context.Context, noteID interface{}) (*Note, error) {
	result, err := c.Call(MethodGetNote, []interface{}{noteID})
	if err != nil {
		return nil, err
	}
	return parseNote(result), nil
}

// UpdateNoteSetting updates note settings by note ID and version.
func (c *Client) UpdateNoteSetting(ctx context.Context, noteID, setting, version interface{}) (*Note, error) {
	result, err := c.Call(MethodUpdateNoteSetting, []interface{}{noteID, setting, version})
	if err != nil {
		return nil, err
	}
	return parseNote(result), nil
}

// DeleteNote deletes a note by ID.
func (c *Client) DeleteNote(ctx context.Context, noteID interface{}) (*DeletedNote, error) {
	result, err := c.Call(MethodDeleteNote, []interface{}{noteID})
	if err != nil {
		return nil, err
	}
	deleted := &DeletedNote{}
	if data, ok := result.(map[string]interface{}); ok {
		deleted.ID = firstMapValue(data, "note_id", "id")
		deleted.Raw = data
	}
	return deleted, nil
}

// LockNote locks a note by ID and version.
func (c *Client) LockNote(ctx context.Context, noteID, version interface{}) error {
	_, err := c.Call(MethodLockNote, []interface{}{noteID, version})
	return err
}

// UnlockNote unlocks a note by ID and version.
func (c *Client) UnlockNote(ctx context.Context, noteID, version interface{}) error {
	_, err := c.Call(MethodUnlockNote, []interface{}{noteID, version})
	return err
}

func parseNoteStatusesResult(v interface{}) *NoteStatusesResult {
	result := &NoteStatusesResult{Contents: []NoteStatus{}}
	data, ok := v.(map[string]interface{})
	if !ok {
		return result
	}
	result.Raw = data
	result.Marker = data["marker"]
	result.NextMarker = firstMapValue(data, "next_marker", "nextMarker")
	for _, item := range mapSliceFromValue(data["contents"]) {
		result.Contents = append(result.Contents, NoteStatus{
			NoteID: firstMapValue(item, "note_id", "id"),
			Raw:    item,
		})
	}
	return result
}

func parseNote(v interface{}) *Note {
	note := &Note{}
	if data, ok := v.(map[string]interface{}); ok {
		note.ID = firstMapValue(data, "note_id", "id")
		note.Raw = data
	}
	return note
}

func firstMapValue(m map[string]interface{}, keys ...string) interface{} {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return v
		}
	}
	return nil
}
