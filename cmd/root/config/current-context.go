// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type CurrentContext struct {
	BaseCommand
}

func (c *CurrentContext) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "current-context",
		Short: "Display the current-context",
		Examples: []ExampleInfo{
			{
				Description: "Display the current-context",
				Steps: []string{
					"sqlcmd config current-context"},
			},
		},
		Run: c.run,
	}

	c.BaseCommand.DefineCommand()
}

func (c *CurrentContext) run(args []string) {
	output.Infof("%v\n", config.GetCurrentContextName())
}
