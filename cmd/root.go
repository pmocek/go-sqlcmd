// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appName,
	Short: "sqlcmd: a command line interface for SQL Server and Azure SQL Database.",
	Long: `sqlcmd: a command line interface for SQL Server and Azure SQL Database.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(
		&cfgFile,
		"sqlconfig",
		"",
		"config file (default is $HOME/.sqlcmd/sqlconfig).",
	)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(filepath.Join(home, ".sqlcmd"))
		viper.SetConfigType("yaml")
		viper.SetConfigName("sqlconfig")
		err = viper.ReadInConfig()
		cobra.CheckErr(err)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	err := viper.ReadInConfig()
	cobra.CheckErr(err)
	//fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}
