// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config2 "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetContexts struct {
	AbstractBase
}

func (c *GetContexts) DefineCommand() (command *Command) {
	const use = "get-contexts [CONTEXT_NAME]"
	const short = "Display one or many contexts from the sqlconfig file."
	const long = short
	const example = `# List all the contexts in your sqlconfig file
  sqlcmd config get-contexts

  # Describe one context in your sqlconfig file
  sqlcmd config get-contexts my-context`

	var run = func(cmd *Command, args []string) {
		if len(args) > 0 {
			name := args[0]

			if config2.ContextExists(name) {
				context := config2.GetContext(name)
				output.Struct(context)
			} else {
				output.FatalfWithHints(
					[]string{"To view available contexts run `sqlcmd config get-contexts`"},
					"error: no context exists with the name: \"%v\"",
					name)
			}
		} else {
			config2.OutputContexts(output.Struct)
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
