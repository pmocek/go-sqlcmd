package install

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/docker"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	. "github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

type Mssql struct {
	AbstractBase

	installType string
	acceptEula bool
	contextName string
	defaultDatabase string
}

func (c *Mssql) GetCommand() *Command {
	const use = "mssql"
	const short = "Install SQL Server"
	const long = short
	const example = `# Install SQL Server in a docker container
  sqlcmd install mssql

# Install SQL Server Edge in a docker container
  sqlcmd install mssql --type edge

# Install SQL Server in a container and specify a context-name
  sqlcmd install mssql --type --context-name my-server`

	c.Command = &Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Args: MaximumNArgs(0),
		Run: c.run}

	c.Command.PersistentFlags().StringVarP(
		&c.installType,
		"type",
		"t",
		"server",
		"Member of #SQLFamily to install (server, edge)",
	)

	c.Command.PersistentFlags().StringVarP(
		&c.contextName,
		"context-name",
		"c",
		"",
		"Context name (a default context name will be created if not provided)",
	)

	c.Command.PersistentFlags().StringVarP(
		&c.defaultDatabase,
		"user-database",
		"u",
		"",
		"Create a user database and set it as the default for the user login",
	)

	c.Command.PersistentFlags().BoolVar(
		&c.acceptEula,
		"accept-eula",
		false,
		"Accept the SQL Server EULA",
	)

	return c.Command
}

func (c *Mssql) run(cmd *Command, args []string) {
	var imageName string

	if !c.acceptEula && viper.GetString("ACCEPT_EULA") == ""  {
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
	default:
		output.FatalWithHints([]string{
		"Specify one of the valid install types, e.g. --type mssql or --type edge"},
		"'%v' is not a valid install type", c.installType)
	}

	c.installDockerImage(imageName, c.contextName)
}

func (c *Mssql) installDockerImage(imageName string, contextName string) {
	saPassword := secret.Generate(
		100, 2, 2, 2)

	env := []string{"ACCEPT_EULA=Y", fmt.Sprintf("SA_PASSWORD=%s", saPassword)}
	port := config.FindFreePortForTds()
	controller  := docker.NewController()
	output.Infof("Downloading %v", imageName)
	err := controller.EnsureImage(imageName)
	if err != nil {
		output.FatalfErrorWithHints(
			err,
			[]string{
				"Is docker installed on this machine?  If not, download from: https://docs.docker.com/get-docker/",
			"Is docker running. Try `docker ps` (list containers), does it return without error?",
			fmt.Sprintf("If `docker ps` works, try `docker pull %s`", imageName)},
			"Unable to download image %s", imageName)
	}

	output.Infof("Starting %v", imageName)
	id, err := controller.ContainerRun(imageName, env, port, []string{})
	if err != nil {
		// Remove the container, because we haven't persisted to config yet, so
		// uninstall won't work yet
		if id != "" {
			controller.ContainerRemove(id)
		}
		output.FatalErr(err)
	}

	previousContextName := config.GetCurrentContextName()

	userName := os.Getenv("USERNAME")
	password := secret.Generate(
		100, 2, 2, 2)
	// Save the config now, so user can uninstall, even if mssql in the container
	// fails to start
	config.Update(id, imageName, port, userName, password, contextName)

	output.Infof(
		"Created context '%s' in %s",
		config.GetCurrentContextName(),
		config.GetConfigFileUsed(),
	)

	// BUG(stuartpa): SQL Server bug: "SQL Server is now ready for client connections", oh no it isn't!!
	// Wait for "Server name is" instead!  Nope, that doesn't work on edge
	// Wait for "The default language" instead
	controller.ContainerWaitForLogEntry(
		id, "The default language")

	output.Infof("Rotating 'sa' password and disabling account, creating user '%s'", userName)
	endpoint, _ := config.GetCurrentContext()
	s := mssql.Connect(endpoint, sqlconfig.User{
		UserDetails: sqlconfig.UserDetails{
			Username: "sa",
			Password: secret.Encrypt(saPassword),
		},
		Name:        "sa",
	})
	c.createNonSaUser(s, userName, password)

	hints := []string{"To run a query:    sqlcmd query \"SELECT @@version\""}
	if previousContextName != "" {
		hints = append(hints, "To change context: sqlcmd config use-context "+previousContextName)
	}
	hints = append(hints, "To view config:    sqlcmd config view")
	hints = append(hints, "To remove:         sqlcmd uninstall")

	output.InfofWithHints(hints,
		"Now ready for client connections on port %d",
		port,
	)
}

func (c *Mssql) createNonSaUser(s *sqlcmd.Sqlcmd, userName string, password string) {
	defaultDatabase := "master"

	if c.defaultDatabase != "" {
		defaultDatabase = c.defaultDatabase
		output.Infof("Creating default database [%s]", defaultDatabase)
		mssql.Query(s, fmt.Sprintf("CREATE DATABASE [%s]", defaultDatabase))
	}

	const createLogin = `CREATE LOGIN [%s]
WITH PASSWORD=N'%s',
DEFAULT_DATABASE=%s,
CHECK_EXPIRATION=OFF,
CHECK_POLICY=OFF`
	const addSrvRoleMember = `EXEC master..sp_addsrvrolemember
@loginame = N'%s',
@rolename = N'sysadmin'`

	mssql.Query(s, fmt.Sprintf(createLogin, userName, password, defaultDatabase))
	mssql.Query(s, fmt.Sprintf(addSrvRoleMember, userName))
	// Correct safety protocol is to rotate the sa password, because the first
	// sa password has been in the docker environment (as SA_PASSWORD)
	rotateSaPassword := secret.Generate(
		100, 2, 2, 2)
	mssql.Query(s, fmt.Sprintf("ALTER LOGIN [sa] WITH PASSWORD = '%s';", rotateSaPassword))
	mssql.Query(s, "ALTER LOGIN [sa] DISABLE")

	if c.defaultDatabase != "" {
		mssql.Query(s, fmt.Sprintf("ALTER AUTHORIZATION ON DATABASE::%s TO %s",
			defaultDatabase,
			userName))
	}
}
