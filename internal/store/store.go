package store

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type TalkKey struct {
	AccountID string
	TalkID    string
}

type ThreadMapping struct {
	AccountID      string    `json:"account_id"`
	TalkID         string    `json:"talk_id"`
	ChannelAlias   string    `json:"channel_alias"`
	ConversationID string    `json:"conversation_id"`
	ServiceURL     string    `json:"service_url"`
	RootID         string    `json:"root_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type TeamsChannelBinding struct {
	Alias          string    `json:"alias"`
	TeamID         string    `json:"team_id,omitempty"`
	ChannelID      string    `json:"channel_id,omitempty"`
	ConversationID string    `json:"conversation_id"`
	ServiceURL     string    `json:"service_url"`
	TenantID       string    `json:"tenant_id,omitempty"`
	BotID          string    `json:"bot_id,omitempty"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type State struct {
	TalkThreads          map[string]ThreadMapping       `json:"talk_threads"`
	TeamsThreadIndex     map[string]string              `json:"teams_thread_index"`
	TeamsChannelBindings map[string]TeamsChannelBinding `json:"teams_channel_bindings"`
	SentTeamsMessages    map[string]time.Time           `json:"sent_teams_messages"`
	SentDirectMessages   map[string]time.Time           `json:"sent_direct_messages"`
	Subscriptions        map[string]string              `json:"subscriptions"`
}

type Store struct {
	path  string
	mu    sync.Mutex
	state State
}

func Open(path string) (*Store, error) {
	s := &Store{path: path}
	s.state = newState()
	data, err := os.ReadFile(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return s, nil
		}
		return nil, err
	}
	if len(data) == 0 {
		return s, nil
	}
	if err := json.Unmarshal(data, &s.state); err != nil {
		return nil, err
	}
	s.ensure()
	return s, nil
}

func newState() State {
	return State{
		TalkThreads:          map[string]ThreadMapping{},
		TeamsThreadIndex:     map[string]string{},
		TeamsChannelBindings: map[string]TeamsChannelBinding{},
		SentTeamsMessages:    map[string]time.Time{},
		SentDirectMessages:   map[string]time.Time{},
		Subscriptions:        map[string]string{},
	}
}

func (s *Store) ensure() {
	if s.state.TalkThreads == nil {
		s.state.TalkThreads = map[string]ThreadMapping{}
	}
	if s.state.TeamsThreadIndex == nil {
		s.state.TeamsThreadIndex = map[string]string{}
	}
	if s.state.TeamsChannelBindings == nil {
		s.state.TeamsChannelBindings = map[string]TeamsChannelBinding{}
	}
	if s.state.SentTeamsMessages == nil {
		s.state.SentTeamsMessages = map[string]time.Time{}
	}
	if s.state.SentDirectMessages == nil {
		s.state.SentDirectMessages = map[string]time.Time{}
	}
	if s.state.Subscriptions == nil {
		s.state.Subscriptions = map[string]string{}
	}
}

func talkKey(accountID, talkID string) string {
	return accountID + ":" + talkID
}

func threadKey(conversationID, rootID string) string {
	return conversationID + ":" + rootID
}

func (s *Store) GetByTalk(accountID, talkID string) (ThreadMapping, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.state.TalkThreads[talkKey(accountID, talkID)]
	return m, ok
}

func (s *Store) GetByThread(conversationID, rootID string) (ThreadMapping, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key, ok := s.state.TeamsThreadIndex[threadKey(conversationID, rootID)]
	if !ok {
		return ThreadMapping{}, false
	}
	m, ok := s.state.TalkThreads[key]
	return m, ok
}

func (s *Store) PutMapping(m ThreadMapping) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	m.UpdatedAt = now
	s.state.TalkThreads[talkKey(m.AccountID, m.TalkID)] = m
	s.state.TeamsThreadIndex[threadKey(m.ConversationID, m.RootID)] = talkKey(m.AccountID, m.TalkID)
	return s.saveLocked()
}

func (s *Store) Forget(accountID, talkID string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := talkKey(accountID, talkID)
	if m, ok := s.state.TalkThreads[key]; ok {
		delete(s.state.TeamsThreadIndex, threadKey(m.ConversationID, m.RootID))
	}
	delete(s.state.TalkThreads, key)
	return s.saveLocked()
}

func (s *Store) PutChannelBinding(binding TeamsChannelBinding) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	if binding.CreatedAt.IsZero() {
		binding.CreatedAt = now
	}
	binding.UpdatedAt = now
	s.state.TeamsChannelBindings[binding.Alias] = binding
	return s.saveLocked()
}

func (s *Store) GetChannelBinding(alias string) (TeamsChannelBinding, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	b, ok := s.state.TeamsChannelBindings[alias]
	return b, ok
}

func (s *Store) ListChannelBindings() []TeamsChannelBinding {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]TeamsChannelBinding, 0, len(s.state.TeamsChannelBindings))
	for _, b := range s.state.TeamsChannelBindings {
		out = append(out, b)
	}
	return out
}

func (s *Store) ForgetChannelBinding(alias string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.state.TeamsChannelBindings, alias)
	return s.saveLocked()
}

func (s *Store) ListMappings() []ThreadMapping {
	s.mu.Lock()
	defer s.mu.Unlock()
	out := make([]ThreadMapping, 0, len(s.state.TalkThreads))
	for _, m := range s.state.TalkThreads {
		out = append(out, m)
	}
	return out
}

func (s *Store) MarkTeamsMessage(id string) error {
	if id == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state.SentTeamsMessages[id] = time.Now().UTC()
	return s.saveLocked()
}

func (s *Store) IsSentTeamsMessage(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.state.SentTeamsMessages[id]
	return ok
}

func (s *Store) MarkDirectMessage(id string) error {
	if id == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	s.state.SentDirectMessages[id] = time.Now().UTC()
	return s.saveLocked()
}

func (s *Store) IsSentDirectMessage(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.state.SentDirectMessages[id]
	return ok
}

func (s *Store) saveLocked() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0700); err != nil {
		return err
	}
	data, err := json.MarshalIndent(s.state, "", "  ")
	if err != nil {
		return err
	}
	tmp := s.path + ".tmp"
	if err := os.WriteFile(tmp, data, 0600); err != nil {
		return err
	}
	return os.Rename(tmp, s.path)
}
