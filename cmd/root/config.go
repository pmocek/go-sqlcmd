// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/config"
)

type Config struct {
	BaseCommand
}

func (c *Config) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "config",
		Short: `Modify sqlconfig files using subcommands like "sqlcmd config use-context mssql"`,
	}
	c.BaseCommand.DefineCommand()
	c.AddSubCommands(config.SubCommands())
}
