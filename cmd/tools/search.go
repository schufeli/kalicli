package tools

import (
	"github.com/schufeli/kalicli/lib"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type SearchCommand struct{}

func (cmd *SearchCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "search",
		Usage: "[PACKAGE]",
		Desc:  "Search tool in repository",
	}
}

func (cmd *SearchCommand) Run(fl *pflag.FlagSet) {
	if len(fl.Arg(0)) == 0 {
		fl.Usage()
	} else {
		search(fl.Arg(0))
	}
}

func search(packagename string) {
	if !lib.CheckFileExists(lib.KeyFileLocation) || !lib.CheckKeyExists() {
		lib.AddRepositoryKey()
		lib.ExportRepositoryKey()
	}

	if !lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.AddRepository()
		lib.UpdateRepositoryList()
	}

	lib.Search(packagename)

	if lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.RemoveRepository()
	}
}
