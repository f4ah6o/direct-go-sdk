package bridge

import (
	"context"
	"log"
	"strings"
	"sync"
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
	readIn   <-chan model.DirectReadReceipt
	teamsIn  <-chan model.DirectOutbound
	sentIn   <-chan model.DirectSent
	logger   *log.Logger
	mu       sync.Mutex
	pending  map[string]pendingDirect
	reads    map[string]pendingRead
}

type pendingDirect struct {
	outbound model.DirectOutbound
	expiry   time.Time
}

type pendingRead struct {
	receipt model.DirectReadReceipt
	expiry  time.Time
}

func NewService(account AccountLookup, st *store.Store, teamsClient *teams.Client, direct DirectSender, directIn <-chan model.DirectMessage, readIn <-chan model.DirectReadReceipt, teamsIn <-chan model.DirectOutbound, sentIn <-chan model.DirectSent, logger *log.Logger) *Service {
	return &Service{account: account, st: st, teams: teamsClient, direct: direct, directIn: directIn, readIn: readIn, teamsIn: teamsIn, sentIn: sentIn, logger: logger, pending: map[string]pendingDirect{}, reads: map[string]pendingRead{}}
}

func (s *Service) Run(ctx context.Context) {
	go s.runDirectToTeams(ctx)
	go s.runDirectReads(ctx)
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
			if outbound, ok := s.consumePendingDirectMessage(msg); ok {
				s.storeTeamsToDirectMessageRef(ctx, outbound, msg)
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
			return s.teams.CreateRootThread(ctx, binding.ServiceURL, teams.ChannelThreadBinding{
				TeamID:         binding.TeamID,
				ChannelID:      binding.ChannelID,
				ConversationID: binding.ConversationID,
				TenantID:       binding.TenantID,
				BotID:          binding.BotID,
			}, msg)
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
		s.storeDirectToTeamsMessageRef(ctx, store.TeamsMessageRef{
			AccountID:       msg.AccountID,
			TalkID:          msg.TalkID,
			DirectMessageID: msg.MessageID,
			ServiceURL:      binding.ServiceURL,
			ConversationID:  binding.ConversationID,
			ActivityID:      rootID,
			DirectSenderID:  msg.UserID,
		})
		s.logger.Printf("[bridge] mapped account=%s talk=%s to teams root=%s", msg.AccountID, msg.TalkID, rootID)
		return nil
	}
	replyID, err := retry(ctx, func() (string, error) {
		return s.teams.ReplyToThread(ctx, mapping.ServiceURL, mapping.ConversationID, mapping.RootID, msg)
	})
	if err != nil {
		return err
	}
	s.logger.Printf("[bridge] direct->teams reply posted: account=%s direct_message=%s teams_reply=%s conversation=%s", msg.AccountID, msg.MessageID, replyID, teamsThreadConversationID(mapping.ConversationID, mapping.RootID))
	if err := s.st.MarkTeamsMessage(replyID); err != nil {
		return err
	}
	s.storeDirectToTeamsMessageRef(ctx, store.TeamsMessageRef{
		AccountID:       msg.AccountID,
		TalkID:          msg.TalkID,
		DirectMessageID: msg.MessageID,
		ServiceURL:      mapping.ServiceURL,
		ConversationID:  teamsThreadConversationID(mapping.ConversationID, mapping.RootID),
		ActivityID:      replyID,
		DirectSenderID:  msg.UserID,
	})
	return nil
}

func (s *Service) runTeamsToDirect(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case msg := <-s.teamsIn:
			s.markPendingDirectMessage(msg)
			if err := s.direct.Send(ctx, msg); err != nil {
				s.clearPendingDirectMessage(msg)
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
			if sent.Err != nil {
				s.clearPendingDirectMessage(sent.Outbound)
				s.notifyTeamsSendFailure(ctx, sent)
				continue
			}
			if sent.MessageID != "" {
				if err := s.st.MarkDirectMessage(sent.MessageID); err != nil {
					s.logger.Printf("[bridge] failed to mark direct message id=%s err=%v", sent.MessageID, err)
				}
				if sent.Outbound.TeamsSource != nil {
					ref := store.TeamsMessageRef{
						AccountID:       sent.Outbound.AccountID,
						TalkID:          sent.Outbound.TalkID,
						DirectMessageID: sent.MessageID,
						ServiceURL:      sent.Outbound.TeamsSource.ServiceURL,
						ConversationID:  sent.Outbound.TeamsSource.ConversationID,
						ActivityID:      sent.Outbound.TeamsSource.ActivityID,
						DirectSenderID:  sent.SenderID,
					}
					s.storeDirectToTeamsMessageRef(ctx, ref)
				}
			}
			if sent.Outbound.Echo {
				s.echoTeamsSentMessage(ctx, sent.Outbound)
			}
		}
	}
}

