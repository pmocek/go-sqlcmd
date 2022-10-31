// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package mssql

var decryptCallback func(cipherText string) (secret string)

func init() {
	Initialize(
		func(err error) {},
		func(format string, a ...any) {},
		func(cipherText string) (secret string) {return})
}

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a ...any),
	decryptHandler func(cipherText string) (secret string)) {


	errorCallback = errorHandler
	traceCallback = traceHandler
	decryptCallback = decryptHandler
}
