// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import . "github.com/spf13/cobra"

type BaseCommand struct {
	Info CommandInfo

	command     Command
	subCommands []Commander
}

type CommandInfo struct {
	Use string
	Short string
	Long string
	Examples []ExampleInfo
	Aliases []string
	Run func([]string)
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
