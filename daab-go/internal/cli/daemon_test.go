package cli

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestRedirectOutputToLog(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("HOME", tmpDir)

	origStdout, origStderr := os.Stdout, os.Stderr
	var logFile *os.File
	defer func() {
		if logFile != nil {
			logFile.Close()
		}
		os.Stdout = origStdout
		os.Stderr = origStderr
	}()

	if err := RedirectOutputToLog(); err != nil {
		t.Fatalf("RedirectOutputToLog() error = %v", err)
	}

	logFilePath, err := GetLogFile()
	if err != nil {
		t.Fatalf("GetLogFile() error = %v", err)
	}

	logFile = os.Stdout
	if logFile == nil {
		t.Fatalf("stdout is nil after redirection")
	}

	fmt.Fprintln(os.Stdout, "stdout line")
	fmt.Fprintln(os.Stderr, "stderr line")

	if err := logFile.Sync(); err != nil {
		t.Fatalf("failed to sync log file: %v", err)
	}

	data, err := os.ReadFile(logFilePath)
	if err != nil {
		t.Fatalf("failed to read log file: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "stdout line") {
		t.Errorf("log file missing stdout output")
	}
	if !strings.Contains(content, "stderr line") {
		t.Errorf("log file missing stderr output")
	}
}
