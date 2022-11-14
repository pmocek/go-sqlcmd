// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
)

var SubCommands = []Commander{
	NewCommand[*GetTags](),
}
