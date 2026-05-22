package direct

import (
	"context"
	"reflect"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-go/testutil"
)

func TestGetAccountControlRequests(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnSimple("get_account_control_requests", []interface{}{
		map[string]interface{}{"id": "req1", "version": int64(2), "user_id": "user1"},
	})

	requests, err := client.GetAccountControlRequests(context.Background())
	if err != nil {
		t.Fatalf("GetAccountControlRequests failed: %v", err)
	}
	if len(requests) != 1 || requests[0].ID != "req1" || requests[0].Version != int64(2) {
		t.Fatalf("unexpected requests: %+v", requests)
	}
	if got := mockServer.GetCallCount("get_account_control_requests"); got != 1 {
		t.Fatalf("get_account_control_requests calls = %d, want 1", got)
	}
}

func TestGetJoinedAccountControlGroup(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnSimple("get_joined_account_control_group", []interface{}{
		map[string]interface{}{"id": "group1", "name": "Control Group"},
	})

	groups, err := client.GetJoinedAccountControlGroup(context.Background())
	if err != nil {
		t.Fatalf("GetJoinedAccountControlGroup failed: %v", err)
	}
	if len(groups) != 1 || groups[0].ID != "group1" || groups[0].Raw["name"] != "Control Group" {
		t.Fatalf("unexpected groups: %+v", groups)
	}
	if got := mockServer.GetCallCount("get_joined_account_control_group"); got != 1 {
		t.Fatalf("get_joined_account_control_group calls = %d, want 1", got)
	}
}

func TestAcceptAccountControlRequestParams(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("accept_account_control_request", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "req1", int64(3))
		return true, nil
	})

	if err := client.AcceptAccountControlRequest(context.Background(), "req1", int64(3)); err != nil {
		t.Fatalf("AcceptAccountControlRequest failed: %v", err)
	}
}

func TestRejectAccountControlRequestParams(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("reject_account_control_request", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "req1", int64(3))
		return true, nil
	})

	if err := client.RejectAccountControlRequest(context.Background(), "req1", int64(3)); err != nil {
		t.Fatalf("RejectAccountControlRequest failed: %v", err)
	}
}

func TestAccountControlErrorPropagation(t *testing.T) {
	tests := []struct {
		name   string
		method string
		call   func(*Client) error
	}{
		{"get requests", "get_account_control_requests", func(c *Client) error {
			_, err := c.GetAccountControlRequests(context.Background())
			return err
		}},
		{"get groups", "get_joined_account_control_group", func(c *Client) error {
			_, err := c.GetJoinedAccountControlGroup(context.Background())
			return err
		}},
		{"accept", "accept_account_control_request", func(c *Client) error {
			return c.AcceptAccountControlRequest(context.Background(), "req1", int64(1))
		}},
		{"reject", "reject_account_control_request", func(c *Client) error {
			return c.RejectAccountControlRequest(context.Background(), "req1", int64(1))
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

func newConnectedMockClient(t *testing.T) (*testutil.MockServer, *Client) {
	t.Helper()
	mockServer := testutil.NewMockServer()
	t.Cleanup(mockServer.Close)

	client := NewClient(Options{Endpoint: mockServer.URL()})
	if err := client.Connect(); err != nil {
		t.Fatalf("Connect failed: %v", err)
	}
	t.Cleanup(func() {
		if err := client.Close(); err != nil {
			t.Fatalf("Close failed: %v", err)
		}
	})
	return mockServer, client
}

func assertParams(t *testing.T, got []interface{}, want ...interface{}) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("params length = %d, want %d: %#v", len(got), len(want), got)
	}
	for i := range want {
		if !reflect.DeepEqual(got[i], want[i]) {
			t.Fatalf("param[%d] = %#v, want %#v (all params %#v)", i, got[i], want[i], got)
		}
	}
}
