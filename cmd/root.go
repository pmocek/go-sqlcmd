// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helper/pal"
)

type Root struct {
	cmd.Base
}

func (c *Root) DefineCommand(subCommands ...cmd.Command) {
	c.Base.Options = cmd.Options{
		Use:   "sqlcmd",
		Short: "sqlcmd: a command-line interface for the #SQLFamily",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Run a query",
				Steps:       []string{`sqlcmd query "SELECT @@SERVERNAME"`}}},
	}

	c.Base.DefineCommand(subCommands...)
	c.addGlobalFlags()
}

func (c *Root) addGlobalFlags() {
	c.AddFlag(cmd.FlagOptions{
		Bool:      &globalOptions.TrustServerCertificate,
		Name:      "trust-server-certificate",
		Shorthand: "C",
		Usage:     "Whether to trust the certificate presented by the endpoint for encryption",
	})

	c.AddFlag(cmd.FlagOptions{
		String:    &globalOptions.DatabaseName,
		Name:      "database-name",
		Shorthand: "d",
		Usage:     "The initial database for the connection",
	})

	c.AddFlag(cmd.FlagOptions{
		Bool:      &globalOptions.UseTrustedConnection,
		Name:      "use-trusted-connection",
		Shorthand: "E",
		Usage:     "Whether to use integrated security",
	})

	configFilename = pal.FilenameInUserHomeDotDirectory(
		".sqlcmd",
		"sqlconfig")

	c.AddFlag(cmd.FlagOptions{
		String:        &configFilename,
		DefaultString: configFilename,
		Name:          "sqlconfig",
		Usage:         "Configuration file",
	})

	c.AddFlag(cmd.FlagOptions{
		String:        &outputType,
		DefaultString: "yaml",
		Name:          "output",
		Shorthand:     "o",
		Usage:         "output type (yaml, json or xml)",
	})

	c.AddFlag(cmd.FlagOptions{
		Int:        &loggingLevel,
		DefaultInt: 2,
		Name:       "verbosity",
		Shorthand:  "v",
		Usage:      "Log level, error = 0, warn = 1, info = 2, debug = 3, trace = 4",
	})
}
