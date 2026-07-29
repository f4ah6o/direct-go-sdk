package direct

import "context"

// ActionQueryOptions controls action/question retrieval boundaries.
type ActionQueryOptions struct {
	SinceID interface{}
	MaxID   interface{}
}

// RawItem is a conservatively parsed API payload for methods without stable Go models yet.
type RawItem map[string]interface{}

// GetActions retrieves action/question payloads.
func (c *Client) GetActions(ctx context.Context, domainID, talkID interface{}, actionType int, marker interface{}, limit int, opts *ActionQueryOptions) ([]RawItem, error) {
	if opts == nil {
		opts = &ActionQueryOptions{}
	}
	result, err := c.CallWithContext(ctx, MethodGetActions, []interface{}{domainID, talkID, actionType, marker, limit, opts.SinceID, opts.MaxID})
	if err != nil {
		return nil, err
	}
	return rawItemsFromValue(result), nil
}

// GetSolutions retrieves solution payloads for a domain.
func (c *Client) GetSolutions(ctx context.Context, domainID, marker interface{}) ([]RawItem, error) {
	result, err := c.CallWithContext(ctx, MethodGetSolutions, []interface{}{domainID, marker})
	if err != nil {
		return nil, err
	}
	return rawItemsFromValue(result), nil
}

// GetStampsets retrieves original stamp sets for a domain.
func (c *Client) GetStampsets(ctx context.Context, domainID interface{}) ([]RawItem, error) {
	result, err := c.CallWithContext(ctx, MethodGetStampsets, []interface{}{domainID})
	if err != nil {
		return nil, err
	}
	return rawItemsFromValue(result), nil
}

// GetFlowNotificationBadges retrieves flow notification badges for a domain.
func (c *Client) GetFlowNotificationBadges(ctx context.Context, domainID interface{}) ([]RawItem, error) {
	result, err := c.CallWithContext(ctx, MethodGetFlowNotificationBadges, []interface{}{domainID})
	if err != nil {
		return nil, err
	}
	return rawItemsFromValue(result), nil
}

// GetDirectApps retrieves direct app payloads for a domain.
func (c *Client) GetDirectApps(ctx context.Context, domainID interface{}) ([]RawItem, error) {
	result, err := c.CallWithContext(ctx, MethodGetDirectApps, []interface{}{domainID})
	if err != nil {
		return nil, err
	}
	return rawItemsFromValue(result), nil
}

func rawItemsFromValue(v interface{}) []RawItem {
	items := []RawItem{}
	for _, data := range mapSliceFromValue(v) {
		items = append(items, RawItem(data))
	}
	return items
}
