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
	const use = "get-contexts [CONTEXT_NAME]"
	const short = "Display one or many contexts from the sqlconfig file."
	const long = short
	const example = `# List all the contexts in your sqlconfig file
  sqlcmd config get-contexts

  # Describe one context in your sqlconfig file
  sqlcmd config get-contexts my-context`

	var run = func(cmd *cobra.Command, args []string) {
		var config Sqlconfig
		viper.Unmarshal(&config)

		if len(args) > 0 {
			name := args[0]

			if contextExists(config, name) {
				context := getContext(config, name)
				fmt.Println(context)
			} else {
				fmt.Printf("error: no context exists with the name: \"%v\"", name)
			}
		} else {
			for _, v := range config.Contexts {
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
