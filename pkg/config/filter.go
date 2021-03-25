package config

// Get the class names from the given config
func ClassNames(data Data) []string {
	classes := []string{}
	for _, class := range data.Classes {
		classes = append(classes, class.Name)
	}
	return classes
}
