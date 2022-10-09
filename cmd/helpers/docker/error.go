// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package docker

var errorHandlerCallback  func(err error)

func checkErr(err error) {
	if errorHandlerCallback == nil {
		panic("errorHandlerCallback not initialized")
	}

	errorHandlerCallback(err)
}
