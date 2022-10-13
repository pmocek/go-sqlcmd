// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type GetEndpoints struct {
	AbstractBase
}

func (c *GetEndpoints) GetCommand() (command *Command) {
	const use = "get-endpoints [ENDPOINT_NAME]"
	const short = "Display one or many endpoints from the sqlconfig file."
	const long = short
	const example = `# List all the endpoints in your sqlconfig file
  sqlcmd config get-endpoints

  # Describe one endpoint in your sqlconfig file
  sqlcmd config get-endpoints my-endpoint`

	var run = func(cmd *Command, args []string) {
		if len(args) > 0 {
			name := args[0]

			if config.EndpointExists(name) {
				context := config.GetEndpoint(name)
				output.Struct(context)
			} else {
				output.FatalfWithHints(
					[]string{"To view available endpoints run `sqlcmd config get-endpoints"},
					"error: no endpoint exists with the name: \"%v\"",
					name)
			}
		} else {
			config.OutputEndpoints(output.Struct)
		}
	}

	command = c.AddCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(1),
		Run:     run})

	return
}
