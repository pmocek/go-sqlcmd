// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/container"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type GetTags struct {
	cmd.Base
}

func (c *GetTags) DefineCommand(...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use: "get-tags",
		Short: "Get tags available for mssql install",
		Examples: []cmd.ExampleInfo{
			{
				Description: "List tags",
				Steps: []string{"sqlcmd install mssql get-tags"},
			},
		},
		Aliases: []string{"gt", "lt"},
		Run: c.run,
	}

	c.Base.DefineCommand()

}

func (c *GetTags) run() {
	tags := container.ListTags(
		"mssql/server",
		"https://mcr.microsoft.com",
	)
	output.Struct(tags)
}
