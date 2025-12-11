package cli

import (
	"fmt"
	"os"
	"syscall"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the daabgo daemon",
	Long:  `Stop the daabgo daemon by sending SIGTERM signal.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return stopDaemon()
	},
}

func stopDaemon() error {
	// Read PID
	pid, err := ReadPID()
	if err != nil {
		fmt.Println("Daemon is not running")
		return nil
	}

	// Check if process is running
	if !IsProcessRunning(pid) {
		fmt.Println("Daemon is not running (stale PID file)")
		RemovePID()
		return nil
	}

	// Find process
	process, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("failed to find process: %w", err)
	}

	// Send SIGTERM
	if err := process.Signal(syscall.SIGTERM); err != nil {
		return fmt.Errorf("failed to stop daemon: %w", err)
	}

	// Remove PID file
	RemovePID()

	fmt.Printf("Daemon (PID %d) stopped\n", pid)
	return nil
}
