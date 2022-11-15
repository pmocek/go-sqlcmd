// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import . "github.com/spf13/cobra"

type Base struct {
	Info Info

	command     Command
	subCommands []Commander
}

type Info struct {
	Use string
	Short string
	Long string
	Examples []ExampleInfo
	Aliases []string
	Run func()
	FirstArgAlternativeForFlag *AlternativeForFlagInfo
}

type ExampleInfo struct {
	Description string
	Steps       []string
}

type AlternativeForFlagInfo struct {
	Flag string
	Value *string
}
