package git

import (
	"context"
	"os"
	"os/exec"
	"strings"

	"github.com/go-git/go-git/v5"
)

// Git is a struct
type Git struct {
	Repository *git.Repository
	Dir        string
	err        error
	BranchName string
}

// New used to get a new instance of Git struct
func New(dir string) *Git {
	return &Git{
		Dir: dir,
	}
}

// GetCurrentBranchName used to get current active branch name
func (g *Git) GetCurrentBranchName() {
	g.getCurrentRepository()
	g.getCurrentBranchFromRepository()
}

func (g *Git) getCurrentRepository() {
	repo, err := git.PlainOpen(g.Dir)
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

func (g *Git) ExecCmd(ctx context.Context, name string, args []string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Dir = g.Dir
	return cmd
}

// Err used to get the git errors
func (g *Git) Err() error {
	if g.err != nil {
		return g.err
	}
	return nil
}
