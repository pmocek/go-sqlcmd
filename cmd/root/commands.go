// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/config"
	"github.com/microsoft/go-sqlcmd/cmd/root/install"
)

var Commands = []Commander{
	&Config{AbstractBase{SubCommands: config.Commands}},
	&Query{},
	&Install{AbstractBase{SubCommands: install.Commands}},
	&Uninstall{},
}
