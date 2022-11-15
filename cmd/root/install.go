// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/cmd/root/install"
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
)

type Install struct {
	cmd.Base
}

func (c *Install) DefineCommand() {
	c.Base.Info = cmd.Info{
		Use: "install",
		Short: "Install/Create #SQLFamily and Tools",
		Aliases: []string{"create"},
	}
	c.Base.DefineCommand()
	c.AddSubCommands(install.SubCommands)

}
