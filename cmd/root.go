// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
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
	OnInitialize(initializeCobra)

	addFlags()
	addCommands()
	addGlobalOptions(command)
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

	command.PersistentFlags().IntP(
		"verbosity",
		"v",
		2,
		"Logging verbosity. error = 0, warn = 1, info = 2, debug = 3, trace = 4",
	)
}

func initializeCobra() {
	configFile, err := command.Flags().GetString("sqlconfig")
	checkErr(err)
	outputType, err := command.Flags().GetString("output")
	checkErr(err)
	loggingLevel, err := command.Flags().GetInt("verbosity")
	checkErr(err)

	helpers.Initialize(
		checkErr,
		displayHints,
		configFile,
		outputType,
		loggingLevel,
	)
}

func addCommands() {
	for _, c := range root.Commands {
		command.AddCommand(c.GetCommand())
	}
}

func checkErr(err error) {
	CheckErr(err)
}

func displayHints(hints []string) {
	if len(hints) > 0 {
		output.Info()
		output.Info("HINT:")
		for i, hint := range hints {
			output.Infof("  %d. %v\n", i+1, hint)
		}
		output.Info()
	}
}
