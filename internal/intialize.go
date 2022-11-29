// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package internal

import (
	"github.com/microsoft/go-sqlcmd/internal/config"
	"github.com/microsoft/go-sqlcmd/internal/container"
	"github.com/microsoft/go-sqlcmd/internal/file"
	"github.com/microsoft/go-sqlcmd/internal/mssql"
	"github.com/microsoft/go-sqlcmd/internal/net"
	"github.com/microsoft/go-sqlcmd/internal/output"
	"github.com/microsoft/go-sqlcmd/internal/output/verbosity"
	"github.com/microsoft/go-sqlcmd/internal/pal"
	"github.com/microsoft/go-sqlcmd/internal/secret"
	"os"
)

type InitializeOptions struct {
	ErrorHandler func(error)
	HintHandler  func([]string)
	OutputType   string
	LoggingLevel int
}

func Initialize(options InitializeOptions) {
	if options.ErrorHandler == nil {
		panic("ErrorHandler is nil")
	}
	if options.HintHandler == nil {
		panic("HintHandler is nil")
	}
	if options.OutputType == "" {
		panic("OutputType is empty")
	}
	if options.LoggingLevel <= 0 || options.LoggingLevel > 4 {
		panic("LoggingLevel must be between 1 and 4 ")
	}

	file.Initialize(options.ErrorHandler, output.Tracef)
	mssql.Initialize(options.ErrorHandler, output.Tracef, secret.Decode)
	output.Initialize(options.ErrorHandler, output.Tracef, options.HintHandler, os.Stdout, options.OutputType, verbosity.Enum(options.LoggingLevel))
	config.Initialize(options.ErrorHandler, output.Tracef, secret.Encode, secret.Decode, net.IsLocalPortAvailable)
	container.Initialize(options.ErrorHandler, output.Tracef)
	secret.Initialize(options.ErrorHandler)
	net.Initialize(options.ErrorHandler, output.Tracef)
	pal.Initialize(options.ErrorHandler)
}
