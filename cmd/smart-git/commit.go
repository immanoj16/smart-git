package main

import (
	"log"
	"os"

	"github.com/immanoj16/smart-git/pkg/client"

	"github.com/spf13/cobra"

	"github.com/immanoj16/smart-git/internal/commit"
)

func newCommitCmd(client *client.Client) *cobra.Command {
	git := client.Git

	commitCmd := &cobra.Command{
		Use:   "commit",
		Short: "used to commit with ticket ID and #comment flag",
		Long:  `used to commit with ticket ID and #comment flag`,
		PreRun: func(cmd *cobra.Command, args []string) {
			if _, err := os.Stat(".git"); os.IsNotExist(err) {
				log.Fatalf("Not a Git repository %s", err)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			commit.Cmd(git, args)
		},
	}

	commitCmd.Flags().BoolVar(&git.Options.Header, "header", false, "Remove header")
	commitCmd.Flags().BoolVarP(&git.Options.Comment, "comment", "c", false, "Remove comment")
	commitCmd.Flags().BoolVarP(&git.Options.Sign, "sign", "s", false, "Sign")

	return commitCmd
}
