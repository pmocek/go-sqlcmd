// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"encoding/base64"
	"fmt"
	"github.com/billgraziano/dpapi"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"runtime"
)

func init() {
	var raw = false

	const use = "view"
	const short = "Display merged sqlconfig settings or a specified sqlconfig file.."
	const long = short
	const example = `# Show merged sqlconfig settings
  sqlcmd config view

  # Show merged sqlconfig settings and raw certificate data
  sqlcmd config view --raw`

	var run = func(cmd *cobra.Command, args []string) {
		var config Sqlconfig
		err := viper.Unmarshal(&config)
		cobra.CheckErr(err)

		for i, v := range config.Users {
			if raw {
				// Show password encrypted (so it can be used in other tools)
				if runtime.GOOS == "windows" {
					password, err := base64.StdEncoding.DecodeString(
						v.UserDetails.Password)
					cobra.CheckErr(err)

					config.Users[i].UserDetails.Password, err = dpapi.Decrypt(
						string(password))
					cobra.CheckErr(err)
				} else {
					// TODO: MacOS (keychain) and Linux (not sure?)
				}
			} else {
				config.Users[i].UserDetails.Password = "REDACTED"
			}
		}

		y, err := yaml.Marshal(config)
		cobra.CheckErr(err)

		fmt.Println(string(y))
	}

	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Long: long,
		Example: example,
		Run: run}

	cmd.PersistentFlags().BoolVar(
		&raw,
		"raw",
		false,
		"Display raw byte data",
	)

	configCmd.AddCommand(cmd)
}
