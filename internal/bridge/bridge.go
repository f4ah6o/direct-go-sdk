package bridge

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/teams"
)

type DirectSender interface {
	Send(context.Context, model.DirectOutbound) error
}

type AccountLookup func(string) (config.AccountConfig, bool)

type Service struct {
	account  AccountLookup
	st       *store.Store
	teams    *teams.Client
	direct   DirectSender
	directIn <-chan model.DirectMessage
	teamsIn  <-chan model.DirectOutbound
	sentIn   <-chan model.DirectSent
	logger   *log.Logger
}

func NewService(account AccountLookup, st *store.Store, teamsClient *teams.Client, direct DirectSender, directIn <-chan model.DirectMessage, teamsIn <-chan model.DirectOutbound, sentIn <-chan model.DirectSent, logger *log.Logger) *Service {
	return &Service{account: account, st: st, teams: teamsClient, direct: direct, directIn: directIn, teamsIn: teamsIn, sentIn: sentIn, logger: logger}
}

func (s *Service) Run(ctx context.Context) {
	go s.runDirectToTeams(ctx)
	go s.runTeamsToDirect(ctx)
	go s.runDirectSent(ctx)
}

func (s *Service) runDirectToTeams(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.directIn:
			if msg.MessageID != "" && s.st.IsSentDirectMessage(msg.MessageID) {
				continue
			}
			if err := s.handleDirectMessage(ctx, msg); err != nil {
				s.logger.Printf("[bridge] direct->teams failed: account=%s talk=%s err=%v", msg.AccountID, msg.TalkID, err)
			}
		}
	}
}

func (s *Service) handleDirectMessage(ctx context.Context, msg model.DirectMessage) error {
	account, ok := s.account(msg.AccountID)
	if !ok {
		return nil
	}
	binding, ok := s.st.GetChannelBinding(account.TeamsChannel)
	if !ok {
		s.logger.Printf("[bridge] teams channel alias %q is not bound; run @bot bind %s in Teams", account.TeamsChannel, account.TeamsChannel)
		return nil
	}
	mapping, ok := s.st.GetByTalk(msg.AccountID, msg.TalkID)
	if !ok {
		rootID, err := retry(ctx, func() (string, error) {
			return s.teams.CreateRootMessage(ctx, binding.ServiceURL, binding.ConversationID, msg)
		})
		if err != nil {
			return err
		}
		mapping = store.ThreadMapping{
			AccountID:      msg.AccountID,
			TalkID:         msg.TalkID,
			ChannelAlias:   account.TeamsChannel,
			ConversationID: binding.ConversationID,
			ServiceURL:     binding.ServiceURL,
			RootID:         rootID,
		}
		if err := s.st.PutMapping(mapping); err != nil {
			return err
		}
		if err := s.st.MarkTeamsMessage(rootID); err != nil {
			return err
		}
		s.logger.Printf("[bridge] mapped account=%s talk=%s to teams root=%s", msg.AccountID, msg.TalkID, rootID)
		return nil
	}
	replyID, err := retry(ctx, func() (string, error) {
		return s.teams.ReplyToThread(ctx, mapping.ServiceURL, mapping.ConversationID, mapping.RootID, msg)
	})
	if err != nil {
		return err
	}
	return s.st.MarkTeamsMessage(replyID)
}

func (s *Service) runTeamsToDirect(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.teamsIn:
			if err := s.direct.Send(ctx, msg); err != nil {
				s.logger.Printf("[bridge] teams->direct failed: account=%s talk=%s err=%v", msg.AccountID, msg.TalkID, err)
				s.notifyTeamsSendFailure(ctx, model.DirectSent{Outbound: msg, Err: err})
			}
		}
	}
}

func (s *Service) runDirectSent(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case sent := <-s.sentIn:
			if sent.MessageID != "" {
				if err := s.st.MarkDirectMessage(sent.MessageID); err != nil {
					s.logger.Printf("[bridge] failed to mark direct message id=%s err=%v", sent.MessageID, err)
				}
			}
			if sent.Err != nil {
				s.notifyTeamsSendFailure(ctx, sent)
				continue
			}
			if sent.Outbound.Echo {
				s.echoTeamsSentMessage(ctx, sent.Outbound)
			}
		}
	}
}

func (s *Service) notifyTeamsSendFailure(ctx context.Context, sent model.DirectSent) {
	mapping, ok := s.st.GetByTalk(sent.Outbound.AccountID, sent.Outbound.TalkID)
	if !ok {
		return
	}
	text := "❌ failed to send to direct: " + sent.Err.Error()
	if _, err := s.teams.SendText(ctx, mapping.ServiceURL, teamsThreadConversationID(mapping.ConversationID, mapping.RootID), "", text); err != nil {
		s.logger.Printf("[bridge] failed to notify teams send failure: account=%s talk=%s err=%v", sent.Outbound.AccountID, sent.Outbound.TalkID, err)
	}
}

func (s *Service) echoTeamsSentMessage(ctx context.Context, msg model.DirectOutbound) {
	mapping, ok := s.st.GetByTalk(msg.AccountID, msg.TalkID)
	if !ok {
		return
	}
	text := "echo: " + msg.Text
	if _, err := s.teams.SendText(ctx, mapping.ServiceURL, teamsThreadConversationID(mapping.ConversationID, mapping.RootID), "", text); err != nil {
		s.logger.Printf("[bridge] failed to echo teams sent message: account=%s talk=%s err=%v", msg.AccountID, msg.TalkID, err)
	}
}

func teamsThreadConversationID(conversationID, rootID string) string {
	if rootID == "" || strings.Contains(conversationID, ";messageid=") {
		return conversationID
	}
	return conversationID + ";messageid=" + rootID
}

func retry[T any](ctx context.Context, fn func() (T, error)) (T, error) {
	var zero T
	var last error
	backoff := time.Second
	for attempt := 0; attempt < 5; attempt++ {
		v, err := fn()
		if err == nil {
			return v, nil
		}
		last = err
		timer := time.NewTimer(backoff)
		select {
		case <-ctx.Done():
			timer.Stop()
			return zero, ctx.Err()
		case <-timer.C:
		}
		backoff *= 2
		if backoff > 30*time.Second {
			backoff = 30 * time.Second
		}
	}
	return zero, last
}
