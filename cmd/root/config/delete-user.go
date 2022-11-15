// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type DeleteUser struct {
	cmd.Base

	name string
}

func (c *DeleteUser) DefineCommand() {
	c.Base.Info = cmd.Info{
		Use: "delete-user",
		Short: "Delete a user",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Delete a user",
				Steps: []string{
					"sqlcmd config delete-user --name user1",
					"sqlcmd config delete-user user1"}},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &cmd.AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagInfo{
		String: &c.name,
		Name: "name",
		Usage: "Name of user to delete"})
}

func (c *DeleteUser) run() {
	config.DeleteUser(c.name)
	output.Infof("User '%v' deleted", c.name)
}
