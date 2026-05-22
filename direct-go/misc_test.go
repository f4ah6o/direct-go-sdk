package direct

import (
	"context"
	"testing"
)

func TestGetActionsParamsAndParsing(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("get_actions", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "domain1", "talk1", int8(15), "marker1", int8(20), "since1", "max1")
		return []interface{}{map[string]interface{}{"id": "action1"}}, nil
	})

	actions, err := client.GetActions(context.Background(), "domain1", "talk1", 15, "marker1", 20, &ActionQueryOptions{SinceID: "since1", MaxID: "max1"})
	if err != nil {
		t.Fatalf("GetActions failed: %v", err)
	}
	if len(actions) != 1 || actions[0]["id"] != "action1" {
		t.Fatalf("unexpected actions: %+v", actions)
	}
}

func TestMiscGetterParamsAndParsing(t *testing.T) {
	tests := []struct {
		name   string
		method string
		call   func(*Client) ([]RawItem, error)
	}{
		{"solutions", "get_solutions", func(c *Client) ([]RawItem, error) {
			return c.GetSolutions(context.Background(), "domain1", "marker1")
		}},
		{"stampsets", "get_stampsets", func(c *Client) ([]RawItem, error) {
			return c.GetStampsets(context.Background(), "domain1")
		}},
		{"flow badges", "get_flow_notification_badges", func(c *Client) ([]RawItem, error) {
			return c.GetFlowNotificationBadges(context.Background(), "domain1")
		}},
		{"direct apps", "get_direct_apps", func(c *Client) ([]RawItem, error) {
			return c.GetDirectApps(context.Background(), "domain1")
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer, client := newConnectedMockClient(t)
			mockServer.OnSimple(tt.method, []interface{}{map[string]interface{}{"id": tt.method}})
			items, err := tt.call(client)
			if err != nil {
				t.Fatalf("%s failed: %v", tt.method, err)
			}
			if len(items) != 1 || items[0]["id"] != tt.method {
				t.Fatalf("unexpected items: %+v", items)
			}
			if got := mockServer.GetCallCount(tt.method); got != 1 {
				t.Fatalf("%s calls = %d, want 1", tt.method, got)
			}
		})
	}
}

func TestMiscGetterErrorPropagation(t *testing.T) {
	tests := []struct {
		name   string
		method string
		call   func(*Client) error
	}{
		{"actions", "get_actions", func(c *Client) error {
			_, err := c.GetActions(context.Background(), "domain1", "talk1", 15, nil, 20, nil)
			return err
		}},
		{"solutions", "get_solutions", func(c *Client) error {
			_, err := c.GetSolutions(context.Background(), "domain1", nil)
			return err
		}},
		{"stampsets", "get_stampsets", func(c *Client) error {
			_, err := c.GetStampsets(context.Background(), "domain1")
			return err
		}},
		{"flow badges", "get_flow_notification_badges", func(c *Client) error {
			_, err := c.GetFlowNotificationBadges(context.Background(), "domain1")
			return err
		}},
		{"direct apps", "get_direct_apps", func(c *Client) error {
			_, err := c.GetDirectApps(context.Background(), "domain1")
			return err
		}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer, client := newConnectedMockClient(t)
			mockServer.OnError(tt.method, "boom")
			if err := tt.call(client); err == nil {
				t.Fatalf("expected error")
			}
		})
	}
}
