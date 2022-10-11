// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/mssql"
	. "github.com/spf13/cobra"
)

type Query struct {
	AbstractBase
}

func (c *Query) GetCommand() *Command {
	const short = "Run a query against the current context"

	c.Command = &Command{
		Use:   "query COMMAND_TEXT",
		Short: short,
		Long: short,
		Example: `Run a query
  # sqlcmd query "SELECT @@SERVERNAME"`,
		ArgAliases: []string{"text"},
		Args: ExactArgs(1),
		Run: runQuery,
	}

	return c.Command
}

func runQuery(cmd *Command, args []string) {
	endpoint, user := config.GetCurrentContext()

	s := mssql.Connect(endpoint, user)
	mssql.Query(s, args[0])
}
