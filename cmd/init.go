package cmd

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/setroot"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a repo with texsch",
	Run: func(cmd *cobra.Command, args []string) {
		setroot.Set(location.GlobalConfigPath)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
