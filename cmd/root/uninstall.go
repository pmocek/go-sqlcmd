// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/docker"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	. "github.com/spf13/cobra"
)

type Uninstall struct {}

func (c *Uninstall) GetCommand() (command *Command) {
	const short = "Uninstall/Delete the current context"
	command = &Command{
		Use:   "uninstall",
		Short: short,
		Long:  short,
		Args: MaximumNArgs(0),
		Aliases: []string{"delete"},
		Run: runUninstall,
	}

	return
}

func runUninstall(cmd *Command, args []string) {
	if currentContextEndPointExists() {
		controller := docker.NewController()
		id := config.GetContainerId()
		shortId := config.GetContainerShortId()
		endpoint, _ := config.GetCurrentContext()
		output.Infof(
			"Stopping %s (id: %s)",
			endpoint.DockerDetails.Image,
			shortId,
		)
		controller.ContainerStop(id)

		output.Infof("Removing context %s", config.GetCurrentContextName())
		controller.ContainerRemove(id)
		config.RemoveCurrentContext()
		config.Save()

		newContextName := config.GetCurrentContextName()
		if  newContextName != "" {
			output.Infof("Current context is now %s", newContextName)
		} else {
			output.Info("Operation completed successfully")
		}
	}
}

func currentContextEndPointExists() (exists bool) {
	exists = true

	if config.GetCurrentContextName() == "" {
		output.FatalWithHints([]string{"To create a context use `sqlcmd install ...`, e.g. `sqlcmd install mssql`"},"No current context")
		exists = false
	}

	if !config.EndpointsExists() {
		output.Fatal("No endpoints to uninstall")
		exists = false
	}

	return
}
