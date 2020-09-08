package setroot

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"gopkg.in/yaml.v3"
)

// Set the project root environment variable
func Set(fPath string) string {
	cwd, err := os.Getwd()
	if err != nil {
		statuser.Error("Failed to get current working directory", err, 1)
	}

	var confirmed bool
	prompt := &survey.Confirm{
		Message: "Is " + cwd + " the project root?",
		Help:    "Project root is the top of your project. Other files commonly located at the root of a project are README and LICENSE files.",
	}
	err = survey.AskOne(prompt, &confirmed)
	if err != nil {
		statuser.Error("Failed to ask user if current directory is project root", err, 1)
	}
	if !confirmed {
		statuser.ErrorMsg("Please change directory into the root of your project", 1)
	}

	yamlContent, err := yaml.Marshal(map[string]string{
		"path": cwd,
	})
	if err != nil {
		statuser.Error("Failed to create yaml from path data", err, 1)
	}

	utils.WriteFileSafely(fPath, []byte(yamlContent), false, false)
	status.Success("Set project root")
	return cwd
}
