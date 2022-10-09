// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type GetUsers struct {
	command Command
}

func (c *GetUsers) GetCommand() (*Command) {
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

			if config.UserExists(name) {
				user := config.GetUser(name)
				output.Struct(user)
			} else {
				output.FatalfWithHints(
					[]string{"To view available users run `sqlcmd config get-users"},
					"error: no user exists with the name: \"%v\"",
					name)
			}
		} else {
			output.Struct(config.GetUsers())
		}
	}

	c.command = Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Args: MaximumNArgs(1),
		Run: run}

	return &c.command
}
