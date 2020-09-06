package cmd

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/create"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a document",
	Run: func(cmd *cobra.Command, args []string) {
		create.Document()
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
}