func (s *Service) runDirectReads(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case receipt := <-s.readIn:
			if err := s.handleDirectReadReceipt(ctx, receipt); err != nil {
				s.logger.Printf("[bridge] direct read receipt failed: account=%s talk=%s err=%v", receipt.AccountID, receipt.TalkID, err)
			}
		}
	}
}

func (s *Service) handleDirectReadReceipt(ctx context.Context, receipt model.DirectReadReceipt) error {
	for _, messageID := range receipt.MessageIDs {
		reacted, err := s.reactToDirectRead(ctx, receipt, messageID)
		if err != nil {
			return err
		}
		if !reacted {
			s.markPendingRead(receipt, messageID)
		}
	}
	return nil
}

func (s *Service) reactToDirectRead(ctx context.Context, receipt model.DirectReadReceipt, messageID string) (bool, error) {
	ref, ok := s.st.GetTeamsMessageRef(receipt.AccountID, messageID)
	if !ok {
		s.logger.Printf("[bridge] read receipt pending without teams ref: account=%s direct_message=%s talk=%s", receipt.AccountID, messageID, receipt.TalkID)
		return false, nil
	}
	if !ref.ReactedAt.IsZero() {
		s.logger.Printf("[bridge] read receipt skipped already reacted: account=%s direct_message=%s teams_activity=%s", ref.AccountID, ref.DirectMessageID, ref.ActivityID)
		return true, nil
	}
	if ref.DirectSenderID == "" {
		s.logger.Printf("[bridge] read receipt skipped without sender id: account=%s direct_message=%s teams_activity=%s", ref.AccountID, ref.DirectMessageID, ref.ActivityID)
		return true, nil
	}
	if !hasReaderOtherThan(receipt.ReadUserIDs, ref.DirectSenderID) {
		s.logger.Printf("[bridge] read receipt skipped self-only: account=%s direct_message=%s sender=%s read_users=%d", ref.AccountID, ref.DirectMessageID, ref.DirectSenderID, len(receipt.ReadUserIDs))
		return true, nil
	}
	if err := s.teams.AddReaction(ctx, ref.ServiceURL, ref.ConversationID, ref.ActivityID, teams.ReactionEyes); err != nil {
		return false, err
	}
	if _, _, err := s.st.MarkTeamsReadReaction(ref.AccountID, ref.DirectMessageID); err != nil {
		return false, err
	}
	s.clearPendingRead(ref.AccountID, ref.DirectMessageID)
	s.logger.Printf("[bridge] added teams read reaction: account=%s direct_message=%s teams_activity=%s", ref.AccountID, ref.DirectMessageID, ref.ActivityID)
	return true, nil
}

func (s *Service) markPendingRead(receipt model.DirectReadReceipt, messageID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.prunePendingReadsLocked(time.Now())
	s.reads[directReadKey(receipt.AccountID, messageID)] = pendingRead{receipt: receipt, expiry: time.Now().Add(2 * time.Minute)}
}

func (s *Service) clearPendingRead(accountID, messageID string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.reads, directReadKey(accountID, messageID))
}

func (s *Service) consumePendingRead(accountID, messageID string) (model.DirectReadReceipt, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	s.prunePendingReadsLocked(now)
	key := directReadKey(accountID, messageID)
	pending, ok := s.reads[key]
	if !ok || now.After(pending.expiry) {
		delete(s.reads, key)
		return model.DirectReadReceipt{}, false
	}
	delete(s.reads, key)
	return pending.receipt, true
}

func (s *Service) prunePendingReadsLocked(now time.Time) {
	for key, pending := range s.reads {
		if now.After(pending.expiry) {
			delete(s.reads, key)
		}
	}
}

func directReadKey(accountID, messageID string) string {
	return accountID + ":" + messageID
}

