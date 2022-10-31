package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type AddEndpoint struct {
	AbstractBase
}

func (c *AddEndpoint) DefineCommand() (command *Command) {
	const use = "add-endpoint"
	const short = "Add an endpoint."
	const long = short
	const example = `Add a default context
	sqlcmd config add-context --name my-context`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run})

	return
}

func (c *AddEndpoint) run(cmd *Command, args []string) {

	output.Info(config.GetCurrentContextName())
}
