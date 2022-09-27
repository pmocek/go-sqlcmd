// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/spf13/viper"
	"os"

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
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		config := Sqlconfig{}

		err := viper.Unmarshal(&config)
		if err != nil {
			fmt.Println(err)
		}

		// I don't understand viper.Unmarhsall at the moment, issue with
		// embedded struct in array
		//fmt.Println(len(config.Users))
		//fmt.Printf("%v\n", config.Users[0])
		//fmt.Println("Username:" + config.Users[0].Name)
		//fmt.Println("Password:" + config.Users[0].Password)
		//fmt.Println("Username2:" + config.Users[0].UserDetails.Username)
		//fmt.Println("Password2:" + config.Users[0].UserDetails.Password)

		password, err := base64.StdEncoding.DecodeString(config.Users[0].UserDetails.Password)
		if err != nil {
			fmt.Println(err)
		}

		v := InitializeVariables(false)
		s := sqlcmd.New(nil, "", v)
		connect := sqlcmd.ConnectSettings{}
		connect.UseTrustedConnection = false
		connect.UserName = config.Users[0].Name //BUGBUG, issue with viper.Unmarshall
		connect.Password = string(password)
		err = s.ConnectDb(&connect, true)
		if err != nil {
			fmt.Println(err)
		}
		s.Query = args[0]
		s.Format = sqlcmd.NewSQLCmdDefaultFormatter(false)
		s.SetOutput(os.Stdout)

		s.Run(true, false)
	},
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
