package configure

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
)

// Commit answers outline
type CommitAnswers struct {
	Email  string
	Emojis bool
}

func Commits() CommitAnswers {
	status.Step("📁", "Automatic Git Commits")
	fmt.Println("texsch can automatically commit your documents for you")
	fmt.Println()
	questions := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "What is the email you use for committing?"},
			Validate: survey.Required,
		},
		{
			Name: "emojis",
			Prompt: &survey.Confirm{
				Message: "Do you want emojis in your commits using gitmoji?",
			},
		},
	}
	var answers CommitAnswers
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask about commit configuration", err, 1)
	}
	return answers
}
