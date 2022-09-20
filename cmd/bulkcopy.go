/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type BcpArguments struct {
	SchemaObjectName     string
	QueryText            string
	BatchSize            int
	ErrorFile            string
	FormatFile           string
	FirstRow             int64
	Hints                string
	InputFile            string
	LastRow              int64
	MaxErrors            int
	OutputFile           string
	RowTerminator        string
	FieldTerminator      string
	CodePage             string
	CharacterMode        bool
	ImportIdentity       bool
	KeepNulls            bool
	NativeMode           bool
	UnicodeNativeMode    bool
	QuotedIdentifiersOn  bool
	ReportVersion        bool
	TypesVersion         string
	UnicodeMode          bool
}

var bcpArguments BcpArguments

// bulkcopyCmd represents the bulkcopy command
var bulkcopyCmd = &cobra.Command{
	Use:   "bulkcopy",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("bulkcopy called")
	},
}

func init() {
	rootCmd.AddCommand(bulkcopyCmd)

	//bulkcopyCmd.PersistentFlags().StringVarP(&bcpArguments.DatabaseName, "database-name", "d", "", "Database name")
}
