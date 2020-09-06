package set_root

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

const EnvKey = "TEXSCH_PATH"

// Set the project root environment variable
func Set() {
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
	err = os.Setenv(EnvKey, cwd)
	if err != nil {
		statuser.Error("Failed to set current working directory in environment variable", err, 1)
	}
}
