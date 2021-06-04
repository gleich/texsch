package commands

import (
	"time"

	"github.com/gleich/texsch/pkg/commands/auto_build"
	"github.com/gleich/texsch/pkg/commands/separate"
	"github.com/gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// autoBuildCmd represents the separate command
var autoBuildCmd = &cobra.Command{
	Use:   "auto-build",
	Short: "Automatically build latex files when changed",
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()
		for {
			files := separate.Files()
			auto_build.BuildFiles(files)
			files = separate.Files()
			separate.MoveFiles(files, false)
			time.Sleep(2 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(autoBuildCmd)
}
