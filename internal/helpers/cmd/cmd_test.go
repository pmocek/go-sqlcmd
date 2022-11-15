package cmd

import (
	"fmt"
	"testing"
)

type TopLevelCommand struct {
	Base
}

func (c *TopLevelCommand) DefineCommand() {
	c.Info = Info{
		Use: "top-level",
		Short: "Hello-World",
		Examples: []cmd.ExampleInfo{
			{	Description: "First example",
				Steps: []string{"This is the example"}},
		},
	}

	c.Base.DefineCommand()
}

type SubCommand1 struct {
	Base
}

func (c *SubCommand1) DefineCommand() {
	c.Info = Info{
		Use: "sub-command1",
		Short: "Sub Command 1",
		Run: c.run,
	}
	c.Base.DefineCommand()
}

func (c *SubCommand1) run() {
	fmt.Println("Sub Command 1")
}

type SubCommand11 struct {
	Base
}

func (c *SubCommand11) DefineCommand() {
	c.Info = Info{
		Use: "sub-command11",
		Short: "Sub Command 11",
		Run: c.run,
	}
	c.Base.DefineCommand()
}

func (c *SubCommand11) run() {
	fmt.Println("Sub Command 11")
}

type SubCommand2 struct {
	Base
}

func (c *SubCommand2) DefineCommand() {
	c.Info = Info{
		Use: "sub-command2",
		Short: "Sub Command 2",
	}
	c.Base.DefineCommand()
}


func Test_EndToEnd(t *testing.T) {
	topLevel := New[*TopLevelCommand]()

	subCmd1 := New[*SubCommand1]()

	topLevel.AddSubCommand(subCmd1)
	topLevel.AddSubCommand(New[*SubCommand2]())
	subCmd1.AddSubCommand(New[*SubCommand11]())

	topLevel.ArgsForUnitTesting([]string{"--help"})
	topLevel.Execute()

	topLevel.ArgsForUnitTesting([]string{"sub-command1", "--help"})
	topLevel.Execute()

	topLevel.ArgsForUnitTesting([]string{"sub-command1 sub-command11"})
	topLevel.Execute()

	topLevel.ArgsForUnitTesting([]string{"sub-command1"})
	topLevel.Execute()
}

func TestAbstractBase_DefineCommand(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	c := Base{}
	c.DefineCommand()
}
