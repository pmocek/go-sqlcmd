// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"os"
	"path/filepath"
)

type Root struct {
	BaseCommand
}

func (c *Root) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "sqlcmd",
		Short: "sqlcmd: a command-line interface for the #SQLFamily",
		Examples: []ExampleInfo{
			{Description: "Run a query", Steps: []string{`sqlcmd query "SELECT @@SERVERNAME"`}}},
	}

	c.BaseCommand.DefineCommand()
	c.AddSubCommands(root.SubCommands())
	c.addGlobalFlags()
}

func (c *Root) addGlobalFlags() {
	c.AddFlag(FlagInfo{
		Bool: &GlobalOptions.TrustServerCertificate,
		Name: "trust-server-certificate",
		Shorthand: "C",
		Usage: "Whether to trust the certificate presented by the endpoint for encryption",
	})

	c.AddFlag(FlagInfo{
		String: &GlobalOptions.DatabaseName,
		Name: "database-name",
		Shorthand: "d",
		Usage: "The initial database for the connection",
	})

	c.AddFlag(FlagInfo{
		Bool: &GlobalOptions.UseTrustedConnection,
		Name: "use-trusted-connection",
		Shorthand: "E",
		Usage: "Whether to use integrated security",
	})

	home, _ := os.UserHomeDir()
	//checkErr(err)
	configFilename = filepath.Join(home, ".sqlcmd", "sqlconfig")

	c.AddFlag(FlagInfo{
		String: &configFilename,
		DefaultString: configFilename,
		Name: "sqlconfig",
		Usage: "Configuration file",
	})

	c.AddFlag(FlagInfo{
		String: &outputType,
		DefaultString: "yaml",
		Name: "output",
		Shorthand: "o",
		Usage: "output type (yaml, json or xml)",
	})

	c.AddFlag(FlagInfo{
		Int: &loggingLevel,
		DefaultInt: 2,
		Name: "verbosity",
		Shorthand: "v",
		Usage: "Log level, error = 0, warn = 1, info = 2, debug = 3, trace = 4",
	})
}
