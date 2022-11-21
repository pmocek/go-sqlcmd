// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package edge

import "github.com/microsoft/go-sqlcmd/internal/helper/cmd"

var SubCommands = []cmd.Command{
	cmd.New[*GetTags](),
}
