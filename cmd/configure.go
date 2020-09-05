package cmd

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure texsch with an interactive prompt",
	Run: func(cmd *cobra.Command, args []string) {
		configure.Write(
			configure.General(),
			configure.Classes(),
			configure.Templates(),
		)
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
