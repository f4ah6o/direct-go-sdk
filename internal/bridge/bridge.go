package bridge

import (
	"context"
	"log"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/teams"
)

type DirectSender interface {
	Send(context.Context, model.DirectOutbound) error
}

type Service struct {
	cfg      *config.Config
	st       *store.Store
	graph    *teams.Client
	direct   DirectSender
	directIn <-chan model.DirectMessage
	teamsIn  <-chan model.DirectOutbound
	logger   *log.Logger
}

func NewService(cfg *config.Config, st *store.Store, graph *teams.Client, direct DirectSender, directIn <-chan model.DirectMessage, teamsIn <-chan model.DirectOutbound, logger *log.Logger) *Service {
	return &Service{cfg: cfg, st: st, graph: graph, direct: direct, directIn: directIn, teamsIn: teamsIn, logger: logger}
}

func (s *Service) Run(ctx context.Context) {
	go s.runDirectToTeams(ctx)
	go s.runTeamsToDirect(ctx)
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
	account, ok := s.cfg.Account(msg.AccountID)
	if !ok {
		return nil
	}
	ch := s.cfg.TeamsChannels[account.TeamsChannel]
	mapping, ok := s.st.GetByTalk(msg.AccountID, msg.TalkID)
	if !ok {
		rootID, err := retry(ctx, func() (string, error) {
			return s.graph.CreateRootMessage(ctx, ch.TeamID, ch.ChannelID, msg)
		})
		if err != nil {
			return err
		}
		mapping = store.ThreadMapping{
			AccountID: msg.AccountID,
			TalkID:    msg.TalkID,
			TeamID:    ch.TeamID,
			ChannelID: ch.ChannelID,
			RootID:    rootID,
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
		return s.graph.ReplyToThread(ctx, mapping.TeamID, mapping.ChannelID, mapping.RootID, msg)
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
			}
		}
	}
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
