package create

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/configuration"
)

type DocumentOutline struct {
	Name  string
	Type  string
	Class string
}

// Create a document
func Document() {
	classNames := []string{}
	classConfig := configuration.GetClasses()
	for _, classConfiguration := range classConfig {
		classNames = append(classNames, classConfiguration.Name)
	}

	questions := []*survey.Question{
		{
			Name:      "name",
			Prompt:    &survey.Input{Message: "What is the name of the document?"},
			Transform: survey.Title,
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What is the type for the document?",
				Options: []string{
					"Paper",
					"Notes",
					"Practice",
					"Assessment",
					"Project",
					"Presentation",
					"Other",
				},
			},
		},
		{
			Name: "class",
			Prompt: &survey.Select{
				Message: "What class is this for?",
				Options: classNames,
			},
		},
	}
	var answers DocumentOutline
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask document questions", err, 1)
	}
}
