// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/config"
	. "github.com/spf13/cobra"
)

type Config struct {
	commander.Commander
}

func (c *Config) GetCommand() (command *Command) {
	const short = "Modify sqlconfig files using subcommands like \"sqlcmd config use-context sa@sql1\""

	command = &Command{
		Use:   "config",
		Short: short,
		Long: short,
		Args: ExactArgs(1),
	}

	for _, subCommand := range config.Commands {
		command.AddCommand(subCommand.GetCommand())
	}

	return
}
