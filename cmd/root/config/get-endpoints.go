// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetEndpoints struct {
	BaseCommand

	detailed bool
}

func (c *GetEndpoints) DefineCommand() (command *Command) {
	const use = "get-endpoints [ENDPOINT_NAME]"
	const short = "Display one or many endpoints from the sqlconfig file."
	const long = short
	const example = `# List all the endpoints in your sqlconfig file
  sqlcmd config get-endpoints

# List all the endpoints in your sqlconfig file
  sqlcmd config get-endpoints --detail

  # Describe one endpoint in your sqlconfig file
  sqlcmd config get-endpoints my-endpoint`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(1),
		Run:     c.run})

	command.PersistentFlags().BoolVar(
		&c.detailed,
		"detailed",
		false,
		"Include endpoint details")

	return
}

func (c *GetEndpoints) run(cmd *Command, args []string) {
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
		config.OutputEndpoints(output.Struct, c.detailed)
	}
}
