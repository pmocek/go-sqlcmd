package install

import (
	"fmt"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/docker"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	. "github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Mssql struct {
	command Command
	installType string
	acceptEula bool
	contextName string
}

func (c *Mssql) GetCommand() *Command {
	const use = "mssql"
	const short = "Install SQL Server"
	const long = short
	const example = `# Install SQL Server in a docker container
  sqlcmd install mssql

# Install the Azure SQL Database emulator in a docker container
  sqlcmd install mssql --type edge

# Install the Azure SQL Database emulator in a docker container and
  set the sqlconfig context name to 'azure-sql-emulator'
  sqlcmd install mssql --type emulator --context-name azure-sql-emulator`

	c.command = Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Run: c.run}

	c.command.PersistentFlags().StringVar(
		&c.installType,
		"type",
		"server",
		"Member of #SQLFamily to install (server, edge, emulator)",
	)

	c.command.PersistentFlags().StringVar(
		&c.contextName,
		"context-name",
		"",
		"Context name (a default context name will be created if not provided)",
	)

	c.command.PersistentFlags().BoolVar(
		&c.acceptEula,
		"accept-eula",
		false,
		"Accept the SQL Server EULA",
	)

	return &c.command
}

func (c *Mssql) run(cmd *Command, args []string) {
	var imageName string

	if !c.acceptEula && viper.Get("ACCEPT_EULA") == ""  {
		output.FatalWithHints(
			[]string{"Either, add the --accept-eula flag to the command-line",
				"Or, set the environment variable SQLCMD_ACCEPT_EULA=YES "},
			"EULA not accepted")
	}

	switch c.installType {
	case "server":
		imageName = "mcr.microsoft.com/mssql/server:2022-latest"
		if c.contextName == "" {
			c.contextName = "mssql"
		}
	case "edge":
		imageName = "mcr.microsoft.com/azure-sql-edge:latest"
		if c.contextName == "" {
			c.contextName = "edge"
		}
	case "emulator":
		// TODO: Emulator will soon have its own image
		imageName = "mcr.microsoft.com/azure-sql-edge:latest"
		if c.contextName == "" {
			c.contextName = "emulator"
		}
	default:
		output.FatalWithHints([]string{"Specify one of the valid install types, e.g. --type mssql or --type edge"},
		"'%v' is not a valid install type", c.installType)
	}

	installDockerImage(imageName, c.contextName)
}

func installDockerImage(imageName string, contextName string) {
	password := secret.Generate(
		100, 2, 2, 2)
	env := []string{"ACCEPT_EULA=Y", fmt.Sprintf("SA_PASSWORD=%s", password)}
	port := config.FindFreePortForTds()
	controller  := docker.NewController()
	output.Infof("Downloading %v", imageName)
	controller.EnsureImage(imageName)

	output.Infof("Starting %v", imageName)
	id, err := controller.ContainerRun(imageName, env, port, []string{})
	if err != nil {
		// Remove the container, because we haven't persisted to config yet, so
		// uninstall won't work yet
		controller.ContainerRemove(id)
	}

	previousContextName := config.GetCurrentContextName()

	// Save the config now, so user can uninstall, even if mssql in the container
	// fails to start
	config.Update(id, imageName, port, password, contextName)

	output.Infof(
		"Context %s starting (id: %s)",
		config.GetCurrentContextName(),
		config.GetContainerShortId(),
	)
	controller.ContainerWaitForLogEntry(
		id, "SQL Server is now ready for client connections")

	hints := []string {"To run a query:    sqlcmd query \"SELECT @@version\""}
	if previousContextName != "" {
		hints = append(hints, "To change context: sqlcmd config use-context " + previousContextName)
	}
	hints = append(hints, "To remove:         sqlcmd uninstall")

	output.InfofWithHints(hints,
		"Now ready for client connections on port %d",
		port,
	)
}
