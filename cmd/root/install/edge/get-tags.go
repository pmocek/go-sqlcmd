// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package edge

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type GetTags struct {
	cmd.Base
}

func (c *GetTags) DefineCommand(subCommands ...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use:   "get-tags",
		Short: "Get tags available for mssql edge install",
		Examples: []cmd.ExampleInfo{
			{
				Description: "List tags",
				Steps:       []string{"sqlcmd install mssql-edge get-tags"},
			},
		},
		Aliases: []string{"gt", "lt"},
		Run:     c.run,
	}

	c.Base.DefineCommand()
}

func (c *GetTags) run() {
	tags := docker.ListTags(
		"azure-sql-edge",
		"https://mcr.microsoft.com",
	)
	output.Struct(tags)
}
