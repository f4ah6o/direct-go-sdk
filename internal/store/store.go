package store

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
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

type TeamsMessageRef struct {
	AccountID       string    `json:"account_id"`
	TalkID          string    `json:"talk_id"`
	DirectMessageID string    `json:"direct_message_id"`
	ServiceURL      string    `json:"service_url"`
	ConversationID  string    `json:"conversation_id"`
	ActivityID      string    `json:"activity_id"`
	DirectSenderID  string    `json:"direct_sender_id,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	SentReactedAt   time.Time `json:"sent_reacted_at,omitempty"`
	ReactedAt       time.Time `json:"reacted_at,omitempty"`
}

type CodexThreadMapping struct {
	QuestionConversationID string    `json:"question_conversation_id"`
	QuestionRootID         string    `json:"question_root_id"`
	QuestionServiceURL     string    `json:"question_service_url"`
	AnswerConversationID   string    `json:"answer_conversation_id,omitempty"`
	AnswerRootID           string    `json:"answer_root_id,omitempty"`
	AnswerServiceURL       string    `json:"answer_service_url,omitempty"`
	CodexThreadID          string    `json:"codex_thread_id"`
	Status                 string    `json:"status"`
	CreatedAt              time.Time `json:"created_at"`
	UpdatedAt              time.Time `json:"updated_at"`
}

type State struct {
	TalkThreads           map[string]ThreadMapping       `json:"talk_threads"`
	TeamsThreadIndex      map[string]string              `json:"teams_thread_index"`
	TeamsChannelBindings  map[string]TeamsChannelBinding `json:"teams_channel_bindings"`
	SentTeamsMessages     map[string]time.Time           `json:"sent_teams_messages"`
	SentDirectMessages    map[string]time.Time           `json:"sent_direct_messages"`
	DirectToTeamsMessages map[string]TeamsMessageRef     `json:"direct_to_teams_messages"`
	Subscriptions         map[string]string              `json:"subscriptions"`
	DirectDevices         map[string]string              `json:"direct_devices"`
	CodexThreads          map[string]CodexThreadMapping  `json:"codex_threads"`
	CodexAnswerIndex      map[string]string              `json:"codex_answer_index"`
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
		TalkThreads:           map[string]ThreadMapping{},
		TeamsThreadIndex:      map[string]string{},
		TeamsChannelBindings:  map[string]TeamsChannelBinding{},
		SentTeamsMessages:     map[string]time.Time{},
		SentDirectMessages:    map[string]time.Time{},
		DirectToTeamsMessages: map[string]TeamsMessageRef{},
		Subscriptions:         map[string]string{},
		DirectDevices:         map[string]string{},
		CodexThreads:          map[string]CodexThreadMapping{},
		CodexAnswerIndex:      map[string]string{},
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
	if s.state.DirectToTeamsMessages == nil {
		s.state.DirectToTeamsMessages = map[string]TeamsMessageRef{}
	}
	if s.state.Subscriptions == nil {
		s.state.Subscriptions = map[string]string{}
	}
	if s.state.DirectDevices == nil {
		s.state.DirectDevices = map[string]string{}
	}
	if s.state.CodexThreads == nil {
		s.state.CodexThreads = map[string]CodexThreadMapping{}
	}
	if s.state.CodexAnswerIndex == nil {
		s.state.CodexAnswerIndex = map[string]string{}
	}
}

func talkKey(accountID, talkID string) string {
	return accountID + ":" + talkID
}

func threadKey(conversationID, rootID string) string {
	return conversationID + ":" + rootID
}

func codexQuestionKey(conversationID, rootID string) string {
	return threadKey(conversationID, rootID)
}

func codexAnswerKey(conversationID, rootID string) string {
	return threadKey(conversationID, rootID)
}

func directMessageKey(accountID, messageID string) string {
	return accountID + ":" + messageID
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
	key := talkKey(m.AccountID, m.TalkID)
	if existing, ok := s.state.TalkThreads[key]; ok {
		delete(s.state.TeamsThreadIndex, threadKey(existing.ConversationID, existing.RootID))
		if m.CreatedAt.IsZero() {
			m.CreatedAt = existing.CreatedAt
		}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	m.UpdatedAt = now
	s.state.TalkThreads[key] = m
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

func (s *Store) UnbindChannel(alias, conversationID string) (int, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.state.TeamsChannelBindings, alias)
	removed := 0
	for key, mapping := range s.state.TalkThreads {
		if mapping.ChannelAlias != alias || mapping.ConversationID != conversationID {
			continue
		}
		delete(s.state.TeamsThreadIndex, threadKey(mapping.ConversationID, mapping.RootID))
		delete(s.state.TalkThreads, key)
		removed++
	}
	return removed, s.saveLocked()
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

func (s *Store) GetCodexByQuestion(conversationID, rootID string) (CodexThreadMapping, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m, ok := s.state.CodexThreads[codexQuestionKey(conversationID, rootID)]
	return m, ok
}

func (s *Store) GetCodexByAnswer(conversationID, rootID string) (CodexThreadMapping, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	qKey, ok := s.state.CodexAnswerIndex[codexAnswerKey(conversationID, rootID)]
	if !ok {
		return CodexThreadMapping{}, false
	}
	m, ok := s.state.CodexThreads[qKey]
	return m, ok
}

func (s *Store) PutCodexMapping(m CodexThreadMapping) error {
	if m.QuestionConversationID == "" || m.QuestionRootID == "" || m.CodexThreadID == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	key := codexQuestionKey(m.QuestionConversationID, m.QuestionRootID)
	if existing, ok := s.state.CodexThreads[key]; ok {
		if m.CreatedAt.IsZero() {
			m.CreatedAt = existing.CreatedAt
		}
		if existing.AnswerConversationID != "" && existing.AnswerRootID != "" {
			delete(s.state.CodexAnswerIndex, codexAnswerKey(existing.AnswerConversationID, existing.AnswerRootID))
		}
	}
	if m.CreatedAt.IsZero() {
		m.CreatedAt = now
	}
	m.UpdatedAt = now
	s.state.CodexThreads[key] = m
	if m.AnswerConversationID != "" && m.AnswerRootID != "" {
		s.state.CodexAnswerIndex[codexAnswerKey(m.AnswerConversationID, m.AnswerRootID)] = key
	}
	return s.saveLocked()
}

func (s *Store) DirectDevice(accountID string) (string, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	deviceID, ok := s.state.DirectDevices[accountID]
	return deviceID, ok
}

func (s *Store) EnsureDirectDevice(accountID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if deviceID := s.state.DirectDevices[accountID]; deviceID != "" {
		return deviceID, nil
	}
	deviceID, err := newDeviceID()
	if err != nil {
		return "", err
	}
	s.state.DirectDevices[accountID] = deviceID
	return deviceID, s.saveLocked()
}

func (s *Store) ResetDirectDevice(accountID string) (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	deviceID, err := newDeviceID()
	if err != nil {
		return "", err
	}
	s.state.DirectDevices[accountID] = deviceID
	return deviceID, s.saveLocked()
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

func (s *Store) PutTeamsMessageRef(ref TeamsMessageRef) error {
	if ref.AccountID == "" || ref.DirectMessageID == "" || ref.ActivityID == "" || ref.ConversationID == "" || ref.ServiceURL == "" {
		return nil
	}
	s.mu.Lock()
	defer s.mu.Unlock()
	now := time.Now().UTC()
	key := directMessageKey(ref.AccountID, ref.DirectMessageID)
	if existing, ok := s.state.DirectToTeamsMessages[key]; ok {
		if ref.CreatedAt.IsZero() {
			ref.CreatedAt = existing.CreatedAt
		}
		if ref.ReactedAt.IsZero() {
			ref.ReactedAt = existing.ReactedAt
		}
		if ref.SentReactedAt.IsZero() {
			ref.SentReactedAt = existing.SentReactedAt
		}
	}
	if ref.CreatedAt.IsZero() {
		ref.CreatedAt = now
	}
	ref.UpdatedAt = now
	s.state.DirectToTeamsMessages[key] = ref
	return s.saveLocked()
}

func (s *Store) GetTeamsMessageRef(accountID, directMessageID string) (TeamsMessageRef, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	ref, ok := s.state.DirectToTeamsMessages[directMessageKey(accountID, directMessageID)]
	return ref, ok
}

func (s *Store) MarkTeamsSentReaction(accountID, directMessageID string) (TeamsMessageRef, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := directMessageKey(accountID, directMessageID)
	ref, ok := s.state.DirectToTeamsMessages[key]
	if !ok {
		return TeamsMessageRef{}, false, nil
	}
	if !ref.SentReactedAt.IsZero() {
		return ref, false, nil
	}
	ref.SentReactedAt = time.Now().UTC()
	ref.UpdatedAt = ref.SentReactedAt
	s.state.DirectToTeamsMessages[key] = ref
	return ref, true, s.saveLocked()
}

func (s *Store) MarkTeamsReadReaction(accountID, directMessageID string) (TeamsMessageRef, bool, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	key := directMessageKey(accountID, directMessageID)
	ref, ok := s.state.DirectToTeamsMessages[key]
	if !ok {
		return TeamsMessageRef{}, false, nil
	}
	if !ref.ReactedAt.IsZero() {
		return ref, false, nil
	}
	ref.ReactedAt = time.Now().UTC()
	ref.UpdatedAt = ref.ReactedAt
	s.state.DirectToTeamsMessages[key] = ref
	return ref, true, s.saveLocked()
}

func newDeviceID() (string, error) {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uint32(b[0])<<24|uint32(b[1])<<16|uint32(b[2])<<8|uint32(b[3]),
		uint16(b[4])<<8|uint16(b[5]),
		uint16(b[6])<<8|uint16(b[7]),
		uint16(b[8])<<8|uint16(b[9]),
		uint64(b[10])<<40|uint64(b[11])<<32|uint64(b[12])<<24|uint64(b[13])<<16|uint64(b[14])<<8|uint64(b[15]),
	), nil
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
