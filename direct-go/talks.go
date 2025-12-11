package direct

import (
	"context"
)

// GroupTalkSettings represents settings for a group talk.
type GroupTalkSettings struct {
	Name                     string
	IconURL                  string
	AllowDisplayPastMessages bool
	Description              string
}

// CreateGroupTalk creates a new group talk/room.
func (c *Client) CreateGroupTalk(ctx context.Context, domainID interface{}, name string, userIDs []interface{}, settings *GroupTalkSettings) (*Talk, error) {
	// Build parameters based on settings
	var params []interface{}
	if settings != nil {
		params = []interface{}{
			domainID,
			name,
			userIDs,
			settings.AllowDisplayPastMessages,
			settings.IconURL,
			settings.Description,
		}
	} else {
		params = []interface{}{
			domainID,
			name,
			userIDs,
			true, // default: allow display past messages
			nil,  // no icon
			"",   // no description
		}
	}

	result, err := c.Call(MethodCreateGroupTalk, params)
	if err != nil {
		return nil, err
	}

	if talkData, ok := result.(map[string]interface{}); ok {
		return parseTalk(talkData), nil
	}

	return nil, nil
}

// CreatePairTalk creates a 1-on-1 talk/room.
func (c *Client) CreatePairTalk(ctx context.Context, domainID, userID interface{}) (*Talk, error) {
	params := []interface{}{domainID, userID}
	result, err := c.Call(MethodCreatePairTalk, params)
	if err != nil {
		return nil, err
	}

	if talkData, ok := result.(map[string]interface{}); ok {
		return parseTalk(talkData), nil
	}

	return nil, nil
}

// UpdateGroupTalk updates a group talk's settings.
func (c *Client) UpdateGroupTalk(ctx context.Context, talkID interface{}, updates map[string]interface{}) (*Talk, error) {
	params := []interface{}{talkID, updates}
	result, err := c.Call(MethodUpdateGroupTalk, params)
	if err != nil {
		return nil, err
	}

	if talkData, ok := result.(map[string]interface{}); ok {
		return parseTalk(talkData), nil
	}

	return nil, nil
}

// AddTalkers adds users to a talk/room.
func (c *Client) AddTalkers(ctx context.Context, talkID interface{}, userIDs []interface{}) error {
	params := []interface{}{talkID, userIDs}
	_, err := c.Call(MethodAddTalkers, params)
	return err
}

// DeleteTalker removes a user from a talk/room.
func (c *Client) DeleteTalker(ctx context.Context, talkID, userID interface{}) error {
	params := []interface{}{talkID, userID}
	_, err := c.Call(MethodDeleteTalker, params)
	return err
}

// AddFavoriteTalk adds a talk to favorites.
func (c *Client) AddFavoriteTalk(ctx context.Context, talkID interface{}) error {
	params := []interface{}{talkID}
	_, err := c.Call(MethodAddFavoriteTalk, params)
	return err
}

// DeleteFavoriteTalk removes a talk from favorites.
func (c *Client) DeleteFavoriteTalk(ctx context.Context, talkID interface{}) error {
	params := []interface{}{talkID}
	_, err := c.Call(MethodDeleteFavoriteTalk, params)
	return err
}

// Helper function to parse talk data
func parseTalk(data map[string]interface{}) *Talk {
	talk := &Talk{}

	if v, ok := data["id"]; ok {
		talk.ID = v
	}
	if v, ok := data["domain_id"]; ok {
		talk.DomainID = v
	}
	if v, ok := data["type"].(int); ok {
		talk.Type = v
	}
	if v, ok := data["name"].(string); ok {
		talk.Name = v
	}
	if v, ok := data["user_ids"].([]interface{}); ok {
		talk.UserIDs = v
	}
	if v, ok := data["allow_display_past_messages"].(bool); ok {
		talk.AllowDisplayPastMessages = v
	}

	return talk
}
