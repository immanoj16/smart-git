package git

import (
	"strings"

	"github.com/go-git/go-git/v5"
)

// Git is a struct
type Git struct {
	Repository *git.Repository
	Path string
	err error
	BranchName string
}

// New used to get a new instance of Git struct
func New(path string) *Git {
	return &Git{
		Path: path,
	}
}

// GetCurrentBranchName used to get currect active branch name
func (g *Git) GetCurrentBranchName() {
	g.getCurrentRepository()
	g.getCurrentBranchFromRepository()
}

func (g *Git) getCurrentRepository() {
	repo, err := git.PlainOpen(g.Path)
	if err != nil {
		g.err = err
	}
	g.Repository = repo
}

func (g *Git) getCurrentBranchFromRepository() {
	headRef, err := g.Repository.Head()
	if err != nil {
		g.err = err
	}

	refBranchName := headRef.Name().String()
	splitedBranchName := strings.Split(refBranchName, "/")
	g.BranchName = splitedBranchName[len(splitedBranchName)-1]
}

// Err used to get the git errors
func (g *Git) Err() error {
	if g.err != nil {
		return g.err
	}
	return nil
}
