package cmd

import (
	"errors"
	"testing"
)

func TestRunCommandLine(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args
	}{
		{ "default", args{[]string{"--help"}}},
		{ "coverDefaultSubCommand", args{[]string{"install", "mssql"}}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(tt.args.args)

			if tt.name == "default" {
				RunCommandLine()
			}
		})
	}
}

func Test_initializeCobra(t *testing.T) {
	initializeCobra()
}

func Test_displayHints(t *testing.T) {
	displayHints([]string{"Test Hint"})
}

func TestIsValidRootCommand(t *testing.T) {
	IsValidRootCommand("install")
	IsValidRootCommand("create")
	IsValidRootCommand("nope")
}

func Test_checkErr(t *testing.T) {
	loggingLevel = 3

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	checkErr(errors.New("Expected error"))
}
