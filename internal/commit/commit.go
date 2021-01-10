package commit

import (
	"fmt"
	"log"

	"github.com/immanoj16/smart-git/pkg/git"
	"github.com/immanoj16/smart-git/pkg/utils"
)

// Cmd is the helper for commit command 
func Cmd() {
	branchName, err := git.GetCurrentBranchName()
	if err != nil {
		log.Fatal(err)
	}

	newBranchName, err := utils.GetJIRATicketFromBranchName(branchName)
	if err != nil {
		fmt.Println(err)
		newBranchName = branchName
	}

	fmt.Println(newBranchName)
}
