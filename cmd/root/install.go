// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Install struct {
	BaseCommand
}

func (c *Install) DefineCommand() (command *Command) {
	const short = "Install/Create #SQLFamliy and Tools"

	command = c.SetCommand(Command{
		Use:     "install",
		Short:   short,
		Long:    short,
		Args:    ExactArgs(1),
		Aliases: []string{"create"},
		Example: `# Install SQL Server in a docker container
  sqlcmd install mssql

# Install SQL Server Edge in a docker container
  sqlcmd install mssql --type edge`})

	return
}
