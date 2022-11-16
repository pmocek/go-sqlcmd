// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"github.com/microsoft/go-sqlcmd/internal/helpers"
	command "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
)

var loggingLevel int
var outputType string
var configFilename string
var rootCmd command.Command

// init initializes the command-line interface
func init() {
	rootCmd = command.New[*Root](root.SubCommands()...)

	helpers.Initialize(
		checkErr,
		displayHints,
		configFilename,
		outputType,
		loggingLevel,
	)
}

// RunCommandLine runs the application based on the command-line
// parameters the user has passed in
func RunCommandLine() {
	rootCmd.Execute()
}

// checkErr uses Cobra to check err, and halts the application if err is not
// nil.  Pass (inject) checkErr into all dependencies (helpers etc.) as an
// errorHandler.
//
// To aid debugging issues, if the logging level is > 2 (e.g. -v 3 or -4), we
// panic which outputs a stacktrace.
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

// IsValidSubCommand is TEMPORARY code, that will be removed when
// we enable the new cobra based CLI by default.  It returns true if the
// command-line provided by the user indicates they want the new cobra
// based CLI, e.g sqlcmd install, or sqlcmd query, or sqlcmd --help etc.
func IsValidSubCommand(command string) bool {
	return rootCmd.IsSubCommand(command)
}
