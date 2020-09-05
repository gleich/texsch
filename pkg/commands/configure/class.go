package configure

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

// Class outline
type Class struct {
	Name        string
	Time        string
	TeacherName string
}

func Classes() []Class {
	fmt.Println(
		`
——————————————————————————————————————————————————
We are now going to configure classes.
Once you are done just enter nothing for the name.
——————————————————————————————————————————————————`,
	)
	return askClasses()
}

func askClasses() []Class {
	classes := []Class{}
	for {
		fmt.Println()
		var className string
		prompt := &survey.Input{
			Message: "What is the name of the class?",
		}
		err := survey.AskOne(prompt, &className)
		if err != nil {
			statuser.Error("Failed to ask what the name of the class is", err, 1)
		}
		if className == "" {
			break
		}
		questions := []*survey.Question{
			{
				Name:     "time",
				Prompt:   &survey.Input{Message: "What time is the class (i.e. period, block)?"},
				Validate: survey.Required,
			},
			{
				Name:     "teacherName",
				Prompt:   &survey.Input{Message: "What is the name of your teacher (with title prefix; i.e. Mr. or Ms.)?"},
				Validate: survey.Required,
			},
		}
		var answers Class
		err = survey.Ask(questions, &answers)
		if err != nil {
			statuser.Error("Failed to ask question about class", err, 1)
		}
		classes = append(classes, Class{className, answers.Time, answers.TeacherName})
	}
	return classes
}