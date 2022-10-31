// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
	"os"
	"path/filepath"
)

type Root struct {
	AbstractBase
}

func (c *Root) DefineCommand() (command *Command) {
	const short = "sqlcmd: a command-line interface for the #SQLFamily."

	command = c.SetCommand(Command{
		Use:   "sqlcmd",
		Short: short,
		Long:  short,
	})

	c.addGlobalFlags(command)

	return
}

func (c *Root) addGlobalFlags(command *Command) {
	command.PersistentFlags().BoolVarP(
		&GlobalOptions.TrustServerCertificate,
		"trust-server-certificate",
		"C",
		false,
		"Whether to trust the certificate presented by the endpoint for encryption",
	)

	command.PersistentFlags().StringVarP(
		&GlobalOptions.DatabaseName,
		"database-name",
		"d",
		"",
		"The initial database for the connection",
	)

	command.PersistentFlags().BoolVarP(
		&GlobalOptions.UseTrustedConnection,
		"use-trusted-connection",
		"E",
		false,
		"Whether to use integrated security",
	)

	home, err := os.UserHomeDir()
	checkErr(err)
	configFile := filepath.Join(home, ".sqlcmd", "sqlconfig")

	command.PersistentFlags().String(
		"sqlconfig",
		configFile,
		"",
	)

	command.PersistentFlags().StringP(
		"output",
		"o",
		"yaml",
		"output type (text, json or yaml)",
	)

	command.PersistentFlags().IntP(
		"verbosity",
		"v",
		2,
		"Logging verbosity. error = 0, warn = 1, info = 2, debug = 3, trace = 4",
	)
}
