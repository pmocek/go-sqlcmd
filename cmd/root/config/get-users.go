// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	config "github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type GetUsers struct {
	BaseCommand

	name string
	detailed bool
}

func (c *GetUsers) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "get-users",
		Short: "Display one or many users from the sqlconfig file",
		Examples: []ExampleInfo{
			{
				Description: "List all the users in your sqlconfig file",
				Steps: []string{"sqlcmd config get-users"},
			},
			{
				Description: "List all the users in your sqlconfig file",
				Steps: []string{"sqlcmd config get-users --detailed"},
			},
			{
				Description: "Describe one user in your sqlconfig file",
				Steps: []string{"sqlcmd config get-users user1"},
			},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "User name to view details of"})

	c.AddFlag(FlagInfo{
		Bool: &c.detailed,
		Name: "detailed",
		Usage: "Include user details"})
}

func (c *GetUsers) run(args []string) {
	if c.name != "" {
		if config.UserExists(c.name) {
			user := config.GetUser(c.name)
			output.Struct(user)
		} else {
			output.FatalfWithHints(
				[]string{"To view available users run `sqlcmd config get-users"},
				"error: no user exists with the name: \"%v\"",
				c.name)
		}
	} else {
		config.OutputUsers(output.Struct, c.detailed)
	}
}
