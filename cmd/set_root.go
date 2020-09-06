package cmd

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/set_root"
	"github.com/spf13/cobra"
)

// setRootCmd represents the setRoot command
var setRootCmd = &cobra.Command{
	Use:   "set-root",
	Short: "Set the current working directory as the project root",
	Run: func(cmd *cobra.Command, args []string) {
		set_root.Set()
	},
}

func init() {
	rootCmd.AddCommand(setRootCmd)
}
