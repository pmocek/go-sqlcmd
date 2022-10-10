// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package file

var errorHandlerCallback  func(err error)

func checkErr(err error) {
	if errorHandlerCallback == nil {
		panic("errorHandlerCallback not initialized")
	}

	errorHandlerCallback(err)
}
