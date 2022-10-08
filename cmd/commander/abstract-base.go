package commander

import (
	"fmt"
	. "github.com/spf13/cobra"
)

type AbstractBase struct {
	BaseCommand Command
}

func (c *AbstractBase) GetCommand() (command *Command) {
	panic("Must implement GetCommand")
}

func (c *AbstractBase) AddSubCommands() {
	fmt.Println("GOT HERE")
	
	return
}
