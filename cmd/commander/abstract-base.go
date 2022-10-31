// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import . "github.com/spf13/cobra"

func (c *AbstractBase) DefineCommand() *Command {
	panic("Must implement definition of your command")
}

func (c *AbstractBase) SetCommand(command Command) *Command {
	c.command = &command
	c.addSubCommands()

	return c.command
}

func (c *AbstractBase) Name() string {
	return c.command.Name()
}

func (c *AbstractBase) Aliases() []string {
	return c.command.Aliases
}

func (c *AbstractBase) addSubCommands() {
	if c.SubCommands != nil {
		for _, sc := range c.SubCommands {
			if c.command != nil {
				c.command.AddCommand(sc.DefineCommand())
			}
		}
	}
}
