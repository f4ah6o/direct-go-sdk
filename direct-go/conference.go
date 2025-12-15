package direct

import (
	"context"
	"time"
)

// Conference represents a video/audio conference.
type Conference struct {
	ID            interface{}
	UserID        interface{}
	DomainID      interface{}
	TalkID        interface{}
	MessageID     interface{}
	CreatedAt     time.Time
	ExpiredAt     time.Time
	Participants  []interface{}
	SkywayVersion int
}

// ConferenceJoinInfo represents information for joining a conference.
type ConferenceJoinInfo struct {
	ConferenceID  interface{}
	RoomName      string
	Credential    string
	Mode          string
	Timestamp     int64
	SkywayVersion int
}

// GetConferences retrieves all active video/audio conferences the user can see.
// Returns a slice of Conference objects with participant lists and metadata.
func (c *Client) GetConferences(ctx context.Context) ([]Conference, error) {
	result, err := c.Call(MethodGetConferences, []interface{}{})
	if err != nil {
		return nil, err
	}

	conferences := []Conference{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if confData, ok := item.(map[string]interface{}); ok {
				conf := parseConference(confData)
				conferences = append(conferences, *conf)
			}
		}
	}

	return conferences, nil
}

// GetConferenceParticipants retrieves the list of users participating in a conference.
// Returns a slice of participant IDs or user objects.
func (c *Client) GetConferenceParticipants(ctx context.Context, conferenceID interface{}) ([]interface{}, error) {
	params := []interface{}{conferenceID}
	result, err := c.Call(MethodGetConferenceParticipants, params)
	if err != nil {
		return nil, err
	}

	if arr, ok := result.([]interface{}); ok {
		return arr, nil
	}

	return []interface{}{}, nil
}

// JoinConference joins an active conference as a participant.
// Returns ConferenceJoinInfo with room name, credentials, and connection details.
func (c *Client) JoinConference(ctx context.Context, conferenceID interface{}) (*ConferenceJoinInfo, error) {
	params := []interface{}{conferenceID}
	result, err := c.Call(MethodJoinConference, params)
	if err != nil {
		return nil, err
	}

	if joinData, ok := result.(map[string]interface{}); ok {
		return parseConferenceJoinInfo(joinData), nil
	}

	return nil, nil
}

// LeaveConference disconnects the current user from an active conference.
func (c *Client) LeaveConference(ctx context.Context, conferenceID interface{}) error {
	params := []interface{}{conferenceID}
	_, err := c.Call(MethodLeaveConference, params)
	return err
}

// RejectConference declines an invitation to join a conference.
func (c *Client) RejectConference(ctx context.Context, conferenceID interface{}) error {
	params := []interface{}{conferenceID}
	_, err := c.Call(MethodRejectConference, params)
	return err
}

// Helper functions

func parseConference(data map[string]interface{}) *Conference {
	conf := &Conference{}

	if v, ok := data["id"]; ok {
		conf.ID = v
	}
	if v, ok := data["conference_id"]; ok {
		conf.ID = v
	}
	if v, ok := data["user_id"]; ok {
		conf.UserID = v
	}
	if v, ok := data["domain_id"]; ok {
		conf.DomainID = v
	}
	if v, ok := data["talk_id"]; ok {
		conf.TalkID = v
	}
	if v, ok := data["message_id"]; ok {
		conf.MessageID = v
	}
	if v, ok := data["created_at"].(int64); ok {
		conf.CreatedAt = time.Unix(v, 0)
	}
	if v, ok := data["expired_at"].(int64); ok {
		conf.ExpiredAt = time.Unix(v, 0)
	}
	if v, ok := data["participants"].([]interface{}); ok {
		conf.Participants = v
	}
	if v, ok := data["skyway_version"].(int); ok {
		conf.SkywayVersion = v
	}

	return conf
}

func parseConferenceJoinInfo(data map[string]interface{}) *ConferenceJoinInfo {
	info := &ConferenceJoinInfo{}

	if v, ok := data["conference_id"]; ok {
		info.ConferenceID = v
	}
	if v, ok := data["room_name"].(string); ok {
		info.RoomName = v
	}
	if v, ok := data["credential"].(string); ok {
		info.Credential = v
	}
	if v, ok := data["mode"].(string); ok {
		info.Mode = v
	}
	if v, ok := data["timestamp"].(int64); ok {
		info.Timestamp = v
	}
	if v, ok := data["skyway_version"].(int); ok {
		info.SkywayVersion = v
	}

	return info
}
