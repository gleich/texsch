package commands

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/commit"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes you've made to your LaTeX and PDF files",
	Run: func(cmd *cobra.Command, args []string) {
		commit.GetChanges()
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
