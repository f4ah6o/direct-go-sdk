package direct

import (
	"context"
)

// UserInfo represents detailed user information.
type UserInfo struct {
	ID                  interface{}
	Name                string
	DisplayName         string
	PhoneticName        string
	Email               string
	IconURL             string
	DomainID            interface{}
	Departments         []interface{}
	Profiles            map[string]interface{}
	CanTalk             bool
	AllowedToCreateTalk bool
}

// ProfileInfo represents user profile details.
type ProfileInfo struct {
	UserID       interface{}
	DomainID     interface{}
	DisplayName  string
	PhoneticName string
	Profiles     map[string]interface{}
}

// PresenceInfo represents user presence/online status.
type PresenceInfo struct {
	UserID interface{}
	Status string // e.g., "online", "offline", "away"
}

// GetUsers retrieves multiple users by their IDs.
func (c *Client) GetUsers(ctx context.Context, domainID interface{}, userIDs []interface{}) ([]UserInfo, error) {
	params := []interface{}{domainID, userIDs}
	result, err := c.Call(MethodGetUsers, params)
	if err != nil {
		return nil, err
	}

	users := []UserInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if userData, ok := item.(map[string]interface{}); ok {
				user := parseUserInfo(userData)
				users = append(users, user)
			}
		}
	}

	return users, nil
}

// GetProfile retrieves detailed profile for a user.
func (c *Client) GetProfile(ctx context.Context, domainID, userID interface{}) (*ProfileInfo, error) {
	params := []interface{}{domainID, userID}
	result, err := c.Call(MethodGetProfile, params)
	if err != nil {
		return nil, err
	}

	if profileData, ok := result.(map[string]interface{}); ok {
		return parseProfileInfo(profileData), nil
	}

	return nil, nil
}

// UpdateProfile updates the current user's profile.
func (c *Client) UpdateProfile(ctx context.Context, domainID interface{}, updates map[string]interface{}) error {
	params := []interface{}{domainID, updates}
	_, err := c.Call(MethodUpdateProfile, params)
	return err
}

// UpdateUser updates user information.
func (c *Client) UpdateUser(ctx context.Context, userID interface{}, updates map[string]interface{}) error {
	params := []interface{}{userID, updates}
	_, err := c.Call(MethodUpdateUser, params)
	return err
}

// GetPresences retrieves presence status for multiple users.
func (c *Client) GetPresences(ctx context.Context, userIDs []interface{}) ([]PresenceInfo, error) {
	params := []interface{}{userIDs}
	result, err := c.Call(MethodGetPresences, params)
	if err != nil {
		return nil, err
	}

	presences := []PresenceInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if presData, ok := item.(map[string]interface{}); ok {
				pres := PresenceInfo{}
				if v, ok := presData["user_id"]; ok {
					pres.UserID = v
				}
				if v, ok := presData["status"].(string); ok {
					pres.Status = v
				}
				presences = append(presences, pres)
			}
		}
	}

	return presences, nil
}

// UserIdentifier represents identifier information for a user.
type UserIdentifier struct {
	UserID     interface{}
	Email      string
	SubEmail   string
	GroupAlias string
	SigninID   string
}

// GetUserIdentifiers retrieves identifiers for multiple users.
func (c *Client) GetUserIdentifiers(ctx context.Context, userIDs []interface{}) ([]UserIdentifier, error) {
	params := []interface{}{userIDs}
	result, err := c.Call(MethodGetUserIdentifiers, params)
	if err != nil {
		return nil, err
	}

	identifiers := []UserIdentifier{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if idData, ok := item.(map[string]interface{}); ok {
				id := UserIdentifier{}
				if v, ok := idData["user_id"]; ok {
					id.UserID = v
				}
				if v, ok := idData["email"].(string); ok {
					id.Email = v
				}
				if v, ok := idData["sub_email"].(string); ok {
					id.SubEmail = v
				}
				if v, ok := idData["group_alias"].(string); ok {
					id.GroupAlias = v
				}
				if v, ok := idData["signin_id"].(string); ok {
					id.SigninID = v
				}
				identifiers = append(identifiers, id)
			}
		}
	}

	return identifiers, nil
}

// GetFriends retrieves the current user's friends list.
func (c *Client) GetFriends(ctx context.Context) ([]UserInfo, error) {
	result, err := c.Call(MethodGetFriends, []interface{}{})
	if err != nil {
		return nil, err
	}

	friends := []UserInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if userData, ok := item.(map[string]interface{}); ok {
				user := parseUserInfo(userData)
				friends = append(friends, user)
			}
		}
	}

	return friends, nil
}

// AddFriend adds a user as a friend.
func (c *Client) AddFriend(ctx context.Context, userID interface{}) error {
	params := []interface{}{userID}
	_, err := c.Call(MethodAddFriend, params)
	return err
}

// DeleteFriend removes a user from friends list.
func (c *Client) DeleteFriend(ctx context.Context, userID interface{}) error {
	params := []interface{}{userID}
	_, err := c.Call(MethodDeleteFriend, params)
	return err
}

// GetAcquaintances retrieves the user's acquaintances list.
func (c *Client) GetAcquaintances(ctx context.Context) ([]UserInfo, error) {
	result, err := c.Call(MethodGetAcquaintances, []interface{}{})
	if err != nil {
		return nil, err
	}

	acquaintances := []UserInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if userData, ok := item.(map[string]interface{}); ok {
				user := parseUserInfo(userData)
				acquaintances = append(acquaintances, user)
			}
		}
	}

	return acquaintances, nil
}

// Helper functions to parse user data

func parseUserInfo(data map[string]interface{}) UserInfo {
	user := UserInfo{}

	if v, ok := data["id"]; ok {
		user.ID = v
	}
	if v, ok := data["name"].(string); ok {
		user.Name = v
	}
	if v, ok := data["display_name"].(string); ok {
		user.DisplayName = v
	}
	if v, ok := data["phonetic_name"].(string); ok {
		user.PhoneticName = v
	}
	if v, ok := data["email"].(string); ok {
		user.Email = v
	}
	if v, ok := data["icon_url"].(string); ok {
		user.IconURL = v
	}
	if v, ok := data["domain_id"]; ok {
		user.DomainID = v
	}
	if v, ok := data["departments"].([]interface{}); ok {
		user.Departments = v
	}
	if v, ok := data["profiles"].(map[string]interface{}); ok {
		user.Profiles = v
	}
	if v, ok := data["can_talk"].(bool); ok {
		user.CanTalk = v
	}
	if v, ok := data["allowed_to_create_talk"].(bool); ok {
		user.AllowedToCreateTalk = v
	}

	return user
}

func parseProfileInfo(data map[string]interface{}) *ProfileInfo {
	profile := &ProfileInfo{}

	if v, ok := data["user_id"]; ok {
		profile.UserID = v
	}
	if v, ok := data["domain_id"]; ok {
		profile.DomainID = v
	}
	if v, ok := data["display_name"].(string); ok {
		profile.DisplayName = v
	}
	if v, ok := data["phonetic_name"].(string); ok {
		profile.PhoneticName = v
	}
	if v, ok := data["profiles"].(map[string]interface{}); ok {
		profile.Profiles = v
	}

	return profile
}
