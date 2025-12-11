package direct

import (
	"context"
	"time"
)

// Announcement represents an announcement message.
type Announcement struct {
	ID              interface{}
	DomainID        interface{}
	Title           string
	Text            string
	CreatedBy       interface{}
	CreatedAt       time.Time
	UpdatedAt       time.Time
	TargetUserIDs   []interface{}
	ReadUserIDs     []interface{}
	UnreadUserCount int
}

// AnnouncementStatus represents the read status of announcements.
type AnnouncementStatus struct {
	DomainID              interface{}
	UnreadCount           int
	MaxAnnouncementID     interface{}
	MaxReadAnnouncementID interface{}
}

// CreateAnnouncement creates a new announcement.
func (c *Client) CreateAnnouncement(ctx context.Context, domainID interface{}, title, text string, targetUserIDs []interface{}) (*Announcement, error) {
	params := []interface{}{domainID, title, text, targetUserIDs}
	result, err := c.Call(MethodCreateAnnouncement, params)
	if err != nil {
		return nil, err
	}

	if announcementData, ok := result.(map[string]interface{}); ok {
		return parseAnnouncement(announcementData), nil
	}

	return nil, nil
}

// GetAnnouncements retrieves announcements for a domain.
func (c *Client) GetAnnouncements(ctx context.Context, domainID interface{}) ([]Announcement, error) {
	params := []interface{}{domainID}
	result, err := c.Call(MethodGetAnnouncements, params)
	if err != nil {
		return nil, err
	}

	announcements := []Announcement{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if announcementData, ok := item.(map[string]interface{}); ok {
				announcement := parseAnnouncement(announcementData)
				announcements = append(announcements, *announcement)
			}
		}
	}

	return announcements, nil
}

// GetAnnouncementStatuses retrieves announcement statuses.
func (c *Client) GetAnnouncementStatuses(ctx context.Context) ([]AnnouncementStatus, error) {
	result, err := c.Call(MethodGetAnnouncementStatuses, []interface{}{})
	if err != nil {
		return nil, err
	}

	statuses := []AnnouncementStatus{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if statusData, ok := item.(map[string]interface{}); ok {
				status := AnnouncementStatus{}
				if v, ok := statusData["domain_id"]; ok {
					status.DomainID = v
				}
				if v, ok := statusData["unread_count"].(int); ok {
					status.UnreadCount = v
				}
				if v, ok := statusData["max_announcement_id"]; ok {
					status.MaxAnnouncementID = v
				}
				if v, ok := statusData["max_read_announcement_id"]; ok {
					status.MaxReadAnnouncementID = v
				}
				statuses = append(statuses, status)
			}
		}
	}

	return statuses, nil
}

// UpdateAnnouncementStatus marks an announcement as read.
func (c *Client) UpdateAnnouncementStatus(ctx context.Context, domainID, announcementID interface{}) error {
	params := []interface{}{domainID, announcementID}
	_, err := c.Call(MethodUpdateAnnouncementStatus, params)
	return err
}

// Helper function

func parseAnnouncement(data map[string]interface{}) *Announcement {
	announcement := &Announcement{}

	if v, ok := data["id"]; ok {
		announcement.ID = v
	}
	if v, ok := data["domain_id"]; ok {
		announcement.DomainID = v
	}
	if v, ok := data["title"].(string); ok {
		announcement.Title = v
	}
	if v, ok := data["text"].(string); ok {
		announcement.Text = v
	}
	if v, ok := data["created_by"]; ok {
		announcement.CreatedBy = v
	}
	if v, ok := data["created_at"].(int64); ok {
		announcement.CreatedAt = time.Unix(v, 0)
	}
	if v, ok := data["updated_at"].(int64); ok {
		announcement.UpdatedAt = time.Unix(v, 0)
	}
	if v, ok := data["target_user_ids"].([]interface{}); ok {
		announcement.TargetUserIDs = v
	}
	if v, ok := data["read_user_ids"].([]interface{}); ok {
		announcement.ReadUserIDs = v
	}
	if v, ok := data["unread_user_count"].(int); ok {
		announcement.UnreadUserCount = v
	}

	return announcement
}
