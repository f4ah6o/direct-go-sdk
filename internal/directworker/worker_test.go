package directworker

import (
	"context"
	"errors"
	"io"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	direct "github.com/f4ah6o/direct-go-sdk/direct-go"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/model"
)

func TestAccountWorkerRecreatesClientOnDirectError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	factory := &fakeDirectClientFactory{}
	in := make(chan model.DirectOutbound)
	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)

	worker := testAccountWorker(runtimeAccount(), in)
	go runAccountWorker(ctx, worker, out, nil, sent, discardLogger(), factory.New, noBackoff, time.Now)

	first := factory.waitForClient(t, 1)
	first.emit(direct.EventError, map[string]string{"error": "read failed"})
	second := factory.waitForClient(t, 2)

	if !first.isClosed() {
		t.Fatalf("first client was not closed after direct error")
	}
	if second == first {
		t.Fatalf("expected a new client after direct error")
	}

	in <- model.DirectOutbound{AccountID: "account-a", TalkID: "talk-a", Text: "hello"}
	got := waitForSent(t, sent)
	if got.Err != nil {
		t.Fatalf("send after reconnect failed: %v", got.Err)
	}
	if got.MessageID != "msg-2" {
		t.Fatalf("message id = %q, want msg-2", got.MessageID)
	}
}

func TestAccountWorkerRecreatesClientOnDone(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	factory := &fakeDirectClientFactory{}
	in := make(chan model.DirectOutbound)
	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)

	worker := testAccountWorker(runtimeAccount(), in)
	go runAccountWorker(ctx, worker, out, nil, sent, discardLogger(), factory.New, noBackoff, time.Now)

	first := factory.waitForClient(t, 1)
	if err := first.Close(); err != nil {
		t.Fatalf("close first client: %v", err)
	}
	second := factory.waitForClient(t, 2)

	if second == first {
		t.Fatalf("expected a new client after Done")
	}

	in <- model.DirectOutbound{AccountID: "account-a", TalkID: "talk-a", Text: "after done"}
	got := waitForSent(t, sent)
	if got.Err != nil {
		t.Fatalf("send after Done reconnect failed: %v", got.Err)
	}
	if got.MessageID != "msg-2" {
		t.Fatalf("message id = %q, want msg-2", got.MessageID)
	}
}

func TestAccountWorkerRetriesOutboundAfterRecoverableSendError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	factory := &fakeDirectClientFactory{
		configure: func(index int, c *fakeDirectClient) {
			if index == 1 {
				c.sendErr = errors.New("websocket: close 1006 abnormal closure")
				return
			}
			c.messageID = "msg-retried"
		},
	}
	in := make(chan model.DirectOutbound)
	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)

	worker := testAccountWorker(runtimeAccount(), in)
	go runAccountWorker(ctx, worker, out, nil, sent, discardLogger(), factory.New, noBackoff, time.Now)

	first := factory.waitForClient(t, 1)
	in <- model.DirectOutbound{AccountID: "account-a", TalkID: "talk-a", Text: "retry me"}

	second := factory.waitForClient(t, 2)
	got := waitForSent(t, sent)
	if got.Err != nil {
		t.Fatalf("send after recoverable error failed: %v", got.Err)
	}
	if got.MessageID != "msg-retried" {
		t.Fatalf("message id = %q, want msg-retried", got.MessageID)
	}
	if !first.isClosed() {
		t.Fatalf("first client was not closed after recoverable send error")
	}
	if first.sendCount() != 1 || second.sendCount() != 1 {
		t.Fatalf("send counts = first:%d second:%d, want 1 and 1", first.sendCount(), second.sendCount())
	}
}

func TestAccountWorkerDoesNotRetryNonRecoverableSendError(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	factory := &fakeDirectClientFactory{
		configure: func(index int, c *fakeDirectClient) {
			c.sendErr = errors.New("RPC error: permission denied")
		},
	}
	in := make(chan model.DirectOutbound)
	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)

	worker := testAccountWorker(runtimeAccount(), in)
	go runAccountWorker(ctx, worker, out, nil, sent, discardLogger(), factory.New, noBackoff, time.Now)

	first := factory.waitForClient(t, 1)
	in <- model.DirectOutbound{AccountID: "account-a", TalkID: "talk-a", Text: "fail"}
	got := waitForSent(t, sent)
	if got.Err == nil {
		t.Fatalf("expected non-recoverable error")
	}
	if first.sendCount() != 1 {
		t.Fatalf("send count = %d, want 1", first.sendCount())
	}
	if factory.clientCount() != 1 {
		t.Fatalf("client count = %d, want 1", factory.clientCount())
	}
}

