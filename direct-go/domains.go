package direct

import (
	"context"
)

// DomainInfo represents detailed domain information.
type DomainInfo struct {
	ID        interface{}
	Name      string
	UpdatedAt int64
	Contract  interface{} // Contract details
	Setting   interface{} // Domain settings
	Role      interface{} // User's role in domain
	Closed    bool
}

// DomainInviteInfo represents a domain invitation.
type DomainInviteInfo struct {
	ID                      interface{}
	Name                    string
	AccountControlRequestID interface{}
	UpdatedAt               int64
}

// GetDomainsWithContext retrieves the list of domains/organizations the user belongs to.
// Returns DomainInfo with domain names, settings, user roles, and contract details.
// This replaces the legacy GetDomains() method.
func (c *Client) GetDomainsWithContext(ctx context.Context) ([]DomainInfo, error) {
	result, err := c.Call(MethodGetDomains, []interface{}{})
	if err != nil {
		return nil, err
	}

	domains := []DomainInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if domainData, ok := item.(map[string]interface{}); ok {
				domain := parseDomainInfo(domainData)
				domains = append(domains, domain)
			}
		}
	}

	return domains, nil
}

// GetDomainInvitesWithContext retrieves pending invitations to join domains/organizations.
// Returns DomainInviteInfo with invitation IDs, domain names, and timestamps.
// This replaces the legacy GetDomainInvites() method.
func (c *Client) GetDomainInvitesWithContext(ctx context.Context) ([]DomainInviteInfo, error) {
	result, err := c.Call(MethodGetDomainInvites, []interface{}{})
	if err != nil {
		return nil, err
	}

	invites := []DomainInviteInfo{}
	if arr, ok := result.([]interface{}); ok {
		for _, item := range arr {
			if inviteData, ok := item.(map[string]interface{}); ok {
				invite := parseDomainInviteInfo(inviteData)
				invites = append(invites, invite)
			}
		}
	}

	return invites, nil
}

// AcceptDomainInviteWithContext accepts a pending domain invitation and joins the domain.
// Returns the newly joined DomainInfo on success.
// This replaces the legacy AcceptDomainInvite() method.
func (c *Client) AcceptDomainInviteWithContext(ctx context.Context, inviteID interface{}) (*DomainInfo, error) {
	params := []interface{}{inviteID}
	result, err := c.Call(MethodAcceptDomainInvite, params)
	if err != nil {
		return nil, err
	}

	if domainData, ok := result.(map[string]interface{}); ok {
		domain := parseDomainInfo(domainData)
		return &domain, nil
	}

	return nil, nil
}

// LeaveDomain removes the current user from the specified domain/organization.
func (c *Client) LeaveDomain(ctx context.Context, domainID interface{}) error {
	params := []interface{}{domainID}
	_, err := c.Call(MethodLeaveDomain, params)
	return err
}

// GetDomainUsers retrieves all users belonging to a specific domain/organization.
// Returns a slice of UserInfo with user profiles, departments, and permissions.
func (c *Client) GetDomainUsers(ctx context.Context, domainID interface{}) ([]UserInfo, error) {
	params := []interface{}{domainID}
	result, err := c.Call(MethodGetDomainUsers, params)
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

// SearchDomainUsers searches for users within a domain using a query string.
// The query matches against user names, display names, and email addresses.
func (c *Client) SearchDomainUsers(ctx context.Context, domainID interface{}, query string) ([]UserInfo, error) {
	params := []interface{}{domainID, query}
	result, err := c.Call(MethodSearchDomainUsers, params)
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

// DeleteDomainInvite rejects and deletes a pending domain invitation.
func (c *Client) DeleteDomainInvite(ctx context.Context, inviteID interface{}) error {
	params := []interface{}{inviteID}
	_, err := c.Call(MethodDeleteDomainInvite, params)
	return err
}

// Helper functions

func parseDomainInfo(data map[string]interface{}) DomainInfo {
	domain := DomainInfo{}

	if v, ok := data["id"]; ok {
		domain.ID = v
	}
	if v, ok := data["domain_id"]; ok {
		domain.ID = v
	}
	if v, ok := data["name"].(string); ok {
		domain.Name = v
	}
	if v, ok := data["domain_name"].(string); ok {
		domain.Name = v
	}
	if v, ok := data["updated_at"].(int64); ok {
		domain.UpdatedAt = v
	}
	if v, ok := data["contract"]; ok {
		domain.Contract = v
	}
	if v, ok := data["setting"]; ok {
		domain.Setting = v
	}
	if v, ok := data["role"]; ok {
		domain.Role = v
	}
	if v, ok := data["closed"].(bool); ok {
		domain.Closed = v
	}

	return domain
}

func parseDomainInviteInfo(data map[string]interface{}) DomainInviteInfo {
	invite := DomainInviteInfo{}

	if v, ok := data["id"]; ok {
		invite.ID = v
	}
	if v, ok := data["domain_id"]; ok {
		invite.ID = v
	}
	if v, ok := data["name"].(string); ok {
		invite.Name = v
	}
	if v, ok := data["domain_name"].(string); ok {
		invite.Name = v
	}
	if v, ok := data["account_control_request_id"]; ok {
		invite.AccountControlRequestID = v
	}
	if v, ok := data["updated_at"].(int64); ok {
		invite.UpdatedAt = v
	}

	return invite
}
