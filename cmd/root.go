// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Root struct {
	AbstractBase
}

type GlobalOptionsType struct {
	TrustServerCertificate bool
	DatabaseName           string
	UseTrustedConnection   bool
	UserName               string
	Endpoint               string
	AuthenticationMethod   string
	UseAad                 bool
	PacketSize             int
	LoginTimeout           int
	WorkstationName        string
	ApplicationIntent      string
	Encrypt                string
	DriverLogLevel         int
}

var GlobalOptions = &GlobalOptionsType{}

func (c *Root) GetCommand() *Command {
	const short = "sqlcmd: a command-line interface for the #SQLFamily."

	c.Command = &Command{
		Use:   appName,
		Short: short,
		Long: short,
	}

	c.addGlobalFlags()
	c.AddSubCommands()

	return c.Command
}

func (c *Root) addGlobalFlags() {
	c.Command.PersistentFlags().BoolVarP(
		&GlobalOptions.TrustServerCertificate,
		"trust-server-certificate",
		"C",
		false,
		"Whether to trust the certificate presented by the endpoint for encryption",
	)
	c.Command.PersistentFlags().StringVarP(
		&GlobalOptions.DatabaseName,
		"database-name",
		"d",
		"",
		"The initial database for the connection",
	)
	c.Command.PersistentFlags().BoolVarP(
		&GlobalOptions.UseTrustedConnection,
		"use-trusted-connection",
		"E",
		false,
		"Whether to use integrated security",
	)

	c.Command.PersistentFlags().String(
		"sqlconfig",
		"",
		"config file (default is ~/.sqlcmd/sqlconfig).",
	)

	c.Command.PersistentFlags().StringP(
		"output",
		"o",
		"yaml",
		"output type (text, json or yaml)",
	)

	c.Command.PersistentFlags().IntP(
		"verbosity",
		"v",
		2,
		"Logging verbosity. error = 0, warn = 1, info = 2, debug = 3, trace = 4",
	)
}
