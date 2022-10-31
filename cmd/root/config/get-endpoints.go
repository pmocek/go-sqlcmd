// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config2 "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetEndpoints struct {
	AbstractBase
}

func (c *GetEndpoints) DefineCommand() (command *Command) {
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

			if config2.EndpointExists(name) {
				context := config2.GetEndpoint(name)
				output.Struct(context)
			} else {
				output.FatalfWithHints(
					[]string{"To view available endpoints run `sqlcmd config get-endpoints"},
					"error: no endpoint exists with the name: \"%v\"",
					name)
			}
		} else {
			config2.OutputEndpoints(output.Struct)
		}
	}

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(1),
		Run:     run})

	return
}
