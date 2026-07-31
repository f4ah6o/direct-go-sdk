package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestExtractGoMethodsFindsLiteralAndMethodConstants(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "z_events.go"), `package direct

const (
	MethodGetUsers = "get_users"
	MethodUnused   = "unused_method"
)
`)
	writeFile(t, filepath.Join(dir, "a_users.go"), `package direct

type Client struct{}

func (c *Client) GetUsers() {
	c.Call(MethodGetUsers, []interface{}{})
	c.Call("get_me", []interface{}{})
}
`)

	methods, err := ExtractGoMethods(dir)
	if err != nil {
		t.Fatalf("ExtractGoMethods returned error: %v", err)
	}

	assertContains(t, methods, "get_users")
	assertContains(t, methods, "get_me")
	assertNotContains(t, methods, "unused_method")
}

func TestExtractGoMethodsFindsContextAwareCalls(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "constants.go"), `package direct

const MethodGetUsers = "get_users"
`)
	writeFile(t, filepath.Join(dir, "users.go"), `package direct

import "context"

type Client struct{}

func (c *Client) GetUsers(ctx context.Context) {
	c.CallWithContext(ctx, MethodGetUsers, nil)
	c.callWithContext(ctx, "get_profile", nil)
}
`)

	methods, err := ExtractGoMethods(dir)
	if err != nil {
		t.Fatalf("ExtractGoMethods returned error: %v", err)
	}

	assertContains(t, methods, "get_users")
	assertContains(t, methods, "get_profile")
}

func TestExtractGoMethodsFindsConnectionBoundCalls(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "constants.go"), `package direct

const MethodStartNotification = "start_notification"
`)
	writeFile(t, filepath.Join(dir, "client.go"), `package direct

type Client struct{}

func (c *Client) initialize(conn interface{}) {
	c.callOnConnection(conn, "create_session", nil, nil, nil)
	c.callOnConnection(conn, MethodStartNotification, nil, nil, nil)
	c.callOnConnection(conn, "reset_notification", nil, nil, nil)
	c.callOnConnection(conn, "update_last_used_at", nil, nil, nil)
}
`)

	methods, err := ExtractGoMethods(dir)
	if err != nil {
		t.Fatalf("ExtractGoMethods returned error: %v", err)
	}

	assertContains(t, methods, "create_session")
	assertContains(t, methods, "start_notification")
	assertContains(t, methods, "reset_notification")
	assertContains(t, methods, "update_last_used_at")
}

func TestExtractGoMethodsSkipsTestsAndTools(t *testing.T) {
	dir := t.TempDir()
	writeFile(t, filepath.Join(dir, "client.go"), `package direct

type Client struct{}

func (c *Client) GetMe() {
	c.Call("get_me", []interface{}{})
}
`)
	writeFile(t, filepath.Join(dir, "client_test.go"), `package direct

func testOnly(c *Client) {
	c.Call("test_only", []interface{}{})
}
`)
	writeFile(t, filepath.Join(dir, "tools", "coverage", "main.go"), `package main

func toolOnly(c interface{ Call(string, []interface{}) }) {
	c.Call("tool_only", nil)
}
`)

	methods, err := ExtractGoMethods(dir)
	if err != nil {
		t.Fatalf("ExtractGoMethods returned error: %v", err)
	}

	assertContains(t, methods, "get_me")
	assertNotContains(t, methods, "test_only")
	assertNotContains(t, methods, "tool_only")
}

func TestResolvePathsDetectsDirectGoRootFromToolDirectory(t *testing.T) {
	dir := t.TempDir()
	root := filepath.Join(dir, "direct-go")
	toolDir := filepath.Join(root, "tools", "coverage")
	writeFile(t, filepath.Join(root, "go.mod"), "module example.com/direct-go\n")
	writeFile(t, filepath.Join(root, "client.go"), "package direct\n")
	writeFile(t, filepath.Join(root, "events.go"), "package direct\n")
	if err := os.MkdirAll(toolDir, 0755); err != nil {
		t.Fatalf("creating tool dir: %v", err)
	}

	previous, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	t.Cleanup(func() {
		if err := os.Chdir(previous); err != nil {
			t.Fatalf("restoring cwd: %v", err)
		}
	})
	if err := os.Chdir(toolDir); err != nil {
		t.Fatalf("Chdir: %v", err)
	}

	paths, err := resolvePaths("", "")
	if err != nil {
		t.Fatalf("resolvePaths returned error: %v", err)
	}
	want, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatalf("EvalSymlinks: %v", err)
	}
	if paths.goPath != root {
		if paths.goPath != want {
			t.Fatalf("goPath = %q, want %q", paths.goPath, want)
		}
	}
	wantJSPath := filepath.Join(want, "direct-js-source")
	if paths.jsPath != wantJSPath {
		t.Fatalf("jsPath = %q, want %q", paths.jsPath, wantJSPath)
	}
}

func writeFile(t *testing.T, path, content string) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatalf("creating parent dir for %s: %v", path, err)
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("writing %s: %v", path, err)
	}
}

func assertContains(t *testing.T, methods []string, want string) {
	t.Helper()
	for _, method := range methods {
		if method == want {
			return
		}
	}
	t.Fatalf("methods %v does not contain %q", methods, want)
}

func assertNotContains(t *testing.T, methods []string, want string) {
	t.Helper()
	for _, method := range methods {
		if method == want {
			t.Fatalf("methods %v unexpectedly contains %q", methods, want)
		}
	}
}
