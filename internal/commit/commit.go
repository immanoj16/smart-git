package commit

import (
	"context"
	"fmt"
	"log"

	"github.com/immanoj16/cli-logger/logger"
	"github.com/immanoj16/smart-git/pkg/git"
	"github.com/immanoj16/smart-git/pkg/utils"
)

// Cmd is the helper for commit command
func Cmd(g *git.Git, args []string) {
	g.GetCurrentBranchName()
	if err := g.Err(); err != nil {
		log.Fatal(err)
	}

	branchName, err := utils.GetJIRATicketFromBranchName(g.BranchName)
	if err != nil {
		logger.Info(err.Error())
		branchName = g.BranchName
	}

	if len(args) > 0 {
		var commitMsg string
		if !g.Options.Header {
			if !g.Options.Comment {
				commitMsg = fmt.Sprintf("%s #comment %s", branchName, args[0])
			} else {
				commitMsg = fmt.Sprintf("%s %s", branchName, args[0])
			}
		} else {
			commitMsg = fmt.Sprintf("%s", args[0])
		}

		var gitArgs []string
		if !g.Options.Sign {
			gitArgs = []string{"commit", "-Sm", commitMsg}
		} else {
			gitArgs = []string{"commit", "-m", commitMsg}
		}

		g.ExecCmd(context.Background(), "echo", gitArgs).Run()
	} else {
		log.Fatal("Pass commit message")
	}
}
