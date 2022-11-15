// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type GetTags struct {
	BaseCommand
}

func (c *GetTags) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "get-tags",
		Short: "Get tags available for mssql install",
		Examples: []ExampleInfo{
			{
				Description: "List tags",
				Steps: []string{"sqlcmd install mssql get-tags"},
			},
		},
		Aliases: []string{"gt", "lt"},
		Run: c.run,
	}

	c.BaseCommand.DefineCommand()

}

func (c *GetTags) run() {
	tags := docker.ListTags(
		"mssql/server",
		"https://mcr.microsoft.com",
	)
	output.Struct(tags)
}
