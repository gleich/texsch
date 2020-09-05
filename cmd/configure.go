package cmd

import (
	"fmt"

	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/spf13/cobra"
)

// configureCmd represents the configure command
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configure texsch with an interactive prompt",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(configure.General())
		fmt.Println(configure.Classes())
	},
}

func init() {
	rootCmd.AddCommand(configureCmd)
}
