package main

import (
	"fmt"

	"github.com/immanoj16/smart-git/pkg/client"

	"github.com/spf13/cobra"
)

func newRootCmd(client *client.Client) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "smart-git",
		Short: "A brief description of your application",
		Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
		Version: client.Version,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("here inside run ")
		},
	}

	cmd.AddCommand(
		newCommitCmd(client),
	)

	return cmd
}
