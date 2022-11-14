// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import (
	"github.com/spf13/cobra"
)

func (c *BaseCommand) DefineCommand() {
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
		Run: func(cmd *cobra.Command, args []string) {
			if c.Info.FirstArgAlternativeForFlag != nil {
				if len(args) > 0 {
					c.Info.FirstArgAlternativeForFlag.Value = &args[0]
				}
			}

			if c.Info.Run != nil {
				c.Info.Run(args)
			}
		},
		Aliases: c.Info.Aliases,
	}

	if c.Info.FirstArgAlternativeForFlag != nil {
		c.command.Args = cobra.MaximumNArgs(1)
	} else {
		c.command.Args = cobra.MaximumNArgs(2)
	}
}

func (c *BaseCommand) ArgsForUnitTesting(args []string) {
	c.command.SetArgs(args)
}

func (c *BaseCommand) Execute() {
	c.command.Execute()
}

func (c *BaseCommand) Name() string {
	return c.command.Name()
}

func (c *BaseCommand) Aliases() []string {
	return c.command.Aliases
}

func (c *BaseCommand) AddSubCommands(commands []Commander) {
	for _, subcmd := range commands {
		c.command.AddCommand(subcmd.Command())
	}
}

func (c *BaseCommand) AddSubCommand(command Commander) {
	c.command.AddCommand(command.Command())
}

func (c *BaseCommand) Command() *cobra.Command {
	return &c.command
}

// BUG(stuartpa): I don't think golang generic support can help here yet
type FlagInfo struct {
	Name string
	Shorthand string
	Usage string

	String *string
	DefaultString string

	Int *int
	DefaultInt int

	Bool *bool
	DefaultBool bool
}

func (c *BaseCommand) CheckErr(err error) {
	cobra.CheckErr(err)
}

func (c *BaseCommand) AddFlag(info FlagInfo) {

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
		c.command.PersistentFlags().IntVar(
			info.Int,
			info.Name,
			info.DefaultInt,
			info.Usage)
	}

	if info.Bool != nil {
		c.command.PersistentFlags().BoolVar(
			info.Bool,
			info.Name,
			info.DefaultBool,
			info.Usage)
	}
}
