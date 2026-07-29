package direct

import (
	"context"
	"fmt"
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

// SearchMessagesAroundDateTimeResult contains messages around a target datetime.
type SearchMessagesAroundDateTimeResult struct {
	DateTime interface{}
	Messages []ReceivedMessage
	Raw      map[string]interface{}
}

// MessageSearchContent represents a search result item.
type MessageSearchContent struct {
	Message    ReceivedMessage
	TalkID     interface{}
	DomainID   interface{}
	MatchScore float64
}

// ReadStatus represents per-message read state.
type ReadStatus struct {
	MessageID     string
	TalkID        string
	ReadUserIDs   []string
	UnreadUserIDs []string
}

// ReadStatusesUpdate represents a read-status notification from direct.
type ReadStatusesUpdate struct {
	TalkID                                string
	MessageIDs                            []string
	MentionMessageIDs                     []string
	ReadUserIDs                           []string
	MessageIDsExcludingUnreadCountTargets []string
}

// GetReadStatus retrieves read status for a message in a talk.
func (c *Client) GetReadStatus(ctx context.Context, talkID, messageID interface{}) (*ReadStatus, error) {
	result, err := c.CallWithContext(ctx, MethodGetReadStatus, []interface{}{talkID, messageID})
	if err != nil {
		return nil, err
	}
	return parseReadStatus(result), nil
}

// UpdateReadStatuses marks messages through maxReadMessageID as read in a talk.
func (c *Client) UpdateReadStatuses(ctx context.Context, talkID, maxReadMessageID interface{}) error {
	_, err := c.CallWithContext(ctx, MethodUpdateReadStatuses, []interface{}{talkID, maxReadMessageID})
	return err
}

// ParseReadStatusesUpdate parses notify_update_read_statuses payloads.
func ParseReadStatusesUpdate(data interface{}) ReadStatusesUpdate {
	return parseReadStatusesUpdate(data)
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
	result, err := c.CallWithContext(ctx, MethodGetMessages, params)
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
	_, err := c.CallWithContext(ctx, MethodDeleteMessage, params)
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
	result, err := c.CallWithContext(ctx, MethodSearchMessages, params)
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

// SearchMessagesAroundDateTime searches for messages around a target datetime.
func (c *Client) SearchMessagesAroundDateTime(ctx context.Context, talkID, datetime interface{}) (*SearchMessagesAroundDateTimeResult, error) {
	result, err := c.CallWithContext(ctx, MethodSearchMessagesAroundDateTime, []interface{}{talkID, datetime})
	if err != nil {
		return nil, err
	}

	searchResult := &SearchMessagesAroundDateTimeResult{DateTime: datetime}
	if resultMap, ok := result.(map[string]interface{}); ok {
		searchResult.Raw = resultMap
		for _, key := range []string{"messages", "contents"} {
			if messages, ok := resultMap[key].([]interface{}); ok {
				for _, item := range messages {
					if msgData, ok := item.(map[string]interface{}); ok {
						searchResult.Messages = append(searchResult.Messages, parseMessage(msgData))
					}
				}
			}
		}
	}

	return searchResult, nil
}

// GetFavoriteMessages retrieves the user's favorite messages.
func (c *Client) GetFavoriteMessages(ctx context.Context) ([]ReceivedMessage, error) {
	result, err := c.CallWithContext(ctx, MethodGetFavoriteMessages, []interface{}{})
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
	_, err := c.CallWithContext(ctx, MethodAddFavoriteMessage, params)
	return err
}

// DeleteFavoriteMessage removes a message from favorites.
func (c *Client) DeleteFavoriteMessage(ctx context.Context, messageID interface{}) error {
	params := []interface{}{messageID}
	_, err := c.CallWithContext(ctx, MethodDeleteFavoriteMessage, params)
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
	result, err := c.CallWithContext(ctx, MethodGetScheduledMessages, []interface{}{})
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
					msg.ScheduledAt = UnixTime(v)
				}
				if v, ok := msgData["created_at"].(int64); ok {
					msg.CreatedAt = UnixTime(v)
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
	result, err := c.CallWithContext(ctx, MethodScheduleMessage, params)
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
			msg.ScheduledAt = UnixTime(v)
		}
		if v, ok := msgData["created_at"].(int64); ok {
			msg.CreatedAt = UnixTime(v)
		}
	}

	return msg, nil
}

