// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Mssql_Edge struct {
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

func (c *Mssql_Edge) DefineCommand() (command *Command) {
	const repo = "azure-sql-edge"

	const use = "mssql-edge"
	const short = "Install SQL Server Edge"
	const long = short
	const example = `# Install SQL Server Edge in a docker container
  sqlcmd install mssql edge`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(2),
		Run:     c.Run})

	c.AddFlags(command, repo, "edge")

	return
}
