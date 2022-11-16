// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/pkg/console"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
)

type Query struct {
	cmd.Base

	text string
}

func (c *Query) DefineCommand(...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use: "query",
		Short: "Run a query against the current context",
		Examples: []cmd.ExampleInfo{
			{Description: "Run a query", Steps: []string{
				`sqlcmd query "SELECT @@SERVERNAME"`,
				`sqlcmd query --text "SELECT @@SERVERNAME"`,
				`sqlcmd query --query "SELECT @@SERVERNAME"`,
			}}},
		Run: c.run,
		FirstArgAlternativeForFlag: &cmd.AlternativeForFlagInfo{
			Flag:  "text",
			Value: &c.text,
		},
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagInfo{
		String: &c.text,
		Name: "text",
		Shorthand: "t",
		Usage: "Command text to run"})

	// BUG(stuartpa): Decide on if --text or --query is best
	c.AddFlag(cmd.FlagInfo{
		String: &c.text,
		Name: "query",
		Shorthand: "q",
		Usage: "Command text to run"})
}

func (c *Query) run() {
	endpoint, user := config.GetCurrentContext()

	var line sqlcmd.Console = nil
	if c.text == "" {
		line = console.NewConsole("")
		defer line.Close()
	}
	s := mssql.Connect(endpoint, user, line)
	if c.text == "" {
		err := s.Run(false, false)
		c.CheckErr(err)
	} else {
		mssql.Query(s, c.text)
	}
}
