// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type View struct {
	AbstractBase

	raw bool
}

func (c *View) DefineCommand() (command *Command) {
	const short = "Display merged sqlconfig settings or a specified sqlconfig file.."
	const long = short
	const example = `# Show merged sqlconfig settings
  sqlcmd config view

  # Show merged sqlconfig settings and raw authentication data
  sqlcmd config view --raw`

	command = c.SetCommand(Command{
		Use:     "view",
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run,
	})

	command.PersistentFlags().BoolVar(
		&c.raw,
		"raw",
		false,
		"Display raw byte data",
	)

	return
}

func (c *View) run(*Command, []string) {
	config := config.GetRedactedConfig(c.raw)
	output.Struct(config)
}
