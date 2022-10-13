// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

func Initialize(handler func(err error)) {
	if handler == nil {
		panic("Please provide an error handler")
	}

	errorHandlerCallback = handler
}
