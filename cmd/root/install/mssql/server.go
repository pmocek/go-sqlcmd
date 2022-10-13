// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	. "github.com/spf13/cobra"
)

type Server struct {
	Base
}

func (c *Server) GetCommand() (command *Command) {
	const repo = "mssql/server"

	const use = "server"
	const short = "Install SQL Server"
	const long = short
	const example = `# Install SQL Server in a docker container
  sqlcmd install mssql server`

	command = c.AddCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(2),
		Run:     c.run})

	c.addFlags(command, repo, "mssql")

	return
}
