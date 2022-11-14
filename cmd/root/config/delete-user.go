// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type DeleteUser struct {
	BaseCommand

	name string
}

func (c *DeleteUser) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "delete-user",
		Short: "Delete a user",
		Examples: []ExampleInfo{
			{
				Description: "Delete a user",
				Steps: []string{
					"sqlcmd config delete-user --name user1",
					"sqlcmd config delete-user user1"}},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "Name of user to delete"})
}

func (c *DeleteUser) run(args []string) {
	config.DeleteUser(c.name)
	output.Infof("User '%v' deleted", c.name)
}
