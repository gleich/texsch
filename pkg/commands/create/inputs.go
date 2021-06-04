package create

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/gleich/statuser/v2"
	"github.com/gleich/texsch/pkg/utils"
	"github.com/spf13/cobra"
)

var DocumentTypes = []string{
	"Notes",
	"Worksheet",
	"Practice",
	"Paper",
	"Assessment",
	"Project",
	"Presentation",
	"Lab",
	"Other",
}

func getInputs(cmd *cobra.Command, classNames []string) DocumentOutline {
	// Setting initial data with data from flags
	data := DocumentOutline{
		Name:  utils.GetStringFlag(cmd, "name"),
		Type:  utils.GetStringFlag(cmd, "type"),
		Class: utils.GetStringFlag(cmd, "class"),
	}

	// Asking questions if needed
	questions := []*survey.Question{}
	if data.Name == "" {
		questions = append(questions, &survey.Question{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is the name of the document?"},
			Validate: survey.Required,
		})
	}
	if data.Type == "" {
		questions = append(questions, &survey.Question{
			Name: "type",
			Prompt: &survey.Select{
				Message:  "What is the type for the document?",
				Options:  DocumentTypes,
				PageSize: 30,
			},
		})
	}
	if data.Class == "" {
		questions = append(questions, &survey.Question{
			Name: "class",
			Prompt: &survey.Select{
				Message:  "What class is this for?",
				Options:  classNames,
				PageSize: 30,
			},
		})
	}

	err := survey.Ask(questions, &data)
	if err != nil {
		statuser.Error("Failed to ask document questions", err, 1)
	}

	// Validating data
	if !utils.ContainsString(data.Class, classNames) {
		statuser.ErrorMsg("Please provide a valid classname", 1)
	}
	if !utils.ContainsString(data.Type, DocumentTypes) {
		statuser.ErrorMsg("Please provide a valid document type", 1)
	}

	return data
}
