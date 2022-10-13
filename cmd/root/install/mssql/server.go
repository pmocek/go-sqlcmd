// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	. "github.com/spf13/cobra"
)

type Server struct {
	Base
}

func (c *Server) GetCommand() *Command {
	const repo = "mssql/server"

	const use = "server"
	const short = "Install SQL Server"
	const long = short
	const example = `# Install SQL Server in a docker container
  sqlcmd install mssql server`

	c.addFlags(use, short, long, example, repo ,"mssql")
	c.AddSubCommands()

	return c.Command
}
