package utils

import (
	"fmt"
	"regexp"
)

const (
	branchFormat = "PLT-123"
)

// GetJIRATicketFromBranchName used to get the jira ticket number from branch name
func GetJIRATicketFromBranchName(branchName string) (string, error) {
	re := regexp.MustCompile(`[a-zA-Z]*-[0-9]*`)
	match := re.FindStringSubmatch(branchName)
	if len(match) > 0 {
		return match[0], nil
	}
	return "", fmt.Errorf("Branch name is not in the format of '%s'", branchFormat)
}
