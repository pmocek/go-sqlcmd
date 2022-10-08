// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"fmt"
	. "github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
)

type UseContext struct {
	command Command
}

func (c *UseContext) GetCommand() (*Command) {
	const use = "use-context CONTEXT_NAME"
	const short = "Set the current-context in a sqlconfig file."
	const long = short
	const example = `# Use the context for the sa@sql1 sql instance
  sqlcmd config use-context sa@sql1`

	var run = func(cmd *Command, args []string) {
		var config Sqlconfig
		var name = args[0]

		viper.Unmarshal(&config)

		if contextExists(config, name) {
			config.CurrentContext = name
			//saveConfig(config)

			fmt.Printf("Switched to context \"%v\".", name)
		} else {
			fmt.Printf("error: no context exists with the name: \"%v\"", name)
		}
	}

	c.command = Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Args: MinimumNArgs(1),
		ArgAliases: []string{"context_name"},
		Aliases: []string{"use"},
		Run: run}

	return &c.command
}

func contextExists(config Sqlconfig, name string) (exists bool) {
	for _, c := range config.Contexts {
		if name == c.Name {
			exists = true
			break
		}
	}
	return
}

func getContext(config Sqlconfig, name string) (context Context) {
	for _, c := range config.Contexts {
		if name == c.Name {
			context = c
			break
		}
	}
	return
}

func endpointExists(config Sqlconfig, name string) (exists bool) {
	for _, c := range config.Endpoints {
		if name == c.Name {
			exists = true
			break
		}
	}
	return
}

func getEndpoint(config Sqlconfig, name string) (endpoint Endpoint) {
	for _, e := range config.Endpoints {
		if name == e.Name {
			endpoint = e
			break
		}
	}
	return
}


func userExists(config Sqlconfig, name string) (exists bool) {
	for _, v := range config.Users {
		if name == v.Name {
			exists = true
			break
		}
	}
	return
}

func getUser(config Sqlconfig, name string) (user User) {
	for _, v := range config.Users {
		if name == v.Name {
			user = v
			break
		}
	}
	return
}
