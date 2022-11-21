// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helper/config"
	"github.com/microsoft/go-sqlcmd/internal/helper/output"
)

type DeleteContext struct {
	cmd.Base

	name    string
	cascade bool
}

func (c *DeleteContext) DefineCommand(...cmd.Command) {
	c.Base.Options = cmd.Options{
		Use:   "delete-context",
		Short: "Delete a context",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Delete a context",
				Steps: []string{
					"sqlcmd config delete-context --name my-context --cascade",
					"sqlcmd config delete-context my-context --cascade"},
			},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &cmd.AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagOptions{
		String: &c.name,
		Name:   "name",
		Usage:  "Name of context to delete"})

	c.AddFlag(cmd.FlagOptions{
		Bool:        &c.cascade,
		Name:        "cascade",
		DefaultBool: true,
		Usage:       "Delete the context's endpoint and user as well"})
}

func (c *DeleteContext) run() {
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

		config.DeleteContext(c.name)

		output.Infof("Context '%v' deleted", c.name)
	} else {
		output.FatalfWithHintExamples([][]string{
			{"View available contexts", "sqlcmd config get-contexts"},
		},
			"Context '%v' does not exist", c.name)
	}
}
