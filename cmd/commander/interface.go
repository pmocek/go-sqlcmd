// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import "github.com/spf13/cobra"

type Commander interface {
	DefineCommand()
	AddSubCommands(command []Commander)
	Command() *cobra.Command
	Name() string
	Aliases() []string
	Execute()
	CheckErr(err error)

	ArgsForUnitTesting(args []string)
}
