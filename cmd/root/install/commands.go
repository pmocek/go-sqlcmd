// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/install/mssql"
	"github.com/microsoft/go-sqlcmd/cmd/root/install/mssql_edge"
)

var Commands = []Commander{
	&Mssql{AbstractBase: AbstractBase{SubCommands: mssql.Commands}},
	&Mssql_Edge{AbstractBase: AbstractBase{SubCommands: mssql_edge.Commands}},
}
