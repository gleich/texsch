package commands

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/Matt-Gleich/texsch/pkg/commands/initialize"
	"github.com/Matt-Gleich/texsch/pkg/commands/setroot"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize a repo with texsch",
	Run: func(cmd *cobra.Command, args []string) {
		setroot.Set(location.RootConfigFile, location.GlobalConfigDir())
		configure.Write(
			configure.General(),
			configure.Classes(),
			configure.Commits(),
			configure.Templates(),
			configure.Create(),
		)
		initialize.CreateFiles()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
