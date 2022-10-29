package mssql

import (
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"github.com/microsoft/go-sqlcmd/pkg/sqlcmd"
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {

	type args struct {
		endpoint Endpoint
		user     User
		console  sqlcmd.Console
	}
	endpoint, user := config.GetCurrentContext()

	if endpoint.Name == "" || user.Name == "" {
		panic("Ensure there is a current context, by installing mssql")
	}

	tests := []struct {
		name string
		args args
		want *sqlcmd.Sqlcmd
	}{
		{
			name: "connect", 
			args: args{endpoint: endpoint, user: user, console:  nil},
			want: nil,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Connect(tt.args.endpoint, tt.args.user, tt.args.console); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Connect() = %v, want %v", got, tt.want)
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
