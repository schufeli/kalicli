package main

import (
	"github.com/schufeli/kalicli/internal"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type search struct{}

func (cmd *search) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "search",
		Usage:   "[package name]",
		Aliases: []string{"s"},
		Desc:    "Search if tool exists in repository",
	}
}

func (cmd *search) Run(fl *pflag.FlagSet) {
	if len(fl.Arg(0)) == 0 {
		fl.Usage()
	} else {
		Run(fl.Arg(0))
	}
}

func Run(packageName string) {
	if !internal.CheckFileExists(internal.KeyFileLocation) || !internal.CheckKeyExists() {
		internal.AddRepositoryKey()
		internal.ExportRepositoryKey()
	}

	if !internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.AddRepository()
		internal.UpdateRepositoryList()
	}

	internal.Search(packageName)

	if internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.RemoveRepository()
	}
}
