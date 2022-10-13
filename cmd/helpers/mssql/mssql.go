// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

import (
	"fmt"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"os"
)

func Connect(endpoint sqlconfig.Endpoint, user sqlconfig.User, console sqlcmd.Console) *sqlcmd.Sqlcmd {
	// BUG(stuartpa): Check what the value of fromEnvironment should be here
	v := sqlcmd.InitializeVariables(false)
	s := sqlcmd.New(console, "", v)
	s.Format = sqlcmd.NewSQLCmdDefaultFormatter(false)
	connect := sqlcmd.ConnectSettings{
		ServerName: fmt.Sprintf(
			"%s,%d",
			endpoint.EndpointDetails.Address,
			endpoint.EndpointDetails.Port),
		UseTrustedConnection: false,
		UserName: user.UserDetails.Username,
		Password: decryptCallback(user.UserDetails.Password),
	}
	err := s.ConnectDb(&connect, true)
	checkErr(err)
	return s
}

func Query(s *sqlcmd.Sqlcmd, text string) {
	s.Query = text
	s.SetOutput(os.Stdout)
	err := s.Run(true, false)
	checkErr(err)
}
