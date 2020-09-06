package cmd

import (
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "texsch",
	Short: "🏫 Automation for techy students that write papers for high school in LaTeX",
	Long: `
🏫 Automation for techy students that write papers for high school in LaTeX

🐙 Repository: https://github.com/Matt-Gleich/texsch
📟 Authors:
	- Matthew Gleich (@Matt-Gleich)

████████╗███████╗██╗  ██╗███████╗ ██████╗██╗  ██╗
╚══██╔══╝██╔════╝╚██╗██╔╝██╔════╝██╔════╝██║  ██║
   ██║   █████╗   ╚███╔╝ ███████╗██║     ███████║
   ██║   ██╔══╝   ██╔██╗ ╚════██║██║     ██╔══██║
   ██║   ███████╗██╔╝ ██╗███████║╚██████╗██║  ██║
   ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝ ╚═════╝╚═╝  ╚═╝`,
}

// Execute the main command
func Execute() {
	location.ChdirProjectRoot()
	if err := rootCmd.Execute(); err != nil {
		statuser.Error("Failed to execute root command", err, 1)
	}
}
