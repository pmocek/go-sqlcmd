// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type GetEndpoints struct {
	BaseCommand

	name string
	detailed bool
}

func (c *GetEndpoints) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "get-endpoints",
		Short: "Display one or many endpoints from the sqlconfig file",
		Examples: []ExampleInfo{
			{
				Description: "List all the endpoints in your sqlconfig file",
				Steps: []string{"sqlcmd config get-endpoints"},
			},
			{
				Description: "List all the endpoints in your sqlconfig file",
				Steps: []string{"sqlcmd config get-endpoints --detailed"},
			},
			{
				Description: "Describe one endpoint in your sqlconfig file",
				Steps: []string{"sqlcmd config get-endpoints my-endpoint"},
			},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "Endpoint name to view details of"})

	c.AddFlag(FlagInfo{
		Bool: &c.detailed,
		Name: "detailed",
		Usage: "Include endpoint details"})
}

func (c *GetEndpoints) run() {
	if c.name != "" {
		if config.EndpointExists(c.name) {
			context := config.GetEndpoint(c.name)
			output.Struct(context)
		} else {
			output.FatalfWithHints(
				[]string{"To view available endpoints run `sqlcmd config get-endpoints"},
				"error: no endpoint exists with the name: \"%v\"",
				c.name)
		}
	} else {
		config.OutputEndpoints(output.Struct, c.detailed)
	}
}
