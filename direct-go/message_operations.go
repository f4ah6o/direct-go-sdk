package direct

import (
	"context"
	"time"
)

// MessageOrder represents the order for retrieving messages.
type MessageOrder int

const (
	// MessageOrderAsc retrieves messages in ascending order (oldest first).
	MessageOrderAsc MessageOrder = 1
	// MessageOrderDesc retrieves messages in descending order (newest first).
	MessageOrderDesc MessageOrder = 2
)

// GetMessagesOptions provides options for retrieving messages.
type GetMessagesOptions struct {
	// SinceID retrieves messages newer than this ID.
	SinceID interface{}
	// MaxID retrieves messages older than this ID.
	MaxID interface{}
	// Order specifies the order of messages (default: MessageOrderDesc).
	Order MessageOrder
}

// MessagesResult contains the result of GetMessages call.
type MessagesResult struct {
	Messages []ReceivedMessage
}

// SearchMessagesResult contains the result of SearchMessages call.
type SearchMessagesResult struct {
	Total      int
	Marker     interface{}
	NextMarker interface{}
	Contents   []MessageSearchContent
}

// MessageSearchContent represents a search result item.
type MessageSearchContent struct {
	Message    ReceivedMessage
	TalkID     interface{}
	DomainID   interface{}
	MatchScore float64
}

// GetMessages retrieves messages from a talk room.
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - domainID: Domain ID
//   - talkID: Talk/Room ID
//   - opts: Options for message retrieval (optional)
//
// Returns a list of messages matching the criteria.
func (c *Client) GetMessages(ctx context.Context, domainID, talkID interface{}, opts *GetMessagesOptions) ([]ReceivedMessage, error) {
	if opts == nil {
		opts = &GetMessagesOptions{Order: MessageOrderDesc}
	}
	if opts.Order == 0 {
		opts.Order = MessageOrderDesc
	}

	params := []interface{}{domainID, talkID, opts.SinceID, opts.MaxID, int(opts.Order)}
	result, err := c.Call(MethodGetMessages, params)
	if err != nil {
		return nil, err
	}

	// Parse result as array of messages
	messages := []ReceivedMessage{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if msgData, ok := item.(map[string]interface{}); ok {
				msg := parseMessage(msgData)
				messages = append(messages, msg)
			}
		}
	}

	return messages, nil
}

// DeleteMessage deletes a message from a talk room.
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - domainID: Domain ID
//   - messageID: Message ID to delete
//
// Returns error if the deletion fails.
func (c *Client) DeleteMessage(ctx context.Context, domainID, messageID interface{}) error {
	params := []interface{}{domainID, messageID}
	_, err := c.Call(MethodDeleteMessage, params)
	return err
}

// SearchMessages searches for messages in a talk room.
// Parameters:
//   - ctx: Context for cancellation and timeout
//   - domainID: Domain ID
//   - talkID: Talk/Room ID (optional, nil for all talks in domain)
//   - keyword: Search keyword
//   - marker: Pagination marker (optional)
//   - limit: Maximum number of results
//
// Returns search results with pagination information.
func (c *Client) SearchMessages(ctx context.Context, domainID, talkID interface{}, keyword string, marker interface{}, limit int) (*SearchMessagesResult, error) {
	params := []interface{}{domainID, talkID, keyword, marker, limit}
	result, err := c.Call(MethodSearchMessages, params)
	if err != nil {
		return nil, err
	}

	// Parse result
	searchResult := &SearchMessagesResult{
		Contents: []MessageSearchContent{},
	}

	if resultMap, ok := result.(map[string]interface{}); ok {
		if total, ok := resultMap["total"].(int); ok {
			searchResult.Total = total
		}
		if v, ok := resultMap["marker"]; ok {
			searchResult.Marker = v
		}
		if v, ok := resultMap["next_marker"]; ok {
			searchResult.NextMarker = v
		}
		if contents, ok := resultMap["contents"].([]interface{}); ok {
			for _, item := range contents {
				if contentMap, ok := item.(map[string]interface{}); ok {
					content := MessageSearchContent{}
					if msgData, ok := contentMap["message"].(map[string]interface{}); ok {
						content.Message = parseMessage(msgData)
					}
					if v, ok := contentMap["talk_id"]; ok {
						content.TalkID = v
					}
					if v, ok := contentMap["domain_id"]; ok {
						content.DomainID = v
					}
					if score, ok := contentMap["match_score"].(float64); ok {
						content.MatchScore = score
					}
					searchResult.Contents = append(searchResult.Contents, content)
				}
			}
		}
	}

	return searchResult, nil
}

// GetFavoriteMessages retrieves the user's favorite messages.
func (c *Client) GetFavoriteMessages(ctx context.Context) ([]ReceivedMessage, error) {
	result, err := c.Call(MethodGetFavoriteMessages, []interface{}{})
	if err != nil {
		return nil, err
	}

	messages := []ReceivedMessage{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if msgData, ok := item.(map[string]interface{}); ok {
				msg := parseMessage(msgData)
				messages = append(messages, msg)
			}
		}
	}

	return messages, nil
}

// AddFavoriteMessage adds a message to favorites.
func (c *Client) AddFavoriteMessage(ctx context.Context, messageID interface{}) error {
	params := []interface{}{messageID}
	_, err := c.Call(MethodAddFavoriteMessage, params)
	return err
}

// DeleteFavoriteMessage removes a message from favorites.
func (c *Client) DeleteFavoriteMessage(ctx context.Context, messageID interface{}) error {
	params := []interface{}{messageID}
	_, err := c.Call(MethodDeleteFavoriteMessage, params)
	return err
}

