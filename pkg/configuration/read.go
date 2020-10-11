package configuration

import (
	"io/ioutil"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
	"gopkg.in/yaml.v3"
)

// Read from the configuration.yaml file
func GetGeneral() configure.GeneralAnswers {
	var data configure.GeneralAnswers
	readYAML(configure.GeneralFile, &data)
	return data
}

// Read from the classes.yaml file
func GetClasses() []configure.Class {
	var data []configure.Class
	readYAML(configure.ClassesFile, &data)
	return data
}

// Read from the commits.yaml file
func GetCommitConfig() configure.CommitAnswers {
	var data configure.CommitAnswers
	readYAML(configure.CommitsFile, &data)
	return data
}

// Read from the create.yaml file
func GetCreateConfig() configure.CreateAnswers {
	var data configure.CreateAnswers
	readYAML(configure.CreateFile, &data)
	return data
}

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
