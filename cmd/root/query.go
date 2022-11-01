// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/pkg/console"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	. "github.com/spf13/cobra"
)

type Query struct {
	AbstractBase

	text string
}

func (c *Query) DefineCommand() (command *Command) {
	const short = "Run a query against the current context"

	command = c.SetCommand(Command{
		Use:   "query COMMAND_TEXT",
		Short: short,
		Long:  short,
		Example: `
Run a query
  # sqlcmd query "SELECT @@SERVERNAME"`,
		ArgAliases: []string{"text"},
		Args: MaximumNArgs(1),
		Run: c.run,
	})

	command.PersistentFlags().StringVarP(
		&c.text,
		"query",
		"q",
		"",
		"Command text to run")

	return
}

func (c *Query) run(cmd *Command, args []string) {
	if len(args) > 0 && args[0] != "" && c.text != "" {
		output.FatalWithHints([]string{"Provide the query command text either as the first argument or using the --query flag"},
		"Two queries have been provided, as an argument '%v' and using the --query flag '%v'", args[0], c.text)
	}

	if len(args) > 0 {
		c.text = args[0]
	}

	endpoint, user := config.GetCurrentContext()

	var line sqlcmd.Console = nil
	if len(args) == 0 {
		line = console.NewConsole("")
		defer line.Close()
	}
	s := mssql.Connect(endpoint, user, line)
	if len(args) == 0 {
		err := s.Run(false, false)
		CheckErr(err)
	} else {
		mssql.Query(s, c.text)
	}
}
