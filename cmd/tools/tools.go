package tools

import (
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type ToolsCommand struct{}

func (r *ToolsCommand) Run(fl *pflag.FlagSet) {
	fl.Usage()
}

func (r *ToolsCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "tools",
		Usage: "COMMAND",
		Desc:  "Manage kali tools and packages",
	}
}

func (r *ToolsCommand) Subcommands() []cli.Command {
	return []cli.Command{
		new(InstallCommand),
	}
}
