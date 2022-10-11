// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Root struct {
	AbstractBase
}

func (c *Root) GetCommand() *Command {
	const short = "sqlcmd: a command-line interface for the #SQLFamily."

	c.Command = &Command{
		Use:   appName,
		Short: short,
		Long: short,
	}

	c.AddSubCommands()

	return c.Command
}
