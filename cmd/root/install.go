// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
)

type Install struct {
	cmd.Base
}

func (c *Install) DefineCommand(subCommands ...cmd.Command) {
	c.Base.Options = cmd.Options{
		Use:     "install",
		Short:   "Install/Create #SQLFamily and Tools",
		Aliases: []string{"create"},
	}
	c.Base.DefineCommand(subCommands...)
}
