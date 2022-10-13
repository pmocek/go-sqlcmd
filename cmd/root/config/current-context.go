// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type CurrentContext struct {
	AbstractBase
}

func (c *CurrentContext) GetCommand() (command *Command) {
	const use = "current-context"
	const short = "Display the current-context."
	const long = short
	const example = `Display the current-context
	sqlcmd config current-context`

	var run = func(cmd *Command, args []string) {
		output.Info(config.GetCurrentContextName())
	}

	command = c.AddCommand(Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Run: run})

	return
}
