package cli

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the daabgo bot as a daemon",
	Long:  `Start the bot as a background daemon process. Logs will be written to ~/.daabgo/daabgo.log`,
	RunE: func(cmd *cobra.Command, args []string) error {
		return startDaemon()
	},
}

func startDaemon() error {
	// Check if already running
	pid, err := ReadPID()
	if err == nil {
		if IsProcessRunning(pid) {
			fmt.Printf("Daemon is already running with PID %d\n", pid)
			return nil
		}
		// PID file exists but process is dead, remove stale PID file
		RemovePID()
	}

	// Start daemon
	if err := Daemonize(); err != nil {
		return fmt.Errorf("failed to start daemon: %w", err)
	}

	pid, _ = ReadPID()
	logFile, _ := GetLogFile()
	fmt.Printf("Daemon started with PID %d\n", pid)
	fmt.Printf("Logs: %s\n", logFile)

	return nil
}
