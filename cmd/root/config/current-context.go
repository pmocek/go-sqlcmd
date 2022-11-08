// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type CurrentContext struct {
	AbstractBase
}

func (c *CurrentContext) DefineCommand() (command *Command) {
	const use = "current-context"
	const short = "Display the current-context."
	const long = short
	const example = `Display the current-context
	sqlcmd config current-context`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run})

	return
}

func (c *CurrentContext) run(cmd *Command, args []string) {
	output.Infof("%v\n", config.GetCurrentContextName())
}
