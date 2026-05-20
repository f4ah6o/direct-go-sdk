package teams

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
)

type Server struct {
	cfg        *config.Config
	graph      *Client
	store      *store.Store
	out        chan<- model.DirectOutbound
	logger     *log.Logger
	httpServer *http.Server
}

type ChangeNotifications struct {
	Value []ChangeNotification `json:"value"`
}

type ChangeNotification struct {
	ClientState  string       `json:"clientState"`
	Resource     string       `json:"resource"`
	ResourceData ResourceData `json:"resourceData"`
}

type ResourceData struct {
	ID string `json:"id"`
}

func NewServer(cfg *config.Config, graph *Client, st *store.Store, out chan<- model.DirectOutbound, logger *log.Logger) *Server {
	s := &Server{cfg: cfg, graph: graph, store: st, out: out, logger: logger}
	mux := http.NewServeMux()
	mux.HandleFunc("/graph/notifications", s.handleNotifications)
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
		s.logger.Printf("[teams] notification server listening on %s", s.cfg.Server.ListenAddr)
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

func (s *Server) handleNotifications(w http.ResponseWriter, r *http.Request) {
	if token := r.URL.Query().Get("validationToken"); token != "" {
		w.Header().Set("Content-Type", "text/plain")
		_, _ = w.Write([]byte(token))
		return
	}
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var payload ChangeNotifications
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusAccepted)
	go s.processNotifications(context.Background(), payload)
}

func (s *Server) processNotifications(ctx context.Context, payload ChangeNotifications) {
	for _, n := range payload.Value {
		if n.ClientState != s.cfg.Graph.ClientState {
			s.logger.Printf("[teams] ignoring notification with invalid clientState")
			continue
		}
		teamID, channelID, rootID, replyID, ok := parseReplyResource(n.Resource)
		if !ok {
			s.logger.Printf("[teams] ignoring unsupported resource=%s", n.Resource)
			continue
		}
		mapping, ok := s.store.GetByThread(teamID, channelID, rootID)
		if !ok {
			s.logger.Printf("[teams] ignoring unmapped thread team=%s channel=%s root=%s", teamID, channelID, rootID)
			continue
		}
		ch, ok := s.cfg.TeamsChannels[channelNameFor(s.cfg, teamID, channelID)]
		if !ok {
			s.logger.Printf("[teams] no config for team=%s channel=%s", teamID, channelID)
			continue
		}
		msg, err := s.graph.GetReply(ctx, teamID, channelID, rootID, replyID)
		if err != nil {
			s.logger.Printf("[teams] get reply failed: %v", err)
			continue
		}
		if s.store.IsSentTeamsMessage(msg.ID) {
			continue
		}
		if !MentionsUser(msg, ch.MentionUserID) {
			continue
		}
		text := StripMentions(msg)
		attachments := AttachmentsFromMessage(ctx, s.graph, teamID, channelID, msg, s.cfg.Attachments.MaxBytes)
		out := model.DirectOutbound{
			AccountID:   mapping.AccountID,
			TalkID:      mapping.TalkID,
			Text:        text,
			Attachments: attachments,
		}
		select {
		case s.out <- out:
		case <-ctx.Done():
			return
		}
	}
}

func channelNameFor(cfg *config.Config, teamID, channelID string) string {
	for name, ch := range cfg.TeamsChannels {
		if ch.TeamID == teamID && ch.ChannelID == channelID {
			return name
		}
	}
	return ""
}

func parseReplyResource(resource string) (teamID, channelID, rootID, replyID string, ok bool) {
	parts := strings.Split(strings.Trim(resource, "/"), "/")
	for i := 0; i < len(parts)-1; i++ {
		switch parts[i] {
		case "teams":
			teamID = parts[i+1]
		case "channels":
			channelID = parts[i+1]
		case "messages":
			if rootID == "" {
				rootID = parts[i+1]
			}
		case "replies":
			replyID = parts[i+1]
		}
	}
	return teamID, channelID, rootID, replyID, teamID != "" && channelID != "" && rootID != "" && replyID != ""
}
