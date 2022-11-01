package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

type AddEndpoint struct {
	AbstractBase

	name string
	address string
	port int
}

func (c *AddEndpoint) DefineCommand() (command *Command) {
	const use = "add-endpoint"
	const short = "Add an endpoint."
	const long = short
	const example = `Add a default endpoint
	sqlcmd config add-endpoint --name my-endpoint --address localhost --port 1433`

	command = c.SetCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Run:     c.run})

	command.PersistentFlags().StringVar(
		&c.name,
		"name",
		"my-endpoint",
		"Display name for the endpoint")

	command.PersistentFlags().StringVar(
		&c.address,
		"address",
		"localhost",
		"The network address to connect to, e.g. 127.0.0.1 etc.")

	command.PersistentFlags().IntVar(
		&c.port,
		"port",
		1433,
		"The network port to connect to, e.g. 1433 etc.")

	return
}

func (c *AddEndpoint) run(cmd *Command, args []string) {

	endpoint := Endpoint{
		ContainerDetails: &ContainerDetails{},
		EndpointDetails:  EndpointDetails{
			Address: c.address,
			Port:    c.port,
		},
		Name:             c.name,
	}

	config.AddEndpoint(endpoint)
	output.Infof("Endpoint '%v' added", endpoint.Name)
}
