package commands

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/create"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a document",
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()
		create.Document()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}
