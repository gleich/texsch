package commit

import (
	"github.com/Matt-Gleich/statuser"
	"github.com/go-git/go-git/v5"
)

func GetChanges() git.Status {
	repo, err := git.PlainOpen(".")
	if err != nil {
		statuser.Error("Failed to read from git repo", err, 1)
	}
	workingTree, err := repo.Worktree()
	if err != nil {
		statuser.Error("Failed to get changes for git repo", err, 1)
	}
	status, err := workingTree.Status()
	if err != nil {
		statuser.Error("Failed to get status of changes in git repo", err, 1)
	}
	return status
}
