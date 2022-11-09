// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql_edge

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetTags struct {
	AbstractBase
}

func (c *GetTags) DefineCommand() (command *Command) {
	const use = "get-tags"
	const short = "Get tags available for mssql-edge install."
	const long = short
	const example = `# List tags
sqlcmd install mssql-edge get-tags`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Aliases: []string{"lt"},
		Args:    MaximumNArgs(0),
		Run:     c.Run})

	return
}

func (c *GetTags) Run(*Command, []string) {
	tags := docker.ListTags(
		"azure-sql-edge",
		"https://mcr.microsoft.com",
	)
	output.Struct(tags)
}
