package create

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/config"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"github.com/dustin/go-humanize"
	"github.com/spf13/cobra"
)

// Create a document
func Document(cmd *cobra.Command, classes []string) string {
	inputs := getInputs(cmd, classes)
	folderPath := createFolder(inputs)
	filePath := createFile(inputs, folderPath)
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
	conf := config.Read()
	var class config.Class
	for _, classConf := range conf.Classes {
		if classConf.Name == answers.Class {
			class = classConf
		}
	}
	today := time.Now()
	filledInDocument := utils.ReplaceAllMapped(
		string(templateContent),
		map[string]string{
			"PATH_ASSIGNMENT_NAME_ESCAPED": utils.ReplaceAllMapped(
				answers.Name,
				map[string]string{
					"#": "\\#",
					"$": "\\$",
					"&": "\\&",
					" ": "-",
				},
			),
			"PATH_ASSIGNMENT_NAME": strings.ReplaceAll(answers.Name, " ", "-"),
			"ASSIGNMENT_NAME": utils.ReplaceAllMapped(
				answers.Name,
				map[string]string{
					"#": "\\#",
					"$": "\\$",
					"&": "\\&",
				},
			),
			"AUTHOR_FULL_NAME": conf.Name,
			"CLASS_NAME":       class.Name,
			"CLASS_TEACHER":    class.Teacher,
			"DATE": fmt.Sprintf(
				"%v, %v %v\\textsuperscript{%v}, %v",
				today.Weekday(),
				today.Month(),
				today.Day(),
				strings.TrimLeft(humanize.Ordinal(today.Day()), fmt.Sprint(today.Day())),
				today.Year(),
			),
			"YEAR_NUMBER":     fmt.Sprint(today.Year()),
			"SCHOOL_YEAR":     conf.School_Year,
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
