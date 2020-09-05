package configure

import (
	"io/ioutil"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
)

// General answers outline
type GeneralAnswers struct {
	FullName string
	Year     string
}

// Configure general information
func General() GeneralAnswers {
	checkProjectRoot()
	return ask()
}

// Make sure the user is at the project root
func checkProjectRoot() {
	// Checking if texsch already exists
	files, err := ioutil.ReadDir("./")
	if err != nil {
		statuser.Error("Failed to read current working directory", err, 1)
	}
	for _, f := range files {
		if f.IsDir() && f.Name() == "texsch" {
			return
		}
	}

	// Asking question
	var projectRoot bool
	prompt := &survey.Confirm{
		Message: "Are currently at the root of your project?",
		Help:    "All of the configuration for your project will be held in a folder called texsch. In order to have that folder be in the correct location you need to make sure it is at the root of your project. Other files located at the root of your project are READMEs or LICENSEs.",
	}
	err = survey.AskOne(prompt, &projectRoot)
	if err != nil {
		statuser.Error("Failed to ask if you are in the root of your project", err, 1)
	}
	if !projectRoot {
		statuser.ErrorMsg("Please change directory into the root of your project before configuring your setup", 1)
	}
}

// Ask the questions
func ask() GeneralAnswers {
	questions := []*survey.Question{
		{
			Name:      "fullname",
			Prompt:    &survey.Input{Message: "What is your first and last name?"},
			Validate:  survey.Required,
			Transform: survey.Title,
		},
		{
			Name: "year",
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
