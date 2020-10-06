package configure

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"github.com/enescakir/emoji"
)

// General answers outline
type GeneralAnswers struct {
	Full_Name   string
	School_Year string
}

// Configure general information
func General() GeneralAnswers {
	status.Step(emoji.Gear, "General Configuration")
	questions := []*survey.Question{
		{
			Name:      "full_name",
			Prompt:    &survey.Input{Message: "What is your first and last name?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "school_year",
			Prompt: &survey.Select{
				Message: "What year are you in?",
				Options: []string{"Freshman", "Sophomore", "Junior", "Senior"},
			},
		},
	}
	var answers GeneralAnswers
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask general questions", err, 1)
	}
	return answers
}
