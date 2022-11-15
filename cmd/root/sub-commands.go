// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"

func SubCommands() []cmd.Commander {
	return []cmd.Commander{
		cmd.New[*Config](),
		cmd.New[*Query](),
		cmd.New[*Install](),
		cmd.New[*Uninstall](),
	}
}
