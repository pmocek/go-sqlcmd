// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package main

import (
	"github.com/microsoft/go-sqlcmd/cmd"
	legacyMain "github.com/microsoft/go-sqlcmd/cmd/sqlcmd"
	"os"
)

func main() {
	if isModernCliEnvVarEnabled() &&
		isAnArgProvided() &&
		isFirstArgAModernCliRootCommand() {
		cmd.RunCommandLine()
	} else {
		legacyMain.BackCompatMode()
	}
}

func isModernCliEnvVarEnabled() (modernCliEnabled bool)  {
	if os.Getenv("SQLCMD_MODERN") != "" {
		modernCliEnabled = true
	}

	return
}

func isFirstArgAModernCliRootCommand() (isNewCliCommand bool) {
	if cmd.IsValidRootCommand(os.Args[1]) {
		isNewCliCommand = true
	} else if os.Args[1] == "--help" {
		isNewCliCommand = true
	} else if os.Args[1] == "completion" {
		isNewCliCommand = true
	}

	return
}

func isAnArgProvided() (isMoreThanOneArg bool) {
	isMoreThanOneArg = len(os.Args) > 1

	return
}
