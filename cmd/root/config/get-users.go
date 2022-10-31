// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config2 "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetUsers struct {
	AbstractBase
}

func (c *GetUsers) DefineCommand() (command *Command) {
	const use = "get-users [USER_NAME]"
	const short = "Display one or many users from the sqlconfig file."
	const long = short
	const example = `# List all the users in your sqlconfig file
  sqlcmd config get-users

  # Describe one user in your sqlconfig file
  sqlcmd config get-users myalias`

	var run = func(cmd *Command, args []string) {
		if len(args) > 0 {
			name := args[0]

			if config2.UserExists(name) {
				user := config2.GetUser(name)
				output.Struct(user)
			} else {
				output.FatalfWithHints(
					[]string{"To view available users run `sqlcmd config get-users"},
					"error: no user exists with the name: \"%v\"",
					name)
			}
		} else {
			config2.OutputUsers(output.Struct)
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
