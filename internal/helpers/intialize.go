// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package helpers

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/container"
	"github.com/microsoft/go-sqlcmd/internal/helpers/file"
	"github.com/microsoft/go-sqlcmd/internal/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/internal/helpers/net"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output/verbosity"
	"github.com/microsoft/go-sqlcmd/internal/helpers/pal"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
	"os"
)

type initInfo struct {
	ErrorHandler func(error)
	TraceHandler func(string, ...any)
}

func Initialize(
	errorHandler func(error),
	hintHandler func([]string),
	sqlconfigFilename string,
	outputType string,
	loggingLevel int,
) {
	info := initInfo{errorHandler, output.Tracef}

	file.Initialize(info.ErrorHandler, info.TraceHandler)
	mssql.Initialize(info.ErrorHandler, info.TraceHandler, secret.Decode)
	output.Initialize(
		info.ErrorHandler,
		info.TraceHandler,
		hintHandler,
		os.Stdout,
		outputType,
		verbosity.Enum(loggingLevel),
	)
	config.Initialize(
		info.ErrorHandler,
		info.TraceHandler,
		secret.Encode,
		secret.Decode,
		net.IsLocalPortAvailable,
		file.CreateEmptyFileIfNotExists,
		sqlconfigFilename,
	)
	container.Initialize(info.ErrorHandler, info.TraceHandler)
	secret.Initialize(info.ErrorHandler)
	net.Initialize(info.ErrorHandler, info.TraceHandler)
	pal.Initialize(info.ErrorHandler)
}
