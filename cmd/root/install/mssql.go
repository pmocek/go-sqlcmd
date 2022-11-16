// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
)

type Mssql struct {
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

func (c *Mssql) DefineCommand(subCommands ...cmd.Command) {
	const repo = "mssql/server"

	c.Base.Info = cmd.Info{
		Use: "mssql",
		Short: "Install SQL Server",
		Examples: []cmd.ExampleInfo{{
				Description: "Install SQL Server in a container",
				Steps: []string{"sqlcmd install mssql"}}},
		Run: c.MssqlBase.Run,
	}

	c.Base.DefineCommand(subCommands...)
	c.AddFlags(c.AddFlag, repo, "mssql")

	return
}
