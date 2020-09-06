package configuration

import (
	"io/ioutil"

	"github.com/Matt-Gleich/statuser"
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"gopkg.in/yaml.v3"
)

// Read from the configuration.yaml file
func GetGeneral() configure.GeneralAnswers {
	var data configure.GeneralAnswers
	readYAML("/texsch/configuration.yaml", &data)
	return data
}

// Read from the classes.yaml file
func GetClasses() []configure.Class {
	var data []configure.Class
	readYAML("/texsch/classes.yaml", &data)
	return data
}

// Read from a yaml file
func readYAML(fPath string, out interface{}) {
	data, err := ioutil.ReadFile(location.GetProjectRoot() + fPath)
	if err != nil {
		statuser.Error("Failed to read from classes.yaml file", err, 1)
	}
	err = yaml.Unmarshal(data, out)
	if err != nil {
		statuser.Error("Failed to parse yaml file", err, 1)
	}
}
