// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/install/edge"
)

type Edge struct {
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

func (c *Edge) DefineCommand() {
	const repo = "azure-sql-edge"

	c.BaseCommand.Info = CommandInfo{
		Use: "mssql-edge",
		Short: "Install SQL Server Edge",
		Examples: []ExampleInfo{{
				Description: "Install SQL Server Edge in a container",
				Steps: []string{"sqlcmd install mssql-edge"}}},
		Run: c.MssqlBase.Run,
	}

	c.BaseCommand.DefineCommand()
	c.AddSubCommands(edge.SubCommands)
	c.AddFlags(c.AddFlag, repo, "edge")
}
