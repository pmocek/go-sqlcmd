package cmd

import (
	"errors"
	"github.com/microsoft/go-sqlcmd/internal/helpers"
	"os"
	"strings"
	"testing"
)

func TestRunCommandLine(t *testing.T) {
	os.Setenv("SQLCMD_PASSWORD", "badpass")

	helpers.Initialize(
		func(err error) {if err != nil {panic(err)}},
		displayHints,
		"sqlconfig-test",
		"yaml",
		4,
	)

	cmds := [][]string{
		{"neg-config-add-context-no-endpoint", "config add-context"},
		{"neg-config-add-context-bad-user", "config add-context --user badbad"},
		{"config-add-endpoint", "config add-endpoint --address localhost --port 1433"},
		{"config-add-endpoint", "config add-endpoint --address localhost --port 1433"},
		{"config-get-endpoints", "config get-endpoints endpoint"},
		{"config-get-endpoints", "config get-endpoints"},
		{"config-get-endpoints", "config get-endpoints --detailed"},
		{"config-add-context", "config add-context --endpoint endpoint"},
		{"config-use-context", "config use-context context"},
		{"config-get-contexts", "config get-contexts context"},
		{"config-get-contexts", "config get-contexts"},
		{"config-get-contexts", "config get-contexts --detailed"},
		{"config-delete-context", "config delete-context context --cascade"},
		{"config-delete-endpoint", "config delete-endpoint endpoint2"},
		{"config-add-user", "config add-user --username foobar"},
		{"config-get-users", "config get-users user"},
		{"config-get-users", "config get-users"},
		{"config-get-users", "config get-users --detailed"},
		{"config-view", "config view"},
		{"config-view", "config view --raw"},
		{"config-delete-user", "config delete-user user"},

		{"neg-config-add-user-bad-auth-type", "config add-user --username foobar --auth-type badbad"},
		{"neg-config-add-user-no-username", "config add-user"},
		{"neg-config-add-user-bad-use-encrypted", "config add-user --type other --encrypted"},

		{"install", "install mssql server --cached --accept-eula"},
		{"config-current-context", "config current-context"},
		{"config-connection-strings", "config connection-strings"},
		{"uninstall", "uninstall --yes"},
	}
	type args struct {
		args []string
	}
	type test struct {
		name string
		args
	}

	tests := []test{
		{"default", args{[]string{"--help"}}},
		{"coverDefaultSubCommand", args{[]string{"install", "mssql"}}},
	}

	for _, cmd := range cmds {
		if len(cmd) != 2 {
			panic("must be length 2")
		}
		tests = append(tests, test{cmd[0], args{strings.Split(cmd[1], " ")}})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootCmd.SetArgs(tt.args.args)

			t.Logf("Running: %v", tt.args.args)

			if tt.name != "coverDefaultSubCommand" {
				// If test name starts with 'neg-' expect a Panic
				if strings.HasPrefix(tt.name, "neg-") {
					defer func() {
						if r := recover(); r == nil {
							t.Errorf("The code did not panic")
						}
					}()

					RunCommandLine(true)

				}

				RunCommandLine(false)
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
