// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetTags struct {
	AbstractBase

	installType string
}

func (c *GetTags) DefineCommand() (command *Command) {
	const use = "get-tags"
	const short = "Get tags available for install."
	const long = short
	const example = `# List tags
sqlcmd install get-tags`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Aliases: []string{"lt"},
		Args:    MaximumNArgs(0),
		Run:     c.Run})

	command.PersistentFlags().StringVarP(
		&c.installType,
		"type",
		"t",
		"server",
		"server = SQL Server, edge = SQL Server Edge",
	)

	return
}

func (c *GetTags) Run(*Command, []string) {
	var path string

	switch c.installType {
	case "server":
		path = "mssql/server"
	case "edge":
		path = "azure-sql-edge"
	default:
		output.Fatal("Unrecognized type, please specify 'server' or 'edge'")
	}
	tags := docker.ListTags(path)
	output.Struct(tags)
}
