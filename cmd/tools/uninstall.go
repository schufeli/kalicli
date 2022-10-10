package tools

import (
	"github.com/schufeli/kalicli/lib"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type UninstallCommand struct{}

func (cmd *UninstallCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "uninstall",
		Usage: "[PACKAGES]",
		Desc:  "Uninstall tools",
	}
}

func (cmd *UninstallCommand) Run(fl *pflag.FlagSet) {
	if len(fl.Args()) == 0 {
		fl.Usage()
	} else {
		uninstall(fl.Args())
	}
}

func uninstall(packages []string) {
	lib.Uninstall(packages)

	if lib.CheckFileExists(lib.RepositoryFileLocation) {
		lib.RemoveRepository()
	}
}
