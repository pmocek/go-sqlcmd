// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/install"
)

type Install struct {
	BaseCommand
}

func (c *Install) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "install",
		Short: "Install/Create #SQLFamily and Tools",
		Aliases: []string{"create"},
	}
	c.BaseCommand.DefineCommand()
	c.AddSubCommands(install.SubCommands)

}
