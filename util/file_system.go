package util

import "os"

// Exists ... Check if a file or exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
