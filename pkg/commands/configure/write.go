package configure

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/configure/templates"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"gopkg.in/yaml.v3"
)

// Write the changes to the files
func Write(generalConfig GeneralAnswers, classes []Class, templatesConfig TemplatesAnswers) {
	status.Step("✍️", "Writing Changes")
	confirm()
	createFolder()
	writeYaml(generalConfig, "texsch/configuration.yaml")
	writeYaml(classes, "texsch/classes.yaml")
	writeTemplates(templatesConfig)
}

func confirm() {
	var permToWrite bool
	prompt := &survey.Confirm{
		Message: "Are you sure you want to write the changes?",
	}
	err := survey.AskOne(prompt, &permToWrite)
	if err != nil {
		statuser.Error("Failed to ask about writing changes", err, 1)
	}
	if !permToWrite {
		fmt.Println("Didn't write any changes")
		os.Exit(0)
	}
}

// Create the texsch folder
func createFolder() {
	err := os.RemoveAll("texsch")
	if err != nil {
		statuser.Error("Failed to remove texsch folder", err, 1)
	}

	err = os.Mkdir("texsch", 0700)
	if err != nil {
		statuser.Error("Failed to create folder called texsch", err, 1)
	}
}

// Write to a yaml file
func writeYaml(data interface{}, fName string) {
	yamlContent, err := yaml.Marshal(data)
	if err != nil {
		statuser.Error("Failed to create yaml for "+fName, err, 1)
	}

	err = ioutil.WriteFile(fName, []byte(yamlContent), 0700)
	if err != nil {
		statuser.Error("Failed to write yaml to "+fName, err, 1)
	}
	status.Success("Wrote to " + fName)
}

// Write the templates
func writeTemplates(templatesConfig TemplatesAnswers) {
	if !templatesConfig.Defaults {
		os.Exit(0)
	}

	err := os.Mkdir("texsch/templates", 0700)
	if err != nil {
		statuser.Error("Failed to create templates folder", err, 1)
	}

	files := map[string][]byte{}
	if templatesConfig.Emojis {
		files["default.txt"] = []byte(templates.DefaultEmojiVersion)
	} else {
		files["default.txt"] = []byte(templates.DefaultEmojiVersion)
	}

	for fileName, fileContent := range files {
		filePath := "texsch/templates/" + fileName
		err = ioutil.WriteFile(filePath, fileContent, 0700)
		if err != nil {
			statuser.Error("Failed to write to "+filePath, err, 1)
		}
		status.Success("Wrote to " + filePath)
	}

}
