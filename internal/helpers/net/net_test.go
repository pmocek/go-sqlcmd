package net

import (
	"testing"
)

func TestIsLocalPortAvailable(t *testing.T) {
	type args struct {
		port int
	}
	tests := []struct {
		name              string
		args              args
		wantPortAvailable bool
	}{
		{name: "expectedToNotBeAvailable", args: args{port: 80}, wantPortAvailable: false},
		{name: "expectedToBeAvailable", args: args{port: 9999}, wantPortAvailable: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotPortAvailable := IsLocalPortAvailable(tt.args.port); gotPortAvailable != tt.wantPortAvailable {
				t.Errorf("IsLocalPortAvailable() = %v, want %v", gotPortAvailable, tt.wantPortAvailable)
			}
		})
	}
}