func TestAccountWorkerEmitsReadReceipts(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	factory := &fakeDirectClientFactory{}
	in := make(chan model.DirectOutbound)
	out := make(chan model.DirectMessage)
	readOut := make(chan model.DirectReadReceipt, 1)
	sent := make(chan model.DirectSent, 1)

	worker := testAccountWorker(runtimeAccount(), in)
	go runAccountWorker(ctx, worker, out, readOut, sent, discardLogger(), factory.New, noBackoff, time.Now)

	client := factory.waitForClient(t, 1)
	client.emit(direct.EventNotifyUpdateReadStatuses, map[string]interface{}{
		"talk_id":       "talk-a",
		"message_ids":   []interface{}{"msg-a"},
		"read_user_ids": []interface{}{"user-a"},
	})

	select {
	case got := <-readOut:
		if got.AccountID != "account-a" || got.TalkID != "talk-a" || got.MessageIDs[0] != "msg-a" || got.ReadUserIDs[0] != "user-a" {
			t.Fatalf("unexpected read receipt: %+v", got)
		}
	case <-time.After(time.Second):
		t.Fatalf("timed out waiting for read receipt")
	}
}

func TestManagerWatchdogRestartsWorkerThatNeverBecomesReady(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)
	factory := &fakeDirectClientFactory{}
	manager := NewManager(out, nil, sent, discardLogger())
	manager.clientFactory = factory.New
	manager.sleepBackoff = noBackoff

	manager.Apply(ctx, []RuntimeAccount{runtimeAccount()})
	first := factory.waitForClient(t, 1)

	manager.mu.Lock()
	status := manager.statuses["account-a"]
	status.StartedAt = time.Now().Add(-time.Minute)
	status.UnreadySince = status.StartedAt
	manager.statuses["account-a"] = status
	manager.mu.Unlock()

	manager.restartStaleWorkers(time.Second)
	second := factory.waitForClient(t, 2)

	if !first.isClosed() {
		t.Fatalf("first client was not closed after watchdog restart")
	}
	if second == first {
		t.Fatalf("expected a new client after watchdog restart")
	}

	statuses := manager.Statuses()
	if len(statuses) != 1 {
		t.Fatalf("statuses len = %d, want 1", len(statuses))
	}
	if statuses[0].Restarts == 0 {
		t.Fatalf("expected restart count to be incremented")
	}
}

func TestManagerWatchdogDoesNotRestartInvalidTokenWorker(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)
	factory := &fakeDirectClientFactory{}
	manager := NewManager(out, nil, sent, discardLogger())
	manager.clientFactory = factory.New
	manager.sleepBackoff = noBackoff

	manager.Apply(ctx, []RuntimeAccount{runtimeAccount()})
	_ = factory.waitForClient(t, 1)

	manager.mu.Lock()
	status := manager.statuses["account-a"]
	status.StartedAt = time.Now().Add(-time.Minute)
	status.UnreadySince = status.StartedAt
	status.AuthInvalid = true
	manager.statuses["account-a"] = status
	manager.mu.Unlock()

	manager.restartStaleWorkers(time.Second)

	if factory.clientCount() != 1 {
		t.Fatalf("client count = %d, want 1", factory.clientCount())
	}
}

func TestManagerIgnoresStatusUpdatesFromOldWorkerGeneration(t *testing.T) {
	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)
	manager := NewManager(out, nil, sent, discardLogger())

	manager.mu.Lock()
	manager.workers["account-a"] = &accountWorker{generation: 2}
	manager.statuses["account-a"] = AccountStatus{AccountID: "account-a", Generation: 2, Running: true, Ready: true}
	manager.mu.Unlock()

	manager.updateStatus("account-a", 1, func(status *AccountStatus) {
		status.Running = false
		status.Ready = false
		status.LastError = "old worker exited"
	})

	statuses := manager.Statuses()
	if len(statuses) != 1 {
		t.Fatalf("statuses len = %d, want 1", len(statuses))
	}
	if !statuses[0].Running || !statuses[0].Ready || statuses[0].LastError != "" {
		t.Fatalf("old generation changed status: %+v", statuses[0])
	}
}

func TestManagerDoesNotResurrectStatusAfterWorkerRemoval(t *testing.T) {
	out := make(chan model.DirectMessage)
	sent := make(chan model.DirectSent, 1)
	manager := NewManager(out, nil, sent, discardLogger())

	manager.updateStatus("account-a", 1, func(status *AccountStatus) {
		status.Running = false
		status.LastError = "removed worker exited"
	})

	if statuses := manager.Statuses(); len(statuses) != 0 {
		t.Fatalf("statuses len = %d, want 0: %+v", len(statuses), statuses)
	}
}

