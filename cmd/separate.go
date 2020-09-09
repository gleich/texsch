package cmd

import (
	"time"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/separate"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// separateCmd represents the separate command
var separateCmd = &cobra.Command{
	Use:   "separate",
	Short: "Separate the pdf files from the latex files",
	Run: func(cmd *cobra.Command, args []string) {
		location.ChdirProjectRoot()
		loop, err := cmd.Flags().GetBool("loop")
		if err != nil {
			statuser.Error("Failed to get loop flag", err, 1)
		}
		if loop {
			for {
				files := separate.Files()
				separate.MoveFiles(files, true)
				time.Sleep(10 * time.Second)
			}
		} else {
			files := separate.Files()
			separate.MoveFiles(files, false)
		}
	},
}

func init() {
	rootCmd.AddCommand(separateCmd)
	separateCmd.Flags().BoolP("loop", "l", false, "Loop the program every 10 seconds")
}
