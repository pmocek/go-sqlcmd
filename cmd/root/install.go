// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Install struct {
	AbstractBase
}

func (c *Install) GetCommand() *Command {
	const use = "install"
	const short = "Install/Create #SQLFamliy and Tools"
	c.Command = &Command{
		Use:   use,
		Short: short,
		Long:  short,
		Example: `# Install SQL Server in a docker container
  sqlcmd install mssql

# Install SQL Server Edge in a docker container
  sqlcmd install mssql --type edge`,
		Args: ExactArgs(1),
  		Aliases: []string{"create"},
	}

	c.AddSubCommands()

	return c.Command
}
