package main

import (
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type root struct{}

func (r *root) Run(fl *pflag.FlagSet) {
	fl.Usage()
}

func (r *root) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "kalicli",
		Usage: "[subcommand]",
		Desc:  "Install tools from the kali repository without destroying your distro.",
	}
}

func (r *root) Subcommands() []cli.Command {
	return []cli.Command{
		new(add),
		new(cleanup),
		new(install),
		new(remove),
		new(search),
	}
}
