// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"github.com/spf13/viper"
	"os"
	"runtime"
)

type MssqlBase struct {
	BaseCommand

	tag             string
	registry        string
	repo            string
	acceptEula      bool
	contextName     string
	defaultDatabase string

	passwordLength         int
	passwordMinSpecial     int
	passwordMinNumber      int
	passwordMinUpper       int
	passwordSpecialCharSet string
	encryptPassword        bool

	useCached              bool
	errorLogEntryToWaitFor string
	defaultContextName     string
	collation              string

	sqlcmdPkg *sqlcmd.Sqlcmd
}

func (c *MssqlBase) AddFlags(
	addFlag func(FlagInfo),
	repo string,
	defaultContextName string,
) {
	c.defaultContextName = defaultContextName

	addFlag(FlagInfo{
		String: &c.registry,
		Name: "registry",
		DefaultString: "mcr.microsoft.com",
		Usage: "Container registry",
	})

	addFlag(FlagInfo{
		String: &c.repo,
		Name: "repo",
		DefaultString: repo,
		Usage: "Container repository",
	})

	addFlag(FlagInfo{
		String: &c.tag,
		Name: "tag",
		DefaultString: "latest",
		Usage: "Tag to use, use get-tags to see list of tags",
	})

	addFlag(FlagInfo{
		String: &c.contextName,
		Name: "context-name",
		Shorthand: "c",
		Usage: "Context name (a default context name will be created if not provided)",
	})

	addFlag(FlagInfo{
		String: &c.defaultDatabase,
		Name: "user-database",
		Shorthand: "u",
		Usage: "Create a user database and set it as the default for login",
	})

	addFlag(FlagInfo{
		Bool: &c.acceptEula,
		Name: "accept-eula",
		Usage: "Accept the SQL Server EULA",
	})

	addFlag(FlagInfo{
		Int: &c.passwordLength,
		DefaultInt: 50,
		Name: "password-length",
		Usage: "Generated password length",
	})

	addFlag(FlagInfo{
		Int: &c.passwordMinSpecial,
		DefaultInt: 10,
		Name: "password-min-special",
		Usage: "Minimum number of special characters",
	})

	addFlag(FlagInfo{
		Int: &c.passwordMinNumber,
		DefaultInt: 10,
		Name: "password-min-number",
		Usage: "Minimum number of numeric characters",
	})

	addFlag(FlagInfo{
		Int: &c.passwordMinUpper,
		DefaultInt: 10,
		Name: "password-min-upper",
		Usage: "Minimum number of upper characters",
	})

	addFlag(FlagInfo{
		String: &c.passwordSpecialCharSet,
		DefaultString: "!@#$%&*",
		Name: "password-special-chars",
		Usage: "Special character set to include in password",
	})

	// Windows has the DPAPI which allows securely encrypting
	if runtime.GOOS == "windows" {
		addFlag(FlagInfo{
			Bool: &c.encryptPassword,
			Name: "encrypt-password",
			Usage: "Encode the generated password in the sqlconfig file",
		})
	}

	addFlag(FlagInfo{
		Bool: &c.useCached,
		Name: "cached",
		Usage: "Don't download image.  Use already downloaded image",
	})

	// BUG(stuartpa): SQL Server bug: "SQL Server is now ready for client connections", oh no it isn't!!
	// Wait for "Server name is" instead!  Nope, that doesn't work on edge
	// Wait for "The default language" instead
	// BUG(stuartpa): This obviously doesn't work for non US LCIDs
	addFlag(FlagInfo{
		String: &c.errorLogEntryToWaitFor,
		DefaultString: "The default language",
		Name: "errorlog-wait-line",
		Usage: "Line in errorlog to wait for before connecting to disable 'sa' account",
	})

	addFlag(FlagInfo{
		String: &c.collation,
		DefaultString: "SQL_Latin1_General_CP1_CI_AS",
		Name: "collation",
		Usage: "The SQL Server collation",
	})
}

