// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import "github.com/spf13/cobra"

type AlternativeForFlagInfo struct {
	Flag string
	Value *string
}

type Base struct {
	Info Info

	command     cobra.Command
	subCommands []Command
}

type ExampleInfo struct {
	Description string
	Steps       []string
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

type Info struct {
	Aliases []string
	Examples []ExampleInfo
	FirstArgAlternativeForFlag *AlternativeForFlagInfo
	Long string
	Run func()
	Short string
	Use string
}
