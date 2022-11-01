package output

import (
	"github.com/microsoft/go-sqlcmd/internal/helpers/output/verbosity"
	"testing"
)

func TestTracef(t *testing.T) {
	type args struct {
		loggingLevel verbosity.Enum
		format string
		a      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{"default", args{
			loggingLevel: verbosity.Trace,
			format: "%v",
			a: []any{"sample trace"},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loggingLevel = tt.args.loggingLevel
			Tracef(tt.args.format, tt.args.a...)
			Trace(tt.args.a)
			Debugf(tt.args.format, tt.args.a...)
			Debug(tt.args.a)
			Infof(tt.args.format, tt.args.a...)
			Info(tt.args.a)
			Warnf(tt.args.format, tt.args.a...)
			Warn(tt.args.a)
			Errorf(tt.args.format, tt.args.a...)
			Error(tt.args.a)
			//Fatalf(tt.args.format, tt.args.a...)
			//Fatal(tt.args.a)

			InfofWithHints([]string{}, tt.args.format, tt.args.a...)
			InfofWithHintExamples([][]string{}, tt.args.format, tt.args.a...)

			//FatalfWithHints([]string{}, tt.args.format, tt.args.a...)
			//FatalfWithHintExamples([][]string{}, tt.args.format, tt.args.a...)

			Initialize(nil, Tracef, nil, "xml", verbosity.Error)
			Initialize(nil, Tracef, nil, "json", verbosity.Error)
		})
	}
}
