// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	"github.com/microsoft/go-sqlcmd/cmd/root/install/edge"
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
)

type Edge struct {
	cmd.Base
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

	c.Base.Info = cmd.Info{
		Use: "mssql-edge",
		Short: "Install SQL Server Edge",
		Examples: []cmd.ExampleInfo{{
				Description: "Install SQL Server Edge in a container",
				Steps: []string{"sqlcmd install mssql-edge"}}},
		Run: c.MssqlBase.Run,
	}

	c.Base.DefineCommand()
	c.AddSubCommands(edge.SubCommands)
	c.AddFlags(c.AddFlag, repo, "edge")
}
