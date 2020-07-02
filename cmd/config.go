package cmd

import (
	"github.com/Matt-Gleich/texsch/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure texsch",
	Run: func(cmd *cobra.Command, args []string) {
		config.Init()
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
