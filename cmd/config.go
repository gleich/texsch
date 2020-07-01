package cmd

import (
	"github.com/Matt-Gleich/texsch/config"
	"github.com/Matt-Gleich/texsch/util"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure texsch",
	Run: func(cmd *cobra.Command, args []string) {
		path := config.Path()
		if !util.Exists(path) {
			config.Create(path)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
