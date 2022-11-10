// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root/config"
	"github.com/microsoft/go-sqlcmd/cmd/root/install"
)

var Commands = []Commander{
	&Config{BaseCommand{SubCommands: config.Commands}},
	&Query{},
	&Install{BaseCommand{SubCommands: install.Commands}},
	&Uninstall{},
}
