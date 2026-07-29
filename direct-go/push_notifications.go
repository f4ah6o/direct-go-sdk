package direct

import "context"

// DisablePushNotification disables push notifications for the specified device.
func (c *Client) DisablePushNotification(ctx context.Context, deviceID, domainID interface{}) error {
	_, err := c.CallWithContext(ctx, MethodDisablePushNotification, []interface{}{deviceID, domainID})
	return err
}

// EnablePushNotification enables push notifications for the specified device.
func (c *Client) EnablePushNotification(ctx context.Context, deviceID, domainID, settings interface{}) error {
	_, err := c.CallWithContext(ctx, MethodEnablePushNotification, []interface{}{deviceID, domainID, settings})
	return err
}
