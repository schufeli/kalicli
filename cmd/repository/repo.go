package repository

import (
	"github.com/spf13/pflag"
	"go.coder.com/cli"
)

type RepoCommand struct{}

func (r *RepoCommand) Run(fl *pflag.FlagSet) {
	fl.Usage()
}

func (r *RepoCommand) Spec() cli.CommandSpec {
	return cli.CommandSpec{
		Name:  "repo",
		Usage: "COMMAND",
		Desc:  "Manage kali repository (key, file)",
	}
}

func (r *RepoCommand) Subcommands() []cli.Command {
	return []cli.Command{
		new(AddCommand),
		new(RemoveCommand),
	}
}
