// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type DeleteEndpoint struct {
	BaseCommand

	name string
}

func (c *DeleteEndpoint) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "delete-endpoint",
		Short: "Delete an endpoint",
		Examples: []ExampleInfo{
			{
				Description: "Delete an endpoint",
				Steps: []string{
				"sqlcmd config delete-endpoint --name my-endpoint",
				"sqlcmd config delete-context endpoint"},
			},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "Name of endpoint to delete"})
}

func (c *DeleteEndpoint) run(args []string) {
	if c.name == "" {
		output.Fatal("Endpoint name must be provided.  Provide endpoint name with --name flag")
	}

	if config.EndpointExists(c.name) {
		config.DeleteEndpoint(c.name)
	} else {
		output.FatalfWithHintExamples([][]string{
			{"View endpoints", "sqlcmd config get-endpoints"},
		},
			fmt.Sprintf("Endpoint '%v' does not exist", c.name))
	}

	output.Infof("Endpoint '%v' deleted", c.name)
}
