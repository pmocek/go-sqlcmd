// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package docker

func init() {
	Initialize(
		func(err error) {},
		func(format string, a ...any) {})
}

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a ...any)) {
	errorCallback = errorHandler
	traceCallback = traceHandler
}
