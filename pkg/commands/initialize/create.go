package initialize

import (
	"fmt"
	"time"

	"github.com/Matt-Gleich/texsch/pkg/config"
	"github.com/Matt-Gleich/texsch/pkg/utils"
)

func CreateFiles() {
	fmt.Println()
	generalConfig := config.Read()

	utils.WriteFileSafely(`.gitignore`, []byte(gitIgnoreTemplate), true, true)

	// LICENSE
	filledInLICENSE := utils.ReplaceAllMapped(
		licenseTemplate,
		map[string]string{
			"CURRENT_YEAR": fmt.Sprint(time.Now().Year()),
			"FULL_NAME":    generalConfig.Name,
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
