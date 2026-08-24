package main

import (
	"context"
	"os"

	"charm.land/fang/v2"
	"github.com/adfinis/baoctx/cmd"
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
