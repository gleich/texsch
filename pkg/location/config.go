package location

import (
	"path/filepath"
	"runtime"
)

var RootConfig = filepath.Join(GlobalConfigDir(), "root.yml")

// Get the global configuration directory
func GlobalConfigDir() string {
	path := HomeDir()
	if (runtime.GOOS == "darwin") || (runtime.GOOS == "linux") {
		path = filepath.Join(path, ".config/texsch/")
	}
	return path
}
