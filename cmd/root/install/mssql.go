// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Mssql struct {
	AbstractBase

	tag             string
	registry        string
	repo            string
	installType     string
	acceptEula      bool
	contextName     string
	defaultDatabase string
}

func (c *Mssql) DefineCommand() (command *Command) {
	const use = "mssql"
	const short = "Install/Create Sql Server"

	command = c.SetCommand(Command{
		Use:   use,
		Short: short,
		Long:  short,
		Example: `# Install SQL Server in a local container
  sqlcmd install mssql`,
	})

	return
}
