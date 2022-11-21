// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/cmd/root/config"
	"github.com/microsoft/go-sqlcmd/cmd/root/install"
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
)

func SubCommands() []cmd.Command {
	return []cmd.Command{
		cmd.New[*Config](config.SubCommands()...),
		cmd.New[*Query](),
		cmd.New[*Install](install.SubCommands...),
		cmd.New[*Uninstall](),
	}
}
