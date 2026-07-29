package codex

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os/exec"
	"sync"
	"sync/atomic"
	"time"

	"github.com/f4ah6o/direct-go-sdk/direct-teams-bridge/internal/config"
)

type Client struct {
	cfg    config.CodexConfig
	logger *log.Logger

	cmd    *exec.Cmd
	stdin  io.WriteCloser
	seq    atomic.Int64
	mu     sync.Mutex
	reqMu  sync.Mutex
	wait   map[int64]chan response
	events chan notification
}

type response struct {
	Result json.RawMessage
	Err    error
}

type notification struct {
	Method string
	Params json.RawMessage
}

type rpcMessage struct {
	JSONRPC string          `json:"jsonrpc,omitempty"`
	ID      json.RawMessage `json:"id,omitempty"`
	Method  string          `json:"method,omitempty"`
	Params  json.RawMessage `json:"params,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
	Error   *rpcError       `json:"error,omitempty"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient(cfg config.CodexConfig, logger *log.Logger) *Client {
	return &Client{cfg: cfg, logger: logger, wait: map[int64]chan response{}, events: make(chan notification, 100)}
}

func (c *Client) Start(ctx context.Context) error {
	c.mu.Lock()
	if c.cmd != nil {
		c.mu.Unlock()
		return nil
	}
	binary := c.cfg.Binary
	if binary == "" {
		binary = "codex"
	}
	cmd := exec.CommandContext(ctx, binary, "app-server", "--listen", "stdio://")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}
	if err := cmd.Start(); err != nil {
		c.mu.Unlock()
		return err
	}
	c.cmd = cmd
	c.stdin = stdin
	c.mu.Unlock()
	go c.readLoop(stdout)
	go c.logStderr(stderr)
	go func() {
		err := cmd.Wait()
		c.failPending(fmt.Errorf("codex app-server exited: %w", err))
	}()
	_, err = c.call(ctx, "initialize", map[string]interface{}{
		"clientInfo": map[string]string{"name": "direct-teams-bridge", "version": "0"},
	})
	return err
}

func (c *Client) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.stdin != nil {
		_ = c.stdin.Close()
	}
	if c.cmd != nil && c.cmd.Process != nil {
		_ = c.cmd.Process.Kill()
	}
	c.cmd = nil
	c.stdin = nil
	return nil
}

func (c *Client) StartThread(ctx context.Context) (string, error) {
	if err := c.Start(ctx); err != nil {
		return "", err
	}
	params := map[string]interface{}{
		"sandbox":        "read-only",
		"approvalPolicy": "never",
	}
	if c.cfg.CWD != "" {
		params["cwd"] = c.cfg.CWD
	}
	if c.cfg.Model != "" {
		params["model"] = c.cfg.Model
	}
	if c.cfg.BaseInstructions != "" {
		params["baseInstructions"] = c.cfg.BaseInstructions
	}
	raw, err := c.call(ctx, "thread/start", params)
	if err != nil {
		return "", err
	}
	var out struct {
		Thread struct {
			ID string `json:"id"`
		} `json:"thread"`
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return "", err
	}
	if out.Thread.ID == "" {
		return "", fmt.Errorf("thread/start response missing thread.id")
	}
	return out.Thread.ID, nil
}

