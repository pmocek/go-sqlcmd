// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/install/mssql"
)

var Commands = []Commander{
	&Mssql{AbstractBase: AbstractBase{SubCommands: mssql.Commands}},
}
