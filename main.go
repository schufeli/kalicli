package main

import (
	"github.com/schufeli/kalicli/cmd/repository"
	"github.com/schufeli/kalicli/cmd/tools"
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

func main() {
	cli.RunRoot(new(rootCommand))
}

type rootCommand struct{}

func (r *rootCommand) Run(fl *pflag.FlagSet) {
	fl.Usage()
}

func (r *rootCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "kalicli",
		Usage: "COMMAND",
		Desc:  "Easily install tools from the official kali repository",
	}
}

func (r *rootCommand) Subcommands() []cli.Command {
	return []cli.Command{
		new(repository.RepoCommand),
		new(tools.ToolsCommand),
	}
}