// ScheduledMessage represents a scheduled message.
type ScheduledMessage struct {
	ID          interface{}
	TalkID      interface{}
	DomainID    interface{}
	Type        MessageType
	Content     interface{}
	ScheduledAt time.Time
	CreatedAt   time.Time
}

// GetScheduledMessages retrieves all scheduled messages.
func (c *Client) GetScheduledMessages(ctx context.Context) ([]ScheduledMessage, error) {
	result, err := c.Call(MethodGetScheduledMessages, []interface{}{})
	if err != nil {
		return nil, err
	}

	messages := []ScheduledMessage{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if msgData, ok := item.(map[string]interface{}); ok {
				msg := ScheduledMessage{}
				if v, ok := msgData["id"]; ok {
					msg.ID = v
				}
				if v, ok := msgData["talk_id"]; ok {
					msg.TalkID = v
				}
				if v, ok := msgData["domain_id"]; ok {
					msg.DomainID = v
				}
				if v, ok := msgData["type"].(int); ok {
					msg.Type = MessageType(v)
				}
				if v, ok := msgData["content"]; ok {
					msg.Content = v
				}
				if v, ok := msgData["scheduled_at"].(int64); ok {
					msg.ScheduledAt = time.Unix(v, 0)
				}
				if v, ok := msgData["created_at"].(int64); ok {
					msg.CreatedAt = time.Unix(v, 0)
				}
				messages = append(messages, msg)
			}
		}
	}

	return messages, nil
}

// ScheduleMessage schedules a message to be sent at a specific time.
func (c *Client) ScheduleMessage(ctx context.Context, talkID interface{}, msgType MessageType, content interface{}, scheduledAt time.Time) (*ScheduledMessage, error) {
	params := []interface{}{talkID, int(msgType), content, scheduledAt.Unix()}
	result, err := c.Call(MethodScheduleMessage, params)
	if err != nil {
		return nil, err
	}

	msg := &ScheduledMessage{}
	if msgData, ok := result.(map[string]interface{}); ok {
		if v, ok := msgData["id"]; ok {
			msg.ID = v
		}
		if v, ok := msgData["talk_id"]; ok {
			msg.TalkID = v
		}
		if v, ok := msgData["domain_id"]; ok {
			msg.DomainID = v
		}
		if v, ok := msgData["type"].(int); ok {
			msg.Type = MessageType(v)
		}
		if v, ok := msgData["content"]; ok {
			msg.Content = v
		}
		if v, ok := msgData["scheduled_at"].(int64); ok {
			msg.ScheduledAt = time.Unix(v, 0)
		}
		if v, ok := msgData["created_at"].(int64); ok {
			msg.CreatedAt = time.Unix(v, 0)
		}
	}

	return msg, nil
}

// DeleteScheduledMessage deletes a scheduled message.
func (c *Client) DeleteScheduledMessage(ctx context.Context, messageID interface{}) error {
	params := []interface{}{messageID}
	_, err := c.Call(MethodDeleteScheduledMessage, params)
	return err
}

// RescheduleMessage changes the scheduled time of a message.
func (c *Client) RescheduleMessage(ctx context.Context, messageID interface{}, newScheduledAt time.Time) error {
	params := []interface{}{messageID, newScheduledAt.Unix()}
	_, err := c.Call(MethodRescheduleMessage, params)
	return err
}

// MessageReaction represents a reaction to a message.
type MessageReaction struct {
	ID       interface{}
	Name     string
	ImageURL string
}

// GetAvailableMessageReactions retrieves all available message reactions.
func (c *Client) GetAvailableMessageReactions(ctx context.Context) ([]MessageReaction, error) {
	result, err := c.Call(MethodGetAvailableMessageReactions, []interface{}{})
	if err != nil {
		return nil, err
	}

	reactions := []MessageReaction{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if data, ok := item.(map[string]interface{}); ok {
				reaction := MessageReaction{}
				if v, ok := data["id"]; ok {
					reaction.ID = v
				}
				if v, ok := data["name"].(string); ok {
					reaction.Name = v
				}
				if v, ok := data["image_url"].(string); ok {
					reaction.ImageURL = v
				}
				reactions = append(reactions, reaction)
			}
		}
	}

	return reactions, nil
}

// SetMessageReaction sets a reaction on a message.
func (c *Client) SetMessageReaction(ctx context.Context, messageID, reactionID interface{}) error {
	params := []interface{}{messageID, reactionID}
	_, err := c.Call(MethodSetMessageReaction, params)
	return err
}

// ResetMessageReaction removes a reaction from a message.
func (c *Client) ResetMessageReaction(ctx context.Context, messageID, reactionID interface{}) error {
	params := []interface{}{messageID, reactionID}
	_, err := c.Call(MethodResetMessageReaction, params)
	return err
}

// MessageReactionUser represents a user who reacted to a message.
type MessageReactionUser struct {
	UserID     interface{}
	ReactionID interface{}
	CreatedAt  time.Time
}

// GetMessageReactionUsers retrieves users who reacted to a message.
func (c *Client) GetMessageReactionUsers(ctx context.Context, messageID interface{}) ([]MessageReactionUser, error) {
	params := []interface{}{messageID}
	result, err := c.Call(MethodGetMessageReactionUsers, params)
	if err != nil {
		return nil, err
	}

	users := []MessageReactionUser{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if data, ok := item.(map[string]interface{}); ok {
				user := MessageReactionUser{}
				if v, ok := data["user_id"]; ok {
					user.UserID = v
				}
				if v, ok := data["reaction_id"]; ok {
					user.ReactionID = v
				}
				if v, ok := data["created_at"].(int64); ok {
					user.CreatedAt = time.Unix(v, 0)
				}
				users = append(users, user)
			}
		}
	}

	return users, nil
}