func (c *MssqlBase) Run() {
	var imageName string

	if !c.acceptEula && viper.GetString("ACCEPT_EULA") == "" {
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

func (c *MssqlBase) installDockerImage(imageName string, contextName string) {
	saPassword := c.generatePassword()

	env := []string{
		"ACCEPT_EULA=Y",
		fmt.Sprintf("MSSQL_SA_PASSWORD=%s", saPassword),
		fmt.Sprintf("MSSQL_COLLATION=%s", c.collation),
	}
	port := config.FindFreePortForTds()
	controller := docker.NewController()

	if !c.useCached {
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
	}

	output.Infof("Starting %v", imageName)
	containerId, err := controller.ContainerRun(imageName, env, port, []string{})
	if err != nil {
		// Remove the container, because we haven't persisted to config yet, so
		// uninstall won't work yet
		if containerId != "" {
			err := controller.ContainerRemove(containerId)
			output.FatalErr(err)
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

	password := c.generatePassword()

	// Save the config now, so user can uninstall, even if mssql in the container
	// fails to start
	config.AddContextWithContainer(
		contextName,
		imageName,
		port,
		containerId,
		userName,
		password,
		c.encryptPassword,
	)

	output.Infof(
		"Created context '%s' in %s",
		config.GetCurrentContextName(),
		config.GetConfigFileUsed(),
	)

	controller.ContainerWaitForLogEntry(
		containerId, c.errorLogEntryToWaitFor)

	output.Infof(
		"Disabling 'sa' account (and rotating 'sa' password). Creating user '%s'",
		userName,
	)
	endpoint, _ := config.GetCurrentContext()
	c.sqlcmdPkg = mssql.Connect(
		endpoint,
		&sqlconfig.User{
			AuthenticationType: "basic",
			BasicAuth: &sqlconfig.BasicAuthDetails{
				Username:          "sa",
				PasswordEncrypted: c.encryptPassword,
				Password:          secret.Encode(saPassword, c.encryptPassword),
			},
			Name: "sa",
		},
		nil,
	)
	c.createNonSaUser(userName, password)

	hints := [][]string{
		{"To run a query", "sqlcmd query \"SELECT @@version\""},
		{"To start interactive session", "sqlcmd query"}}

	if previousContextName != "" {
		hints = append(
			hints,
			[]string{"To change context", fmt.Sprintf(
				"sqlcmd config use-context %v",
				previousContextName,
			)},
		)
	}

	hints = append(hints, []string{"To view config", "sqlcmd config view"})
	hints = append(hints, []string{"To see connection strings", "sqlcmd config connection-strings"})
	hints = append(hints, []string{"To remove", "sqlcmd uninstall"})

	output.InfofWithHintExamples(hints,
		"Now ready for client connections on port %d",
		port,
	)
}

func (c *MssqlBase) createNonSaUser(userName string, password string) {
	defaultDatabase := "master"

	if c.defaultDatabase != "" {
		defaultDatabase = c.defaultDatabase
		output.Infof("Creating default database [%s]", defaultDatabase)
		c.Query(fmt.Sprintf("CREATE DATABASE [%s]", defaultDatabase))
	}

	const createLogin = `CREATE LOGIN [%s]
WITH PASSWORD=N'%s',
DEFAULT_DATABASE=[%s],
CHECK_EXPIRATION=OFF,
CHECK_POLICY=OFF`
	const addSrvRoleMember = `EXEC master..sp_addsrvrolemember
@loginame = N'%s',
@rolename = N'sysadmin'`

	c.Query(fmt.Sprintf(createLogin, userName, password, defaultDatabase))
	c.Query(fmt.Sprintf(addSrvRoleMember, userName))

	// Correct safety protocol is to rotate the sa password, because the first
	// sa password has been in the docker environment (as SA_PASSWORD)
	c.Query(fmt.Sprintf("ALTER LOGIN [sa] WITH PASSWORD = N'%s';",
		c.generatePassword()))
	c.Query("ALTER LOGIN [sa] DISABLE")

	if c.defaultDatabase != "" {
		c.Query(fmt.Sprintf("ALTER AUTHORIZATION ON DATABASE::[%s] TO %s",
			defaultDatabase, userName))
	}
}

func (c *MssqlBase) generatePassword() (password string) {
	password = secret.Generate(
		c.passwordLength,
		c.passwordMinSpecial,
		c.passwordMinNumber,
		c.passwordMinUpper,
		c.passwordSpecialCharSet)

	return
}

func (c *MssqlBase) Query(commandText string) {
	output.Tracef(commandText)

	//BUG(stuartpa): Need to check for errors
	mssql.Query(c.sqlcmdPkg, commandText)
}
