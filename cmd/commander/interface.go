package commander

import (
	. "github.com/spf13/cobra"
)

type Commander interface {
	GetCommand() *Command
	AddSubCommands()
	//Run(cmd *Command, args []string)
}
