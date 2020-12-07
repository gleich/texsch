package auto_build

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/statuser/v2"
	"gopkg.in/djherbis/times.v1"
)

// Build the files that have been changed in the last 10 seconds.
func BuildFiles(files []string) {
	cwd, err := os.Getwd()
	if err != nil {
		statuser.Error("Failed to get current working directory", err, 1)
	}

	for _, path := range files {
		if strings.HasSuffix(path, ".tex") && !strings.HasSuffix(path, "__latexindent_temp.tex") {
			t, err := times.Stat(path)
			if err != nil {
				statuser.Error("Failed to status time for "+path, err, 1)
			}
			timeThreshold := time.Now().Add(time.Duration(-2) * time.Second)

			if t.ModTime().Unix() >= timeThreshold.Unix() {
				fName := filepath.Base(path)

				// Go to the folder when the file lives
				err = os.Chdir(strings.TrimSuffix(path, fName))
				if err != nil {
					logoru.Error("Failed to go to folder for", path, ";", err)
					break
				}

				// Building file with pdflatex
				logoru.Info("Starting to build", fName, "with pdflatex")
				cmd := exec.Command("pdflatex", fName)
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr
				err = cmd.Run()
				if err != nil {
					logoru.Error("Failed to compile", path, "with pdflatex;", err)
					break
				}
				logoru.Success("Built", fName)

				// Cleaning up the auxiliary files
				err = exec.Command("latexmk", "-c", fName).Run()
				if err != nil {
					logoru.Error("Failed to cleanup", path, "with latexmk -c;", err)
					break
				}

				// Going back to original execution poin
				err = os.Chdir(cwd)
				if err != nil {
					statuser.Error("Failed to go to back to original execution directory", err, 1)
					break
				}

			}
		}
	}
}
