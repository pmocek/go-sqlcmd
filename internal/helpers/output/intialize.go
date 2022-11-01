// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package output

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/internal/helpers/output/formatter"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output/verbosity"
)

// init initializes the package for unit testing.  For production, use
// the Initialize method to inject fully functional dependencies
func init() {
	Initialize(
		func(err error){if err != nil {panic(err)}},
		func(format string, a ...any){},
		func(hints []string){},
		"yaml",
		verbosity.Info)
}

func Initialize(
	errorHandler func(err error),
	traceHandler func(format string, a ...any),
	hintHandler func(hints []string),
	format string,
	verbosity verbosity.Enum,
) {
	errorCallback = errorHandler
	traceCallback = traceHandler
	hintCallback = hintHandler
	loggingLevel = verbosity

	trace("Initializing output as '%v'", format)

	switch format {
	case "json":
		formatter = &Json{Base: Base{ErrorHandlerCallback: errorHandler}}
	case "yaml":
		formatter = &Yaml{Base: Base{ErrorHandlerCallback: errorHandler}}
	case "xml":
		formatter = &Xml{Base: Base{ErrorHandlerCallback: errorHandler}}
	default:
		panic(fmt.Sprintf("Format '%v' not supported", format))
	}
}
