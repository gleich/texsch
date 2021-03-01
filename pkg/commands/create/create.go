package create

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/Matt-Gleich/texsch/pkg/configuration"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"github.com/dustin/go-humanize"
)

type DocumentOutline struct {
	Name  string
	Type  string
	Class string
}

// Create a document
func Document() string {
	// Asking information
	classNames := []string{}
	classConfig := configuration.GetClasses()
	for _, classConfiguration := range classConfig {
		classNames = append(classNames, classConfiguration.Name)
	}
	sort.Strings(classNames)

	questions := []*survey.Question{
		{
			Name:     "name",
			Prompt:   &survey.Input{Message: "What is the name of the document?"},
			Validate: survey.Required,
		},
		{
			Name: "type",
			Prompt: &survey.Select{
				Message: "What is the type for the document?",
				Options: []string{
					"Notes",
					"Worksheet",
					"Practice",
					"Paper",
					"Assessment",
					"Project",
					"Presentation",
					"Lab",
					"Other",
				},
				PageSize: 30,
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
	folderPath := createFolder(answers)
	filePath := createFile(answers, folderPath)
	return filePath
}

// Create the folder for a file
func createFolder(answers DocumentOutline) string {
	path := fmt.Sprintf(
		"LaTeX/%v/%v/%v/",
		strings.ReplaceAll(answers.Class, " ", "-"),
		time.Now().Month(),
		answers.Type,
	)
	err := os.MkdirAll(path, 0700)
	if err != nil {
		statuser.Error("Failed to create folder for file", err, 1)
	}
	return path
}

// Create the actual document file
func createFile(answers DocumentOutline, folderPath string) string {
	// Reading from template
	files, err := ioutil.ReadDir("./texsch/templates")
	if err != nil {
		statuser.Error("Failed to get a list of all the templates", err, 1)
	}
	templates := []string{}
	for _, file := range files {
		if !file.IsDir() {
			templates = append(templates, strings.TrimSuffix(file.Name(), ".txt"))
		}
	}
	var templateName string
	if len(templates) > 1 {
		prompt := &survey.Select{
			Message: "What template would you like to use?",
			Options: templates,
			Default: "Default",
		}
		err := survey.AskOne(prompt, &templateName)
		if err != nil {
			statuser.Error("Failed to ask what template you want to use", err, 1)
		}
	}
	templateContent, err := ioutil.ReadFile("./texsch/templates/" + templateName + ".txt")
	if err != nil {
		statuser.Error("Failed to read from template file", err, 1)
	}

	// Replacing document variables
	var class configure.Class
	for _, classInstance := range configuration.GetClasses() {
		if classInstance.Name == answers.Class {
			class = classInstance
		}
	}
	today := time.Now()
	generalConfiguration := configuration.GetGeneral()
	filledInDocument := utils.ReplaceAllMapped(
		string(templateContent),
		map[string]string{
			"PATH_ASSIGNMENT_NAME_ESCAPED": utils.ReplaceAllMapped(
				answers.Name,
				map[string]string{
					"#": "\\#",
					"$": "\\$",
					" ": "-",
				},
			),
			"PATH_ASSIGNMENT_NAME": strings.ReplaceAll(answers.Name, " ", "-"),
			"ASSIGNMENT_NAME": utils.ReplaceAllMapped(
				answers.Name,
				map[string]string{
					"#": "\\#",
					"$": "\\$",
				},
			),
			"AUTHOR_FULL_NAME": generalConfiguration.Full_Name,
			"CLASS_NAME":       class.Name,
			"CLASS_TEACHER":    class.Teacher_Name,
			"DATE": fmt.Sprintf(
				"%v, %v %v\\textsuperscript{%v}, %v",
				today.Weekday(),
				today.Month(),
				today.Day(),
				strings.TrimLeft(humanize.Ordinal(today.Day()), fmt.Sprint(today.Day())),
				today.Year(),
			),
			"YEAR_NUMBER":     fmt.Sprint(today.Year()),
			"SCHOOL_YEAR":     generalConfiguration.Full_Name,
			"DOCUMENT_TYPE":   answers.Type,
			"PATH_CLASS-NAME": strings.ReplaceAll(answers.Class, " ", "-"),
			"MONTH_NAME":      today.Month().String(),
		},
	)

	// Creating the actual file
	filePath := folderPath + strings.ReplaceAll(answers.Name, " ", "-") + ".tex"
	utils.WriteFileSafely(filePath, []byte(filledInDocument), true, true)
	return filePath
}
