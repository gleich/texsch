package config

import (
	"os"
	"runtime"
	"strings"

	"github.com/Matt-Gleich/statuser"
	"github.com/manifoldco/promptui"
)

const fileName = "config.yml"

// Path ...  Get the path to the config file
func Path() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home directory in order to load config", err, 1)
	}
	if runtime.GOOS == "windows" {
		return homeDir + "\\texsch\\" + fileName
	}
	return homeDir + "/.config/texsch/" + fileName
}

// Create ... Create the config file
func Create(path string, askUser bool) {
	if askUser {
		_, err := os.Stat(path)
		if !os.IsNotExist(err) {
			prompt := promptui.Select{
				Label: "The config already exits. Do you want to reset it?",
				Items: []string{"Yes", "No"},
			}
			_, result, err := prompt.Run()
			if err != nil {
				statuser.Error("Failed to load prompt", err, 1)
			}
			if result == "No" {
				os.Exit(0)
			}
		}
	}
	folderPath := strings.TrimSuffix(path, fileName)
	err := os.MkdirAll(folderPath, 0700)
	if err != nil {
		statuser.Error("Failed to create folder for config", err, 1)
	}
	_, err = os.Create(path)
	if err != nil {
		statuser.Error("Failed to create config file", err, 1)
	}
}
