package utils

// Check if a list of strings contains a string
func ContainsString(str string, list []string) bool {
	for _, item := range list {
		if item == str {
			return true
		}
	}
	return false
}
