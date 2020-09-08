package configure

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
)

// Class outline
type Class struct {
	Name         string
	Teacher_Name string
}

// Configure the classes
func Classes() []Class {
	status.Step("📚", "Class Configuration")
	fmt.Println(
		`We are now going to configure classes.
Once you are done just enter nothing for the class name.`,
	)
	return askClasses()
}

// Ask about the classes
func askClasses() []Class {
	classes := []Class{}
	for {
		fmt.Println()
		var (
			className string
			prompt    []*survey.Question
			message   string = "What is the name of the class?"
		)
		if len(classes) == 0 {
			prompt = []*survey.Question{
				{
					Prompt:   &survey.Input{Message: message},
					Validate: survey.Required,
				},
			}
		} else {
			prompt = []*survey.Question{
				{
					Prompt: &survey.Input{Message: message},
				},
			}
		}
		err := survey.Ask(prompt, &className)
		if err != nil {
			statuser.Error("Failed to ask what the name of the class is", err, 1)
		}
		if className == "" {
			break
		}

		question := []*survey.Question{
			{
				Name:     "teacher_name",
				Prompt:   &survey.Input{Message: "What is the name of your teacher (with title prefix; i.e. Mr. or Ms.)?"},
				Validate: survey.Required,
			},
		}
		var answers Class
		err = survey.Ask(question, &answers)
		if err != nil {
			statuser.Error("Failed to ask question about class", err, 1)
		}
		classes = append(classes, Class{className, answers.Teacher_Name})
	}
	return classes
}
