// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package helper

import (
	"github.com/microsoft/go-sqlcmd/internal/helper/config"
	"github.com/microsoft/go-sqlcmd/internal/helper/container"
	"github.com/microsoft/go-sqlcmd/internal/helper/file"
	"github.com/microsoft/go-sqlcmd/internal/helper/mssql"
	"github.com/microsoft/go-sqlcmd/internal/helper/net"
	"github.com/microsoft/go-sqlcmd/internal/helper/output"
	"github.com/microsoft/go-sqlcmd/internal/helper/output/verbosity"
	"github.com/microsoft/go-sqlcmd/internal/helper/pal"
	"github.com/microsoft/go-sqlcmd/internal/helper/secret"
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
		file.CreateEmptyIfNotExists,
		sqlconfigFilename,
	)
	container.Initialize(info.ErrorHandler, info.TraceHandler)
	secret.Initialize(info.ErrorHandler)
	net.Initialize(info.ErrorHandler, info.TraceHandler)
	pal.Initialize(info.ErrorHandler)
}
