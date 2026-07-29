package teams

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
)

type Server struct {
	cfg        *config.Config
	client     *Client
	store      *store.Store
	out        chan<- model.DirectOutbound
	logger     *log.Logger
	validator  AuthValidator
	httpServer *http.Server
	hasChannel func(string) bool
	account    func(string) (config.AccountConfig, bool)
	token      func(string) (string, bool)
	health     func() (bool, interface{})
	codex      CodexHandler
}

type CodexActivity struct {
	ServiceURL     string
	ConversationID string
	RootID         string
	ActivityID     string
	Text           string
	FromID         string
	FromName       string
	Attachments    []model.Attachment
}

type CodexHandler interface {
	HandleQuestion(context.Context, CodexActivity)
	HandleAnswer(context.Context, CodexActivity)
}

func NewServer(cfg *config.Config, client *Client, st *store.Store, out chan<- model.DirectOutbound, logger *log.Logger, opts ...func(*Server)) *Server {
	validator := AuthValidator(NewBotFrameworkValidator(cfg.Bot))
	s := &Server{
		cfg:        cfg,
		client:     client,
		store:      st,
		out:        out,
		logger:     logger,
		validator:  validator,
		hasChannel: func(alias string) bool { _, ok := cfg.TeamsChannels[alias]; return ok },
		account:    cfg.Account,
		token:      func(accountID string) (string, bool) { return "", false },
		health:     func() (bool, interface{}) { return true, nil },
	}
	for _, opt := range opts {
		opt(s)
	}
	mux := http.NewServeMux()
	mux.HandleFunc(cfg.Bot.EndpointPath, s.handleActivity)
	mux.HandleFunc("/files/direct", s.handleDirectFile)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		ok, details := s.health()
		if details == nil {
			if ok {
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte("ok\n"))
				return
			}
			http.Error(w, "unhealthy", http.StatusServiceUnavailable)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		if !ok {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": ok, "direct_accounts": details})
	})
	s.httpServer = &http.Server{
		Addr:              cfg.Server.ListenAddr,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
	}
	return s
}

func WithHealthCheck(health func() (bool, interface{})) func(*Server) {
	return func(s *Server) {
		s.health = health
	}
}

func WithRuntimeLookups(hasChannel func(string) bool, account func(string) (config.AccountConfig, bool), token func(string) (string, bool)) func(*Server) {
	return func(s *Server) {
		s.hasChannel = hasChannel
		s.account = account
		s.token = token
	}
}

func WithCodexHandler(handler CodexHandler) func(*Server) {
	return func(s *Server) {
		s.codex = handler
	}
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
	r.Body = http.MaxBytesReader(w, r.Body, 1<<20)
	var activity Activity
	if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	if !s.authValidationBypassed(r) {
		if err := s.validator.Validate(r.Context(), r, activity); err != nil {
			s.logger.Printf("[teams] unauthorized activity: %v", err)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
	}
	w.WriteHeader(http.StatusAccepted)
	go s.processActivity(context.Background(), activity)
}

func (s *Server) authValidationBypassed(r *http.Request) bool {
	if !s.cfg.Bot.DisableAuthValidation {
		return false
	}
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}
	ip := net.ParseIP(strings.Trim(host, "[]"))
	return ip != nil && ip.IsLoopback()
}

func (s *Server) handleDirectFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	accountID := r.URL.Query().Get("account")
	rawURL := r.URL.Query().Get("url")
	exp, err := strconv.ParseInt(r.URL.Query().Get("exp"), 10, 64)
	if err != nil {
		http.Error(w, "invalid signed file url", http.StatusBadRequest)
		return
	}
	if err := validateDirectFileSignature(s.cfg.Bot, accountID, rawURL, exp, r.URL.Query().Get("sig"), time.Now()); err != nil {
		http.Error(w, "invalid signed file url", http.StatusForbidden)
		return
	}
	account, ok := s.account(accountID)
	if !ok {
		http.Error(w, "unknown account", http.StatusNotFound)
		return
	}
	fileURL, err := url.Parse(rawURL)
	if err != nil || fileURL.Scheme != "https" || fileURL.Host != "api.direct4b.com" || !strings.HasPrefix(fileURL.Path, "/albero-app-server/files/") {
		http.Error(w, "invalid direct file url", http.StatusBadRequest)
		return
	}
	token, ok := s.token(account.ID)
	if !ok || token == "" {
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
	if alias, ok := ParseUnbindAlias(activity); ok {
		s.unbindChannel(ctx, alias, activity)
		return
	}
	if s.handleCodexAnswerActivity(ctx, activity) {
		return
	}
	if !MentionsRecipient(activity) {
		return
	}
	if s.handleCodexActivity(ctx, activity) {
		return
	}
	if s.handleCommand(ctx, activity) {
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
	trimmedText := strings.TrimSpace(text)
	fields := strings.Fields(trimmedText)
	if len(fields) == 0 {
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, threadReplyUsageText())
		return
	}
	switch strings.ToLower(fields[0]) {
	case "reply":
		text = strings.TrimSpace(trimmedText[len(fields[0]):])
	default:
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, threadReplyUsageText())
		return
	}
	if strings.TrimSpace(text) == "" && len(activity.Attachments) == 0 {
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, threadReplyUsageText())
		return
	}
	text = appendTeamsSenderName(text, activity.From.Name)
	attachments := s.attachmentsFromActivity(ctx, activity)
	out := model.DirectOutbound{
		AccountID:   mapping.AccountID,
		TalkID:      mapping.TalkID,
		Text:        text,
		Attachments: attachments,
		TeamsSource: &model.TeamsSource{
			ServiceURL:     activity.ServiceURL,
			ConversationID: activity.Conversation.ID,
			ActivityID:     activity.ID,
		},
	}
	select {
	case s.out <- out:
	case <-ctx.Done():
	}
}

