// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type DeleteUser struct {
	AbstractBase

	name string
}

func (c *DeleteUser) DefineCommand() (command *Command) {
	const use = "delete-user"
	const short = "Delete a user"
	const long = short
	const example = `Delete a user
	sqlcmd config delete-user --name my-user`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run})

	command.PersistentFlags().StringVar(
		&c.name,
		"name",
		"",
		"Name of user to delete")

	return
}

func (c *DeleteUser) run(cmd *Command, args []string) {
	config.DeleteUser(c.name)

	output.Infof("User '%v' deleted", c.name)
}
