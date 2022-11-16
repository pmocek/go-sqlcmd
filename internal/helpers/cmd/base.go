// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/spf13/cobra"

	"strings"
)

func (c *Base) AddFlag(info FlagInfo) {
	// BUG(stuartpa): verify info

	if info.String != nil {
		if info.Shorthand == "" {
			c.command.PersistentFlags().StringVar(
				info.String,
				info.Name,
				info.DefaultString,
				info.Usage)
		} else {
			c.command.PersistentFlags().StringVarP(
				info.String,
				info.Name,
				info.Shorthand,
				info.DefaultString,
				info.Usage)
		}
	}

	if info.Int != nil {
		if info.Shorthand == "" {
			c.command.PersistentFlags().IntVar(
				info.Int,
				info.Name,
				info.DefaultInt,
				info.Usage)
		} else {
			c.command.PersistentFlags().IntVarP(
				info.Int,
				info.Name,
				info.Shorthand,
				info.DefaultInt,
				info.Usage)
		}
	}

	if info.Bool != nil {
		if info.Shorthand == "" {
			c.command.PersistentFlags().BoolVar(
				info.Bool,
				info.Name,
				info.DefaultBool,
				info.Usage)
		} else {
			c.command.PersistentFlags().BoolVarP(
				info.Bool,
				info.Name,
				info.Shorthand,
				info.DefaultBool,
				info.Usage)
		}
	}
}

func (c *Base) ArgsForUnitTesting(args []string) {
	c.command.SetArgs(args)
}

func (c *Base) DefineCommand(subCommands ...Command) {
	if c.Info.Use == "" {
		panic("Must implement command definition")
	}

	if c.Info.Long == "" {
		c.Info.Long = c.Info.Short
	}

	c.command = cobra.Command{
		Use: c.Info.Use,
		Short: c.Info.Short,
		Long: c.Info.Long,
		Aliases: c.Info.Aliases,
		Example: c.generateExamples(),
		Run: c.run,
	}

	if c.Info.FirstArgAlternativeForFlag != nil {
		c.command.Args = cobra.MaximumNArgs(1)
	} else {
		c.command.Args = cobra.MaximumNArgs(0)
	}

	c.addSubCommands(subCommands)
}

func (c *Base) CheckErr(err error) {
	cobra.CheckErr(err)
}

func (c *Base) Command() *cobra.Command {
	return &c.command
}

func (c *Base) Execute() {
	err := c.command.Execute()
	c.CheckErr(err)
}

func (c *Base) IsSubCommand(command string) (valid bool) {

	if command == "--help" {
		valid = true
	} else if command == "completion" {
		valid = true
	} else {

	outer:
		for _, subCommand := range c.command.Commands() {
			if command == subCommand.Name() {
				valid = true
				break
			}
			for _, alias := range subCommand.Aliases {
				if alias == command {
					valid = true
					break outer
				}
			}
		}
	}
	return
}

func (c *Base) addSubCommands(commands []Command) {
	for _, subCommand := range commands {
		c.command.AddCommand(subCommand.Command())
	}
}

func (c *Base) generateExamples() string {
	var sb strings.Builder

	for _, e := range c.Info.Examples {
		sb.WriteString(fmt.Sprintf("# %v\n", e.Description))
		for _, s := range e.Steps {
			sb.WriteString(fmt.Sprintf("  %v\n", s))
		}
	}

	return sb.String()
}

func (c *Base) run(_ *cobra.Command, args []string) {
	if c.Info.FirstArgAlternativeForFlag != nil {
		if len(args) > 0 {
			flag, err := c.command.PersistentFlags().GetString(
				c.Info.FirstArgAlternativeForFlag.Flag)
			c.CheckErr(err)
			if flag != "" {
				output.Fatal(
					fmt.Sprintf(
						"Both an argument and the --%v flag have been provided. " +
							"Please provide either an argument or the --%v flag",
						c.Info.FirstArgAlternativeForFlag.Flag,
						c.Info.FirstArgAlternativeForFlag.Flag))
			}
			if c.Info.FirstArgAlternativeForFlag.Value == nil {
				panic ("Must set Value")
			}
			*c.Info.FirstArgAlternativeForFlag.Value = args[0]
		}
	}

	if c.Info.Run != nil {
		c.Info.Run()
	}
}
