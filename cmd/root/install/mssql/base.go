// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

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
	"runtime"
)

type Base struct {
	AbstractBase

	tag string
	registry string
	repo string
	acceptEula bool
	contextName string
	defaultDatabase string

	defaultContextName string
}

func (c *Base) addFlags(
	use string,
	short string,
	long string,
	example string,
	repo string,
	defaultContextName string,
) {
	c.defaultContextName = defaultContextName

	c.Command = &Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Args:    MaximumNArgs(2),
		Run:     c.run}

	c.Command.PersistentFlags().StringVar(
		&c.registry,
		"registry",
		"mcr.microsoft.com",
		"Docker registry",
	)

	c.Command.PersistentFlags().StringVar(
		&c.repo,
		"repo",
		repo,
		"Registry repository",
	)

	c.Command.PersistentFlags().StringVar(
		&c.tag,
		"tag",
		"latest",
		"Use `sqlcmd install mssql get-tags` to see full list of options",
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
}

func (c *Base) run(cmd *Command, args []string) {
	var imageName string

	if !c.acceptEula && viper.GetString("ACCEPT_EULA") == ""  {
		output.FatalWithHints(
			[]string{"Either, add the --accept-eula flag to the command-line",
				"Or, set the environment variable SQLCMD_ACCEPT_EULA=YES "},
			"EULA not accepted")
	}

	imageName = fmt.Sprintf(
		"%s/%s:%s",
		c.registry,
		c.repo,
		c.tag)

	if c.contextName == "" {
		c.contextName = c.defaultContextName
	}

	c.installDockerImage(imageName, c.contextName)
}

func (c *Base) installDockerImage(imageName string, contextName string) {
	saPassword := secret.Generate(
		100, 10, 10, 10)

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

	var userName string
	if runtime.GOOS == "windows" {
		userName = os.Getenv("USERNAME")
	} else {
		userName = os.Getenv("USER")
	}
	if userName == "" {
		panic("Unable to get username, set env var USERNAME or USER")
	}

		password := secret.Generate(
		100, 10, 10, 10)
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
		Name: "sa",
	}, nil)
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

func (c *Base) createNonSaUser(s *sqlcmd.Sqlcmd, userName string, password string) {
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
		100, 10, 10, 10)
	mssql.Query(s, fmt.Sprintf("ALTER LOGIN [sa] WITH PASSWORD = '%s';", rotateSaPassword))
	mssql.Query(s, "ALTER LOGIN [sa] DISABLE")

	if c.defaultDatabase != "" {
		mssql.Query(s, fmt.Sprintf("ALTER AUTHORIZATION ON DATABASE::%s TO %s",
			defaultDatabase,
			userName))
	}
}
