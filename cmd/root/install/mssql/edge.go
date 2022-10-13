// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	. "github.com/spf13/cobra"
)

type Edge struct {
	Base
}

func (c *Edge) GetCommand() (command *Command) {
	const repo = "azure-sql-edge"

	const use = "edge"
	const short = "Install SQL Server Edge"
	const long = short
	const example = `# Install SQL Server Edge in a docker container
  sqlcmd install mssql edge`

	command = c.AddCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(2),
		Run:     c.run})

	c.addFlags(command, repo, "edge")

	return
}
