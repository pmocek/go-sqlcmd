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
		"my-context",
		"Display name for the context")

	command.PersistentFlags().StringVar(
		&c.endpointName,
		"endpoint-name",
		"my-endpoint",
		"Name of endpoint this context will use, use `sqlcmd config get-endpoints` to see list")

	command.PersistentFlags().StringVar(
		&c.userName,
		"user-name",
		"my-user",
		"Name of user this context will use, use `sqlcmd config get-users` to see list")

	return
}

func (c *AddContext) run(cmd *Command, args []string) {
	context := sqlconfig.Context{
		ContextDetails: sqlconfig.ContextDetails{
			Endpoint: c.endpointName,
			User:     c.userName,
		},
		Name:           c.name,
	}

	if !config.EndpointExists(c.endpointName) {
		output.FatalfWithHints([]string{
			"Use `sqlcmd config get-endpoints` to view endpoint list",
			fmt.Sprintf("Use `sqlcmd config add-endpoint --name %v` to add an endpoint", c.endpointName),
			"Add an endpoint using `sqlcmd install`"},
		"Endpoint '%v' does not exist", c.endpointName)
	}

	if !config.UserExists(c.userName) {
		output.FatalfWithHints([]string{
			"Use `sqlcmd config get-users` to view user list",
			fmt.Sprintf("Use `sqlcmd config add-user --name %v` to add a user", c.userName),
			"Add an endpoint using `sqlcmd install`"},
			"User '%v' does not exist", c.userName)
	}
	config.AddContext(context)
	output.Infof("Context '%v' added", context.Name)
}
