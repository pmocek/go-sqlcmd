// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type DeleteContext struct {
	BaseCommand

	name    string
	cascade bool
}

func (c *DeleteContext) DefineCommand() (command *Command) {
	const use = "delete-context NAME"
	const short = "Delete a context"
	const long = short
	const example = `Delete a context
	sqlcmd config delete-context --name my-context --cascade`

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
		"Name of context to delete")

	command.PersistentFlags().BoolVar(
		&c.cascade,
		"cascade",
		true,
		"Delete the context's endpoint and user as well")

	return
}

func (c *DeleteContext) run(cmd *Command, args []string) {
	if len(args) == 1 {
		c.name = args[0]
	}

	if c.name == "" {
		output.FatalWithHints([]string{"Use the --name flag to pass in a context name to delete"},
			"A 'name' is required")
	}

	if config.ContextExists(c.name) {
		context := config.GetContext(c.name)

		if c.cascade {
			config.DeleteEndpoint(context.Endpoint)
			if *context.User != "" {
				config.DeleteUser(*context.User)
			}
		}

		config.DeleteContext(c.name, c.cascade)

		output.Infof("Context '%v' deleted", c.name)
	} else {
		output.FatalfWithHintExamples([][]string{
			{"View available contexts", "sqlcmd config get-contexts"},
		},
			"Context '%v' does not exist", c.name)
	}
}
