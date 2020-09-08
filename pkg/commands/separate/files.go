package separate

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/status"
)

// Getting all the tex and pdf files for the project
func Files() (filePaths []string) {
	// Getting files in project recursively
	allFiles := []string{}
	err := filepath.Walk(".", func(path string, f os.FileInfo, err error) error {
		allFiles = append(allFiles, path)
		return nil
	})
	if err != nil {
		statuser.Error("Failed to get all tex and pdf files", err, 1)
	}

	for _, filePath := range allFiles {
		if strings.HasSuffix(filePath, ".tex") || strings.HasSuffix(filePath, ".pdf") {
			filePaths = append(filePaths, filePath)
		}
	}

	return filePaths
}

// Move the files
func MoveFiles(filePaths []string, loop bool) {
	for _, filePath := range filePaths {
		var prefix string
		if strings.HasSuffix(filePath, ".tex") {
			latexPrefix := "LaTeX/"
			if strings.HasPrefix(filePath, latexPrefix) {
				continue
			}
			prefix = latexPrefix
		} else if strings.HasSuffix(filePath, ".pdf") {
			pdfPrefix := "PDF/"
			if strings.HasPrefix(filePath, pdfPrefix) {
				continue
			}
			prefix = pdfPrefix
		}
		fileFolderPathChunks := strings.Split(filePath, "/")
		fileFolderPath := prefix + strings.Join(fileFolderPathChunks[:len(fileFolderPathChunks)-1], "/") + "/"
		fileName := fileFolderPathChunks[len(fileFolderPathChunks)-1]

		_, err := os.Stat(fileFolderPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(fileFolderPath, 0655)
			if err != nil {
				statuser.Error("Failed to make folders", err, 1)
			}
		}

		err = os.Rename(filePath, fileFolderPath+fileName)
		if err != nil {
			statuser.Error("Failed to move file", err, 1)
		}
		if loop {
			logoru.Success("Moved", filePath)
		} else {
			status.Success("Moved " + filePath)
		}
	}
}
