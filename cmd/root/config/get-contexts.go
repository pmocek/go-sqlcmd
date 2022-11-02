// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetContexts struct {
	AbstractBase

	detailed bool
}

func (c *GetContexts) DefineCommand() (command *Command) {
	const use = "get-contexts [CONTEXT_NAME]"
	const short = "Display one or many contexts from the sqlconfig file."
	const long = short
	const example = `# List all the context names in your sqlconfig file
  sqlcmd config get-contexts

# List all the contexts in your sqlconfig file
  sqlcmd config get-contexts --detail

  # Describe one context in your sqlconfig file
  sqlcmd config get-contexts my-context`

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
		"Include context details")

	return
}

func (c *GetContexts) run(cmd *Command, args []string) {
	if len(args) > 0 {
		name := args[0]

		if config.ContextExists(name) {
			context := config.GetContext(name)
			output.Struct(context)
		} else {
			output.FatalfWithHints(
				[]string{"To view available contexts run `sqlcmd config get-contexts`"},
				"error: no context exists with the name: \"%v\"",
				name)
		}
	} else {
		config.OutputContexts(output.Struct, c.detailed)
	}
}
