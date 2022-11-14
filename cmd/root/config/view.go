// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type View struct {
	BaseCommand

	raw bool
}

func (c *View) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "view",
		Short: "Display merged sqlconfig settings or a specified sqlconfig file",
		Examples: []ExampleInfo{
			{
				Description: "Show merged sqlconfig settings",
				Steps: []string{"sqlcmd config view"},
			},
			{
				Description: "Show merged sqlconfig settings and raw authentication data",
				Steps: []string{"sqlcmd config view --raw"},
			},

		},
		Aliases: []string{"use", "change-context", "set-context"},
		Run: c.run,
	}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		Name: "raw",
		Bool: &c.raw,
		Usage: "Display raw byte data",
	})
}

func (c *View) run([]string) {
	config := config.GetRedactedConfig(c.raw)
	output.Struct(config)
}
