package config

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

type AddContext struct {
	BaseCommand

	name         string
	endpointName string
	userName     string
}

func (c *AddContext) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{Use: "add-context",
		Short: "Add a context",
		Examples: []ExampleInfo{
			{
				Description: "Add a default context",
				Steps: []string{
					"sqlcmd config add-context --name my-context"},
			},
		},
		Run: c.run}

	c.BaseCommand.DefineCommand()

	c.AddFlag(FlagInfo{
		String: &c.name,
		Name: "name",
		DefaultString: "context",
		Usage: "Display name for the context"})

	c.AddFlag(FlagInfo{
		String: &c.endpointName,
		Name: "endpoint",
		Usage: "Name of endpoint this context will use, use `sqlcmd config get-endpoints` to see list"})

	c.AddFlag(FlagInfo{
			String: &c.userName,
			Name: "user",
			Usage: "Name of user this context will use, use `sqlcmd config get-users` to see list"})
}

func (c *AddContext) run(args []string) {
	context := sqlconfig.Context{
		ContextDetails: sqlconfig.ContextDetails{
			Endpoint: c.endpointName,
			User:     &c.userName,
		},
		Name: c.name,
	}

	if c.endpointName == "" || !config.EndpointExists(c.endpointName) {
		output.FatalfWithHintExamples([][]string{
			{"View existing endpoints to choose from", "sqlcmd config get-endpoints"},
			{"Add a new local endpoint", "sqlcmd install"},
			{"Add an already existing endpoint", "sqlcmd config add-endpoint --address localhost --port 1433"}},
			"Endpoint required to add context.  Endpoint '%v' does not exist.  Use --endpoint flag", c.endpointName)
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
