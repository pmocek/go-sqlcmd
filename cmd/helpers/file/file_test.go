package file

import (
	"fmt"
	"testing"
)

func TestGetwd(t *testing.T) {
	wd := getwd()

	fmt.Printf("The current working directory is: %s\n", wd)

	if wd == "" {
		t.Fatal("Getwd returns empty string")
	}
}

func TestMkdir(t *testing.T) {
	type args struct {
		folder string
	}
	var tests []struct {
		name string
		args args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
