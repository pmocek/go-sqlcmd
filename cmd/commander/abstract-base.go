// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import . "github.com/spf13/cobra"

func (c *AbstractBase) GetCommand() *Command {
	panic("Must implement")
}

func  (c *AbstractBase)  AddCommand(command Command) *Command {
	c.command = &command
	c.addSubCommands()

	return c.command
}

func (c *AbstractBase) addSubCommands() {
	if c.SubCommands != nil {
		for _, sc := range c.SubCommands {
			if c.command != nil {
				c.command.AddCommand(sc.GetCommand())
			}
		}
	}
}
