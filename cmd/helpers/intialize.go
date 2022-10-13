// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package helpers

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/docker"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/file"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/folder"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/mssql"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/net"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output/verbosity"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
)

func Initialize(
	errorHandler func(error),
	hintHandler func([]string),
	sqlconfigFilename string,
	outputType string,
	loggingLevel int,
) {
	file.Initialize(errorHandler, output.Tracef)
	folder.Initialize(errorHandler, output.Tracef)
	mssql.Initialize(errorHandler, output.Tracef, secret.Decrypt,)
	output.Initialize(
		errorHandler,
		output.Tracef,
		hintHandler,
		outputType,
		verbosity.Enum(loggingLevel),
	)
	config.Initialize(
		errorHandler,
		output.Tracef,
		secret.Encrypt,
		secret.Decrypt,
		net.IsLocalPortAvailable,
		file.CreateEmptyFileIfNotExists,
		sqlconfigFilename,
	)
	docker.Initialize(errorHandler, output.Tracef)
	secret.Initialize(errorHandler)
	net.Initialize(errorHandler, output.Tracef)
}
