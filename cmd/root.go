// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/docker"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	. "github.com/spf13/cobra"
)

const short = "sqlcmd: a command line interface for SQL Server and Azure SQL Database."
var command = &Command{
	Use:   appName,
	Short: short,
	Long: short,
}

func Execute() {
	err := command.Execute()
	checkErr(err)
}

func init() {
	addFlags()
	OnInitialize(initializeCobra)
	addCommands()
}

func addFlags() {
	command.PersistentFlags().StringP(
		"output",
		"o",
		"yaml",
		"output type (text, json or yaml)",
	)

	command.PersistentFlags().String(
		"sqlconfig",
		"",
		"config file (default is ~/.sqlcmd/sqlconfig).",
	)
}

func initializeCobra() {
	configFile, err := command.Flags().GetString("sqlconfig")
	checkErr(err)
	outputType, err := command.Flags().GetString("output")
	checkErr(err)

	output.Initialize(outputType, checkErr)
	config.Initialize(configFile, checkErr)
	docker.Initialize(checkErr)
	secret.Initialize(checkErr)

	addGlobalOptions(command)
}

func addCommands() {
	for _, c := range root.Commands {
		command.AddCommand(c.GetCommand())
	}
}

func checkErr(err error) {
	CheckErr(err)
}
