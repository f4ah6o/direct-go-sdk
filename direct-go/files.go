package direct

import (
	"context"
	"time"
)

// UploadAuth represents authentication credentials for file upload.
type UploadAuth struct {
	FileID   interface{}
	PostURL  string
	PostForm map[string]string
	PutURL   string
}

// Attachment represents a file attachment.
type Attachment struct {
	ID          interface{}
	MessageID   interface{}
	TalkID      interface{}
	FileID      interface{}
	Name        string
	ContentType string
	ContentSize int64
	URL         string
	CreatedAt   time.Time
}

// FilePreview represents a preview of a file.
type FilePreview struct {
	FileID            interface{}
	Status            string
	FilePreviewFileID interface{}
	URL               string
	Key               string
}

// CreateUploadAuth creates authentication credentials for uploading a file.
// Returns UploadAuth with a file ID and either a POST URL with form data or a PUT URL.
// The useType parameter specifies how the file will be used (e.g., "message", "profile").
func (c *Client) CreateUploadAuth(ctx context.Context, filename, contentType string, size int64, useType string) (*UploadAuth, error) {
	params := []interface{}{filename, contentType, size, 0, useType}
	result, err := c.Call(MethodCreateUploadAuth, params)
	if err != nil {
		return nil, err
	}

	auth := &UploadAuth{}
	if authData, ok := result.(map[string]interface{}); ok {
		if v, ok := authData["file_id"]; ok {
			auth.FileID = v
		}
		if v, ok := authData["post_url"].(string); ok {
			auth.PostURL = v
		}
		if v, ok := authData["put_url"].(string); ok {
			auth.PutURL = v
		}
		if v, ok := authData["post_form"].(map[string]interface{}); ok {
			auth.PostForm = make(map[string]string)
			for k, val := range v {
				if str, ok := val.(string); ok {
					auth.PostForm[k] = str
				}
			}
		}
	}

	return auth, nil
}

// GetAttachments retrieves file attachments from a talk/conversation.
// The limit parameter controls how many attachments to return (most recent first).
func (c *Client) GetAttachments(ctx context.Context, talkID interface{}, limit int) ([]Attachment, error) {
	params := []interface{}{talkID, limit}
	result, err := c.Call(MethodGetAttachments, params)
	if err != nil {
		return nil, err
	}

	attachments := []Attachment{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if attachData, ok := item.(map[string]interface{}); ok {
				attachment := parseAttachment(attachData)
				attachments = append(attachments, attachment)
			}
		}
	}

	return attachments, nil
}

// DeleteAttachment removes a file attachment from the system.
func (c *Client) DeleteAttachment(ctx context.Context, attachmentID interface{}) error {
	params := []interface{}{attachmentID}
	_, err := c.Call(MethodDeleteAttachment, params)
	return err
}

// SearchAttachments searches for file attachments within a talk by filename or content.
// Returns matching Attachment objects with file metadata and download URLs.
func (c *Client) SearchAttachments(ctx context.Context, query string, talkID interface{}) ([]Attachment, error) {
	params := []interface{}{query, talkID}
	result, err := c.Call(MethodSearchAttachments, params)
	if err != nil {
		return nil, err
	}

	attachments := []Attachment{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if attachData, ok := item.(map[string]interface{}); ok {
				attachment := parseAttachment(attachData)
				attachments = append(attachments, attachment)
			}
		}
	}

	return attachments, nil
}

// CreateFilePreview generates a preview (thumbnail/image) for a file.
// This is useful for displaying image or document previews in the UI.
func (c *Client) CreateFilePreview(ctx context.Context, fileID interface{}) (*FilePreview, error) {
	params := []interface{}{fileID}
	result, err := c.Call(MethodCreateFilePreview, params)
	if err != nil {
		return nil, err
	}

	if previewData, ok := result.(map[string]interface{}); ok {
		return parseFilePreview(previewData), nil
	}

	return nil, nil
}

// GetFilePreview retrieves an existing file preview if one has been generated.
// Returns FilePreview with the preview URL and status.
func (c *Client) GetFilePreview(ctx context.Context, fileID interface{}) (*FilePreview, error) {
	params := []interface{}{fileID}
	result, err := c.Call(MethodGetFilePreview, params)
	if err != nil {
		return nil, err
	}

	if previewData, ok := result.(map[string]interface{}); ok {
		return parseFilePreview(previewData), nil
	}

	return nil, nil
}

// Helper functions

func parseAttachment(data map[string]interface{}) Attachment {
	attachment := Attachment{}

	if v, ok := data["id"]; ok {
		attachment.ID = v
	}
	if v, ok := data["message_id"]; ok {
		attachment.MessageID = v
	}
	if v, ok := data["talk_id"]; ok {
		attachment.TalkID = v
	}
	if v, ok := data["file_id"]; ok {
		attachment.FileID = v
	}
	if v, ok := data["name"].(string); ok {
		attachment.Name = v
	}
	if v, ok := data["content_type"].(string); ok {
		attachment.ContentType = v
	}
	if v, ok := data["content_size"].(int64); ok {
		attachment.ContentSize = v
	}
	if v, ok := data["url"].(string); ok {
		attachment.URL = v
	}
	if v, ok := data["created_at"].(int64); ok {
		attachment.CreatedAt = time.Unix(v, 0)
	}

	return attachment
}

func parseFilePreview(data map[string]interface{}) *FilePreview {
	preview := &FilePreview{}

	if v, ok := data["file_id"]; ok {
		preview.FileID = v
	}
	if v, ok := data["status"].(string); ok {
		preview.Status = v
	}
	if v, ok := data["file_preview_file_id"]; ok {
		preview.FilePreviewFileID = v
	}
	if v, ok := data["url"].(string); ok {
		preview.URL = v
	}
	if v, ok := data["key"].(string); ok {
		preview.Key = v
	}

	return preview
}
