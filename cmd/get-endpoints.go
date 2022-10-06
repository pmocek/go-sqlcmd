// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
)

func init() {
	const use = "get-endpoints [ENDPOINT_NAME]"
	const short = "Display one or many endpoints from the sqlconfig file."
	const long = short
	const example = `# List all the endpoints in your sqlconfig file
  sqlcmd config get-endpoints

  # Describe one endpoint in your sqlconfig file
  sqlcmd config get-endpoints my-endpoint`

	var run = func(cmd *cobra.Command, args []string) {
		var config Sqlconfig
		viper.Unmarshal(&config)

		if len(args) > 0 {
			name := args[0]

			if endpointExists(config, name) {
				context := getEndpoint(config, name)
				fmt.Println(context)
			} else {
				fmt.Printf("error: no endpoint exists with the name: \"%v\"", name)
			}
		} else {
			for _, v := range config.Endpoints {
				fmt.Println(v)
			}
		}
	}

	configCmd.AddCommand(&cobra.Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Args: cobra.MaximumNArgs(1),
		Run: run})
}
