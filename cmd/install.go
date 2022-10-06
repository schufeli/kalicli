package main

import (
	"github.com/schufeli/kalicli/internal"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type install struct{}

func (cmd *install) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "install",
		Usage:   "[package names]",
		Aliases: []string{"i"},
		Desc:    "Install tools",
	}
}

func (cmd *install) Run(fl *pflag.FlagSet) {
	if len(fl.Args()) == 0 {
		fl.Usage()
	} else {
		Install(fl.Args())
	}
}

func Install(packages []string) {
	if !internal.CheckFileExists(internal.KeyFileLocation) || !internal.CheckKeyExists() {
		internal.AddRepositoryKey()
		internal.ExportRepositoryKey()
	}

	if !internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.AddRepository()
		internal.UpdateRepositoryList()
	}

	internal.Install(packages)

	if internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.RemoveRepository()
	}
}