func (s *Server) handleCodexAnswerActivity(ctx context.Context, activity Activity) bool {
	if s.codex == nil || s.cfg == nil || s.store == nil || !s.cfg.Codex.Enabled {
		return false
	}
	conversationID, rootID := threadReference(activity)
	if rootID == "" {
		return false
	}
	if _, ok := s.store.GetCodexByAnswer(conversationID, rootID); !ok {
		return false
	}
	in := CodexActivity{
		ServiceURL:     activity.ServiceURL,
		ConversationID: conversationID,
		RootID:         rootID,
		ActivityID:     activity.ID,
		Text:           StripRecipientMention(activity),
		FromID:         activity.From.ID,
		FromName:       activity.From.Name,
		Attachments:    s.attachmentsFromActivity(ctx, activity),
	}
	s.codex.HandleAnswer(ctx, in)
	return true
}

func (s *Server) handleCodexActivity(ctx context.Context, activity Activity) bool {
	if s.codex == nil || s.cfg == nil || s.store == nil || !s.cfg.Codex.Enabled {
		return false
	}
	conversationID, rootID := threadReference(activity)
	if rootID == "" {
		rootID = activity.ID
	}
	bindingAlias := s.boundAliasForConversation(conversationID)
	if bindingAlias == "" {
		return false
	}
	in := CodexActivity{
		ServiceURL:     activity.ServiceURL,
		ConversationID: conversationID,
		RootID:         rootID,
		ActivityID:     activity.ID,
		Text:           StripRecipientMention(activity),
		FromID:         activity.From.ID,
		FromName:       activity.From.Name,
		Attachments:    s.attachmentsFromActivity(ctx, activity),
	}
	switch bindingAlias {
	case s.cfg.Codex.QuestionAlias:
		if !s.codexUserAllowed(activity.From.ID) {
			_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "このユーザーは Codex bridge を利用できません。")
			return true
		}
		s.codex.HandleQuestion(ctx, in)
		return true
	case s.cfg.Codex.AnswerAlias:
		s.codex.HandleAnswer(ctx, in)
		return true
	default:
		return false
	}
}

func (s *Server) boundAliasForConversation(conversationID string) string {
	for _, binding := range s.store.ListChannelBindings() {
		if binding.ConversationID == conversationID {
			return binding.Alias
		}
	}
	return ""
}

func (s *Server) codexUserAllowed(userID string) bool {
	if len(s.cfg.Codex.AllowedUserIDs) == 0 {
		return true
	}
	for _, allowed := range s.cfg.Codex.AllowedUserIDs {
		if allowed == userID {
			return true
		}
	}
	return false
}

func appendTeamsSenderName(text, displayName string) string {
	name := strings.TrimSpace(displayName)
	if name == "" {
		return text
	}
	if strings.TrimSpace(text) == "" {
		return "（" + name + "）"
	}
	return strings.TrimSpace(text) + "（" + name + "）"
}

func (s *Server) sendWelcome(ctx context.Context, activity Activity) {
	text := "direct bridge は利用可能です。対象チャネルで `@direct bind <alias>` を送ると、この Teams チャネルを Direct アカウントの転送先に紐付けます。"
	if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, "", text); err != nil {
		s.logger.Printf("[teams] welcome message failed conversation=%s err=%v", activity.Conversation.ID, err)
	}
}

func (s *Server) handleCommand(ctx context.Context, activity Activity) bool {
	command := ParseCommand(activity)
	switch command {
	case "bind":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "使い方: `@direct bind <alias>`")
		return true
	case "unbind":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "使い方: `@direct unbind <alias>`")
		return true
	case "hi", "hello":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, rootOnlyHelpText())
		return true
	case "help":
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, rootOnlyHelpText())
		return true
	default:
		if IsNewThreadCommand(command) {
			s.resetThreadMapping(ctx, activity)
			return true
		}
		return false
	}
}

