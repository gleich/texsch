package commit

import (
	"fmt"
	"strings"
	"time"

	"github.com/Matt-Gleich/statuser/v2"
	"github.com/Matt-Gleich/texsch/pkg/configuration"
	"github.com/Matt-Gleich/texsch/pkg/utils"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

// Commit the changes
func CommitChanges(changes git.Status, workingTree *git.Worktree) {
	commitDocuments(changes, workingTree)
}

// Commit all LaTeX and PDF files
func commitDocuments(changes git.Status, workingTree *git.Worktree) {
	// Sorting files
	latexChanges, PDFChanges, _ := sortChanges(changes, true)

	// Commiting and staging changes
	var (
		authorFullName = configuration.GetGeneral().Full_Name
		commitConfig   = configuration.GetCommitConfig()
		committed      int
	)
	for latexPath := range latexChanges {
		var (
			pathChunks = strings.Split(latexPath, "/")
			docName    = strings.TrimSuffix(pathChunks[len(pathChunks)-1], ".tex")
			docType    = pathChunks[len(pathChunks)-2]
			changeMsg  string
			PDFPath    = strings.Replace(
				strings.TrimSuffix(latexPath, ".tex")+".pdf",
				"LaTeX",
				"PDF",
				1,
			)
		)
		if utils.ContainsString(PDFPath, utils.MapKeys(PDFChanges)) {
			// Staging changes
			stageFile(PDFPath, workingTree)
			stageFile(latexPath, workingTree)

			// Getting new status for staged files
			var (
				newChanges, _         = GetChanges()
				newLatexChanges, _, _ = sortChanges(newChanges, false)
				newLatexChangeStatus  = newLatexChanges[latexPath]
			)
			switch newLatexChangeStatus {
			case " ":
				break
			case "?":
				statuser.ErrorMsg("Failed to stage changes for "+latexPath, 1)
			case "M":
				changeMsg = "update"
			case "A":
				changeMsg = "create"
			case "D":
				changeMsg = "remove"
			case "R":
				changeMsg = "rename"
			case "C":
				changeMsg = "copy"
			case "U":
				changeMsg = "update"
			default:
				statuser.ErrorMsg("Unrecognized change status "+newLatexChangeStatus, 1)
			}

			// Committing changes
			_, err := workingTree.Commit(
				fmt.Sprintf(
					"%v[%v] %v",
					changeMsg,
					docType,
					strings.ReplaceAll(docName, "-", " "),
				),
				&git.CommitOptions{
					Author: &object.Signature{
						Name:  authorFullName,
						Email: commitConfig.Email,
						When:  time.Now(),
					},
				},
			)
			if err != nil {
				statuser.Error("Failed to commit changes for "+docName, err, 1)
			}
			statuser.Success("Committed " + docName)
			committed++
		}
	}
	statuser.Success("Committed " + fmt.Sprint(committed) + " documents")
}

// Sort the changes into latex, pdf, and other files
func sortChanges(
	changes git.Status,
	workTree bool,
) (map[string]string, map[string]string, map[string]string) {
	var (
		latexFiles = map[string]string{}
		PDFFiles   = map[string]string{}
		otherFiles = map[string]string{}
	)
	for filePath, rawStatus := range changes {
		status := string(rawStatus.Staging)
		if workTree {
			status = string(rawStatus.Worktree)
		}
		if strings.HasSuffix(filePath, ".tex") {
			latexFiles[filePath] = status
		} else if strings.HasSuffix(filePath, ".pdf") {
			PDFFiles[filePath] = status
		} else {
			otherFiles[filePath] = status
		}
	}
	return latexFiles, PDFFiles, otherFiles
}

// Stage a file
func stageFile(filePath string, workingTree *git.Worktree) {
	_, err := workingTree.Add(filePath)
	if err != nil {
		statuser.Error("Failed to stage "+filePath, err, 1)
	}
}
