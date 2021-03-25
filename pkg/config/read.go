package config

import (
	"os"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/pelletier/go-toml"
)

// Read from the configuration file
func Read() Data {
	bin, err := os.ReadFile("texsch/config.toml")
	if err != nil {
		statuser.Error("Failed to read from config", err, 1)
	}

	var config Data
	err = toml.Unmarshal(bin, &config)
	if err != nil {
		statuser.Error("Failed to parse config", err, 1)
	}

	return config
}
