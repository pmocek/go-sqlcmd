// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type GetUsers struct {
	AbstractBase

	detailed bool
}

func (c *GetUsers) DefineCommand() (command *Command) {
	const use = "get-users [USER_NAME]"
	const short = "Display one or many users from the sqlconfig file."
	const long = short
	const example = `# List all the users in your sqlconfig file
  sqlcmd config get-users

# List all the users in your sqlconfig file
  sqlcmd config get-users --detail

  # Describe one user in your sqlconfig file
  sqlcmd config get-users myalias`

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
		"Include user details")

	return
}

func (c *GetUsers) run(cmd *Command, args []string) {
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
		config.OutputUsers(output.Struct, c.detailed)
	}
}
