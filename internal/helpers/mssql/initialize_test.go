package mssql

import (
	"strings"
	"testing"
)

func TestInitialize(t *testing.T) {
	type args struct {
		errorHandler func(err error)
		traceHandler func(format string, a ...any)
		decryptHandler func(cipherText string) (secret string)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "noErrorHandlerPanic",
			args: args{
				nil,
				func(format string, a ...any) {},
				func(cipherText string) (secret string){return},
			}},
		{
			name: "noTraceHandlerPanic",
			args: args{
				func(err error){},
				nil,
				func(cipherText string) (secret string) {return},
			}},
		{
			name: "noDecryptHandlerPanic",
			args: args{
				func(err error){},
				func(format string, a ...any) {},
				nil,
			}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// If test name ends in 'Panic' expect a Panic
			if strings.HasSuffix(tt.name, "Panic") {
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("The code did not panic")
					}
				}()
			}

			Initialize(
				tt.args.errorHandler,
				tt.args.traceHandler,
				tt.args.decryptHandler,
			)
		})
	}
}
