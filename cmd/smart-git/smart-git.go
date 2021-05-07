package main

import (
	"fmt"
	"os"

	"github.com/immanoj16/smart-git/pkg/client"
	"github.com/immanoj16/smart-git/pkg/git"
)

var (
	version = ""
)

func main() {
	g := git.New(".")
	c := client.New(g, version)
	rootCmd := newRootCmd(c)

	if err := rootCmd.Execute(); err != nil {
		// Execute adds all child commands to the root command and sets flags appropriately.
		// This is called by main.main(). It only needs to happen once to the rootCmd.
		fmt.Println(err)
		os.Exit(1)
	}
}
