package mssql

import (
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"testing"
)

func TestConnect(t *testing.T) {

	type args struct {
		endpoint Endpoint
		user     *User
		console  sqlcmd.Console
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "connect", 
			args: args{endpoint: Endpoint{
				EndpointDetails:  EndpointDetails{
					Address: "localhost",
					Port:    1433,
				},
				Name:             "local-default-instance",
			}, user: nil, console:  nil},
			want: 0,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Connect(tt.args.endpoint, tt.args.user, tt.args.console); got.Exitcode != tt.want {
				t.Errorf("ExitCode = %v, want %v", got.Exitcode, tt.want)
			}
		})
	}
}

func TestQuery(t *testing.T) {
	type args struct {
		s    *sqlcmd.Sqlcmd
		text string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Query(tt.args.s, tt.args.text)
		})
	}
}
