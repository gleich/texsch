package configure

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
	"github.com/enescakir/emoji"
)

type TemplatesAnswers struct {
	Defaults bool
	Emojis   bool
}

// Configure the templates
func Templates() TemplatesAnswers {
	status.Step(emoji.Printer, "Template Configuration")
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
