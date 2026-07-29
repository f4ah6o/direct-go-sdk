package codexbridge

import (
	"context"
	"fmt"
	"log"
	"strings"
	"sync"

	"github.com/f4ah6o/direct-go-sdk/direct-go/debuglog"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/teams"
)

type CodexClient interface {
	StartThread(context.Context) (string, error)
	Turn(context.Context, string, string) (string, error)
}

type Service struct {
	cfg    config.CodexConfig
	store  *store.Store
	teams  *teams.Client
	codex  CodexClient
	logger *log.Logger
	mu     sync.Mutex
}

func NewService(cfg config.CodexConfig, st *store.Store, teamsClient *teams.Client, codexClient CodexClient, logger *log.Logger) *Service {
	return &Service{cfg: cfg, store: st, teams: teamsClient, codex: codexClient, logger: logger}
}

func (s *Service) HandleQuestion(ctx context.Context, in teams.CodexActivity) {
	go s.handleQuestion(context.Background(), in)
}

func (s *Service) HandleAnswer(ctx context.Context, in teams.CodexActivity) {
	go s.handleAnswer(context.Background(), in)
}

func (s *Service) handleQuestion(ctx context.Context, in teams.CodexActivity) {
	text := strings.TrimSpace(in.Text)
	if text == "" {
		_, _ = s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), in.ActivityID, "質問本文を入力してください。")
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	mapping, ok := s.store.GetCodexByQuestion(in.ConversationID, in.RootID)
	if !ok {
		threadID, err := s.codex.StartThread(ctx)
		if err != nil {
			s.replyQuestionError(ctx, in, err)
			return
		}
		mapping = store.CodexThreadMapping{
			QuestionConversationID: in.ConversationID,
			QuestionRootID:         in.RootID,
			QuestionServiceURL:     in.ServiceURL,
			CodexThreadID:          threadID,
			Status:                 "awaiting_codex",
		}
		if err := s.store.PutCodexMapping(mapping); err != nil {
			s.replyQuestionError(ctx, in, err)
			return
		}
	}
	answer, err := s.codex.Turn(ctx, mapping.CodexThreadID, questionPrompt(in))
	if err != nil {
		s.replyQuestionError(ctx, in, err)
		return
	}
	if escalation, ok := parseEscalation(answer); ok {
		if err := s.escalate(ctx, mapping, in, escalation); err != nil {
			s.replyQuestionError(ctx, in, err)
		}
		return
	}
	if _, err := s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), "", strings.TrimSpace(answer)); err != nil {
		s.logger.Printf("[codex-bridge] failed to reply question: %s", debuglog.SummarizePayload(err))
	}
	mapping.Status = "answered"
	_ = s.store.PutCodexMapping(mapping)
}

func (s *Service) handleAnswer(ctx context.Context, in teams.CodexActivity) {
	text := strings.TrimSpace(in.Text)
	if text == "" {
		return
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	mapping, ok := s.store.GetCodexByAnswer(in.ConversationID, in.RootID)
	if !ok {
		_, _ = s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), in.ActivityID, "対応する Codex 質問が見つかりません。")
		return
	}
	answer, err := s.codex.Turn(ctx, mapping.CodexThreadID, answerPrompt(in))
	if err != nil {
		_, _ = s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), in.ActivityID, "Codex への回答反映に失敗しました: "+err.Error())
		return
	}
	if _, err := s.teams.SendText(ctx, mapping.QuestionServiceURL, teamsThreadConversationID(mapping.QuestionConversationID, mapping.QuestionRootID), "", strings.TrimSpace(answer)); err != nil {
		s.logger.Printf("[codex-bridge] failed to post final answer: %s", debuglog.SummarizePayload(err))
		return
	}
	_, _ = s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), in.ActivityID, "質問 channel へ回答しました。")
	mapping.Status = "answered"
	_ = s.store.PutCodexMapping(mapping)
}

func (s *Service) escalate(ctx context.Context, mapping store.CodexThreadMapping, in teams.CodexActivity, escalation string) error {
	binding, ok := s.store.GetChannelBinding(s.cfg.AnswerAlias)
	if !ok {
		return fmt.Errorf("codex answer alias %q is not bound", s.cfg.AnswerAlias)
	}
	body := strings.Join([]string{
		"# Codex 確認依頼",
		"",
		"## 質問",
		strings.TrimSpace(in.Text),
		"",
		"## Codex の確認事項",
		strings.TrimSpace(escalation),
		"",
		"この thread に回答を reply してください。回答は Codex が質問者向けに整えて返します。",
	}, "\n")
	rootID, err := s.teams.CreateRootThreadText(ctx, binding.ServiceURL, teams.ChannelThreadBinding{
		TeamID:         binding.TeamID,
		ChannelID:      binding.ChannelID,
		ConversationID: binding.ConversationID,
		TenantID:       binding.TenantID,
		BotID:          binding.BotID,
	}, body)
	if err != nil {
		return err
	}
	mapping.AnswerConversationID = binding.ConversationID
	mapping.AnswerRootID = rootID
	mapping.AnswerServiceURL = binding.ServiceURL
	mapping.Status = "awaiting_human"
	if err := s.store.PutCodexMapping(mapping); err != nil {
		return err
	}
	_, _ = s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), "", "回答者 channel に確認依頼を出しました。")
	return nil
}

func (s *Service) replyQuestionError(ctx context.Context, in teams.CodexActivity, err error) {
	s.logger.Printf("[codex-bridge] question failed: %s", debuglog.SummarizePayload(err))
	_, _ = s.teams.SendText(ctx, in.ServiceURL, teamsThreadConversationID(in.ConversationID, in.RootID), in.ActivityID, "Codex bridge error: "+err.Error())
}

func questionPrompt(in teams.CodexActivity) string {
	return strings.TrimSpace(fmt.Sprintf("Teams の質問者 %s からの質問です。\n\n%s", displayName(in), in.Text))
}

func answerPrompt(in teams.CodexActivity) string {
	return strings.TrimSpace(fmt.Sprintf("回答者 %s から補足回答がありました。この内容を使って、元の質問者向けに簡潔で正確な最終回答を作ってください。\n\n%s", displayName(in), in.Text))
}

func displayName(in teams.CodexActivity) string {
	if strings.TrimSpace(in.FromName) != "" {
		return strings.TrimSpace(in.FromName)
	}
	return strings.TrimSpace(in.FromID)
}

func parseEscalation(answer string) (string, bool) {
	text := strings.TrimSpace(answer)
	if strings.HasPrefix(strings.ToUpper(text), "ESCALATE:") {
		return strings.TrimSpace(text[len("ESCALATE:"):]), true
	}
	return "", false
}

func teamsThreadConversationID(conversationID, rootID string) string {
	if rootID == "" || strings.Contains(conversationID, ";messageid=") {
		return conversationID
	}
	return conversationID + ";messageid=" + rootID
}
