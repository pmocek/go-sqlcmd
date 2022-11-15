// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
)

func SubCommands() []Commander {
	return []Commander{
		NewCommand[*Config](),
		NewCommand[*Query](),
		NewCommand[*Install](),
		NewCommand[*Uninstall](),
	}
}
