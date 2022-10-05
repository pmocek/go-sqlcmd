// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"encoding/base64"
	"github.com/spf13/viper"
	"os"
	"strconv"

	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"github.com/spf13/cobra"

	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
)

type QueryArguments struct {
	QueryText string
}

var queryArguments QueryArguments

// queryCmd represents the query command
var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Run a query against the current context",
	Long: `TODO.`,
	Run: func(cmd *cobra.Command, args []string) {
		config := Sqlconfig{}

		err := viper.Unmarshal(&config)
		cobra.CheckErr(err)

		endpoint, user := getCurrentContext(config)

		password, err := base64.StdEncoding.DecodeString(user.UserDetails.Password)
		cobra.CheckErr(err)

		v := InitializeVariables(false)
		s := sqlcmd.New(nil, "", v)
		connect := sqlcmd.ConnectSettings{}
		connect.ServerName = endpoint.EndpointDetails.Address + "," + strconv.Itoa(endpoint.EndpointDetails.Port)
		connect.UseTrustedConnection = false
		connect.UserName = user.UserDetails.Username
		connect.Password = string(password)
		err = s.ConnectDb(&connect, true)
		cobra.CheckErr(err)

		s.Query = args[0]
		s.Format = sqlcmd.NewSQLCmdDefaultFormatter(false)
		s.SetOutput(os.Stdout)

		s.Run(true, false)
	},
}

func getCurrentContext(config Sqlconfig) (endpoint Endpoint, user User){
	currentContextName := config.CurrentContext

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					endpoint = e
					break
				}
			}

			for _, u := range config.Users {
				if u.Name == c.User {
					user = u
					break
				}
			}
		}
	}
	return
}

func init() {
	// queryCmd.Flags().StringVarP(&queryArguments.QueryText, "text", "t", "Command text", "Command/Query text.")
	rootCmd.AddCommand(queryCmd)
}

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
