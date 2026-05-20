package teams

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
)

type AuthValidator interface {
	Validate(*http.Request) error
}

type BearerValidator struct{}

func (BearerValidator) Validate(r *http.Request) error {
	auth := r.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") || strings.TrimSpace(strings.TrimPrefix(auth, "Bearer ")) == "" {
		return fmt.Errorf("missing bearer token")
	}
	return nil
}

type Server struct {
	cfg        *config.Config
	client     *Client
	store      *store.Store
	out        chan<- model.DirectOutbound
	logger     *log.Logger
	validator  AuthValidator
	httpServer *http.Server
}

func NewServer(cfg *config.Config, client *Client, st *store.Store, out chan<- model.DirectOutbound, logger *log.Logger) *Server {
	validator := AuthValidator(BearerValidator{})
	s := &Server{cfg: cfg, client: client, store: st, out: out, logger: logger, validator: validator}
	mux := http.NewServeMux()
	mux.HandleFunc(cfg.Bot.EndpointPath, s.handleActivity)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	s.httpServer = &http.Server{
		Addr:              cfg.Server.ListenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	return s
}

func (s *Server) Run(ctx context.Context) error {
	errCh := make(chan error, 1)
	go func() {
		s.logger.Printf("[teams] bot endpoint listening on %s%s", s.cfg.Server.ListenAddr, s.cfg.Bot.EndpointPath)
		errCh <- s.httpServer.ListenAndServe()
	}()
	select {
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		return s.httpServer.Shutdown(shutdownCtx)
	case err := <-errCh:
		if err == http.ErrServerClosed {
			return nil
		}
		return err
	}
}

func (s *Server) handleActivity(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	if !s.cfg.Bot.DisableAuthValidation {
		if err := s.validator.Validate(r); err != nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	}
	var activity Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	go s.processActivity(context.Background(), activity)
}

func (s *Server) processActivity(ctx context.Context, activity Activity) {
	if activity.Type != "message" {
		return
	}
	if activity.From.ID == activity.Recipient.ID {
		return
	}
	if alias, ok := ParseBindAlias(activity); ok {
		s.bindChannel(ctx, alias, activity)
		return
	}
	if !MentionsRecipient(activity) {
		return
	}
	rootID := activity.ReplyToID
	if rootID == "" {
		s.logger.Printf("[teams] ignoring mentioned message without replyToId activity=%s", activity.ID)
		return
	}
	mapping, ok := s.store.GetByThread(activity.Conversation.ID, rootID)
	if !ok {
		s.logger.Printf("[teams] ignoring unmapped thread conversation=%s root=%s", activity.Conversation.ID, rootID)
		return
	}
	text := StripRecipientMention(activity)
	attachments := s.attachmentsFromActivity(ctx, activity)
	out := model.DirectOutbound{
		AccountID:   mapping.AccountID,
		TalkID:      mapping.TalkID,
		Text:        text,
		Attachments: attachments,
	}
	select {
	case s.out <- out:
	case <-ctx.Done():
	}
}

func (s *Server) bindChannel(ctx context.Context, alias string, activity Activity) {
	if _, ok := s.cfg.TeamsChannels[alias]; !ok {
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "unknown channel alias: "+alias)
		return
	}
	binding := store.TeamsChannelBinding{
		Alias:          alias,
		TeamID:         activity.ChannelData.Team.ID,
		ChannelID:      activity.ChannelData.Channel.ID,
		ConversationID: activity.Conversation.ID,
		ServiceURL:     activity.ServiceURL,
		TenantID:       firstNonEmpty(activity.ChannelData.Tenant.ID, activity.Conversation.TenantID),
		BotID:          activity.Recipient.ID,
	}
	if err := s.store.PutChannelBinding(binding); err != nil {
		s.logger.Printf("[teams] bind failed alias=%s err=%v", alias, err)
		return
	}
	_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "bound channel alias: "+alias)
}

func (s *Server) attachmentsFromActivity(ctx context.Context, activity Activity) []model.Attachment {
	out := make([]model.Attachment, 0, len(activity.Attachments))
	for _, att := range activity.Attachments {
		item := model.Attachment{Name: att.Name, ContentType: att.ContentType, URL: att.ContentURL}
		if data, contentType, err := s.client.DownloadAttachment(ctx, att, s.cfg.Attachments.MaxBytes); err == nil {
			item.Data = data
			item.ContentType = firstNonEmpty(contentType, item.ContentType)
		}
		out = append(out, item)
	}
	return out
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if strings.TrimSpace(value) != "" {
			return value
		}
	}
	return ""
}
