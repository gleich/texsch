package cmd

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/set_root"
	"github.com/spf13/cobra"
)

// setRootCmd represents the setRoot command
var setRootCmd = &cobra.Command{
	Use:   "set-root",
	Short: "Set the current working directory as the project root",
	Long: `
Set the environment variable for the project root. The project root is where your
README is for example.

You will need to do this if you ever move the location of your project on your computer
or plan to run texsch commands on another project.
`,
	Run: func(cmd *cobra.Command, args []string) {
		set_root.Set()
	},
}

func init() {
	rootCmd.AddCommand(setRootCmd)
}
