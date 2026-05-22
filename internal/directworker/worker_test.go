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

	go runAccountWorker(ctx, runtimeAccount(), in, out, sent, discardLogger(), factory.New, noBackoff)

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

	go runAccountWorker(ctx, runtimeAccount(), in, out, sent, discardLogger(), factory.New, noBackoff)

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

	go runAccountWorker(ctx, runtimeAccount(), in, out, sent, discardLogger(), factory.New, noBackoff)

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

	go runAccountWorker(ctx, runtimeAccount(), in, out, sent, discardLogger(), factory.New, noBackoff)

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

func runtimeAccount() RuntimeAccount {
	return RuntimeAccount{
		Config: config.AccountConfig{ID: "account-a", TeamsChannel: "support"},
		Token:  "token-a",
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

func (c *fakeDirectClient) Call(string, []interface{}) (interface{}, error) {
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
