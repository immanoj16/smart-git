package commit

import (
	"fmt"
	"log"

	"github.com/immanoj16/cli-logger/logger"
	"github.com/immanoj16/smart-git/pkg/git"
	"github.com/immanoj16/smart-git/pkg/utils"
)

// Cmd is the helper for commit command
func Cmd(args []string) {
	g := git.New(".")

	g.GetCurrentBranchName()
	if err := g.Err(); err != nil {
		log.Fatal(err)
	}

	branchName, err := utils.GetJIRATicketFromBranchName(g.BranchName)
	if err != nil {
		logger.Info(err.Error())
		branchName = g.BranchName
	}

	logger.Info(branchName)
	fmt.Println(args)
}
