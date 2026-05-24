package direct

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
)

const (
	// NoteContentTypeText is the Direct API content type for plain text notes.
	NoteContentTypeText = 1
	// NoteContentTypeRichText is the Direct API content type for rich-text notes.
	NoteContentTypeRichText = 13
)

var ErrUnsupportedNoteContentType = errors.New("unsupported note content type")

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

// CreateNoteInput describes a high-level note creation request.
type CreateNoteInput struct {
	Title         string
	Content       string
	ContentType   int
	CreateMessage *bool
}

// UpdateNoteInput describes a high-level note update request.
type UpdateNoteInput struct {
	Title                     string
	Content                   string
	ContentType               int
	NotifyTalkMembersOfUpdate *bool
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

// GetNoteDecompressed retrieves a note and expands rich_text_compressed content when present.
func (c *Client) GetNoteDecompressed(ctx context.Context, noteID interface{}) (*Note, error) {
	result, err := c.Call(MethodGetNote, []interface{}{noteID})
	if err != nil {
		return nil, err
	}
	decompressed, err := decompressNotePayload(result)
	if err != nil {
		return nil, err
	}
	return parseNote(decompressed), nil
}

// CreateNote creates a text or rich-text note.
func (c *Client) CreateNote(ctx context.Context, talkID interface{}, input CreateNoteInput) (*Note, error) {
	note, err := c.CreateNoteRaw(ctx, talkID, input)
	if err != nil {
		return nil, err
	}
	return decompressNote(note)
}

// CreateNoteRaw creates a text or rich-text note and preserves compressed response content.
func (c *Client) CreateNoteRaw(ctx context.Context, talkID interface{}, input CreateNoteInput) (*Note, error) {
	content, err := buildNoteContentPayload(input.ContentType, input.Content)
	if err != nil {
		return nil, err
	}
	createMessage := true
	if input.CreateMessage != nil {
		createMessage = *input.CreateMessage
	}
	result, err := c.Call(MethodCreateNote, []interface{}{talkID, input.Title, input.ContentType, content, createMessage})
	if err != nil {
		return nil, err
	}
	return parseNote(result), nil
}

// UpdateNote updates a text or rich-text note.
func (c *Client) UpdateNote(ctx context.Context, noteID, currentRevision interface{}, input UpdateNoteInput) (*Note, error) {
	note, err := c.UpdateNoteRaw(ctx, noteID, currentRevision, input)
	if err != nil {
		return nil, err
	}
	return decompressNote(note)
}

// UpdateNoteRaw updates a text or rich-text note and preserves compressed response content.
func (c *Client) UpdateNoteRaw(ctx context.Context, noteID, currentRevision interface{}, input UpdateNoteInput) (*Note, error) {
	content, err := buildNoteContentPayload(input.ContentType, input.Content)
	if err != nil {
		return nil, err
	}
	notifyTalkMembers := true
	if input.NotifyTalkMembersOfUpdate != nil {
		notifyTalkMembers = *input.NotifyTalkMembersOfUpdate
	}
	result, err := c.Call(MethodUpdateNote, []interface{}{noteID, currentRevision, input.Title, input.ContentType, content, notifyTalkMembers})
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

// CompressNoteRichText gzip-compresses rich-text note content using UTF-8 bytes.
func CompressNoteRichText(text string) ([]byte, error) {
	var buf bytes.Buffer
	zw := gzip.NewWriter(&buf)
	if _, err := zw.Write([]byte(text)); err != nil {
		_ = zw.Close()
		return nil, err
	}
	if err := zw.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecompressNoteRichText expands gzip-compressed rich-text note content.
func DecompressNoteRichText(data []byte) (string, error) {
	zr, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", err
	}
	plain, err := io.ReadAll(zr)
	if err != nil {
		_ = zr.Close()
		return "", err
	}
	if err := zr.Close(); err != nil {
		return "", err
	}
	return string(plain), nil
}

func buildNoteContentPayload(contentType int, content string) (interface{}, error) {
	switch contentType {
	case NoteContentTypeText:
		return content, nil
	case NoteContentTypeRichText:
		compressed, err := CompressNoteRichText(content)
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{
			"files":                []interface{}{},
			"rich_text_compressed": compressed,
		}, nil
	default:
		return nil, fmt.Errorf("%w: %d", ErrUnsupportedNoteContentType, contentType)
	}
}

func decompressNotePayload(v interface{}) (interface{}, error) {
	data, ok := v.(map[string]interface{})
	if !ok {
		return v, nil
	}
	revision, ok := data["note_revision"].(map[string]interface{})
	if !ok {
		return v, nil
	}
	content, ok := revision["content"].(map[string]interface{})
	if !ok {
		return v, nil
	}
	compressed, ok := bytesFromValue(content["rich_text_compressed"])
	if !ok {
		return v, nil
	}
	richText, err := DecompressNoteRichText(compressed)
	if err != nil {
		return nil, err
	}

	contentCopy := copyStringMap(content)
	delete(contentCopy, "rich_text_compressed")
	contentCopy["rich_text"] = richText

	revisionCopy := copyStringMap(revision)
	revisionCopy["content"] = contentCopy

	dataCopy := copyStringMap(data)
	dataCopy["note_revision"] = revisionCopy
	return dataCopy, nil
}

func bytesFromValue(v interface{}) ([]byte, bool) {
	switch data := v.(type) {
	case []byte:
		return data, true
	case string:
		// Some MessagePack decoders expose bin payloads as strings.
		return []byte(data), true
	default:
		return nil, false
	}
}

func decompressNote(note *Note) (*Note, error) {
	if note == nil {
		return nil, nil
	}
	decompressed, err := decompressNotePayload(note.Raw)
	if err != nil {
		return nil, err
	}
	return parseNote(decompressed), nil
}

func copyStringMap(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
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
