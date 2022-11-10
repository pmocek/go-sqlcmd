// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	. "github.com/spf13/cobra"
	"path/filepath"
	"strings"
)

type Uninstall struct {
	BaseCommand

	force bool
	yes   bool
}

// systemDatabases are the list of non-user databases, used to do a safety check
// when doing a delete/drop/uninstall
var systemDatabases = [...]string{
	"/var/opt/mssql/data/msdbdata.mdf",
	"/var/opt/mssql/data/tempdb.mdf",
	"/var/opt/mssql/data/model.mdf",
	"/var/opt/mssql/data/model_msdbdata.mdf",
	"/var/opt/mssql/data/model_replicatedmaster.mdf",
	"/var/opt/mssql/data/master.mdf",
}

func (c *Uninstall) DefineCommand() (command *Command) {
	const short = "Uninstall/Delete the current context"

	command = c.SetCommand(Command{
		Use:   "uninstall",
		Short: short,
		Long:  short,
		Example: `# Uninstall/Delete the current context (includes the endpoint and user)
  sqlcmd uninstall

# Uninstall/Delete the current context, no user prompt
  sqlcmd uninstall --yes

# Uninstall/Delete the current context, no user prompt and override safety check for user databases
  sqlcmd uninstall --yes --force`,
		Args:    MaximumNArgs(0),
		Aliases: []string{"delete", "drop"},
		Run:     c.run,
	})

	command.PersistentFlags().BoolVar(
		&c.yes,
		"yes",
		false,
		"Quiet mode (do not stop for user input to confirm the operation)",
	)

	command.PersistentFlags().BoolVar(
		&c.force,
		"force",
		false,
		"Complete the operation even if non-system (user) database files are present",
	)

	return
}

func (c *Uninstall) run(*Command, []string) {
	if config.GetCurrentContextName() == "" {
		output.FatalfWithHintExamples([][]string{
			{"To view available contexts", "sqlcmd config get-contexts"},
		}, "No current context")
	}
	if currentContextEndPointExists() {
		if config.CurrentContextEndpointHasContainer() {
			controller := docker.NewController()
			id := config.GetContainerId()
			endpoint, _ := config.GetCurrentContext()

			var input string
			if !c.yes {
				output.Infof(
					"Current context is '%s'. Do you want to continue? (Y/N)",
					config.GetCurrentContextName(),
				)
				_, err := fmt.Scanln(&input)
				CheckErr(err)

				if strings.ToLower(input) != "yes" && strings.ToLower(input) != "y" {
					output.Fatal("Operation cancelled.")
				}
			}
			if !c.force {
				output.Infof("Verifying no user (non-system) database (.mdf) files")
				userDatabaseSafetyCheck(controller, id)
			}

			output.Infof(
				"Stopping %s",
				endpoint.ContainerDetails.Image,
			)
			err := controller.ContainerStop(id)
			CheckErr(err)

			output.Infof("Removing context %s", config.GetCurrentContextName())
			err = controller.ContainerRemove(id)
			CheckErr(err)
		}

		config.RemoveCurrentContext()
		config.Save()

		newContextName := config.GetCurrentContextName()
		if newContextName != "" {
			output.Infof("Current context is now %s", newContextName)
		} else {
			output.Infof("%v\n", "Operation completed successfully")
		}
	}
}

func userDatabaseSafetyCheck(controller *docker.Controller, id string) {
	files := controller.ContainerFiles(id, "*.mdf")
	for _, databaseFile := range files {
		if strings.HasSuffix(databaseFile, ".mdf") {
			isSystemDatabase := false
			for _, systemDatabase := range systemDatabases {
				if databaseFile == systemDatabase {
					isSystemDatabase = true
					break
				}
			}

			if !isSystemDatabase {
				output.FatalfWithHints([]string{
					fmt.Sprintf(
						"If the database is mounted, run `sqlcmd query \"use master; DROP DATABASE [%s]\"`",
						strings.TrimSuffix(filepath.Base(databaseFile), ".mdf")),
					"Pass in the flag --force to override this safety check for user (non-system) databases"},
					"Unable to continue, a user (non-system) database (%s) is present", databaseFile)
			}
		}
	}
}

func currentContextEndPointExists() (exists bool) {
	exists = true

	if !config.EndpointsExists() {
		output.Fatal("No endpoints to uninstall")
		exists = false
	}

	return
}
