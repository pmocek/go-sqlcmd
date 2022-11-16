// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type GetContexts struct {
	cmd.Base

	name string
	detailed bool
}

func (c *GetContexts) DefineCommand(...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use: "get-contexts",
		Short: "Display one or many contexts from the sqlconfig file",
		Examples: []cmd.ExampleInfo{
			{
				Description: "List all the context names in your sqlconfig file",
				Steps: []string{"sqlcmd config get-contexts"},
			},
			{
				Description: "List all the contexts in your sqlconfig file",
				Steps: []string{"sqlcmd config get-contexts --detailed"},
			},
			{
				Description: "Describe one context in your sqlconfig file",
				Steps: []string{"sqlcmd config get-contexts my-context"},
			},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &cmd.AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "Context name to view details of"})

	c.AddFlag(cmd.FlagInfo{
		Bool: &c.detailed,
		Name: "detailed",
		Usage: "Include context details"})
}

func (c *GetContexts) run() {
	if c.name != "" {
		if config.ContextExists(c.name) {
			context := config.GetContext(c.name)
			output.Struct(context)
		} else {
			output.FatalfWithHints(
				[]string{"To view available contexts run `sqlcmd config get-contexts`"},
				"error: no context exists with the name: \"%v\"",
				c.name)
		}
	} else {
		config.OutputContexts(output.Struct, c.detailed)
	}
}
