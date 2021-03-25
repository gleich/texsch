package config

// Configuration outline for a class
type Class struct {
	Name    string
	Teacher string
}

// Outline for the config
type Data struct {
	Name        string
	School_Year string
	Classes     []Class
	Commit      struct {
		Email string
	}
	Create struct {
		Clipboard bool
		Editor    string
	}
}
