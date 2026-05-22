package direct

import "fmt"

const DefaultBotOS = "bot"

func (c *Client) CreateAccessToken(email, password, deviceID, osName string) (string, error) {
	if osName == "" {
		osName = DefaultBotOS
	}
	result, err := c.Call(MethodCreateAccessToken, []interface{}{email, password, deviceID, osName, ""})
	if err != nil {
		return "", err
	}
	token := accessTokenFromResult(result)
	if token == "" {
		return "", fmt.Errorf("direct did not return an access token")
	}
	return token, nil
}

func accessTokenFromResult(result interface{}) string {
	if token, ok := result.(string); ok {
		return token
	}
	if m, ok := result.(map[string]interface{}); ok {
		if token, ok := m["access_token"].(string); ok {
			return token
		}
	}
	if arr, ok := result.([]interface{}); ok && len(arr) > 0 {
		if token, ok := arr[0].(string); ok {
			return token
		}
	}
	return ""
}
