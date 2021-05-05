package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/immanoj16/smart-git/internal/commit"
	"github.com/immanoj16/smart-git/pkg/git"
)

func newCommitCmd() *cobra.Command {
	g := git.New(".")

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
			commit.Cmd(g, args)
		},
	}

	commitCmd.Flags().BoolVar(&g.Options.Header, "header", false, "Remove header")
	commitCmd.Flags().BoolVarP(&g.Options.Comment, "comment", "c", false, "Remove comment")
	commitCmd.Flags().BoolVarP(&g.Options.Sign, "sign", "s", false, "Sign")

	return commitCmd
}
