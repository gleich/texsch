package configure

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
)

// General answers outline
type CreateAnswers struct {
	Clipboard bool
	Editor    string
}

// Configure document creation
func Create() CreateAnswers {
	status.Step("✨", "Document Creation Configuration")
	questions := []*survey.Question{
		{
			Name:     "clipboard",
			Prompt:   &survey.Confirm{Message: "Do you want the path of the created document to be copied to you clipboard?"},
			Validate: survey.Required,
		},
		{
			Name:   "editor",
			Prompt: &survey.Input{Message: "What CLI tool do you want to open the file with? (blank for nothing)"},
		},
	}
	var answers CreateAnswers
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask general questions", err, 1)
	}
	return answers
}
