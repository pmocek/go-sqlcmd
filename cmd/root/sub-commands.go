// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
)

var SubCommands = []Commander{
	NewCommand[*Config](),
	NewCommand[*Query](),
	NewCommand[*Install](),
	NewCommand[*Uninstall](),
}
