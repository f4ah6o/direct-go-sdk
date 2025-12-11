package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "daabgo",
	Short: "daabgo - Go implementation of daab bot framework",
	Long: `daabgo is a Go implementation of daab (direct agent assist bot).
	
It allows you to create and run bots for the direct chat service.

Available Commands:
  init      Setup a new daabgo bot project
  login     Login to direct as a bot account
  logout    Logout from the service
  run       Run the bot
  version   Show version information`,
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(invitesCmd)
	rootCmd.AddCommand(versionCmd)
}
