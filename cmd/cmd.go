// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"github.com/microsoft/go-sqlcmd/internal/helpers"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
	"os"
)

var rootCmd *Command
var loggingLevel int

// init initializes the command-line interface
func init() {
	r := Root{AbstractBase{SubCommands: root.Commands}}
	rootCmd = r.DefineCommand()

	OnInitialize(initializeCobra)
}

// ExecuteCommandLine runs the application based on the command-line
// parameters the user has passed in
func ExecuteCommandLine() {

	setDefaultSubCommandForInstallMssql()

	err := rootCmd.Execute()
	checkErr(err)
}

// setDefaultSubCommandForInstallMssql runs `sqlcmd install mssql server` if
// no subcommand is added for 'mssql'
//
// BUG(stuartpa): Workout a way to  encapsulate this code in the install 'mssql' command
func setDefaultSubCommandForInstallMssql() {
	if len(os.Args) == 3 {
		if os.Args[1] == "install" || os.Args[1] == "create" {
			if os.Args[2] == "mssql" {
				args := append(os.Args[1:], "server")
				rootCmd.SetArgs(args)
			}
		}
	}
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
