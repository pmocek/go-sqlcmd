// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/helpers"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	. "github.com/spf13/cobra"
)

var rootCmd  *Command
var loggingLevel int

// init initializes the command-line interface
func init() {
	r := Root{AbstractBase{SubCommands: root.Commands}}
	rootCmd = r.GetCommand()

	OnInitialize(initializeCobra)
}

// ExecuteCommandLine runs the application based on the command-line
// parameters the user has passed in
func ExecuteCommandLine() {
	err := rootCmd.Execute()
	checkErr(err)
}

func initializeCobra() {
	configFile, err := rootCmd.Flags().GetString("sqlconfig")
	checkErr(err)
	outputType, err := rootCmd.Flags().GetString("output")
	checkErr(err)
	loggingLevel, err = rootCmd.Flags().GetInt("verbosity")
	checkErr(err)

	helpers.Initialize(
		checkErr,
		displayHints,
		configFile,
		outputType,
		loggingLevel,
	)
}

// checkErr uses Cobra to checks the err, and halts the application if err is not
// nil.  Pass (inject) checkErr into all dependencies (helpers etc.) as an
// errorHandler
func checkErr(err error) {
	if loggingLevel > 2 {
		if err != nil {
			panic(err)
		}
	}
	CheckErr(err)
}

// displayHints displays helpful information on what the user should do next
// to make progress.  displayHints is injected into dependencies (helpers etc.)
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

func IsValidRootCommand(command string) (valid bool) {
	for _, c := range root.Commands {
		if command == c.GetCommand().Name() {
			valid = true
			break
		}
		for _, alias := range c.GetCommand().Aliases {
			if alias == command {
				valid = true
				break
			}
		}
	}
	return
}
