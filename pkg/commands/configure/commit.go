package configure

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"github.com/enescakir/emoji"
)

// Commit answers outline
type CommitAnswers struct {
	Email string
}

func Commits() CommitAnswers {
	status.Step(emoji.FileFolder, "Automatic Git Commits")
	fmt.Println("texsch can automatically commit your documents for you")
	fmt.Println()
	questions := []*survey.Question{
		{
			Name:     "email",
			Prompt:   &survey.Input{Message: "What is the email you use for committing?"},
			Validate: survey.Required,
		},
	}
	var answers CommitAnswers
	err := survey.Ask(questions, &answers)
	if err != nil {
		statuser.Error("Failed to ask about commit configuration", err, 1)
	}
	return answers
}
