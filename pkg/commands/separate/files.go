package separate

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/Matt-Gleich/logoru"
	"github.com/Matt-Gleich/statuser/v2"
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

		const (
			latexPrefix = "LaTeX/"
			pdfPrefix   = "PDF/"
		)
		if strings.HasSuffix(filePath, ".tex") {
			if strings.HasPrefix(filePath, latexPrefix) {
				continue
			}
			prefix = latexPrefix
		} else if strings.HasSuffix(filePath, ".pdf") {
			if strings.HasPrefix(filePath, pdfPrefix) {
				continue
			}
			prefix = pdfPrefix
		}

		cleanedFilePath := filePath
		cleanedFilePath = strings.TrimPrefix(cleanedFilePath, latexPrefix)
		cleanedFilePath = strings.TrimPrefix(cleanedFilePath, pdfPrefix)

		pathChunks := strings.Split(cleanedFilePath, "/")
		fileFolderPath := prefix + strings.Join(pathChunks[:len(pathChunks)-1], "/") + "/"
		fileName := pathChunks[len(pathChunks)-1]
		outputPath := strings.Join(pathChunks[2:], "/")

		_, err := os.Stat(fileFolderPath)
		if os.IsNotExist(err) {
			err := os.MkdirAll(fileFolderPath, 0700)
			if err != nil {
				statuser.Error("Failed to make folders", err, 1)
			}
		}

		if loop {
			logoru.Info("Detected", outputPath)
		}

		err = os.Rename(filePath, fileFolderPath+fileName)
		if err != nil {
			statuser.Error("Failed to move file", err, 1)
		}

		if loop {
			logoru.Success("Moved", outputPath)
		} else {
			statuser.Success("Moved " + outputPath)
		}
	}
}
