package slackcompat

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
)

type Server struct {
	clients map[string]DirectAPI
	mapper  Mapper
	teamID  string
	team    string
	botUser string
	token   string
	logger  *log.Logger
}

func NewServer(clients []DirectAPI, opts ...func(*Server)) *Server {
	clientMap := map[string]DirectAPI{}
	for _, client := range clients {
		clientMap[client.AccountID()] = client
	}
	s := &Server{
		clients: clientMap,
		mapper:  NewMapper(),
		teamID:  "Tdirect",
		team:    "Direct4B",
		logger:  log.New(log.Writer(), "", log.LstdFlags),
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

func WithTeam(teamID, teamName string) func(*Server) {
	return func(s *Server) {
		if strings.TrimSpace(teamID) != "" {
			s.teamID = teamID
		}
		if strings.TrimSpace(teamName) != "" {
			s.team = teamName
		}
	}
}

func WithBotUserID(botUserID string) func(*Server) {
	return func(s *Server) {
		s.botUser = strings.TrimSpace(botUserID)
	}
}

func WithBearerToken(token string) func(*Server) {
	return func(s *Server) {
		s.token = strings.TrimSpace(token)
	}
}

func WithLogger(logger *log.Logger) func(*Server) {
	return func(s *Server) {
		if logger != nil {
			s.logger = logger
		}
	}
}

func (s *Server) HTTPHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/auth.test", s.handleAuthTest)
	mux.HandleFunc("/api/chat.postMessage", s.handlePostMessage)
	mux.HandleFunc("/api/conversations.list", s.handleConversationsList)
	mux.HandleFunc("/api/conversations.history", s.handleConversationsHistory)
	mux.HandleFunc("/api/users.list", s.handleUsersList)
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})
	return s.authMiddleware(mux)
}

