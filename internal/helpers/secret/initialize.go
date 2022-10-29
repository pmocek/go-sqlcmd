// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package secret

func init() {
	Initialize(func(err error){})
}

func Initialize(handler func(err error)) {
	errorHandlerCallback = handler
}
