// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package edge

import "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"

var SubCommands = []cmd.Commander{
	cmd.New[*GetTags](),
}
