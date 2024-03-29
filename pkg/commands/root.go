package commands

import (
	"github.com/gleich/statuser/v2"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "texsch",
	Short: "🏫 Automation for techy students that write papers for high school in LaTeX",
	Long: `
🏫 Automation for techy students that write papers for high school in LaTeX

🐙 Repository: https://github.com/gleich/texsch
📟 Authors:
	- Matthew Gleich (@gleich)

████████╗███████╗██╗  ██╗███████╗ ██████╗██╗  ██╗
╚══██╔══╝██╔════╝╚██╗██╔╝██╔════╝██╔════╝██║  ██║
   ██║   █████╗   ╚███╔╝ ███████╗██║     ███████║
   ██║   ██╔══╝   ██╔██╗ ╚════██║██║     ██╔══██║
   ██║   ███████╗██╔╝ ██╗███████║╚██████╗██║  ██║
   ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝`,
}

// Execute the main command
func Execute() {
	statuser.Emojis = false
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
