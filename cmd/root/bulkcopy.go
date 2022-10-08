// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"fmt"

	. "github.com/spf13/cobra"
)

type Bulkcopy struct {}

func (c *Bulkcopy) GetCommand() (command *Command) {

	command = &Command{
		Use:   "bulkcopy",
		Short: "Bulk import or export data",
		Long:  `TODO.`,
		Run: func(cmd *Command, args []string) {
			fmt.Println("bulkcopy called")
		},
	}

	return
}

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
