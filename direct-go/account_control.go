package direct

import "context"

// AccountControlRequest represents an account-control request payload.
type AccountControlRequest struct {
	ID      interface{}
	Version interface{}
	Raw     map[string]interface{}
}

// AccountControlGroup represents an account-control group payload.
type AccountControlGroup struct {
	ID  interface{}
	Raw map[string]interface{}
}

// GetAccountControlRequests retrieves pending account-control requests.
func (c *Client) GetAccountControlRequests(ctx context.Context) ([]AccountControlRequest, error) {
	result, err := c.CallWithContext(ctx, MethodGetAccountControlRequests, []interface{}{})
	if err != nil {
		return nil, err
	}

	requests := []AccountControlRequest{}
	for _, data := range mapSliceFromValue(result) {
		requests = append(requests, AccountControlRequest{
			ID:      data["id"],
			Version: data["version"],
			Raw:     data,
		})
	}
	return requests, nil
}

// GetJoinedAccountControlGroup retrieves account-control groups joined by the current user.
func (c *Client) GetJoinedAccountControlGroup(ctx context.Context) ([]AccountControlGroup, error) {
	result, err := c.CallWithContext(ctx, MethodGetJoinedAccountControlGroup, []interface{}{})
	if err != nil {
		return nil, err
	}

	groups := []AccountControlGroup{}
	for _, data := range mapSliceFromValue(result) {
		groups = append(groups, AccountControlGroup{
			ID:  data["id"],
			Raw: data,
		})
	}
	return groups, nil
}

// AcceptAccountControlRequest accepts an account-control request by id and version.
func (c *Client) AcceptAccountControlRequest(ctx context.Context, id, version interface{}) error {
	_, err := c.CallWithContext(ctx, MethodAcceptAccountControlRequest, []interface{}{id, version})
	return err
}

// RejectAccountControlRequest rejects an account-control request by id and version.
func (c *Client) RejectAccountControlRequest(ctx context.Context, id, version interface{}) error {
	_, err := c.CallWithContext(ctx, MethodRejectAccountControlRequest, []interface{}{id, version})
	return err
}
