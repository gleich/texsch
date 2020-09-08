package location

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/setroot"
	"gopkg.in/yaml.v3"
)

var GlobalConfigPath = GetHomeDir() + "/.texsch.yaml"

// Change directory to configured project root
func ChdirProjectRoot() {
	path := GetProjectRoot()
	err := os.Chdir(path)
	if err != nil {
		statuser.Error("Failed to change directory to project root", err, 1)
	}
}

// Get the project root path from the global config
func GetProjectRoot() string {
	var path string
	_, err := ioutil.ReadFile(GlobalConfigPath)
	if err != nil {
		path = setroot.Set(GlobalConfigPath)
	} else {
		data, err := ioutil.ReadFile(GlobalConfigPath)
		if err != nil {
			statuser.Error("Failed to read from global config file", err, 1)
		}
		globalConfig := struct {
			Path string `yaml:"path"`
		}{}
		err = yaml.Unmarshal(data, &globalConfig)
		if err != nil {
			statuser.Error("Failed to parse yaml from global config file", err, 1)
		}
		path = globalConfig.Path
	}
	return path
}

// Get user's home directory
func GetHomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home working directory", err, 1)
	}
	return homeDir
}
