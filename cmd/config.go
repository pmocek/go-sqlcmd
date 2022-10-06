// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the context command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Modify sqlconfig files using subcommands like \"sqlcmd config use-context my-context\"",
	Long: `TODO.`,
}

func init() {
	rootCmd.AddCommand(configCmd)
}
