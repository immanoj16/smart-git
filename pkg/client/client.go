package client

import (
	"fmt"
	"runtime/debug"

	"github.com/immanoj16/smart-git/pkg/git"
)

type Client struct {
	Git     *git.Git
	Version string
}

func New(git *git.Git, version string) *Client {
	return &Client{git, version}
}

func (c *Client) GetVersion() string {
	if c.Version != "" {
		return c.Version
	}

	info, ok := debug.ReadBuildInfo()
	if !ok || info.Main.Version == "" {
		return "unknown"
	}

	version := info.Main.Version
	if info.Main.Sum != "" {
		version += fmt.Sprintf(" (%s)", info.Main.Sum)
	}

	return version
}
