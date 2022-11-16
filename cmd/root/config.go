// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
)

type Config struct {
	cmd.Base
}

func (c *Config) DefineCommand(subCommands ...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use: "config",
		Short: `Modify sqlconfig files using subcommands like "sqlcmd config use-context mssql"`,
	}
	c.Base.DefineCommand(subCommands...)
}
