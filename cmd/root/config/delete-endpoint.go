// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type DeleteEndpoint struct {
	AbstractBase

	name string
}

func (c *DeleteEndpoint) DefineCommand() (command *Command) {
	const use = "delete-endpoint"
	const short = "Delete an endpoint"
	const long = short
	const example = `Delete an endpoint
	sqlcmd config delete-endpoint --name my-endpoint`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run})

	command.PersistentFlags().StringVar(
		&c.name,
		"name",
		"",
		"Name of endpoint to delete")

	return
}

func (c *DeleteEndpoint) run(cmd *Command, args []string) {
	if c.name == "" {
		output.Fatal("Endpoint name must be provided.  Provide endpoint name with --name flag")
	}
	config.DeleteEndpoint(c.name)

	output.Infof("Endpoint '%v' deleted", c.name)
}
