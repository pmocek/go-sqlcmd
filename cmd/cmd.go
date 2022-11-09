// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"github.com/microsoft/go-sqlcmd/internal/helpers"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
)

var rootCmd *Command
var loggingLevel int
var panicOnFailure bool

// init initializes the command-line interface
func init() {
	r := Root{AbstractBase{SubCommands: root.Commands}}
	rootCmd = r.DefineCommand()

	OnInitialize(initializeCobra)
}

// RunCommandLine runs the application based on the command-line
// parameters the user has passed in
func RunCommandLine(negativeUnitTest bool) {
	panicOnFailure = negativeUnitTest
	err := rootCmd.Execute()
	checkErr(err)
}

func initializeCobra() {
	var configFilename, outputType string
	var err error

	configFilename, err = rootCmd.Flags().GetString("sqlconfig")
	checkErr(err)
	outputType, err = rootCmd.Flags().GetString("output")
	checkErr(err)
	loggingLevel, err = rootCmd.Flags().GetInt("verbosity")
	checkErr(err)

	helpers.Initialize(
		checkErr,
		displayHints,
		configFilename,
		outputType,
		loggingLevel,
	)
}

// checkErr uses Cobra to check err, and halts the application if err is not
// nil.  Pass (inject) checkErr into all dependencies (helpers etc.) as an
// errorHandler
//
// DEVNOTE: cobra.CheckErr (last line of function), goes on to call os.Exit(1)
// if error != nil, this will be an issue for negative Unit Tests (it will close
// down the test executor, so you'll need inject a checkErr handler that doesn't
// call os.Exit (probably one that just calls Panic(), which you recover from as
// an expected panic)
func checkErr(err error) {
	if loggingLevel > 2 {
		if err != nil {
			panic(err)
		}
	}
	if panicOnFailure && err != nil {
		panic(err)
	} else {
		CheckErr(err)
	}
}

// displayHints displays helpful information on what the user should do next
// to make progress.  displayHints is injected into dependencies (helpers etc.)
func displayHints(hints []string) {
	if len(hints) > 0 {
		output.Infof("\nHINT:")
		for i, hint := range hints {
			output.Infof("  %d. %v", i+1, hint)
		}
		//output.Infof("%v", "\n")
	}
}

func IsValidRootCommand(command string) (valid bool) {
	for _, c := range root.Commands {
		if command == c.Name() {
			valid = true
			return
		}
		for _, alias := range c.Aliases() {
			if alias == command {
				valid = true
				return
			}
		}
	}
	return
}
