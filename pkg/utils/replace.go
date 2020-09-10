package utils

import "strings"

// Replace all keys in the original string with the given values of the map
func ReplaceAllMapped(original string, toReplace map[string]string) string {
	filledInDocument := original
	for key, value := range toReplace {
		filledInDocument = strings.ReplaceAll(filledInDocument, key, value)
	}
	return filledInDocument
}
