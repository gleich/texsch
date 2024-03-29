package setroot

import (
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/statuser/v2"
	"gopkg.in/yaml.v3"
)

// Set the project root environment variable
func Set(fPath string, configFolder string) string {
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

	// Creating the config folder if it doesn't already exist
	err = os.MkdirAll(configFolder, 0777)
	if err != nil {
		statuser.Error("Failed to create config folder", err, 1)
	}

	yamlContent, err := yaml.Marshal(map[string]string{
		"path": cwd,
	})
	if err != nil {
		statuser.Error("Failed to create yaml from path data", err, 1)
	}

	err = ioutil.WriteFile(fPath, []byte(yamlContent), 0666)
	if err != nil {
		statuser.Error("Failed to write project root path to global config", err, 1)
	}
	statuser.Success("Set project root")
	return cwd
}
