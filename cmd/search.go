package main

import (
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type search struct{}

func (cmd *search) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "search",
		Aliases: []string{"s"},
		Desc:    "Search tool in kali repository",
	}
}

func (cmd *search) Run(fl *pflag.FlagSet) {
	Search()
}

func Search() {

}
