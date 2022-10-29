// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/file"
	"github.com/microsoft/go-sqlcmd/internal/helpers/net"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"github.com/microsoft/go-sqlcmd/internal/helpers/secret"
)

var encryptCallback func(plainText string) (cipherText string)
var decryptCallback func(cipherText string) (secret string)
var isLocalPortAvailableCallback func(port int) (portAvailable bool)
var createEmptyFileIfNotExistsCallback func(filename string)

func init() {
	Initialize(
		func(err error) {if err != nil {panic(err)}},
		output.Tracef,
		secret.Encrypt,
		secret.Decrypt,
		net.IsLocalPortAvailable,
		file.CreateEmptyFileIfNotExists,
		"sqlconfig-test",
	)
}

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a ...any),
	encryptHandler func(plainText string) (cipherText string),
	decryptHandler func(cipherText string) (secret string),
	isLocalPortAvailableCallbackHandler func(port int) (portAvailable bool),
	createEmptyFileIfNotExistsHandler func(filename string),
	configFile string,
) {
	if errorHandler == nil {
		panic("Please provide an errorHandler")
	}
	if traceHandler == nil {
		panic("Please provide an traceHandler")
	}
	if encryptHandler == nil {
		panic("Please provide an encryptHandler")
	}
	if decryptHandler == nil {
		panic("Please provide an decryptHandler")
	}
	if isLocalPortAvailableCallbackHandler == nil {
		panic("Please provide an  isLocalPortAvailableCallbackHandler")
	}
	if createEmptyFileIfNotExistsHandler == nil {
		panic("Please provide an  createEmptyFileIfNotExistsHandler")
	}
	errorCallback = errorHandler
	traceCallback = traceHandler
	encryptCallback = encryptHandler
	decryptCallback = decryptHandler
	isLocalPortAvailableCallback = isLocalPortAvailableCallbackHandler
	createEmptyFileIfNotExistsCallback = createEmptyFileIfNotExistsHandler

	configureViper(configFile)
	load()
}
