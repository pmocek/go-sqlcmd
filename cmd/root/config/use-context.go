// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type UseContext struct {
	BaseCommand

	name string
}

func (c *UseContext) DefineCommand() (command *Command) {
	const use = "use-context CONTEXT_NAME"
	const short = "Set the current-context in a sqlconfig file."
	const long = short
	const example = `# Use the context for the sa@sql1 sql instance
  sqlcmd config use-context sa@sql1`

	var run = func(cmd *Command, args []string) {
		if c.name != "" && len(args) == 1 && args[0] != "" {
			output.Fatal("Both an argument and the --name flag have been provided.  Please provide either an argument or the --name flag")
		}
		if c.name == "" {
			if len(args) == 1 {
				c.name = args[0]
			}
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

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(1),
		Aliases: []string{"use", "change-context", "set-context"},
		Run:     run})

	command.PersistentFlags().StringVar(
		&c.name,
		"name",
		"",
		"Name of context to set as current context")

	return
}
