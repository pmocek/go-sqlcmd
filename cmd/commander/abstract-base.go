package commander

import 	. "github.com/spf13/cobra"

func (c *AbstractBase) GetCommand() *Command {
	panic("Must implement")
}

func (c *AbstractBase) AddSubCommands() {
	if c.SubCommands != nil {
		for _, sc := range c.SubCommands {
			if c.Command != nil { // BUG(stuartpa): why would this be nil?
				c.Command.AddCommand(sc.GetCommand())
			}
		}
	}
}
