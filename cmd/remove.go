package main

import (
	"github.com/schufeli/kalicli/internal"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type remove struct{}

func (cmd *remove) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "remove",
		Usage:   "[package names]",
		Aliases: []string{"r"},
		Desc:    "Uninstall tools",
	}
}

func (cmd *remove) Run(fl *pflag.FlagSet) {
	if len(fl.Args()) == 0 {
		fl.Usage()
	} else {
		Remove(fl.Args())
	}
}

func Remove(packages []string) {

	internal.Remove(packages)

	if internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.RemoveRepository()
	}
}
