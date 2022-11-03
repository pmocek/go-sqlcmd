package cmd

import (
	"errors"
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	"github.com/microsoft/go-sqlcmd/cmd/root"
	"github.com/microsoft/go-sqlcmd/internal/helpers"
	"github.com/microsoft/go-sqlcmd/internal/helpers/config"
	"os"
	"strings"
	"testing"
)

var offlineMode = false

func TestRunCommandLine(t *testing.T) {
	os.Setenv("SQLCMD_PASSWORD", "badpass")

	useCached := "--cached"
	if !offlineMode {
		useCached = ""
	}

	helpers.Initialize(
		func(err error) {if err != nil {panic(err)}},
		displayHints,
		"sqlconfig-test-cmd-line",
		"yaml",
		4,
	)

	config.Clean()

	type args struct {
		args []string
	}
	split := func(cmd string) (args) {
		return args{strings.Split(cmd, " ")}
	}
	tests := []struct {
		name string
		args args
	}{
		{"default",
			split("--help")},
		{"neg-config-use-context-double-name",
			split("config use-context badbad --name andbad")},
		{"neg-config-use-context-bad-name",
			split("config use-context badbad")},
		{"neg-config-get-contexts-bad-context",
			split("config get-contexts badbad")},
		{"neg-config-get-endpoints-bad-endpoint",
			split("config get-endpoints badbad")},
		{"coverDefaultSubCommand",
			split("install mssql")},
		{"neg-config-add-context-no-endpoint",
			split("config add-context")},
		{"config-add-endpoint",
			split("config add-endpoint --address localhost --port 1433")},
		{"config-add-endpoint",
			split("config add-endpoint --address localhost --port 1433")},
		{"neg-config-add-context-bad-user",
			split("config add-context --endpoint endpoint --user badbad")},
		{"config-get-endpoints",
			split("config get-endpoints endpoint")},
		{"config-get-endpoints",
			split("config get-endpoints")},
		{"config-get-endpoints",
			split("config get-endpoints --detailed")},
		{"config-add-context",
			split("config add-context --endpoint endpoint")},
		{"neg-uninstall-but-context-has-no-container",
			split("uninstall --force --yes")},
		{"config-use-context",
			split("config use-context context")},
		{"config-get-contexts",
			split("config get-contexts context")},
		{"config-get-contexts",
			split("config get-contexts")},
		{"config-get-contexts",
			split("config get-contexts --detailed")},
		{"config-delete-context",
			split("config delete-context context --cascade")},
		{"neg-config-delete-context",
			split("config delete-context")},
		{"neg-config-delete-context",
			split("config delete-context badbad-name")},
		{"neg-config-get-users-bad-user",
			split("config get-users badbad")},
		{"config-add-user",
			split("config add-user --username foobar")},
		{"config-delete-endpoint",
			split("config delete-endpoint endpoint2")},
		{"neg-config-delete-endpoint-no-name",
			split("config delete-endpoint")},
		{"config-add-endpoint",
			split("config add-endpoint --address localhost --port 1433")},
		{"config-add-context",
			split("config add-context --user user --endpoint endpoint --name my-context")},
		{"config-delete-context-cascade",
			split("config delete-context my-context --cascade")},
		{"config-add-user",
			split("config add-user --username foobar")},
		{"config-get-users",
			split("config get-users user")},
		{"config-get-users",
			split("config get-users")},
		{"config-get-users",
			split("config get-users --detailed")},
		{"neg-config-add-user-no-username",
			split("config add-user")},
		{"neg-config-add-user-no-password",
			split("config add-user --username foobar")},
		{"config-view",
			split("config view")},
		{"config-view",
			split("config view --raw")},
		{"config-delete-user",
			split("config delete-user user")},

		{"neg-config-add-user-bad-auth-type",
			split("config add-user --username foobar --auth-type badbad")},
		{"neg-config-add-user-bad-use-encrypted",
			split("config add-user --username foobar  --auth-type other --encrypt-password")},

		{"get-tags",
			split("install mssql get-tags")},
		{"neg-install-no-eula",
			split("install mssql server")},
		{"install",
			split(fmt.Sprintf("install mssql server %v --user-database my-database --accept-eula", useCached))},
		{"config-current-context",
			split("config current-context")},
		{"config-connection-strings",
			split("config connection-strings")},
		{"query",
			split("query GO")},
		{"query",
			split("query")},
		{"neg-query-two-queries",
			split("query bad --query bad")},

		/* How to get code coverage for user input
		{"neg-uninstall-no-yes",
			split("uninstall")},*/
		{"uninstall",
			split("uninstall --yes --force")},

	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Root{AbstractBase{SubCommands: root.Commands}}
			rootCmd = r.DefineCommand()

			rootCmd.SetArgs(tt.args.args)

			t.Logf("Running: %v", tt.args.args)

			if tt.name == "neg-config-add-user-no-password" {
				os.Setenv("SQLCMD_PASSWORD", "")
			}

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
