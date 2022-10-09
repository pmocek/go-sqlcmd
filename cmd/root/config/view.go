// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type View struct {
	command Command
	raw bool
}

func (c *View) GetCommand() (*Command) {
	const short = "Display merged sqlconfig settings or a specified sqlconfig file.."
	const long = short
	const example = `# Show merged sqlconfig settings
  sqlcmd config view

  # Show merged sqlconfig settings and raw authentication data
  sqlcmd config view --raw`

	c.command = Command{
		Use:   "view",
		Short: short,
		Long: long,
		Example: example,
		Run: c.run,
	}

	c.command.PersistentFlags().BoolVar(
		&c.raw,
		"raw",
		false,
		"Display raw byte data",
	)

	return &c.command
}

func (c *View) run(cmd *Command, args []string) {
	config := config.GetRedactedConfig(c.raw)
	output.Struct(config)
}
