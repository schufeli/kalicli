package repository

import (
	"fmt"

	"github.com/schufeli/kalicli/lib"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type RemoveCommand struct{}

func (cmd *RemoveCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name: "remove",
		Desc: "Remove repository-key and file",
	}
}

func (cmd *RemoveCommand) Run(fl *pflag.FlagSet) {
	if len(fl.Args()) == 0 {
		fl.Usage()
	} else {
		remove(fl.Args())
	}
}

func remove(packages []string) {
	if lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.RemoveRepository()
	} else {
		fmt.Printf("%s[-]%s Skipped removing repository file: not found\n", lib.ColorGreen, lib.ColorNone)
	}

	if lib.CheckKeyExists() {
		lib.RemoveRepositoryKey()
	} else {
		fmt.Printf("%s[-]%s Skipped removing repository key: not found\n", lib.ColorGreen, lib.ColorNone)
	}

	if lib.CheckFileExists(lib.KeyFileLocation) {
		lib.RemoveRepositoryKeyFile()
	} else {
		fmt.Printf("%s[-]%s Skipped removing repository key file: not found\n", lib.ColorGreen, lib.ColorNone)
	}
}
