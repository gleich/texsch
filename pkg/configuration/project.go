package configuration

import (
	"github.com/Matt-Gleich/texsch/pkg/commands/configure"
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
