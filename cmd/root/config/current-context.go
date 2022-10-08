// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type CurrentContext struct {
	command Command
}

func (c *CurrentContext) GetCommand() *Command {
	const use = "current-context"
	const short = "Display the current-context."
	const long = short
	const example = `Display the current-context
	sqlcmd config current-context`

	var run = func(cmd *Command, args []string) {
		output.PrintString(config.GetCurrentContextName())
	}

	c.command = Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Run: run}

	return &c.command
}
