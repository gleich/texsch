package utils

// Get a list of all the keys in map
func MapKeys(m map[string]string) []string {
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
