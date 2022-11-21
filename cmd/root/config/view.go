// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helper/config"
	"github.com/microsoft/go-sqlcmd/internal/helper/output"
)

type View struct {
	cmd.Base

	raw bool
}

func (c *View) DefineCommand(...cmd.Command) {
	c.Base.Options = cmd.Options{
		Use:   "view",
		Short: "Display merged sqlconfig settings or a specified sqlconfig file",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Show merged sqlconfig settings",
				Steps:       []string{"sqlcmd config view"},
			},
			{
				Description: "Show merged sqlconfig settings and raw authentication data",
				Steps:       []string{"sqlcmd config view --raw"},
			},
		},
		Aliases: []string{"use", "change-context", "set-context"},
		Run:     c.run,
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagOptions{
		Name:  "raw",
		Bool:  &c.raw,
		Usage: "Display raw byte data",
	})
}

func (c *View) run() {
	contents := config.GetRedactedConfig(c.raw)
	output.Struct(contents)
}
