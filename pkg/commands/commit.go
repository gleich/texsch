package commands

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/commit"
	"github.com/Matt-Gleich/texsch/pkg/commands/separate"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"github.com/spf13/cobra"
)

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit the changes you've made to your LaTeX and PDF files",
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()

		files := separate.Files()
		separate.MoveFiles(files, false)

		changes, workingTree := commit.GetChanges()
		status.Success("Got changes")
		commit.CommitChanges(changes, workingTree)
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)
}
