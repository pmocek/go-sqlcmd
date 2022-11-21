// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package edge

import (
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helper/container"
	"github.com/microsoft/go-sqlcmd/internal/helper/output"
)

type GetTags struct {
	cmd.Base
}

func (c *GetTags) DefineCommand(...cmd.Command) {
	c.Base.Options = cmd.Options{
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
	tags := container.ListTags(
		"azure-sql-edge",
		"https://mcr.microsoft.com",
	)
	output.Struct(tags)
}
