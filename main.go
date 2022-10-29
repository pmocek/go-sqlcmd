// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package main

import (
	"github.com/microsoft/go-sqlcmd/cmd"
	legacyMain "github.com/microsoft/go-sqlcmd/cmd/sqlcmd"
	"os"
)

func main() {
	if len(os.Args) > 1 &&
		(cmd.IsValidRootCommand(os.Args[1]) || os.Args[1] == "--help") {
		cmd.ExecuteCommandLine()
	} else {
		legacyMain.BackCompatMode()
	}
}
