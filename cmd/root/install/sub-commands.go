// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	"github.com/microsoft/go-sqlcmd/cmd/root/install/edge"
	"github.com/microsoft/go-sqlcmd/cmd/root/install/mssql"
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
)

var SubCommands = []cmd.Command{
	cmd.New[*Mssql](mssql.SubCommands...),
	cmd.New[*Edge](edge.SubCommands...),
}
