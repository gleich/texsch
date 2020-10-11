package location

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/setroot"
	"gopkg.in/yaml.v3"
)

// Change directory to configured project root
func ChdirProjectRoot() {
	path := ProjectRoot()
	err := os.Chdir(path)
	if err != nil {
		statuser.Error("Failed to change directory to project root", err, 1)
	}
}

// Get the project root path from the global config
func ProjectRoot() (path string) {
	_, err := ioutil.ReadFile(RootConfigFile)
	if err != nil {
		path = setroot.Set(RootConfigFile, GlobalConfigDir())
	} else {
		data, err := ioutil.ReadFile(RootConfigFile)
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
