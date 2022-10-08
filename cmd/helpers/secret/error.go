// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

type errorHandlerService func(err error)
var errorHandlerCallback errorHandlerService

func checkErr(err error) {
	if errorHandlerCallback == nil {
		panic("errorHandlerCallback not initialized")
	}

	errorHandlerCallback(err)
}
