// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"fmt"
	"github.com/microsoft/go-sqlcmd/internal/helper/cmd"
	"github.com/microsoft/go-sqlcmd/internal/helper/config"
	"github.com/microsoft/go-sqlcmd/internal/helper/output"
)

type DeleteEndpoint struct {
	cmd.Base

	name string
}

func (c *DeleteEndpoint) DefineCommand(...cmd.Command) {
	c.Base.Options = cmd.Options{
		Use:   "delete-endpoint",
		Short: "Delete an endpoint",
		Examples: []cmd.ExampleInfo{
			{
				Description: "Delete an endpoint",
				Steps: []string{
					"sqlcmd config delete-endpoint --name my-endpoint",
					"sqlcmd config delete-context endpoint"},
			},
		},
		Run: c.run,

		FirstArgAlternativeForFlag: &cmd.AlternativeForFlagInfo{Flag: "name", Value: &c.name},
	}

	c.Base.DefineCommand()

	c.AddFlag(cmd.FlagOptions{
		String: &c.name,
		Name:   "name",
		Usage:  "Name of endpoint to delete"})
}

func (c *DeleteEndpoint) run() {
	if c.name == "" {
		output.Fatal("Endpoint name must be provided.  Provide endpoint name with --name flag")
	}

	if config.EndpointExists(c.name) {
		config.DeleteEndpoint(c.name)
	} else {
		output.FatalfWithHintExamples([][]string{
			{"View endpoints", "sqlcmd config get-endpoints"},
		},
			fmt.Sprintf("Endpoint '%v' does not exist", c.name))
	}

	output.Infof("Endpoint '%v' deleted", c.name)
}
