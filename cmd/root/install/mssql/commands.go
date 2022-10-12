package mssql

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
)

var Commands = []Commander{
	&GetTags{},
	&Server{},
	&Edge{},
}
