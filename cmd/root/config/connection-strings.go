// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
	"runtime"
	"strconv"
)

type ConnectionStrings struct {
	BaseCommand
}

func (c *ConnectionStrings) DefineCommand() {
	c.BaseCommand.Info = CommandInfo{
		Use: "connection-strings",
		Short: "Display connections strings for the current context",
		Examples: []ExampleInfo{
			{
				Description: "List connection strings for all client drivers",
				Steps: []string{
					"sqlcmd config connection-strings",
					"sqlcmd config cs"},
			},
		},
		Run: c.run,
		Aliases: []string{"cs"},
	}
}

func (c *ConnectionStrings) run([]string) {
	// connectionStringFormats borrowed from "portal.azure.com" "connection strings" pane
	var connectionStringFormats = map[string]string{
		"ADO.NET": "Server=tcp:%s,%s;Initial Catalog=%s;Persist Security Info=False;User ID=%s;Password=%s;MultipleActiveResultSets=False;Encode=True;TrustServerCertificate=False;Connection Timeout=30;",
		"JDBC":    "jdbc:sqlserver://%s:%s;database=%s;user=%s;password=%s;encrypt=true;trustServerCertificate=false;loginTimeout=30;",
		"ODBC":    "Driver={ODBC Driver 13 for SQL Server};Server=tcp:%s,%s;Database=%s;Uid=%s;Pwd=%s;Encode=yes;TrustServerCertificate=no;Connection Timeout=30;",
	}

	endpoint, user := config.GetCurrentContext()
	for k, v := range connectionStringFormats {
		connectionStringFormats[k] = fmt.Sprintf(v,
			endpoint.EndpointDetails.Address,
			strconv.Itoa(endpoint.EndpointDetails.Port),
			"master",
			user.BasicAuth.Username,
			secret.Decode(user.BasicAuth.Password, user.BasicAuth.PasswordEncrypted))
	}

	var format string
	if runtime.GOOS == "windows" {
		format = "SET \"SQLCMDPASSWORD=%s\" & sqlcmd -S %s,%s -U %s"
	} else {
		format = "export 'SQLCMDPASSWORD=%s'; sqlcmd -S %s,%s -U %s"
	}

	connectionStringFormats["SQLCMD"] = fmt.Sprintf(format,
		secret.Decode(user.BasicAuth.Password, user.BasicAuth.PasswordEncrypted),
		endpoint.EndpointDetails.Address,
		strconv.Itoa(endpoint.EndpointDetails.Port),
		user.BasicAuth.Username)

	output.Struct(connectionStringFormats)
}
