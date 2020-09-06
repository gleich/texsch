package location

import (
	"io/ioutil"
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/commands/set_root"
)

// Change directory to configured project root
func ChdirProjectRoot() {
	// Checking if texsch folder already exists
	files, err := ioutil.ReadDir("./")
	if err != nil {
		statuser.Error("Failed to read current working directory", err, 1)
	}
	for _, f := range files {
		if f.IsDir() && f.Name() == "texsch" {
			return
		}
	}
	envPath := os.Getenv(set_root.EnvKey)
	if envPath == "" {
		set_root.Set()
	}

	err = os.Chdir(envPath)
	if err != nil {
		statuser.Error("Failed to change directory to project root", err, 1)
	}
}
