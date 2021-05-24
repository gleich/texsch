package commands

import (
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	DisableFlagsInUseLine: true,
	Use:                   "completion [bash|zsh|fish|powershell]",
	Short:                 "Generate shell completion for texsch",
	Hidden:                true,
	Long: `To load completions

Bash:

	$ source <(texsch completion bash)

	# To load completions for each session, execute once:
	Linux:
		$ texsch completion bash > /etc/bash_completion.d/texsch
	MacOS:
		$ texsch completion bash > /usr/local/etc/bash_completion.d/texsch

Zsh:

	# If shell completion is not already enabled in your environment you will need
	# to enable it.  You can execute the following once:

	$ echo "autoload -U compinit; compinit" >> ~/.zshrc

	# To load completions for each session, execute once:
	$ texsch completion zsh > "${fpath[1]}/_texsch"

	# You will need to start a new shell for this setup to take effect.

Fish:

	$ texsch completion fish | source

	# To load completions for each session, execute once:
	$ texsch completion fish > ~/.config/fish/completions/texsch.fish

	`,
	ValidArgs: []string{"bash", "zsh", "fish", "powershell"},
	Args:      cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletion(os.Stdout)
		}
		if err != nil {
			statuser.Error("Failed to generate completion for "+args[0], err, 1)
		}
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
