package initialize

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/texsch/pkg/configuration"
	"github.com/Matt-Gleich/texsch/pkg/utils"
)

func CreateFiles() {
	fmt.Println()
	generalConfig := configuration.GetGeneral()

	utils.WriteFileSafely(`.gitignore`, []byte(gitIgnoreTemplate), true, true)

	// LICENSE
	filledInLICENSE := utils.ReplaceAllMapped(
		licenseTemplate,
		map[string]string{
			"CURRENT_YEAR": fmt.Sprint(time.Now().Year()),
			"FULL_NAME":    generalConfig.Full_Name,
		},
	)
	utils.WriteFileSafely("LICENSE", []byte(filledInLICENSE), true, true)

	// README
	filledInREADME := utils.ReplaceAllMapped(
		readmeTemplate,
		map[string]string{
			"SCHOOL_YEAR": generalConfig.School_Year,
		},
	)
	utils.WriteFileSafely("README.md", []byte(filledInREADME), true, true)
}
