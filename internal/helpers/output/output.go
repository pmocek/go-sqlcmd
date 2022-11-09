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
	. "github.com/microsoft/go-sqlcmd/internal/helpers/output/formatter"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output/verbosity"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"github.com/pkg/errors"
	"io"
	"regexp"
	"strings"
)

var formatter Formatter
var loggingLevel verbosity.Enum
var runningUnitTests bool

var standardWriteCloser io.WriteCloser
var errorWriteCloser io.WriteCloser

func Struct(in interface{}) (bytes []byte) {
	bytes = formatter.Serialize(in)

	return
}

func Tracef(format string, a ...any) {
	if loggingLevel >= verbosity.Trace {
		format = ensureEol(format)
		printf("TRACE: " + format, a...)
	}
}

func printf(format string, a ...any) {
	text := fmt.Sprintf(format, a...)
	text = maskSecrets(text)
	_, err := standardWriteCloser.Write([]byte(text))
	checkErr(err)
}

func maskSecrets(text string) string {

	// Mask password from T/SQL e.g. ALTER LOGIN [sa] WITH PASSWORD = N'foo';
	r := regexp.MustCompile("(PASSWORD.*\\s?=.*\\s?N?')(.*)(')")
	text = r.ReplaceAllString(text, "$1********$3")
	return text
}

func ensureEol(format string) string {
	if len(format) >= len(sqlcmd.SqlcmdEol) {
		if !strings.HasSuffix(format, sqlcmd.SqlcmdEol) {
			format = format + sqlcmd.SqlcmdEol
		}
	} else {
		format = sqlcmd.SqlcmdEol
	}
	return format
}

func Debugf(format string, a ...any) {
	if loggingLevel >= verbosity.Debug {
		format = ensureEol(format)
		printf("DEBUG: " + format, a...)
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
			format = "INFO:  " + format
		}
		printf(format, a...)
		displayHints(hints)
	}
}

func InfofWithHintExamples(hintExamples [][]string, format string, a ...any) {
	if loggingLevel >= verbosity.Info || runningUnitTests {
		format = ensureEol(format)
		if loggingLevel >= verbosity.Debug {
			format = "INFO:  " + format
		}
		printf(format, a...)
		displayHintExamples(hintExamples)
	}
}

func displayHintExamples(hintExamples [][]string) {
	var hints []string

	maxLengthHintText := 0
	for _, hintExample := range hintExamples {
		if len(hintExample) != 2 {
			panic("Hintexample must be 2 elements, a description, and an example")
		}

		if len(hintExample[0]) > maxLengthHintText {
			maxLengthHintText = len(hintExample[0])
		}
	}

	for _, hintExample := range hintExamples {
		padLength := maxLengthHintText - len(hintExample[0])
		hints = append(hints, fmt.Sprintf(
			"%v: %v%s",
			hintExample[0],
			strings.Repeat(" ", padLength),
			hintExample[1],
		))
	}
	displayHints(hints)
}

func Warnf(format string, a ...any) {
	if loggingLevel >= verbosity.Warn {
		format = ensureEol(format)
		if loggingLevel >= verbosity.Debug {
			format = "WARN:  " + format
		}
		printf(format, a...)
	}
}

func Errorf(format string, a ...any) {
	if loggingLevel >= verbosity.Error {
		format = ensureEol(format)
		if loggingLevel >= verbosity.Debug {
			format = "ERROR: " + format
		}
		printf(format, a...)
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

func FatalfWithHintExamples(hintExamples [][]string, format string, a ...any) {
	err := errors.New(fmt.Sprintf(format, a...))
	displayHintExamples(hintExamples)
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
