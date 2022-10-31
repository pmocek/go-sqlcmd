// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import (
	. "github.com/spf13/cobra"
)

type Commander interface {
	SetCommand(command Command) *Command
	DefineCommand() *Command
	Name() string
	Aliases() []string
}
