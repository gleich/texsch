package cmd

import (
	"fmt"
	"time"

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
		generalConfig := configuration.GetGeneral()

		utils.WriteFileSafely(`.gitignore`, []byte(initialize.GitIgnore), true, true)

		// LICENSE
		filledInLICENSE := utils.ReplaceAllMapped(
			initialize.LICENSE,
			map[string]string{
				"CURRENT_YEAR": fmt.Sprint(time.Now().Year()),
				"FULL_NAME":    generalConfig.Full_Name,
			},
		)
		utils.WriteFileSafely("LICENSE", []byte(filledInLICENSE), true, true)

		// README
		filledInREADME := utils.ReplaceAllMapped(
			initialize.README,
			map[string]string{
				"SCHOOL_YEAR": generalConfig.School_Year,
			},
		)
		utils.WriteFileSafely("README.md", []byte(filledInREADME), true, true)
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
