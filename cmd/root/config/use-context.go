// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type UseContext struct {
	command Command
}

func (c *UseContext) GetCommand() (*Command) {
	const use = "use-context CONTEXT_NAME"
	const short = "Set the current-context in a sqlconfig file."
	const long = short
	const example = `# Use the context for the sa@sql1 sql instance
  sqlcmd config use-context sa@sql1`

	var run = func(cmd *Command, args []string) {
		var name = args[0]
		if config.ContextExists(name) {
			config.SetCurrentContextName(name)
			output.Infof("Switched to context \"%v\".", name)
		} else {
			output.FatalfWithHints([]string{"To view available contexts run `sqlcmd config get-contexts`"},
			"No context exists with the name: \"%v\"", name)
		}
	}

	c.command = Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Args: ExactArgs(1),
		ArgAliases: []string{"context_name"},
		Aliases: []string{"use", "change-context"},
		Run: run}

	return &c.command
}
