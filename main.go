package main

import (
	"context"
	"os"

	"github.com/adfinis/baoctx/cmd"
	"github.com/charmbracelet/fang"
)

var (
	// Version is the current version of bssh.
	Version = "devel"
	// Commit is the git commit hash of the current version.
	Commit = "none"
)

func main() {
	if err := fang.Execute(
		context.Background(),
		cmd.Root(),
		fang.WithCommit(Commit),
		fang.WithVersion(Version),
	); err != nil {
		os.Exit(1)
	}
}
