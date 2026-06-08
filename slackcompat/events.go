package slackcompat

import (
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

func ConvertMessageEvent(mapper Mapper, teamID, accountID string, msg direct.ReceivedMessage) EventsEnvelope {
	created := msg.Timestamp
	if created.IsZero() && msg.Created > 0 {
		created = time.Unix(msg.Created, 0).UTC()
	}
	ts := mapper.Timestamp(msg.ID, created)
	eventTime := time.Now().Unix()
	if parsed, ok := parseSlackTS(ts); ok {
		eventTime = parsed.Unix()
	}
	return EventsEnvelope{
		TeamID:    teamID,
		Type:      "event_callback",
		EventID:   "Ev" + encodeID(accountID+"|"+msg.ID+"|"+ts),
		EventTime: eventTime,
		Event: SlackEvent{
			Type:    "message",
			Channel: mapper.ChannelID(accountID, msg.TalkID),
			User:    mapper.UserID(msg.UserID),
			Text:    msg.Text,
			TS:      ts,
		},
	}
}
