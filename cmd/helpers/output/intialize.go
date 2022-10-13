// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package output

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/helpers/output/formatter"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output/verbosity"
)

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a...any),
	hintHandler func(hints []string),
	format string,
	verbosity verbosity.Enum,
) {
	if errorHandler == nil {
		panic("Please provide an errorHandler")
	}
	if traceHandler == nil {
		panic("Please provide an traceHandler")
	}
	if hintHandler == nil {
		panic("Please provide a hintHandler")
	}

	errorCallback = errorHandler
	traceCallback = traceHandler
	hintCallback = hintHandler
	loggingLevel = verbosity

	switch format {
	case "json":
		formatter = &Json{Base{errorHandler}}
	case "yaml":
		formatter = &Yaml{Base{errorHandler}}
	case "xml":
		formatter = &Xml{Base{errorHandler}}
	default:
		panic(fmt.Sprintf("Format '%v' not supported", format))
	}
}
