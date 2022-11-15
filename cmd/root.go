// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"os"
	"path/filepath"
)

type Root struct {
	cmd.Base
}

func (c *Root) DefineCommand() {
	c.Base.Info = cmd.Info{
		Use: "sqlcmd",
		Short: "sqlcmd: a command-line interface for the #SQLFamily",
		Examples: []cmd.ExampleInfo{
			{Description: "Run a query", Steps: []string{`sqlcmd query "SELECT @@SERVERNAME"`}}},
	}

	c.Base.DefineCommand()
	c.AddSubCommands(root.SubCommands())
	c.addGlobalFlags()
}

func (c *Root) addGlobalFlags() {
	c.AddFlag(cmd.FlagInfo{
		Bool: &GlobalOptions.TrustServerCertificate,
		Name: "trust-server-certificate",
		Shorthand: "C",
		Usage: "Whether to trust the certificate presented by the endpoint for encryption",
	})

	c.AddFlag(cmd.FlagInfo{
		String: &GlobalOptions.DatabaseName,
		Name: "database-name",
		Shorthand: "d",
		Usage: "The initial database for the connection",
	})

	c.AddFlag(cmd.FlagInfo{
		Bool: &GlobalOptions.UseTrustedConnection,
		Name: "use-trusted-connection",
		Shorthand: "E",
		Usage: "Whether to use integrated security",
	})

	home, _ := os.UserHomeDir()
	//checkErr(err)
	configFilename = filepath.Join(home, ".sqlcmd", "sqlconfig")

	c.AddFlag(cmd.FlagInfo{
		String: &configFilename,
		DefaultString: configFilename,
		Name: "sqlconfig",
		Usage: "Configuration file",
	})

	c.AddFlag(cmd.FlagInfo{
		String: &outputType,
		DefaultString: "yaml",
		Name: "output",
		Shorthand: "o",
		Usage: "output type (yaml, json or xml)",
	})

	c.AddFlag(cmd.FlagInfo{
		Int: &loggingLevel,
		DefaultInt: 2,
		Name: "verbosity",
		Shorthand: "v",
		Usage: "Log level, error = 0, warn = 1, info = 2, debug = 3, trace = 4",
	})
}
