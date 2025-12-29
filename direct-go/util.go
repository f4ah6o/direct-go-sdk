package direct

import "time"

// UnixTime converts an int64 Unix timestamp to time.Time.
// This is a convenience function to reduce code duplication when parsing
// MessagePack timestamps.
func UnixTime(v int64) time.Time {
	return time.Unix(v, 0)
}

// UnixTimeFromInterface extracts an int64 timestamp from an interface{}
// and converts it to time.Time. Returns zero time if the value is not an int64.
func UnixTimeFromInterface(v interface{}) time.Time {
	if i, ok := v.(int64); ok {
		return time.Unix(i, 0)
	}
	return time.Time{}
}

// ToInt64 extracts an int64 value from an interface{}.
// Supports int64, int, and float64 types.
// Returns the value and true if successful, 0 and false otherwise.
func ToInt64(v interface{}) (int64, bool) {
	switch t := v.(type) {
	case int64:
		return t, true
	case int:
		return int64(t), true
	case float64:
		return int64(t), true
	default:
		return 0, false
	}
}
