// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package root

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
)

type Bulk struct {
	AbstractBase
}

func (c *Bulk) GetCommand() (command *Command) {
	const short = "Bulk import or export data"

	command = c.AddCommand(Command{
		Use:   "bulk",
		Short: short,
		Long:  short,
		Run: func(cmd *Command, args []string) {
			fmt.Println("bulk called")
		},
	})

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
