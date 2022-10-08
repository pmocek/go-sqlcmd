package install

import . "github.com/microsoft/go-sqlcmd/cmd/commander"

var Commands = []Commander{
	&Mssql{},
}
