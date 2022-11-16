// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type CurrentContext struct {
	cmd.Base
}

func (c *CurrentContext) DefineCommand(...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use: "current-context",
		Short: "Display the current-context",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Display the current-context",
				Steps: []string{
					"sqlcmd config current-context"},
			},
		},
		Run: c.run,
	}

	c.Base.DefineCommand()
}

func (c *CurrentContext) run() {
	output.Infof("%v\n", config.GetCurrentContextName())
}