func (s *Service) storeDirectToTeamsMessageRef(ctx context.Context, ref store.TeamsMessageRef) {
	if ref.DirectMessageID == "" {
		return
	}
	if ref.ActivityID == "" || ref.ConversationID == "" || ref.ServiceURL == "" {
		s.logger.Printf("[bridge] direct->teams message ref incomplete: account=%s direct_message=%s service_url=%t conversation_id=%t activity_id=%t", ref.AccountID, ref.DirectMessageID, ref.ServiceURL != "", ref.ConversationID != "", ref.ActivityID != "")
		return
	}
	if err := s.st.PutTeamsMessageRef(ref); err != nil {
		s.logger.Printf("[bridge] failed to store direct->teams message ref: account=%s direct_message=%s teams_activity=%s err=%v", ref.AccountID, ref.DirectMessageID, ref.ActivityID, err)
		return
	}
	s.logger.Printf("[bridge] stored direct->teams message ref: account=%s direct_message=%s teams_activity=%s", ref.AccountID, ref.DirectMessageID, ref.ActivityID)
	if receipt, ok := s.consumePendingRead(ref.AccountID, ref.DirectMessageID); ok {
		if reacted, err := s.reactToDirectRead(ctx, receipt, ref.DirectMessageID); err != nil {
			s.logger.Printf("[bridge] pending read receipt reaction failed: account=%s direct_message=%s err=%v", ref.AccountID, ref.DirectMessageID, err)
		} else if !reacted {
			s.markPendingRead(receipt, ref.DirectMessageID)
		}
	}
}

func hasReaderOtherThan(readUserIDs []string, senderID string) bool {
	for _, userID := range readUserIDs {
		if userID != "" && userID != senderID {
			return true
		}
	}
	return false
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

func pendingDirectMessageKey(msg model.DirectOutbound) string {
	return msg.AccountID + ":" + msg.TalkID + ":" + normalizePendingDirectMessageText(msg.Text)
}

func pendingDirectMessageKeyFromDirect(msg model.DirectMessage) string {
	return msg.AccountID + ":" + msg.TalkID + ":" + normalizePendingDirectMessageText(msg.Text)
}

func normalizePendingDirectMessageText(text string) string {
	return strings.TrimSpace(strings.ReplaceAll(text, "\r\n", "\n"))
}

func (s *Service) markPendingDirectMessage(msg model.DirectOutbound) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.prunePendingDirectMessagesLocked(time.Now())
	s.pending[pendingDirectMessageKey(msg)] = pendingDirect{outbound: msg, expiry: time.Now().Add(2 * time.Minute)}
}

func (s *Service) clearPendingDirectMessage(msg model.DirectOutbound) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.pending, pendingDirectMessageKey(msg))
}

func (s *Service) consumePendingDirectMessage(msg model.DirectMessage) (model.DirectOutbound, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now()
	s.prunePendingDirectMessagesLocked(now)
	key := pendingDirectMessageKeyFromDirect(msg)
	pending, ok := s.pending[key]
	if !ok || now.After(pending.expiry) {
		delete(s.pending, key)
		return model.DirectOutbound{}, false
	}
	delete(s.pending, key)
	return pending.outbound, true
}

func (s *Service) prunePendingDirectMessagesLocked(now time.Time) {
	for key, pending := range s.pending {
		if now.After(pending.expiry) {
			delete(s.pending, key)
		}
	}
}

func (s *Service) storeTeamsToDirectMessageRef(ctx context.Context, outbound model.DirectOutbound, msg model.DirectMessage) {
	if msg.MessageID == "" {
		return
	}
	if err := s.st.MarkDirectMessage(msg.MessageID); err != nil {
		s.logger.Printf("[bridge] failed to mark pending direct message id=%s err=%v", msg.MessageID, err)
	}
	if outbound.TeamsSource == nil {
		s.logger.Printf("[bridge] pending direct message has no teams source: account=%s direct_message=%s", msg.AccountID, msg.MessageID)
		return
	}
	s.storeDirectToTeamsMessageRef(ctx, store.TeamsMessageRef{
		AccountID:       outbound.AccountID,
		TalkID:          outbound.TalkID,
		DirectMessageID: msg.MessageID,
		ServiceURL:      outbound.TeamsSource.ServiceURL,
		ConversationID:  outbound.TeamsSource.ConversationID,
		ActivityID:      outbound.TeamsSource.ActivityID,
		DirectSenderID:  msg.UserID,
	})
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
