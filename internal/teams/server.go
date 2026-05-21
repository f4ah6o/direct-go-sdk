package teams

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
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
	mux.HandleFunc("/files/direct", s.handleDirectFile)
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

func (s *Server) handleDirectFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	accountID := r.URL.Query().Get("account")
	rawURL := r.URL.Query().Get("url")
	if accountID == "" || rawURL == "" {
		http.Error(w, "missing account or url", http.StatusBadRequest)
		return
	}
	account, ok := s.cfg.Account(accountID)
	if !ok {
		http.Error(w, "unknown account", http.StatusNotFound)
		return
	}
	fileURL, err := url.Parse(rawURL)
	if err != nil || fileURL.Scheme != "https" || fileURL.Host != "api.direct4b.com" || !strings.HasPrefix(fileURL.Path, "/albero-app-server/files/") {
		http.Error(w, "invalid direct file url", http.StatusBadRequest)
		return
	}
	token := os.Getenv(account.TokenEnv)
	if token == "" {
		http.Error(w, "direct token is not available", http.StatusInternalServerError)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), 60*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, fileURL.String(), nil)
	if err != nil {
		http.Error(w, "invalid direct file request", http.StatusBadRequest)
		return
	}
	req.Header.Set("Authorization", "ALB "+token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		s.logger.Printf("[teams] direct file proxy failed account=%s err=%v", accountID, err)
		http.Error(w, "direct file fetch failed", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		b, _ := io.ReadAll(io.LimitReader(resp.Body, 1024))
		s.logger.Printf("[teams] direct file proxy status=%d account=%s body=%s", resp.StatusCode, accountID, string(b))
		http.Error(w, "direct file fetch failed", http.StatusBadGateway)
		return
	}
	if ct := resp.Header.Get("Content-Type"); ct != "" {
		w.Header().Set("Content-Type", ct)
	}
	if cd := resp.Header.Get("Content-Disposition"); cd != "" {
		w.Header().Set("Content-Disposition", cd)
	}
	if cl := resp.Header.Get("Content-Length"); cl != "" {
		w.Header().Set("Content-Length", cl)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = io.Copy(w, resp.Body)
}

func (s *Server) processActivity(ctx context.Context, activity Activity) {
	if activity.Type == "conversationUpdate" {
		if BotWasAdded(activity) {
			s.sendWelcome(ctx, activity)
		}
		return
	}
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
	if s.handleCommand(ctx, activity) {
		return
	}
	if !MentionsRecipient(activity) {
		return
	}
	conversationID, rootID := threadReference(activity)
	if rootID == "" {
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, rootOnlyHelpText())
		return
	}
	mapping, ok := s.store.GetByThread(conversationID, rootID)
	if !ok {
		s.logger.Printf("[teams] ignoring unmapped thread conversation=%s root=%s", conversationID, rootID)
		return
	}
	text := StripRecipientMention(activity)
	echo := false
	trimmedText := strings.TrimSpace(text)
	if fields := strings.Fields(trimmedText); len(fields) > 0 && strings.EqualFold(fields[0], "echo") {
		echo = true
		text = strings.TrimSpace(trimmedText[len(fields[0]):])
	}
	attachments := s.attachmentsFromActivity(ctx, activity)
	out := model.DirectOutbound{
		AccountID:   mapping.AccountID,
		TalkID:      mapping.TalkID,
		Text:        text,
		Attachments: attachments,
		Echo:        echo,
	}
	select {
	case s.out <- out:
	case <-ctx.Done():
	}
}

func (s *Server) sendWelcome(ctx context.Context, activity Activity) {
	text := "direct bridge is ready. In the target channel, mention me with `bind <alias>` to connect this Teams channel to a configured direct account route."
	if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, "", text); err != nil {
		s.logger.Printf("[teams] welcome message failed conversation=%s err=%v", activity.Conversation.ID, err)
	}
}

func (s *Server) handleCommand(ctx context.Context, activity Activity) bool {
	command := ParseCommand(activity)
	switch command {
	case "bind":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "usage: @direct bind <alias>")
		return true
	case "hi", "hello":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "Hi. Use `@direct bind <alias>` in a channel, or reply in a bridged thread with `@direct <message>`.")
		return true
	case "help":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, rootOnlyHelpText())
		return true
	default:
		return false
	}
}

func rootOnlyHelpText() string {
	return "This bridge accepts `@direct bind <alias>` in a Teams channel. To send a Teams reply back to direct, reply inside a thread created from direct and mention `@direct`."
}

func (s *Server) bindChannel(ctx context.Context, alias string, activity Activity) {
	if activity.ReplyToID != "" {
		if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ bind must be sent as a new channel message, not as a thread reply."); err != nil {
			s.logger.Printf("[teams] bind placement response failed alias=%s err=%v", alias, err)
		}
		return
	}
	if _, ok := s.cfg.TeamsChannels[alias]; !ok {
		if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ unknown channel alias: "+alias); err != nil {
			s.logger.Printf("[teams] unknown alias response failed alias=%s err=%v", alias, err)
		}
		return
	}
	binding := store.TeamsChannelBinding{
		Alias:          alias,
		TeamID:         activity.ChannelData.Team.ID,
		ChannelID:      activity.ChannelData.Channel.ID,
		ConversationID: channelConversationID(activity),
		ServiceURL:     activity.ServiceURL,
		TenantID:       firstNonEmpty(activity.ChannelData.Tenant.ID, activity.Conversation.TenantID),
		BotID:          activity.Recipient.ID,
	}
	if err := s.store.PutChannelBinding(binding); err != nil {
		s.logger.Printf("[teams] bind failed alias=%s err=%v", alias, err)
		return
	}
	if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "✅ bound channel alias: "+alias); err != nil {
		s.logger.Printf("[teams] bind response failed alias=%s err=%v", alias, err)
	}
}

func channelConversationID(activity Activity) string {
	if activity.ChannelData.Channel.ID != "" {
		return activity.ChannelData.Channel.ID
	}
	return strings.SplitN(activity.Conversation.ID, ";messageid=", 2)[0]
}

func threadReference(activity Activity) (conversationID, rootID string) {
	conversationID = channelConversationID(activity)
	if activity.ReplyToID != "" {
		return conversationID, activity.ReplyToID
	}
	_, rootID, ok := strings.Cut(activity.Conversation.ID, ";messageid=")
	if !ok {
		return conversationID, ""
	}
	return conversationID, rootID
}

func (s *Server) attachmentsFromActivity(ctx context.Context, activity Activity) []model.Attachment {
	out := make([]model.Attachment, 0, len(activity.Attachments))
	for _, att := range activity.Attachments {
		item := model.Attachment{Name: att.Name, ContentType: att.ContentType, URL: att.ContentURL}
		if data, contentType, err := s.client.DownloadAttachment(ctx, att, s.cfg.Attachments.MaxBytes); err == nil {
			item.Data = data
			item.ContentType = firstNonEmpty(contentType, item.ContentType)
		}
		if strings.TrimSpace(item.Name) == "" && strings.TrimSpace(item.URL) == "" && len(item.Data) == 0 {
			continue
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
