package git

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

// GetCurrentBranchName used to get currect active branch name
func GetCurrentBranchName() (string, error) {
	repo, err := getCurrentRepository()
	if err != nil {
		return "", err
	}

	branchName, err := getCurrentBranchFromRepository(repo)
	if err != nil {
		return "", err
	}
	
	return branchName, nil
}

func getCurrentRepository() (*git.Repository, error) {
	repo, err := git.PlainOpen(".")
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func getCurrentBranchFromRepository(repo *git.Repository) (string, error) {
	headRef, err := repo.Head()
	if err != nil {
		return "", err
	}

	refBranchName := headRef.Name().String()
	splitedBranchName := strings.Split(refBranchName, "/")
	return splitedBranchName[len(splitedBranchName)-1], nil
}
