// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/cmd/root/install"
	. "github.com/spf13/cobra"
)

type Install struct {
}

func (c *Install) GetCommand() (command *Command) {
	const use = "install"
	const short = "Install the #SQLFamliy and Tools"
	command = &Command{
		Use:   use,
		Short: short,
		Long:  short,
		Aliases: []string{"create"},
	}

	// TODO: Push into base class
	for _, subCommand := range install.Commands {
		command.AddCommand(subCommand.GetCommand())
	}

	return
}
