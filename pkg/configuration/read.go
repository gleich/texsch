package configuration

import (
	"io/ioutil"

	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/yaml.v3"
)

// Read from a yaml file
func readYAML(fPath string, out interface{}) {
	data, err := ioutil.ReadFile(fPath)
	if err != nil {
		statuser.Error("Failed to read from classes.yaml file", err, 1)
	}
	err = yaml.Unmarshal(data, out)
	if err != nil {
		statuser.Error("Failed to parse yaml file", err, 1)
	}
}
