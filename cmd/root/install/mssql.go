// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/install/mssql"
)

type Mssql struct {
	BaseCommand
	MssqlBase

	tag             string
	registry        string
	repo            string
	installType     string
	acceptEula      bool
	contextName     string
	defaultDatabase string
}

func (c *Mssql) DefineCommand() {
	const repo = "mssql/server"

	c.BaseCommand.Info = CommandInfo{
		Use: "mssql",
		Short: "Install SQL Server",
		Examples: []ExampleInfo{{
				Description: "Install SQL Server in a container",
				Steps: []string{"sqlcmd install mssql"}}},
		Run: c.MssqlBase.Run,
	}

	c.BaseCommand.DefineCommand()
	c.AddSubCommands(mssql.SubCommands)
	c.AddFlags(c.AddFlag, repo, "mssql")

	return
}
