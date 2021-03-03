package utils

import (
	"sort"
	"strings"
)

// Replace all keys in the original string with the given values of the map
func ReplaceAllMapped(original string, toReplace map[string]string) string {
	var (
		filledInDocument = original
		keys             = MapKeys(toReplace)
	)

	sort.SliceStable(keys, func(i, j int) bool {
		return len(keys[i]) > len(keys[j])
	})

	for _, key := range keys {
		filledInDocument = strings.ReplaceAll(filledInDocument, key, toReplace[key])
	}
	return filledInDocument
}
