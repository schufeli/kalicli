package tools

import (
	"github.com/schufeli/kalicli/lib"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type InstallCommand struct{}

func (cmd *InstallCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "install",
		Usage: "[PACKAGES]",
		Desc:  "Install tools",
	}
}

func (cmd *InstallCommand) Run(fl *pflag.FlagSet) {
	if len(fl.Args()) == 0 {
		fl.Usage()
	} else {
		install(fl.Args())
	}
}

func install(packages []string) {
	if !lib.CheckFileExists(lib.KeyFileLocation) || !lib.CheckKeyExists() {
		lib.AddRepositoryKey()
		lib.ExportRepositoryKey()
	}

	if !lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.AddRepository()
		lib.UpdateRepositoryList()
	}

	lib.Install(packages)

	if lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.RemoveRepository()
	}
}