func runtimeAccount() RuntimeAccount {
	return RuntimeAccount{
		Config: config.AccountConfig{ID: "account-a", TeamsChannel: "support"},
		Token:  "token-a",
	}
}

func testAccountWorker(account RuntimeAccount, in chan model.DirectOutbound) *accountWorker {
	statuses := map[string]AccountStatus{}
	var mu sync.Mutex
	return &accountWorker{
		account:    account,
		generation: 1,
		queue:      in,
		restart:    make(chan struct{}, 1),
		status: func(update func(*AccountStatus)) {
			mu.Lock()
			defer mu.Unlock()
			status := statuses[account.Config.ID]
			if status.AccountID == "" {
				status.AccountID = account.Config.ID
			}
			update(&status)
			statuses[account.Config.ID] = status
		},
	}
}

func discardLogger() *log.Logger {
	return log.New(io.Discard, "", 0)
}

func noBackoff(context.Context, *time.Duration) {}

func waitForSent(t *testing.T, sent <-chan model.DirectSent) model.DirectSent {
	t.Helper()
	select {
	case got := <-sent:
		return got
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for direct sent result")
	}
	return model.DirectSent{}
}

type fakeDirectClientFactory struct {
	mu        sync.Mutex
	clients   []*fakeDirectClient
	configure func(int, *fakeDirectClient)
}

func (f *fakeDirectClientFactory) New(direct.Options) directClient {
	f.mu.Lock()
	defer f.mu.Unlock()
	client := newFakeDirectClient(len(f.clients) + 1)
	if f.configure != nil {
		f.configure(len(f.clients)+1, client)
	}
	f.clients = append(f.clients, client)
	return client
}

func (f *fakeDirectClientFactory) waitForClient(t *testing.T, n int) *fakeDirectClient {
	t.Helper()
	deadline := time.After(2 * time.Second)
	ticker := time.NewTicker(5 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case <-deadline:
			t.Fatalf("timed out waiting for client %d; got %d", n, f.clientCount())
		case <-ticker.C:
			f.mu.Lock()
			if len(f.clients) >= n {
				client := f.clients[n-1]
				f.mu.Unlock()
				return client
			}
			f.mu.Unlock()
		}
	}
}

func (f *fakeDirectClientFactory) clientCount() int {
	f.mu.Lock()
	defer f.mu.Unlock()
	return len(f.clients)
}

type fakeDirectClient struct {
	mu        sync.Mutex
	done      chan struct{}
	handlers  map[string][]direct.EventHandler
	closed    bool
	sendErr   error
	messageID string
	sends     int
}

func newFakeDirectClient(index int) *fakeDirectClient {
	return &fakeDirectClient{
		done:      make(chan struct{}),
		handlers:  map[string][]direct.EventHandler{},
		messageID: "msg-" + strconv.Itoa(index),
	}
}

func (c *fakeDirectClient) Connect() error {
	return nil
}

func (c *fakeDirectClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if !c.closed {
		c.closed = true
		close(c.done)
	}
	return nil
}

func (c *fakeDirectClient) On(event string, handler direct.EventHandler) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.handlers[event] = append(c.handlers[event], handler)
}

func (c *fakeDirectClient) OnMessage(func(direct.ReceivedMessage)) {}

func (c *fakeDirectClient) CreateTextMessageWithContext(context.Context, string, string) (string, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.sends++
	if c.sendErr != nil {
		return "", c.sendErr
	}
	return c.messageID, nil
}

func (c *fakeDirectClient) CreateUploadAuth(context.Context, string, string, int64, string) (*direct.UploadAuth, error) {
	return nil, errors.New("not implemented")
}

func (c *fakeDirectClient) Call(method string, _ []interface{}) (interface{}, error) {
	if method == direct.MethodGetMe {
		return map[string]interface{}{"id": "direct-self"}, nil
	}
	return nil, errors.New("not implemented")
}

func (c *fakeDirectClient) Done() <-chan struct{} {
	return c.done
}

func (c *fakeDirectClient) emit(event string, data interface{}) {
	c.mu.Lock()
	handlers := append([]direct.EventHandler(nil), c.handlers[event]...)
	c.mu.Unlock()
	for _, handler := range handlers {
		handler(data)
	}
}

func (c *fakeDirectClient) isClosed() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.closed
}

func (c *fakeDirectClient) sendCount() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.sends
}
