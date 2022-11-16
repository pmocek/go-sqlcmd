// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"

func SubCommands() []cmd.Command {
	return []cmd.Command{
		cmd.New[*AddContext](),
		cmd.New[*AddEndpoint](),
		cmd.New[*AddUser](),
		cmd.New[*ConnectionStrings](),
		cmd.New[*CurrentContext](),
		cmd.New[*DeleteContext](),
		cmd.New[*DeleteEndpoint](),
		cmd.New[*DeleteUser](),
		cmd.New[*GetContexts](),
		cmd.New[*GetEndpoints](),
		cmd.New[*GetUsers](),
		cmd.New[*UseContext](),
		cmd.New[*View](),
	}
}
