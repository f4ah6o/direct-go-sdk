package codexbridge

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/store"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/teams"
)

func TestQuestionImmediateAnswerPostsBackToQuestionThread(t *testing.T) {
	connector, posted := fakeConnector(t)
	codex := &fakeCodex{answers: []string{"回答です"}}
	st, err := store.Open(t.TempDir() + "/state.json")
	if err != nil {
		t.Fatal(err)
	}
	s := NewService(config.CodexConfig{}, st, teamsClient(connector.URL), codex, discardLogger())

	s.handleQuestion(t.Context(), teams.CodexActivity{
		ServiceURL:     connector.URL,
		ConversationID: "question-conversation",
		RootID:         "question-root",
		ActivityID:     "activity-id",
		Text:           "質問です",
		FromName:       "User",
	})

	if got := strings.Join(posted.texts(), "\n"); !strings.Contains(got, "回答です") {
		t.Fatalf("answer was not posted: %q", got)
	}
	m, ok := st.GetCodexByQuestion("question-conversation", "question-root")
	if !ok || m.CodexThreadID != "thread-1" || m.Status != "answered" {
		t.Fatalf("unexpected mapping: %+v ok=%v", m, ok)
	}
}

func TestEscalationCreatesAnswerThreadAndHumanAnswerPostsFinal(t *testing.T) {
	connector, posted := fakeConnector(t)
	codex := &fakeCodex{answers: []string{"ESCALATE: 担当者確認が必要です", "最終回答です"}}
	st, err := store.Open(t.TempDir() + "/state.json")
	if err != nil {
		t.Fatal(err)
	}
	if err := st.PutChannelBinding(store.TeamsChannelBinding{
		Alias:          "codex-answer",
		ConversationID: "answer-conversation",
		ServiceURL:     connector.URL,
		ChannelID:      "answer-channel",
		BotID:          "bot-id",
	}); err != nil {
		t.Fatal(err)
	}
	s := NewService(config.CodexConfig{AnswerAlias: "codex-answer"}, st, teamsClient(connector.URL), codex, discardLogger())

	s.handleQuestion(t.Context(), teams.CodexActivity{
		ServiceURL:     connector.URL,
		ConversationID: "question-conversation",
		RootID:         "question-root",
		ActivityID:     "activity-id",
		Text:           "質問です",
		FromName:       "User",
	})
	m, ok := st.GetCodexByQuestion("question-conversation", "question-root")
	if !ok || m.AnswerRootID != "created-root" || m.Status != "awaiting_human" {
		t.Fatalf("unexpected escalation mapping: %+v ok=%v", m, ok)
	}
	s.handleAnswer(t.Context(), teams.CodexActivity{
		ServiceURL:     connector.URL,
		ConversationID: "answer-conversation",
		RootID:         "created-root",
		ActivityID:     "answer-activity",
		Text:           "人間の回答です",
		FromName:       "Responder",
	})

	if got := strings.Join(posted.texts(), "\n"); !strings.Contains(got, "最終回答です") {
		t.Fatalf("final answer was not posted: %q", got)
	}
}

type fakeCodex struct {
	mu      sync.Mutex
	answers []string
}

func (f *fakeCodex) StartThread(context.Context) (string, error) {
	return "thread-1", nil
}

func (f *fakeCodex) Turn(_ context.Context, _ string, _ string) (string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	answer := f.answers[0]
	f.answers = f.answers[1:]
	return answer, nil
}

type postedActivities struct {
	mu    sync.Mutex
	items []teams.Activity
}

func (p *postedActivities) texts() []string {
	p.mu.Lock()
	defer p.mu.Unlock()
	out := make([]string, 0, len(p.items))
	for _, item := range p.items {
		out = append(out, item.Text)
	}
	return out
}

func fakeConnector(t *testing.T) (*httptest.Server, *postedActivities) {
	t.Helper()
	posted := &postedActivities{}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.URL.Path == "/token":
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"access_token":"token","expires_in":3600}`))
		case r.URL.Path == "/v3/conversations":
			var activity struct {
				Activity teams.Activity `json:"activity"`
			}
			if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
				t.Fatalf("decode create conversation: %v", err)
			}
			posted.mu.Lock()
			posted.items = append(posted.items, activity.Activity)
			posted.mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"activityId":"created-root"}`))
		case strings.Contains(r.URL.Path, "/activities"):
			var activity teams.Activity
			if err := json.NewDecoder(r.Body).Decode(&activity); err != nil {
				t.Fatalf("decode activity: %v", err)
			}
			posted.mu.Lock()
			posted.items = append(posted.items, activity)
			posted.mu.Unlock()
			w.Header().Set("Content-Type", "application/json")
			_, _ = w.Write([]byte(`{"id":"reply-id"}`))
		default:
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
	}))
	t.Cleanup(server.Close)
	return server, posted
}

func teamsClient(baseURL string) *teams.Client {
	return teams.NewClient(config.BotConfig{AppID: "app-id", AppPassword: "secret", TokenURL: baseURL + "/token", ConnectorScope: "scope"})
}

func discardLogger() *log.Logger {
	return log.New(io.Discard, "", 0)
}
