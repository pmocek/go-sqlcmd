// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql_edge

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
)

var Commands = []Commander{
	&GetTags{},
}
