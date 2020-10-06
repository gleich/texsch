package create

import (
	"os/exec"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/configuration"
	"github.com/atotto/clipboard"
)

func ClipboardAndOpen(path string) {
	config := configuration.GetCreateConfig()

	// Clipboard
	if config.Clipboard {
		err := clipboard.WriteAll(path)
		if err != nil {
			statuser.Error("Failed to copy path to clipboard", err, 1)
		}
	}

	// Editor open
	if config.Editor != "" {
		err := exec.Command(config.Editor, path).Run()
		if err != nil {
			statuser.Error("Failed to open file in editor", err, 1)
		}
	}
}
