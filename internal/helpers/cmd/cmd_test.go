package cmd

import (
	"fmt"
	"testing"
)

type TopLevelCommand struct {
	Base
}

func (c *TopLevelCommand) DefineCommand(subCommands ...Command) {
	c.Info = Info{
		Use: "top-level",
		Short: "Hello-World",
		Examples: []ExampleInfo{
			{	Description: "First example",
				Steps: []string{"This is the example"}},
		},
	}

	c.Base.DefineCommand(subCommands...)
}

type SubCommand1 struct {
	Base
}

func (c *SubCommand1) DefineCommand(subCommands ...Command) {
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

func (c *SubCommand11) DefineCommand(subCommands ...Command) {
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

func (c *SubCommand2) DefineCommand(subCommands ...Command) {
	c.Info = Info{
		Use: "sub-command2",
		Short: "Sub Command 2",
		Aliases: []string{"sub-command2-alias"},
	}
	c.Base.DefineCommand()
}


func Test_EndToEnd(t *testing.T) {
	subCmd11 := New[*SubCommand11]()
	subCmd1 := New[*SubCommand1](subCmd11)
	subCmd2 := New[*SubCommand2]()

	topLevel := New[*TopLevelCommand](subCmd1, subCmd2)

	topLevel.IsSubCommand("sub-command2")
	topLevel.IsSubCommand(	"sub-command2-alias")

	topLevel.IsSubCommand("--help")
	topLevel.IsSubCommand("completion")

	var s string
	topLevel.AddFlag(FlagInfo{
		String: &s,
		Name: "string",
	})
	topLevel.AddFlag(FlagInfo{
		String: &s,
		Shorthand: "s",
		Name: "string2",
	})

	var i int
	topLevel.AddFlag(FlagInfo{
		Int: &i,
		Name: "int",
	})
	topLevel.AddFlag(FlagInfo{
		Int: &i,
		Shorthand: "i",
		Name: "int2",
	})

	var b bool
	topLevel.AddFlag(FlagInfo{
		Bool: &b,
		Name: "bool",
	})
	topLevel.AddFlag(FlagInfo{
		Bool: &b,
		Shorthand: "b",
		Name: "bool2",
	})

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