func (s *Server) authMiddleware(next http.Handler) http.Handler {
	if s.token == "" {
		return next
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/healthz" {
			next.ServeHTTP(w, r)
			return
		}
		auth := strings.TrimSpace(r.Header.Get("Authorization"))
		typ, token, ok := strings.Cut(auth, " ")
		if !ok || !strings.EqualFold(typ, "Bearer") || token != s.token {
			w.Header().Set("WWW-Authenticate", "Bearer")
			writeJSON(w, SlackResponse{OK: false, Error: "not_authed"})
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (s *Server) handleAuthTest(w http.ResponseWriter, r *http.Request) {
	if !allowMethod(w, r, http.MethodGet, http.MethodPost) {
		return
	}
	client, err := s.defaultClient()
	if err != nil {
		writeJSON(w, AuthTestResponse{OK: false, Error: err.Error()})
		return
	}
	me, err := client.GetMe(r.Context())
	if err != nil {
		writeJSON(w, AuthTestResponse{OK: false, Error: "direct_error"})
		return
	}
	userID := s.botUser
	if userID == "" {
		userID = s.mapper.UserID(fmt.Sprint(me.ID))
	}
	writeJSON(w, AuthTestResponse{
		OK:     true,
		URL:    "https://direct4b.com/",
		Team:   s.team,
		TeamID: s.teamID,
		User:   displayUserName(*me),
		UserID: userID,
		BotID:  s.mapper.BotID(client.AccountID()),
	})
}

func (s *Server) handlePostMessage(w http.ResponseWriter, r *http.Request) {
	if !allowMethod(w, r, http.MethodPost) {
		return
	}
	req, err := parseParams(r)
	if err != nil {
		writeJSON(w, SlackResponse{OK: false, Error: "invalid_request"})
		return
	}
	channel := strings.TrimSpace(req["channel"])
	text := req["text"]
	accountID, talkID, ok := s.mapper.TalkID(channel)
	if !ok {
		writeJSON(w, SlackResponse{OK: false, Error: "channel_not_found"})
		return
	}
	client, ok := s.clients[accountID]
	if !ok {
		writeJSON(w, SlackResponse{OK: false, Error: "channel_not_found"})
		return
	}
	messageID, err := client.SendText(r.Context(), talkID, text)
	if err != nil {
		s.logger.Printf("[slackcompat] chat.postMessage failed account=%s talk=%s err=%v", accountID, talkID, err)
		writeJSON(w, SlackResponse{OK: false, Error: "direct_error"})
		return
	}
	ts := s.mapper.Timestamp(messageID, time.Time{})
	writeJSON(w, PostMessageResponse{
		OK:      true,
		Channel: channel,
		TS:      ts,
		Message: SlackMessage{Type: "message", Text: text, TS: ts},
	})
}

func (s *Server) handleConversationsList(w http.ResponseWriter, r *http.Request) {
	if !allowMethod(w, r, http.MethodGet, http.MethodPost) {
		return
	}
	channels := []SlackChannel{}
	for _, client := range s.clients {
		talks, err := client.GetTalks(r.Context())
		if err != nil {
			s.logger.Printf("[slackcompat] conversations.list failed account=%s err=%v", client.AccountID(), err)
			writeJSON(w, SlackResponse{OK: false, Error: "direct_error"})
			return
		}
		for _, talk := range talks {
			channels = append(channels, s.slackChannel(client.AccountID(), talk))
		}
	}
	writeJSON(w, ConversationsListResponse{OK: true, Channels: channels})
}

func (s *Server) handleConversationsHistory(w http.ResponseWriter, r *http.Request) {
	if !allowMethod(w, r, http.MethodGet, http.MethodPost) {
		return
	}
	req, err := parseParams(r)
	if err != nil {
		writeJSON(w, SlackResponse{OK: false, Error: "invalid_request"})
		return
	}
	accountID, talkID, ok := s.mapper.TalkID(strings.TrimSpace(req["channel"]))
	if !ok {
		writeJSON(w, SlackResponse{OK: false, Error: "channel_not_found"})
		return
	}
	client, ok := s.clients[accountID]
	if !ok {
		writeJSON(w, SlackResponse{OK: false, Error: "channel_not_found"})
		return
	}
	domainID, err := s.domainIDForTalk(r.Context(), client, talkID)
	if err != nil {
		writeJSON(w, SlackResponse{OK: false, Error: "history_unavailable"})
		return
	}
	messages, err := client.GetMessages(r.Context(), domainID, talkID, &direct.GetMessagesOptions{Order: direct.MessageOrderDesc})
	if err != nil {
		s.logger.Printf("[slackcompat] conversations.history failed account=%s talk=%s err=%v", accountID, talkID, err)
		writeJSON(w, SlackResponse{OK: false, Error: "direct_error"})
		return
	}
	limit := parseLimit(req["limit"], len(messages))
	out := make([]SlackMessage, 0, min(limit, len(messages)))
	for i, msg := range messages {
		if i >= limit {
			break
		}
		out = append(out, s.slackMessage(accountID, msg))
	}
	writeJSON(w, ConversationsHistoryResponse{OK: true, Messages: out, HasMore: len(messages) > limit})
}

func (s *Server) handleUsersList(w http.ResponseWriter, r *http.Request) {
	if !allowMethod(w, r, http.MethodGet, http.MethodPost) {
		return
	}
	members := []SlackUser{}
	for _, client := range s.clients {
		users, err := client.GetUsers(r.Context())
		if err != nil {
			s.logger.Printf("[slackcompat] users.list failed account=%s err=%v", client.AccountID(), err)
			writeJSON(w, SlackResponse{OK: false, Error: "direct_error"})
			return
		}
		for _, user := range users {
			members = append(members, s.slackUser(user))
		}
	}
	writeJSON(w, UsersListResponse{OK: true, Members: members})
}

func (s *Server) slackChannel(accountID string, talk direct.Talk) SlackChannel {
	id := s.mapper.ChannelID(accountID, fmt.Sprint(talk.ID))
	name := talk.Name
	if strings.TrimSpace(name) == "" {
		name = "direct-" + sanitizeName(fmt.Sprint(talk.ID))
	}
	isIM := talk.Type == int(direct.RoomTypePair)
	return SlackChannel{
		ID:         id,
		Name:       sanitizeName(name),
		IsChannel:  !isIM,
		IsGroup:    !isIM,
		IsIM:       isIM,
		IsPrivate:  true,
		IsArchived: false,
		NumMembers: len(talk.UserIDs),
	}
}

func (s *Server) slackUser(user direct.UserInfo) SlackUser {
	id := s.mapper.UserID(fmt.Sprint(user.ID))
	name := user.Name
	if strings.TrimSpace(name) == "" {
		name = user.DisplayName
	}
	if strings.TrimSpace(name) == "" {
		name = fmt.Sprint(user.ID)
	}
	return SlackUser{
		ID:      id,
		Name:    sanitizeName(name),
		Deleted: false,
		IsBot:   false,
		Profile: SlackUserProfile{
			RealName: displayUserName(user),
			Email:    user.Email,
			Image48:  user.IconURL,
		},
	}
}

func (s *Server) slackMessage(accountID string, msg direct.ReceivedMessage) SlackMessage {
	created := msg.Timestamp
	if created.IsZero() && msg.Created > 0 {
		created = time.Unix(msg.Created, 0).UTC()
	}
	return SlackMessage{
		Type: "message",
		User: s.mapper.UserID(msg.UserID),
		Text: msg.Text,
		TS:   s.mapper.Timestamp(msg.ID, created),
	}
}

func (s *Server) domainIDForTalk(ctx context.Context, client DirectAPI, talkID string) (interface{}, error) {
	talks, err := client.GetTalks(ctx)
	if err != nil {
		return nil, err
	}
	for _, talk := range talks {
		if fmt.Sprint(talk.ID) == talkID && talk.DomainID != nil {
			return talk.DomainID, nil
		}
	}
	return nil, errors.New("domain id not found")
}

func (s *Server) defaultClient() (DirectAPI, error) {
	for _, client := range s.clients {
		return client, nil
	}
	return nil, errors.New("not_authed")
}

func parseParams(r *http.Request) (map[string]string, error) {
	out := map[string]string{}
	ct := r.Header.Get("Content-Type")
	if strings.Contains(ct, "application/json") {
		var raw map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&raw); err != nil {
			return nil, err
		}
		for k, v := range raw {
			out[k] = fmt.Sprint(v)
		}
		return out, nil
	}
	if err := r.ParseForm(); err != nil {
		return nil, err
	}
	for k, v := range r.Form {
		if len(v) > 0 {
			out[k] = v[0]
		}
	}
	return out, nil
}

func writeJSON(w http.ResponseWriter, value interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(value)
}

func allowMethod(w http.ResponseWriter, r *http.Request, methods ...string) bool {
	for _, method := range methods {
		if r.Method == method {
			return true
		}
	}
	http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	return false
}

func displayUserName(user direct.UserInfo) string {
	if strings.TrimSpace(user.DisplayName) != "" {
		return user.DisplayName
	}
	if strings.TrimSpace(user.Name) != "" {
		return user.Name
	}
	return fmt.Sprint(user.ID)
}

func sanitizeName(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	value = strings.ReplaceAll(value, " ", "-")
	var b strings.Builder
	for _, r := range value {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '-' || r == '_' {
			b.WriteRune(r)
		}
	}
	out := strings.Trim(b.String(), "-_")
	if out == "" {
		return "direct"
	}
	return out
}

func parseLimit(raw string, fallback int) int {
	if fallback <= 0 {
		fallback = 100
	}
	if raw == "" {
		return fallback
	}
	n, err := strconv.Atoi(raw)
	if err != nil || n <= 0 {
		return fallback
	}
	if n > 1000 {
		return 1000
	}
	return n
}
