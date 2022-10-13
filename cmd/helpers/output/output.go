// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

/*
	output.Trace("Something very low level.")
	output.Debug("Useful debugging information.")
	output.Info("Something noteworthy happened!")
	output.Warn("You should probably take a look at this.")
	output.Error("Something failed but I'm not quitting.")
	// Calls os.Exit(1) after logging
	output.Fatal("Bye.")
	// Calls panic() after logging
	output.Panic("I'm bailing.")
*/

package output

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/helpers/output/formatter"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output/verbosity"
	"github.com/pkg/errors"
)

var formatter Formatter
var loggingLevel verbosity.Enum

func Struct(in interface{}) {
	if formatter == nil {
		panic("formatter not initialized")
	}

	formatter.Serialize(in)
}

func Tracef(format string, a ...any) {
	if loggingLevel >= verbosity.Trace {
		format = ensureEol(format)
		_, err := fmt.Printf("TRACE: ")
		checkErr(err)
		_, err = fmt.Printf(format, a...)
		checkErr(err)
	}
}

func ensureEol(format string) string {
	if format[len(format)-1] != '\n' {
		format = format + "\n"
	}
	return format
}

func Trace(a ...any) {
	if loggingLevel >= verbosity.Trace {
		_, err := fmt.Printf("TRACE: ")
		checkErr(err)
		_, err = fmt.Println(a...)
		checkErr(err)
	}
}

func Debugf(format string, a ...any) {
	if loggingLevel >= verbosity.Debug {
		format = ensureEol(format)
		_, err := fmt.Printf("DEBUG: ")
		checkErr(err)
		_, err = fmt.Printf(format, a...)
		checkErr(err)
	}
}

func Debug(a ...any) {
	if loggingLevel >= verbosity.Debug {
		_, err := fmt.Printf("DEBUG: ")
		checkErr(err)
		_, err = fmt.Println(a...)
		checkErr(err)
	}
}

func Infof(format string, a ...any) {
	infofWithHints([]string{}, format, a...)
}

func InfofWithHints(hints []string, format string, a ...any) {
	infofWithHints(hints, format, a...)
}

func infofWithHints(hints []string, format string, a ...any) {
	if loggingLevel >= verbosity.Info {
		format = ensureEol(format)
		if loggingLevel >= verbosity.Debug {
			_, err := fmt.Printf("INFO:  ")
			checkErr(err)
		}
		_, err := fmt.Printf(format, a...)
		checkErr(err)
		displayHints(hints)
	}
}

func Info(a ...any) {
	if loggingLevel >= verbosity.Info {
		if loggingLevel >= verbosity.Debug {
			_, err := fmt.Printf("INFO:  ")
			checkErr(err)
		}
		_, err := fmt.Println(a...)
		checkErr(err)
	}
}

func Warnf(format string, a ...any) {
	if loggingLevel >= verbosity.Warn {
		format = ensureEol(format)
		if loggingLevel >= verbosity.Debug {
			_, err := fmt.Printf("WARN:  ")
			checkErr(err)
		}
		_, err := fmt.Printf(format, a...)
		checkErr(err)
	}
}

func Warn(a ...any) {
	if loggingLevel >= verbosity.Warn {
		if loggingLevel >= verbosity.Debug {
			_, err := fmt.Printf("WARN:  ")
			checkErr(err)
		}
		_, err := fmt.Println(a...)
		checkErr(err)
	}
}

func Errorf(format string, a ...any) {
	if loggingLevel >= verbosity.Error {
		format = ensureEol(format)
		if loggingLevel >= verbosity.Debug {
			_, err := fmt.Printf("ERROR:  ")
			checkErr(err)
		}
		_, err := fmt.Printf(format, a...)
		checkErr(err)
	}
}

func Error(a ...any) {
	if loggingLevel >= verbosity.Error {
		if loggingLevel >= verbosity.Debug {
			_, err := fmt.Printf("ERROR:  ")
			checkErr(err)
		}
		_, err := fmt.Println(a...)
		checkErr(err)
	}
}

func Fatal(a ...any) {
	fatal([]string{}, a...)
}

func FatalWithHints(hints []string, a ...any) {
	fatal(hints, a...)
}

func fatal(hints []string, a ...any) {
	err := errors.New(fmt.Sprintf("%v", a...))
	displayHints(hints)
	checkErr(err)
}

func Fatalf(format string, a ...any) {
	fatalf([]string{}, format, a...)
}

func FatalErr(err error) {
	checkErr(err)
}

func FatalfWithHints(hints []string, format string, a ...any) {
	fatalf(hints, format, a...)
}

func FatalfErrorWithHints(err error, hints []string, format string, a ...any) {
	fatalf(hints, format, a...)
	checkErr(err)
}

func fatalf(hints []string, format string, a ...any) {
	err := errors.New(fmt.Sprintf(format, a...))
	displayHints(hints)
	checkErr(err)
}

func Panicf(format string, a ...any) {
	panic(fmt.Sprintf(format, a...))
}

func Panic(a ...any) {
	panic(a)
}
