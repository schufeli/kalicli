package tools

import (
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
	install()
}

func install() {

}
