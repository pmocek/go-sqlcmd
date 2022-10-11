package mssql

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"os"
	"strconv"
)

func Connect(endpoint sqlconfig.Endpoint, user sqlconfig.User) *sqlcmd.Sqlcmd {
	// BUG(stuartpa): Check what the value of fromEnvironment should be here
	v := sqlcmd.InitializeVariables(false)
	s := sqlcmd.New(nil, "", v)
	connect := sqlcmd.ConnectSettings{}
	connect.ServerName = endpoint.EndpointDetails.Address + "," +
		strconv.Itoa(endpoint.EndpointDetails.Port)
	connect.UseTrustedConnection = false
	connect.UserName = user.UserDetails.Username
	connect.Password = secret.Decrypt(user.UserDetails.Password)

	err := s.ConnectDb(&connect, true)
	checkErr(err)
	return s
}

func Query(s *sqlcmd.Sqlcmd, args []string) {
	s.Query = args[0]
	s.Format = sqlcmd.NewSQLCmdDefaultFormatter(false)
	s.SetOutput(os.Stdout)
	s.Run(true, false)
}
