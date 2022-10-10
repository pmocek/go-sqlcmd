package helpers

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/config"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/docker"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/file"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/folder"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/mssql"
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
	file.Initialize(errorHandler)
	folder.Initialize(errorHandler)
	mssql.Initialize(errorHandler)
	output.Initialize(errorHandler, hintHandler, outputType, verbosity.Enum(loggingLevel))
	config.Initialize(errorHandler, sqlconfigFilename)
	docker.Initialize(errorHandler)
	secret.Initialize(errorHandler)
}
