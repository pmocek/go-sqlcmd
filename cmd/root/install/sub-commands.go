// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"

var SubCommands = []cmd.Commander{
	cmd.New[*Mssql](),
	cmd.New[*Edge](),
}
