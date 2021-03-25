package configuration

import (
	"sort"

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

// Get the names of all the classes in alphabetical order
func GetClassNames() []string {
	classNames := []string{}
	classes := GetClasses()
	for _, class := range classes {
		classNames = append(classNames, class.Name)
	}
	sort.Strings(classNames)
	return classNames
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
