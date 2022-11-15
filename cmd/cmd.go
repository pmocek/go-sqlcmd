// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"github.com/microsoft/go-sqlcmd/internal/helpers"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

var loggingLevel int
var outputType string
var configFilename string
var rootCmd commander.Commander

// init initializes the command-line interface
func init() {
	rootCmd = commander.NewCommand[*Root]()

	helpers.Initialize(
		checkErr,
		displayHints,
		configFilename,
		outputType,
		loggingLevel,
	)
}

func SetArgsForUnitTests(args []string) {
	rootCmd = commander.NewCommand[*Root]()
	rootCmd.ArgsForUnitTesting(args)
}

// RunCommandLine runs the application based on the command-line
// parameters the user has passed in
func RunCommandLine(negativeUnitTest bool) {
	err := rootCmd.Execute()
	checkErr2(err, negativeUnitTest)
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
	rootCmd.CheckErr(err)
}

func checkErr2(err error, panicOnFailure bool) {
	if loggingLevel > 2 {
		if err != nil {
			panic(err)
		}
	}
	fmt.Println("GOT HERE2")
	fmt.Println(err)
	fmt.Println(panicOnFailure)

	if panicOnFailure && err != nil {
		fmt.Println("GOT HERE")
		panic(err)
	} else {
		rootCmd.CheckErr(err)
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
	}
}

func IsValidRootCommand(command string) (valid bool) {
	for _, c := range root.SubCommands() {
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
