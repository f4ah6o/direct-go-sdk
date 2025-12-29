package cli

import (
	"os"

	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate shell completion script",
	Long: `To load completions:

Bash:
  $ source <(daabgo completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ daabgo completion bash > /etc/bash_completion.d/daabgo
  # macOS:
  $ daabgo completion bash > /usr/local/etc/bash_completion.d/daabgo

Zsh:
  # If shell completion is not already enabled in your environment,
  # you will need to enable it. You can execute the following once:
  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ daabgo completion zsh > "${fpath[1]}/_daabgo"

  # You will need to start a new shell for this setup to take effect.

fish:
  $ daabgo completion fish | source

  # To load completions for each session, execute once:
  $ daabgo completion fish > ~/.config/fish/completions/daabgo.fish

PowerShell:
  PS> daabgo completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> daabgo completion powershell > daabgo.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		switch args[0] {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		}
	},
}
