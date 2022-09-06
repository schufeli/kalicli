package main

import (
	"fmt"

	"github.com/schufeli/kalicli/internal"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type add struct{}

func (cmd *add) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:    "add",
		Aliases: []string{"a"},
		Desc:    "Add kali repository and key",
	}
}

func (cmd *add) Run(fl *pflag.FlagSet) {
	Add()
}

func Add() {
	if !internal.CheckKeyExists() {
		internal.AddRepositoryKey()
	} else {
		fmt.Printf("%s[+]%s Skipping adding key: already exists\n", internal.ColorGreen, internal.ColorNone)
	}

	if !internal.CheckFileExists(internal.KeyFileLocation) {
		internal.ExportRepositoryKey()
	} else {
		fmt.Printf("%s[+]%s Skipping exporting key: already exists\n", internal.ColorGreen, internal.ColorNone)
	}

	if !internal.CheckFileExists(internal.RepositoryFileLocation) {
		internal.AddRepository()
	} else {
		fmt.Printf("%s[+]%s Skipping creating repository file: already exists\n", internal.ColorGreen, internal.ColorNone)
	}

	internal.UpdateRepositoryList()
}
