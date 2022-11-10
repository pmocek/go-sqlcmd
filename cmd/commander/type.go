// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package commander

import . "github.com/spf13/cobra"

type BaseCommand struct {
	command     *Command
	SubCommands []Commander
}
