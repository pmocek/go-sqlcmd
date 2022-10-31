// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Config struct {
	AbstractBase
}

func (c *Config) DefineCommand() (command *Command) {
	const short = "Modify sqlconfig files using subcommands like \"sqlcmd config use-context mssql\""

	command = c.SetCommand(Command{
		Use:   "config",
		Short: short,
		Long:  short,
		Args:  ExactArgs(1),
	})

	return
}
