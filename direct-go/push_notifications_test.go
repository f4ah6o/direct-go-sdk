package direct

import (
	"context"
	"testing"
)

func TestDisablePushNotificationParams(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	mockServer.OnDynamic("disable_push_notification", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "device1", "domain1")
		return true, nil
	})

	if err := client.DisablePushNotification(context.Background(), "device1", "domain1"); err != nil {
		t.Fatalf("DisablePushNotification failed: %v", err)
	}
}

func TestEnablePushNotificationParams(t *testing.T) {
	mockServer, client := newConnectedMockClient(t)
	settings := map[string]interface{}{"sound": true}
	mockServer.OnDynamic("enable_push_notification", func(params []interface{}) (interface{}, error) {
		assertParams(t, params, "device1", "domain1", settings)
		return true, nil
	})

	if err := client.EnablePushNotification(context.Background(), "device1", "domain1", settings); err != nil {
		t.Fatalf("EnablePushNotification failed: %v", err)
	}
}

func TestPushNotificationErrorPropagation(t *testing.T) {
	tests := []struct {
		name   string
		method string
		call   func(*Client) error
	}{
		{"disable", "disable_push_notification", func(c *Client) error {
			return c.DisablePushNotification(context.Background(), "device1", "domain1")
		}},
		{"enable", "enable_push_notification", func(c *Client) error {
			return c.EnablePushNotification(context.Background(), "device1", "domain1", map[string]interface{}{})
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
