package mssql

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"os"
	"strconv"
)

// TODO: Copy and Paste below, refactor

// Built-in scripting variables
const (
	SQLCMDHEADERS           = "SQLCMDHEADERS"
	SQLCMDCOLWIDTH          = "SQLCMDCOLWIDTH"
	SQLCMDMAXVARTYPEWIDTH   = "SQLCMDMAXVARTYPEWIDTH"
	SQLCMDMAXFIXEDTYPEWIDTH = "SQLCMDMAXFIXEDTYPEWIDTH"
)

// defaultVariables defines variables that cannot be removed from the map, only reset
// to their default values.
var defaultVariables = sqlcmd.Variables{
	SQLCMDCOLWIDTH:          "0",
	SQLCMDHEADERS:           "0",
	SQLCMDMAXFIXEDTYPEWIDTH: "0",
	SQLCMDMAXVARTYPEWIDTH:   "256",
}

// InitializeVariables initializes variables with default values.
// When fromEnvironment is true, then loads from the runtime environment
func InitializeVariables(fromEnvironment bool) *sqlcmd.Variables {
	variables := sqlcmd.Variables{
		SQLCMDCOLWIDTH:          defaultVariables[SQLCMDCOLWIDTH],
		SQLCMDHEADERS:           defaultVariables[SQLCMDHEADERS],
		SQLCMDMAXFIXEDTYPEWIDTH: defaultVariables[SQLCMDMAXFIXEDTYPEWIDTH],
		SQLCMDMAXVARTYPEWIDTH:   defaultVariables[SQLCMDMAXVARTYPEWIDTH],
	}

	return &variables
}

func Connect(endpoint sqlconfig.Endpoint, user sqlconfig.User) *sqlcmd.Sqlcmd {
	v := InitializeVariables(false)
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
