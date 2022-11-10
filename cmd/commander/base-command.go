// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import . "github.com/spf13/cobra"

func (c *BaseCommand) DefineCommand() *Command {
	panic("Must implement command definition")
}

func (c *BaseCommand) SetCommand(command Command) *Command {
	c.command = &command
	c.addSubCommands()

	return c.command
}

func (c *BaseCommand) Name() string {
	return c.command.Name()
}

func (c *BaseCommand) Aliases() []string {
	return c.command.Aliases
}

func (c *BaseCommand) addSubCommands() {
	if c.SubCommands != nil {
		for _, sc := range c.SubCommands {
			if c.command != nil {
				c.command.AddCommand(sc.DefineCommand())
			}
		}
	}
}
