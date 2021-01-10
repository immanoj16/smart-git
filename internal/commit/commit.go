package commit

import (
	"fmt"
	"log"

	"github.com/immanoj16/smart-git/pkg/git"
	"github.com/immanoj16/smart-git/pkg/utils"
)

// Cmd is the helper for commit command 
func Cmd() {
	g := git.New(".")

	g.GetCurrentBranchName()
	if err := g.Err(); err != nil {
		log.Fatal(err)
	}
	

	newBranchName, err := utils.GetJIRATicketFromBranchName(g.BranchName)
	if err != nil {
		fmt.Println(err)
		newBranchName = g.BranchName
	}

	fmt.Println(newBranchName)
}