// DeleteScheduledMessage deletes a scheduled message.
func (c *Client) DeleteScheduledMessage(ctx context.Context, messageID interface{}) error {
	params := []interface{}{messageID}
	_, err := c.CallWithContext(ctx, MethodDeleteScheduledMessage, params)
	return err
}

// RescheduleMessage changes the scheduled time of a message.
func (c *Client) RescheduleMessage(ctx context.Context, messageID interface{}, newScheduledAt time.Time) error {
	params := []interface{}{messageID, newScheduledAt.Unix()}
	_, err := c.CallWithContext(ctx, MethodRescheduleMessage, params)
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
	result, err := c.CallWithContext(ctx, MethodGetAvailableMessageReactions, []interface{}{})
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
	_, err := c.CallWithContext(ctx, MethodSetMessageReaction, params)
	return err
}

// ResetMessageReaction removes a reaction from a message.
func (c *Client) ResetMessageReaction(ctx context.Context, messageID, reactionID interface{}) error {
	params := []interface{}{messageID, reactionID}
	_, err := c.CallWithContext(ctx, MethodResetMessageReaction, params)
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
	result, err := c.CallWithContext(ctx, MethodGetMessageReactionUsers, params)
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
					user.CreatedAt = UnixTime(v)
				}
				users = append(users, user)
			}
		}
	}

	return users, nil
}

func parseReadStatus(data interface{}) *ReadStatus {
	status := &ReadStatus{}
	m, ok := data.(map[string]interface{})
	if !ok {
		return status
	}
	status.MessageID = stringFromMap(m, "message_id", "id")
	status.TalkID = stringFromMap(m, "talk_id")
	status.ReadUserIDs = stringSliceFromValue(m["read_user_ids"])
	status.UnreadUserIDs = stringSliceFromValue(m["unread_user_ids"])
	return status
}

func parseReadStatusesUpdate(data interface{}) ReadStatusesUpdate {
	update := ReadStatusesUpdate{}
	m, ok := data.(map[string]interface{})
	if !ok {
		return update
	}
	update.TalkID = stringFromMap(m, "talk_id")
	update.MessageIDs = stringSliceFromValue(m["message_ids"])
	update.MentionMessageIDs = stringSliceFromValue(m["mention_message_ids"])
	update.ReadUserIDs = stringSliceFromValue(m["read_user_ids"])
	update.MessageIDsExcludingUnreadCountTargets = stringSliceFromValue(m["message_ids_excluding_unread_count_targets"])
	return update
}

func stringFromMap(m map[string]interface{}, keys ...string) string {
	for _, key := range keys {
		if v, ok := m[key]; ok {
			return stringFromValue(v)
		}
	}
	return ""
}

func stringFromValue(v interface{}) string {
	if v == nil {
		return ""
	}
	return fmt.Sprintf("%v", v)
}

func stringSliceFromValue(v interface{}) []string {
	arr, ok := v.([]interface{})
	if !ok {
		return nil
	}
	out := make([]string, 0, len(arr))
	for _, item := range arr {
		out = append(out, stringFromValue(item))
	}
	return out
}

func mapSliceFromValue(v interface{}) []map[string]interface{} {
	arr, ok := v.([]interface{})
	if !ok {
		return nil
	}
	out := make([]map[string]interface{}, 0, len(arr))
	for _, item := range arr {
		if data, ok := item.(map[string]interface{}); ok {
			out = append(out, data)
		}
	}
	return out
}
