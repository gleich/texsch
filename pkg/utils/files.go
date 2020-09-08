package utils

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
)

// Write to a file safely.
// Checks if a file already exists or a folder and confirms with the user if they want to overide them. If the user doesn't want to override then the program exits with a 0 exit code.
// If the path is not relative the current working directory will be added before the path
func WriteFileSafely(path string, content []byte, relative bool, output bool) {
	if relative {
		cwd, err := os.Getwd()
		if err != nil {
			statuser.Error("Failed to get current working directory", err, 1)
		}
		if strings.HasPrefix(path, "/") {
			path = cwd + path
		} else {
			path = cwd + "/" + path
		}
	}

	info, err := os.Stat(path)
	if !os.IsNotExist(err) {
		if info.IsDir() {
			var overrideFolder bool
			prompt := &survey.Confirm{
				Message: "A folder exists at " + path + " do you want to override it?",
			}
			err := survey.AskOne(prompt, &overrideFolder)
			if err != nil {
				statuser.Error("Failed ask if you want to override the file", err, 1)
			}
			if !overrideFolder {
				statuser.ErrorMsg("Please go remove the folder at "+path+" and then rerun the command", 0)
			}
		}
	}
	if !os.IsNotExist(err) {
		var overrideFolder bool
		prompt := &survey.Confirm{
			Message: "A file exists at " + path + " do you want to override it?",
		}
		err := survey.AskOne(prompt, &overrideFolder)
		if err != nil {
			statuser.Error("Failed ask if you want to override the file", err, 1)
		}
		if !overrideFolder {
			statuser.ErrorMsg("Please go remove "+path+" and then rerun the command", 0)
		}
	}

	err = ioutil.WriteFile(path, content, 0655)
	if err != nil {
		statuser.Error("Failed to write to file at "+path, err, 1)
	}
	if output {
		status.Success("Created file at " + path)
	}
}
