package location

import (
	"os"

	"github.com/gleich/statuser/v2"
)

// Get user's home directory
func HomeDir() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		statuser.Error("Failed to get home working directory", err, 1)
	}
	return homeDir
}
