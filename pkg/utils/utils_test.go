package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetJIRATicketFromBranchName(t *testing.T) {
	branchName := "PLT-123"
	matched, err := GetJIRATicketFromBranchName(branchName)
	require.NoError(t, err)
	require.Equal(t, matched, branchName)

	branchName = "branch"
	matched, err = GetJIRATicketFromBranchName(branchName)
	require.Error(t, err)
	require.Equal(t, matched, "")
}
