package cmd

import (
	"fmt"
	"strings"

	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/Matt-Gleich/texsch/pkg/commands/initialize"
	"github.com/Matt-Gleich/texsch/pkg/commands/setroot"
	"github.com/Matt-Gleich/texsch/pkg/configuration"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "initialize",
	Short: "Initialize a repo with texsch",
	Run: func(cmd *cobra.Command, args []string) {
		setroot.Set(location.GlobalConfigPath)
		configure.Write(
			configure.General(),
			configure.Classes(),
			configure.Templates(),
		)
		fmt.Println()
		utils.WriteFileSafely(`.gitignore`, []byte(initialize.GitIgnore), true, true)
		generalConfig := configuration.GetGeneral()
		utils.WriteFileSafely(`LICENSE`, []byte(strings.ReplaceAll(strings.ReplaceAll(initialize.LICENSE, "CURRENT_YEAR", generalConfig.Year), "FULL_NAME", generalConfig.Full_Name)), true, true)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
