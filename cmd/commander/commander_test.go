package commander

import (
	"fmt"
	"testing"
)

type TopLevelCommand struct {
	BaseCommand
}

func (c *TopLevelCommand) DefineCommand() {
	c.Info = CommandInfo{
		Use: "top-level",
		Short: "Hello-World",
		Examples: []ExampleInfo{
			{	Description: "First example",
				Steps: []string{"This is the example"}},
		},
	}

	c.BaseCommand.DefineCommand()
}

type SubCommand1 struct {
	BaseCommand
}

func (c *SubCommand1) DefineCommand() {
	c.Info = CommandInfo{
		Use: "sub-command1",
		Short: "Sub Command 1",
		Run: c.run,
	}
	c.BaseCommand.DefineCommand()
}

func (c *SubCommand1) run(args []string) {
	fmt.Println("Sub Command 1")
}

type SubCommand11 struct {
	BaseCommand
}

func (c *SubCommand11) DefineCommand() {
	c.Info = CommandInfo{
		Use: "sub-command11",
		Short: "Sub Command 11",
		Run: c.run,
	}
	c.BaseCommand.DefineCommand()
}

func (c *SubCommand11) run(args []string) {
	fmt.Println("Sub Command 11")
}

type SubCommand2 struct {
	BaseCommand
}

func (c *SubCommand2) DefineCommand() {
	c.Info = CommandInfo{
		Use: "sub-command2",
		Short: "Sub Command 2",
	}
	c.BaseCommand.DefineCommand()
}


func Test_EndToEnd(t *testing.T) {
	topLevel := NewCommand[*TopLevelCommand]()

	subCmd1 := NewCommand[*SubCommand1]()

	topLevel.AddSubCommand(subCmd1)
	topLevel.AddSubCommand(NewCommand[*SubCommand2]())
	subCmd1.AddSubCommand(NewCommand[*SubCommand11]())

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

	c := BaseCommand{}
	c.DefineCommand()
}
