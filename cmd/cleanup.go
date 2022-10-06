package main

import (
	"fmt"

	"github.com/schufeli/kalicli/internal"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type cleanup struct{}

func (cmd *cleanup) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "cleanup",
		Aliases: []string{"c"},
		Desc:    "Remove repositories, files and keys created by this tool",
	}
}

func (cmd *cleanup) Run(fl *pflag.FlagSet) {
	Cleanup()
}

func Cleanup() {
	if internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.RemoveRepository()
	} else {
		fmt.Printf("%s[-]%s Skipping removing repository file: file not found\n", internal.ColorGreen, internal.ColorNone)
	}

	if internal.CheckKeyExists() {
		internal.RemoveRepositoryKey()
	} else {
		fmt.Printf("%s[-]%s Skipping removing repository key: key not found\n", internal.ColorGreen, internal.ColorNone)
	}

	if internal.CheckFileExists(internal.KeyFileLocation) {
		internal.RemoveRepositoryKeyFile()
	} else {
		fmt.Printf("%s[-]%s Skipping removing repository key file: file not found\n", internal.ColorGreen, internal.ColorNone)
	}
}
