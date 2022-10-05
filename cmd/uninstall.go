// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type UninstallArguments struct {
	Name string
}

var uninstallArguments UninstallArguments

// installCmd represents the install command
var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstall SQL Server and Tools",
	Long: `TODO.`,
	Run: func(cmd *cobra.Command, args []string) {
		var config Sqlconfig

		viper.Unmarshal(&config)

		if config.CurrentContext == "" {
			fmt.Println("No current context")
			return
		}

		if len(config.Endpoints) == 0 {
			fmt.Println("No endpoints to uninstall")
			return
		}

		containerId := getContainerId(config)
		containerIdShort := containerId[len(containerId)-12:]

		c, err := NewController()
		cobra.CheckErr(err)

		fmt.Printf("Stopping SQL Server\n")

		err = c.ContainerStop(containerId)
		cobra.CheckErr(err)

		 err = c.ContainerRemove(containerId)
		cobra.CheckErr(err)

		removeContext(&config)
		saveConfig(config)

		fmt.Printf("SQL Server uninstalled (id: '%s')\n", containerIdShort)
	},
}

func getContainerId(config Sqlconfig) (containerId string) {
	currentContextName := config.CurrentContext

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					containerId = e.DockerDetails.ContainerId
					return
				}
			}
		}
	}
	panic("ContainerId not found")
}

func removeContext(config *Sqlconfig) {
	currentContextName := config.CurrentContext

	for i, c := range config.Contexts {
		if c.Name == currentContextName {
			for ei, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					config.Endpoints = append(
						config.Endpoints[:ei],
						config.Endpoints[ei+1:]...)
					break
				}
			}

			for ii, u := range config.Users {
				if u.Name == c.User {
					config.Users = append(
						config.Users[:ii],
						config.Users[ii+1:]...)
					break
				}
			}

			config.Contexts = append(
				config.Contexts[:i],
				config.Contexts[i+1:]...)
			break
		}
	}

	if len(config.Contexts) > 0 {
		config.CurrentContext = config.Contexts[0].Name
	} else {
		config.CurrentContext = ""
	}

	return
}

func init() {
	rootCmd.AddCommand(uninstallCmd)
}
