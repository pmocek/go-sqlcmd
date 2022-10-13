// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package folder

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a...any)) {
	if errorHandler == nil {
		panic("Please provide an errorHandler")
	}
	if traceHandler == nil {
		panic("Please provide an traceHandler")
	}

	errorCallback = errorHandler
	traceCallback = traceHandler
}
