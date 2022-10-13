// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

var decryptCallback func(cipherText string) (secret string)

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a...any),
	decryptHandler func(cipherText string) (secret string)) {
	if errorHandler == nil {
		panic("Please provide an errorHandler")
	}
	if traceHandler == nil {
		panic("Please provide an traceHandler")
	}
	if decryptHandler == nil {
		panic("Please provide an decryptHandler")
	}

	errorCallback = errorHandler
	traceCallback = traceHandler
	decryptCallback = decryptHandler
}
