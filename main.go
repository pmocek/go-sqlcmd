// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package main

import (
	"github.com/microsoft/go-sqlcmd/cmd"
	legacyMain "github.com/microsoft/go-sqlcmd/cmd/sqlcmd"
	"os"
)

func main() {
	if isModernCliEnabled() && isFirstArgValidSubCommand() {
		cmd.RunCommandLine()
	} else {
		legacyMain.BackCompatMode()
	}
}

func isModernCliEnabled() (modernCliEnabled bool) {
	if os.Getenv("SQLCMD_MODERN") != "" {
		modernCliEnabled = true
	}
	return
}

func isFirstArgValidSubCommand() (isNewCliCommand bool) {
	if len(os.Args) > 0 {
		if cmd.IsValidSubCommand(os.Args[1]) {
			isNewCliCommand = true
		}
	}
	return
}
