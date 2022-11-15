// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import . "github.com/microsoft/go-sqlcmd/cmd/commander"

func SubCommands() []Commander {
	return []Commander{
		NewCommand[*AddContext](),
		NewCommand[*AddEndpoint](),
		NewCommand[*AddUser](),
		NewCommand[*ConnectionStrings](),
		NewCommand[*CurrentContext](),
		NewCommand[*DeleteContext](),
		NewCommand[*DeleteEndpoint](),
		NewCommand[*DeleteUser](),
		NewCommand[*GetContexts](),
		NewCommand[*GetEndpoints](),
		NewCommand[*GetUsers](),
		NewCommand[*UseContext](),
		NewCommand[*View](),
	}
}
