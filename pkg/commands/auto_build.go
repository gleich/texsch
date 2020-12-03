package commands

import (
	"time"

	"github.com/Matt-Gleich/texsch/pkg/commands/auto_build"
	"github.com/Matt-Gleich/texsch/pkg/commands/separate"
	"github.com/Matt-Gleich/texsch/pkg/location"
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
			separate.MoveFiles(files, false)
			time.Sleep(5 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(autoBuildCmd)
}
