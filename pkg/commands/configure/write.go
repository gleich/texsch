package configure

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/configure/templates"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"gopkg.in/yaml.v3"
)

// Where the configuration files live
const (
	GeneralFile = "texsch/configuration.yaml"
	ClassesFile = "texsch/classes.yaml"
	CommitsFile = "texsch/commits.yaml"
	CreateFile  = "texsch/create.yaml"
)

// Write the changes to the files
func Write(generalConfig GeneralAnswers, classes []Class, commitConfig CommitAnswers, templatesConfig TemplatesAnswers, createConfig CreateAnswers) {
	status.Step("✍️", "Writing Changes")
	confirm()
	createFolder()
	writeYaml(generalConfig, GeneralFile)
	writeYaml(classes, ClassesFile)
	writeYaml(commitConfig, CommitsFile)
	writeYaml(createConfig, CreateFile)
	writeTemplates(templatesConfig)
}

// Get confirmation from the user to the write the changes
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
	utils.WriteFileSafely(fName, []byte(yamlContent), true, true)
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
		files["Default.txt"] = []byte(templates.DefaultEmojiVersion)
	} else {
		files["Default.txt"] = []byte(templates.DefaultPlainVersion)
	}
	for fileName, fileContent := range files {
		filePath := "texsch/templates/" + fileName
		utils.WriteFileSafely(filePath, fileContent, true, true)
	}

}
