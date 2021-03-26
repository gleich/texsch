package create

import (
	"os"
	"os/exec"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/config"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"github.com/atotto/clipboard"
	"github.com/spf13/cobra"
)

// Run Actions after creating the file
func Post(cmd *cobra.Command, path string) {
	conf := config.Read().Create

	// Clipboard
	if conf.Clipboard && !utils.GetBoolFlag(cmd, "no-clipboard") {
		err := clipboard.WriteAll(path)
		if err != nil {
			statuser.Error("Failed to copy path to clipboard", err, 1)
		}
	}

	// Editor open
	if conf.Editor != "" && !utils.GetBoolFlag(cmd, "no-editor") {
		cmd := exec.Command(conf.Editor, path)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			statuser.Error("Failed to open file in editor", err, 1)
		}
	}
}