func (s *Server) resetThreadMapping(ctx context.Context, activity Activity) {
	conversationID, rootID := threadReference(activity)
	if rootID == "" {
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ new-thread must be sent as a reply in a bridged thread.")
		return
	}
	mapping, ok := s.store.GetByThread(conversationID, rootID)
	if !ok {
		s.logger.Printf("[teams] new-thread ignored for unmapped thread conversation=%s root=%s", conversationID, rootID)
		_, _ = s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ this thread is not mapped to a Direct room.")
		return
	}
	if err := s.store.Forget(mapping.AccountID, mapping.TalkID); err != nil {
		s.logger.Printf("[teams] new-thread failed account=%s talk=%s err=%v", mapping.AccountID, mapping.TalkID, err)
		return
	}
	text := "✅ reset thread mapping. The next Direct message for this room will start a new Teams thread."
	if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, text); err != nil {
		s.logger.Printf("[teams] new-thread response failed account=%s talk=%s err=%v", mapping.AccountID, mapping.TalkID, err)
	}
}

func (s *Server) unbindChannel(ctx context.Context, alias string, activity Activity) {
	if activity.ReplyToID != "" {
		if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ unbind must be sent as a new channel message, not as a thread reply."); err != nil {
			s.logger.Printf("[teams] unbind placement response failed alias=%s err=%v", alias, err)
		}
		return
	}
	if !s.hasChannel(alias) {
		if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ unknown channel alias: "+alias); err != nil {
			s.logger.Printf("[teams] unknown unbind alias response failed alias=%s err=%v", alias, err)
		}
		return
	}
	binding, ok := s.store.GetChannelBinding(alias)
	conversationID := channelConversationID(activity)
	if !ok || binding.ConversationID != conversationID {
		if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ channel alias is not bound here: "+alias); err != nil {
			s.logger.Printf("[teams] unbind not-bound response failed alias=%s err=%v", alias, err)
		}
		return
	}
	removed, err := s.store.UnbindChannel(alias, conversationID)
	if err != nil {
		s.logger.Printf("[teams] unbind failed alias=%s err=%v", alias, err)
		return
	}
	text := "✅ unbound channel alias: " + alias + "; removed " + strconv.Itoa(removed) + " thread mappings"
	if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, text); err != nil {
		s.logger.Printf("[teams] unbind response failed alias=%s err=%v", alias, err)
	}
}

func rootOnlyHelpText() string {
	return strings.Join([]string{
		"Direct と Teams をつなぐブリッジです。",
		"",
		"チャネルで使うコマンド:",
		"- `@direct bind <alias>`: この Teams チャネルを設定済みの Direct 転送先に紐付けます。",
		"- `@direct unbind <alias>`: この Teams チャネルの紐付けを解除します。",
		"",
		"Direct から作られた thread の返信で使うコマンド:",
		"- `@direct reply <本文>`: Direct の同じトークルームへ返信します。",
		"- `@direct new-thread`: この Direct トークルームとの紐付けをリセットし、次の Direct メッセージから新しい Teams thread を作ります。",
		"",
		"使い方をもう一度見るには `@direct help` を送ってください。",
	}, "\n")
}

func threadReplyUsageText() string {
	return "Direct に返信するには `@direct reply <本文>` を使ってください。"
}

func (s *Server) bindChannel(ctx context.Context, alias string, activity Activity) {
	if activity.ReplyToID != "" {
		if _, err := s.client.SendText(ctx, activity.ServiceURL, activity.Conversation.ID, activity.ID, "❌ bind must be sent as a new channel message, not as a thread reply."); err != nil {
			s.logger.Printf("[teams] bind placement response failed alias=%s err=%v", alias, err)
		}
		return
	}
	if !s.hasChannel(alias) {
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
	_, rootID, ok := strings.Cut(activity.Conversation.ID, ";messageid=")
	if ok {
		return conversationID, rootID
	}
	if activity.ReplyToID != "" {
		return conversationID, activity.ReplyToID
	}
	return conversationID, ""
}

func (s *Server) attachmentsFromActivity(ctx context.Context, activity Activity) []model.Attachment {
	out := make([]model.Attachment, 0, len(activity.Attachments))
	for _, att := range activity.Attachments {
		item := model.Attachment{Name: att.Name, ContentType: att.ContentType, URL: firstNonEmpty(att.DownloadURL(), att.ContentURL)}
		if data, contentType, err := s.client.DownloadAttachment(ctx, att, s.cfg.Attachments.MaxBytes); err == nil {
			item.Data = data
			item.ContentType = firstNonEmpty(contentType, item.ContentType)
		} else {
			s.logger.Printf("[teams] attachment download failed content_type=%s name=%s url=%s err=%v", att.ContentType, att.Name, item.URL, err)
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
