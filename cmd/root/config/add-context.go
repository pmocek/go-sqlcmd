package config

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type AddContext struct {
	AbstractBase

	name string
	endpointName string
	userName string
}

func (c *AddContext) DefineCommand() (command *Command) {
	const use = "add-context"
	const short = "Add a context"
	const long = short
	const example = `Add a default context
	sqlcmd config add-context --name my-context`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run})

	command.PersistentFlags().StringVar(
		&c.name,
		"name",
		"context",
		"Display name for the context")

	command.PersistentFlags().StringVar(
		&c.endpointName,
		"endpoint",
		"",
		"Name of endpoint this context will use, use `sqlcmd config get-endpoints` to see list")

	command.PersistentFlags().StringVar(
		&c.userName,
		"user",
		"",
		"Name of user this context will use, use `sqlcmd config get-users` to see list")

	return
}

func (c *AddContext) run(cmd *Command, args []string) {
	context := sqlconfig.Context{
		ContextDetails: sqlconfig.ContextDetails{
			Endpoint: c.endpointName,
			User:     &c.userName,
		},
		Name:           c.name,
	}

	if c.endpointName == "" || !config.EndpointExists(c.endpointName) {
		output.FatalfWithHintExamples([][]string{
			{"Use to view endpoints to choose from", "sqlcmd config get-endpoints"},
			{"Add a local endpoint", "sqlcmd install"},
			{"Add an already existing endpoint", "sqlcmd config add-endpoint --address localhost --port 1433"}},
		"An endpoint is required to add a context.  Endpoint '%v' does not exist", c.endpointName)
	}

	if c.userName != "" {
		if !config.UserExists(c.userName) {
			output.FatalfWithHintExamples([][]string{
				{"View list of users", "sqlcmd config get-users"},
				{"Add the user", fmt.Sprintf("sqlcmd config add-user --name %v", c.userName)},
				{"Add an endpoint", "sqlcmd install"}},
				"User '%v' does not exist", c.userName)
		}
	}

	config.AddContext(context)
	config.SetCurrentContextName(context.Name)
	output.InfofWithHintExamples([][]string{
			{"To start interactive query session", "sqlcmd query"},
			{"To run a query", "sqlcmd query \"SELECT @@version\""},
		},
	"Current Context '%v'", context.Name)
}
