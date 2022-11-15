// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import "github.com/spf13/cobra"

type Commander interface {
	DefineCommand()
	AddSubCommands([]Commander)
	Command() *cobra.Command
	Name() string
	Aliases() []string
	Execute() error
	CheckErr(error)
	IsSubCommand(command string) bool

	ArgsForUnitTesting(args []string)
}
