// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
)

var SubCommands = []Commander{
	NewCommand[*Mssql](),
	NewCommand[*Edge](),
}
