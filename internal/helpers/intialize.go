// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package helpers

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/internal/helpers/docker"
	"github.com/microsoft/go-sqlcmd/internal/helpers/file"
	"github.com/microsoft/go-sqlcmd/internal/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/internal/helpers/net"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output/verbosity"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
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
	mssql.Initialize(info.ErrorHandler, info.TraceHandler, secret.Decrypt)
	output.Initialize(
		info.ErrorHandler,
		info.TraceHandler,
		hintHandler,
		outputType,
		verbosity.Enum(loggingLevel),
	)
	config.Initialize(
		info.ErrorHandler,
		info.TraceHandler,
		secret.Encrypt,
		secret.Decrypt,
		net.IsLocalPortAvailable,
		file.CreateEmptyFileIfNotExists,
		sqlconfigFilename,
	)
	docker.Initialize(info.ErrorHandler, info.TraceHandler)
	secret.Initialize(info.ErrorHandler)
	net.Initialize(info.ErrorHandler, info.TraceHandler)
}
