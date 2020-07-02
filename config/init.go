package config

import "fmt"

// Init ... Init the config file
func Init() {
	path := Path()
	Create(path, true)
	fmt.Println("Created the config file at " + path + " please go fill it in")
}
