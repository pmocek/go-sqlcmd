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
	command = &Command{
		Use:   "uninstall",
		Short: "Uninstall SQL Server and Tools",
		Long:  `TODO.`,
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
		output.Line("Stopping SQL Server")
		controller.ContainerStop(id)
		controller.ContainerRemove(id)
		config.RemoveCurrentContext()
		config.Save()
		output.Linef(
			"SQL Server uninstalled (id: %s). Current context is now: '%s'\n",
			shortId,
			config.GetCurrentContextName(),
		)
	}
}

func currentContextEndPointExists() (exists bool) {
	exists = true

	if config.GetCurrentContextName() == "" {
		output.Line("No current context")
		exists = false
	}

	if !config.EndpointsExists() {
		output.Line("No endpoints to uninstall")
		exists = false
	}

	return
}
