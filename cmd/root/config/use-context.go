// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type UseContext struct {
	BaseCommand

	name string
}

func (c *UseContext) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "use-context",
		Short: "Display one or many users from the sqlconfig file",
		Examples: []ExampleInfo{
			{
				Description: "Use the context for the user@mssql sql instance",
				Steps: []string{"sqlcmd config use-context user@mssql"},
			},
		},
		Aliases: []string{"use", "change-context", "set-context"},
		Run: c.run,

		FirstArgAlternativeForFlag: &AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "Name of context to set as current context"})
}

func (c *UseContext) run(args []string) {
	if c.name != "" {
		output.Fatal("Both an argument and the --name flag have been provided.  Please provide either an argument or the --name flag")
	}
	if config.ContextExists(c.name) {
		config.SetCurrentContextName(c.name)
		output.InfofWithHints([]string{
			"To run a query:    sqlcmd query \"SELECT @@SERVERNAME\"",
			"To remove:         sqlcmd uninstall"},
			"Switched to context \"%v\".", c.name)
	} else {
		output.FatalfWithHints([]string{"To view available contexts run `sqlcmd config get-contexts`"},
			"No context exists with the name: \"%v\"", c.name)
	}
}
