// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Mssql struct {
	AbstractBase
	MssqlBase

	tag             string
	registry        string
	repo            string
	installType     string
	acceptEula      bool
	contextName     string
	defaultDatabase string
}

func (c *Mssql) DefineCommand() (command *Command) {
	const repo = "mssql/server"

	const use = "mssql"
	const short = "Install SQL Server"
	const long = short
	const example = `# Install SQL Server in a docker container
  sqlcmd install mssql server`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(2),
		Run:     c.Run})

	c.AddFlags(command, repo, "mssql")

	return
}
