package configure

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

type TemplatesAnswers struct {
	Defaults bool
	Emojis   bool
}

func Templates() TemplatesAnswers {
	fmt.Println()
	var defaults bool
	prompt := &survey.Confirm{
		Message: "Do you want to use the default templates?",
	}
	err := survey.AskOne(prompt, &defaults)
	if err != nil {
		statuser.Error("Failed to ask about default template use", err, 1)
	}
	if !defaults {
		return TemplatesAnswers{}
	}

	var emojis bool
	prompt = &survey.Confirm{
		Message: "Do you want the default templates to have emojis?",
	}
	err = survey.AskOne(prompt, &emojis)
	if err != nil {
		statuser.Error("Failed to ask about emoji use for the template", err, 1)
	}
	return TemplatesAnswers{
		Defaults: defaults,
		Emojis:   emojis,
	}
}
