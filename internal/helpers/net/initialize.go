// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package net

func init() {
	Initialize(
		func(err error) {if err != nil {panic(err)}},
		func(format string, a ...any) {})
}

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a ...any)) {
	if errorHandler == nil {
		panic("Please provide an errorHandler")
	}
	if traceHandler == nil {
		panic("Please provide an traceHandler")
	}

	errorCallback = errorHandler
	traceCallback = traceHandler
}
