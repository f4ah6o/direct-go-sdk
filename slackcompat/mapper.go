package slackcompat

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	channelPrefix = "C"
	userPrefix    = "U"
	botPrefix     = "B"
)

type Mapper struct{}

func NewMapper() Mapper {
	return Mapper{}
}

func (m Mapper) ChannelID(accountID, talkID string) string {
	return channelPrefix + encodeID(accountID+"|"+talkID)
}

func (m Mapper) TalkID(channelID string) (accountID, talkID string, ok bool) {
	raw, ok := decodeID(channelID, channelPrefix)
	if !ok {
		return "", "", false
	}
	accountID, talkID, ok = strings.Cut(raw, "|")
	return accountID, talkID, ok && accountID != "" && talkID != ""
}

func (m Mapper) UserID(directUserID string) string {
	return userPrefix + encodeID(directUserID)
}

func (m Mapper) DirectUserID(slackUserID string) (string, bool) {
	return decodeID(slackUserID, userPrefix)
}

func (m Mapper) BotID(accountID string) string {
	return botPrefix + encodeID(accountID)
}

func (m Mapper) Timestamp(messageID string, createdAt time.Time) string {
	if !createdAt.IsZero() {
		return formatSlackTS(createdAt)
	}
	if messageID != "" {
		return messageID
	}
	return formatSlackTS(time.Now().UTC())
}

func formatSlackTS(t time.Time) string {
	t = t.UTC()
	return fmt.Sprintf("%d.%06d", t.Unix(), t.Nanosecond()/1000)
}

func parseSlackTS(ts string) (time.Time, bool) {
	sec, frac, ok := strings.Cut(ts, ".")
	if !ok {
		return time.Time{}, false
	}
	s, err := strconv.ParseInt(sec, 10, 64)
	if err != nil {
		return time.Time{}, false
	}
	if len(frac) > 6 {
		frac = frac[:6]
	}
	for len(frac) < 6 {
		frac += "0"
	}
	us, err := strconv.ParseInt(frac, 10, 64)
	if err != nil {
		return time.Time{}, false
	}
	return time.Unix(s, us*1000).UTC(), true
}

func encodeID(raw string) string {
	return base64.RawURLEncoding.EncodeToString([]byte(raw))
}

func decodeID(value, prefix string) (string, bool) {
	if !strings.HasPrefix(value, prefix) {
		return "", false
	}
	raw, err := base64.RawURLEncoding.DecodeString(strings.TrimPrefix(value, prefix))
	if err != nil {
		return "", false
	}
	out := string(raw)
	return out, out != ""
}
