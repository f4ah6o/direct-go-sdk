package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// outputFormat controls the format of command output.
	// Valid values: "text", "json".
	outputFormat string
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
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// Validate output format
		if outputFormat != "" && outputFormat != "text" && outputFormat != "json" {
			return fmt.Errorf("invalid output format: %s (must be 'text' or 'json')", outputFormat)
		}
		return nil
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		printError(err)
		os.Exit(1)
	}
}

// printError prints an error message with helpful suggestions.
func printError(err error) {
	fmt.Fprintf(os.Stderr, "Error: %v\n\n", err)
	fmt.Fprintln(os.Stderr, "Suggestions:")
	fmt.Fprintln(os.Stderr, "  - Run 'daabgo login' to authenticate")
	fmt.Fprintln(os.Stderr, "  - Check your network connection")
	fmt.Fprintln(os.Stderr, "  - Run 'daabgo --help' for usage information")
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringVarP(&outputFormat, "output", "o", "text", "Output format: text, json")

	// Add subcommands
	rootCmd.AddCommand(completionCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(loginCmd)
	rootCmd.AddCommand(logoutCmd)
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(stopCmd)
	rootCmd.AddCommand(invitesCmd)
	rootCmd.AddCommand(versionCmd)
}
