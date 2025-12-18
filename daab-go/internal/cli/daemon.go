package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"syscall"
)

// GetPIDFile returns the path to the PID file.
func GetPIDFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	daabDir := filepath.Join(homeDir, ".daabgo")
	if err := os.MkdirAll(daabDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(daabDir, "daabgo.pid"), nil
}

// GetLogFile returns the path to the log file.
func GetLogFile() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	daabDir := filepath.Join(homeDir, ".daabgo")
	if err := os.MkdirAll(daabDir, 0755); err != nil {
		return "", err
	}
	return filepath.Join(daabDir, "daabgo.log"), nil
}

// ReadPID reads the PID from the PID file.
func ReadPID() (int, error) {
	pidFile, err := GetPIDFile()
	if err != nil {
		return 0, err
	}

	data, err := os.ReadFile(pidFile)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, fmt.Errorf("daemon is not running")
		}
		return 0, err
	}

	pid, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, fmt.Errorf("invalid PID file: %w", err)
	}

	return pid, nil
}

// WritePID writes the PID to the PID file.
func WritePID(pid int) error {
	pidFile, err := GetPIDFile()
	if err != nil {
		return err
	}

	return os.WriteFile(pidFile, []byte(fmt.Sprintf("%d", pid)), 0644)
}

// RemovePID removes the PID file.
func RemovePID() error {
	pidFile, err := GetPIDFile()
	if err != nil {
		return err
	}
	return os.Remove(pidFile)
}

// IsProcessRunning checks if a process with the given PID is running.
func IsProcessRunning(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		return false
	}
	// Send signal 0 to check if process exists
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

// Daemonize starts the current program as a background daemon.
func Daemonize() error {
	// Get executable path
	executable, err := os.Executable()
	if err != nil {
		return err
	}

	// Get log file
	logFile, err := GetLogFile()
	if err != nil {
		return err
	}

	// Open log file
	log, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer log.Close()

	// Prepare command
	cmd := exec.Command(executable, "run", "--daemon")
	cmd.Stdout = log
	cmd.Stderr = log
	cmd.Stdin = nil
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true, // Create new session
	}

	// Start the process
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("failed to start daemon: %w", err)
	}

	// Write PID
	if err := WritePID(cmd.Process.Pid); err != nil {
		return fmt.Errorf("failed to write PID: %w", err)
	}

	// Detach from parent
	if err := cmd.Process.Release(); err != nil {
		return fmt.Errorf("failed to release process: %w", err)
	}

	return nil
}

// RedirectOutputToLog redirects stdout and stderr to the log file.
func RedirectOutputToLog() error {
	logFile, err := GetLogFile()
	if err != nil {
		return err
	}

	log, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}

	// Redirect stdout and stderr to the log file
	os.Stdout = log
	os.Stderr = log

	return nil
}
