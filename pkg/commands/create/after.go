package create

import (
	"os"
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
		cmd := exec.Command(config.Editor, path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			statuser.Error("Failed to open file in editor", err, 1)
		}
	}
}
