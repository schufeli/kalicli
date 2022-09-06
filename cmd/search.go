package main

import (
	"log"

	"github.com/schufeli/kalicli/internal"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type search struct{}

func (cmd *search) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "search",
		Usage:   "[package]",
		Aliases: []string{"s"},
		Desc:    "Search tool in kali repository",
	}
}

func (cmd *search) Run(fl *pflag.FlagSet) {
	if len(fl.Arg(0)) == 0 {
		log.Fatalf("%s[x]%s Provide a package name to search for\n", internal.ColorRed, internal.ColorNone)
	} else {
		Run(fl.Arg(0))
	}
}

func Run(packageName string) {
	if !internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.AddRepository()
	}

	internal.Search(packageName)

	if internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.RemoveRepository()
	}
}
