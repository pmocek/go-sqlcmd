// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import "github.com/spf13/cobra"

type AlternativeForFlagInfo struct {
	Flag string
	Value *string
}

type Base struct {
	Options Options

	command     cobra.Command
	subCommands []Command
}

type ExampleInfo struct {
	Description string
	Steps       []string
}
