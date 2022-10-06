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
	const use = "current-context"
	const short = "Display the current-context."
	const long = short
	const example = `Display the current-context
	sqlcmd config current-context`

	var run = func(cmd *cobra.Command, args []string) {
		var config Sqlconfig
		viper.Unmarshal(&config)
		fmt.Println(config.CurrentContext)
	}

	configCmd.AddCommand(&cobra.Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Run: run})
}
