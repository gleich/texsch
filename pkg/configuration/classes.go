package configuration

import (
	"io/ioutil"

	"github.com/Matt-Gleich/statuser"
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"github.com/Matt-Gleich/texsch/pkg/location"
	"gopkg.in/yaml.v3"
)

// Read from the classes.yaml file
func GetClasses() []configure.Class {
	data, err := ioutil.ReadFile(location.GetProjectRoot() + "/texsch/classes.yaml")
	if err != nil {
		statuser.Error("Failed to read from classes.yaml file", err, 1)
	}
	var content []configure.Class
	err = yaml.Unmarshal(data, &content)
	if err != nil {
		statuser.Error("Failed to parse yaml file", err, 1)
	}
	return content
}
