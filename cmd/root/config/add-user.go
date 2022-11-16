package config

import (
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
	"os"
)

type AddUser struct {
	cmd.Base

	name            string
	authType        string
	username        string
	encryptPassword bool
}

func (c *AddUser) DefineCommand(...cmd.Command) {
	c.Base.Info = cmd.Info{
		Use: "add-user",
		Short: "Add a user",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Add a user",
				Steps: []string{
					`SET SQLCMD_PASSWORD="AComp!exPa$$w0rd"`,
					"sqlcmd config add-user --name my-user --name user1"},
			},
		},
		Run: c.run,
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagInfo{
		String: &c.name,
		Name: "name",
		DefaultString: "user",
		Usage: "Display name for the user (this is not the username)",
	})

	c.AddFlag(cmd.FlagInfo{
		String: &c.authType,
		Name: "auth-type",
		DefaultString: "basic",
		Usage: "Authentication type this user will use (basic | other)",
	})

	c.AddFlag(cmd.FlagInfo{
		String: &c.username,
		Name: "username",
		Usage: "The username (provide password in SQLCMD_PASSWORD environment variable)",
	})

	c.encryptPasswordFlag()
}

func (c *AddUser) run() {
	if c.authType != "basic" &&
		c.authType != "other" {
		output.FatalfWithHints([]string{"Authentication type must be 'basic' or 'other'"},
			"Authentication type '' is not valid %v'", c.authType)
	}

	if c.authType != "basic" && c.encryptPassword {
		output.FatalWithHints([]string{
			"Remove the --encrypt-password flag",
			"Pass in the --auth-type basic"},
			"The --encrypt-password flag can only be used when authentication type is 'basic'")
	}

	user := sqlconfig.User{
		Name:               c.name,
		AuthenticationType: c.authType,
	}

	if c.authType == "basic" {
		if os.Getenv("SQLCMD_PASSWORD") == "" {
			output.FatalWithHints([]string{
				"Provide password in the SQLCMD_PASSWORD environment variable"},
				"Authentication Type 'basic' requires a password")
		}

		if c.username == "" {
			output.FatalfWithHintExamples([][]string{
				{"Provide a username with the --username flag",
					"sqlcmd config add-user --username stuartpa"},
			},
				"Username not provider")
		}

		user.BasicAuth = &sqlconfig.BasicAuthDetails{
			Username:          c.username,
			PasswordEncrypted: c.encryptPassword,
			Password:          secret.Encode(os.Getenv("SQLCMD_PASSWORD"), c.encryptPassword),
		}
	}

	config.AddUser(user)
	output.Infof("User '%v' added", user.Name)
}
