package config

import (
	"os"
	"runtime"
	"strings"

	"github.com/Matt-Gleich/statuser"
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
func Create(path string) {
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
