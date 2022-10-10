package repository

import (
	"fmt"

	"github.com/schufeli/kalicli/lib"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type AddCommand struct{}

func (cmd *AddCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name: "add",
		Desc: "Add repository-key and file",
	}
}

func (cmd *AddCommand) Run(fl *pflag.FlagSet) {
	add()
}

func add() {
	if !lib.CheckKeyExists() {
		lib.AddRepositoryKey()
	} else {
		fmt.Printf("%s[+]%s Skipped adding key: already exists\n", lib.ColorGreen, lib.ColorNone)
	}

	if !lib.CheckFileExists(lib.KeyFileLocation) {
		lib.ExportRepositoryKey()
	} else {
		fmt.Printf("%s[+]%s Skipped exporting key: already exists\n", lib.ColorGreen, lib.ColorNone)
	}

	if !lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.AddRepository()
	} else {
		fmt.Printf("%s[+]%s Skipping creating repository file: already exists\n", lib.ColorGreen, lib.ColorNone)
	}

	lib.UpdateRepositoryList()
}