func (c *Client) Turn(ctx context.Context, threadID, text string) (string, error) {
	c.reqMu.Lock()
	defer c.reqMu.Unlock()
	if _, err := c.call(ctx, "turn/start", map[string]interface{}{
		"threadId": threadID,
		"input":    []map[string]string{{"type": "text", "text": text}},
		"sandboxPolicy": map[string]interface{}{
			"type":          "readOnly",
			"networkAccess": false,
		},
		"approvalPolicy": "never",
	}); err != nil {
		return "", err
	}
	var answer string
	timer := time.NewTimer(10 * time.Minute)
	defer timer.Stop()
	for {
		select {
		case <-ctx.Done():
			return "", ctx.Err()
		case <-timer.C:
			return "", fmt.Errorf("codex turn timed out")
		case ev := <-c.events:
			if evThreadID(ev.Params) != threadID {
				continue
			}
			switch ev.Method {
			case "item/agentMessage/delta":
				var p struct {
					Delta string `json:"delta"`
				}
				if json.Unmarshal(ev.Params, &p) == nil {
					answer += p.Delta
				}
			case "item/completed":
				if text := completedAgentText(ev.Params); text != "" {
					answer = text
				}
			case "turn/completed":
				return answer, nil
			case "error":
				return "", fmt.Errorf("codex error: %s", string(ev.Params))
			}
		}
	}
}

func (c *Client) call(ctx context.Context, method string, params interface{}) (json.RawMessage, error) {
	id := c.seq.Add(1)
	body, err := json.Marshal(map[string]interface{}{"jsonrpc": "2.0", "id": id, "method": method, "params": params})
	if err != nil {
		return nil, err
	}
	ch := make(chan response, 1)
	c.mu.Lock()
	if c.stdin == nil {
		c.mu.Unlock()
		return nil, fmt.Errorf("codex app-server is not running")
	}
	c.wait[id] = ch
	_, err = c.stdin.Write(append(body, '\n'))
	c.mu.Unlock()
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		c.mu.Lock()
		delete(c.wait, id)
		c.mu.Unlock()
		return nil, ctx.Err()
	case resp := <-ch:
		return resp.Result, resp.Err
	}
}

func (c *Client) readLoop(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Buffer(make([]byte, 0, 64*1024), 10*1024*1024)
	for scanner.Scan() {
		var msg rpcMessage
		if err := json.Unmarshal(scanner.Bytes(), &msg); err != nil {
			c.logger.Printf("[codex] invalid json-rpc message: %v", err)
			continue
		}
		if len(msg.ID) > 0 && msg.Method == "" {
			var id int64
			_ = json.Unmarshal(msg.ID, &id)
			c.mu.Lock()
			ch := c.wait[id]
			delete(c.wait, id)
			c.mu.Unlock()
			if ch != nil {
				if msg.Error != nil {
					ch <- response{Err: fmt.Errorf("codex rpc error %d: %s", msg.Error.Code, msg.Error.Message)}
				} else {
					ch <- response{Result: msg.Result}
				}
			}
			continue
		}
		if len(msg.ID) > 0 && msg.Method != "" {
			c.rejectServerRequest(msg.ID)
			continue
		}
		if msg.Method != "" {
			select {
			case c.events <- notification{Method: msg.Method, Params: msg.Params}:
			default:
				c.logger.Printf("[codex] dropping notification method=%s", msg.Method)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		c.failPending(err)
	}
}

func (c *Client) rejectServerRequest(id json.RawMessage) {
	body, _ := json.Marshal(map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      json.RawMessage(id),
		"error":   map[string]interface{}{"code": -32000, "message": "Teams bridge runs Codex in read-only mode"},
	})
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.stdin != nil {
		_, _ = c.stdin.Write(append(body, '\n'))
	}
}

func (c *Client) failPending(err error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for id, ch := range c.wait {
		delete(c.wait, id)
		ch <- response{Err: err}
	}
}

func (c *Client) logStderr(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		c.logger.Printf("[codex] %s", scanner.Text())
	}
}

func evThreadID(raw json.RawMessage) string {
	var p struct {
		ThreadID string `json:"threadId"`
	}
	_ = json.Unmarshal(raw, &p)
	return p.ThreadID
}

func completedAgentText(raw json.RawMessage) string {
	var p struct {
		Item struct {
			Type string `json:"type"`
			Text string `json:"text"`
		} `json:"item"`
	}
	if json.Unmarshal(raw, &p) != nil || p.Item.Type != "agentMessage" {
		return ""
	}
	return p.Item.Text
}
