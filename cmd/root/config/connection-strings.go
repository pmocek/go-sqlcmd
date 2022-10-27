// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
	. "github.com/spf13/cobra"
	"runtime"
	"strconv"
)

type ConnectionStrings struct {
	AbstractBase
}

func (c *ConnectionStrings) GetCommand() (command *Command) {
	const use = "connection-strings"
	const short = "Display connections strings for the current context."
	const long = short
	const example = `# List connection strings for all client drivers
  sqlcmd config connection-strings

# Or
  sqlcmd config cs`

	command = c.AddCommand(Command{
		Use:     use,
		Short:   short,
		Long:    long,
		Example: example,
		Aliases: []string{"cs"},
		Args:    MaximumNArgs(1),
		Run:     c.Run})

	return
}

func (c *ConnectionStrings) Run(*Command, []string) {
	// connectionStringFormats borrowed from "portal.azure.com" "connection strings" pane
	var connectionStringFormats = map[string]string{
		"ADO.NET": "Server=tcp:%s,%s;Initial Catalog=%s;Persist Security Info=False;User ID=%s;Password=%s;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
		"JDBC":    "jdbc:sqlserver://%s:%s;database=%s;user=%s;password=%s;encrypt=true;trustServerCertificate=false;loginTimeout=30;",
		"ODBC":    "Driver={ODBC Driver 13 for SQL Server};Server=tcp:%s,%s;Database=%s;Uid=%s;Pwd=%s;Encrypt=yes;TrustServerCertificate=no;Connection Timeout=30;",
	}

	endpoint, user := config.GetCurrentContext()
	for k, v := range connectionStringFormats {
		connectionStringFormats[k] = fmt.Sprintf(v,
			endpoint.EndpointDetails.Address,
			strconv.Itoa(endpoint.EndpointDetails.Port),
			"master",
			user.UserDetails.Username,
			secret.Decrypt(user.UserDetails.Password))
	}

	var format string
	if runtime.GOOS == "windows" {
		format = "SET \"SQLCMDPASSWORD=%s\" & sqlcmd -S %s,%s -U %s"
	} else {
		format = "export 'SQLCMDPASSWORD=%s'; sqlcmd -S %s,%s -U %s"
	}

	connectionStringFormats["SQLCMD"] = fmt.Sprintf(format,
		secret.Decrypt(user.UserDetails.Password),
		endpoint.EndpointDetails.Address,
		strconv.Itoa(endpoint.EndpointDetails.Port),
		user.UserDetails.Username)

	output.Struct(connectionStringFormats)
}
