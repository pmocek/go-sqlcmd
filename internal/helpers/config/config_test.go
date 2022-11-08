package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/microsoft/go-sqlcmd/internal/helpers/output"
	"reflect"
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	type args struct {
		Config Sqlconfig
	}
	tests := []struct {
		name string
		args args
	}{
		{ "config",
			args{
				Config: Sqlconfig{
					Users: []User{{
						Name:               "user1",
						AuthenticationType: "basic",
						BasicAuth:          &BasicAuthDetails{
							Username:          "user",
							PasswordEncrypted: false,
							Password:          "weak",
						},
					}}}}},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config = tt.args.Config
			GetConfigFileUsed()
			GetRedactedConfig(false)
			GetRedactedConfig(true)

			AddEndpoint(Endpoint{
				ContainerDetails: &ContainerDetails{
					Id: strings.Repeat("9", 64),
					Image: "www.image.url"},
				EndpointDetails:  EndpointDetails{
					Address: "localhost",
					Port:    1433,
				},
				Name:             "endpoint",
			})

			AddEndpoint(Endpoint{
				EndpointDetails:  EndpointDetails{
					Address: "localhost",
					Port:    1434,
				},
				Name:             "endpoint",
			})

			EndpointsExists()
			EndpointExists("endpoint")
			GetEndpoint("endpoint")
			OutputEndpoints(output.Struct, true)
			OutputEndpoints(output.Struct, false)
			FindFreePortForTds()
			DeleteEndpoint("endpoint2")

			user := User{
				Name:               "user",
				AuthenticationType: "basic",
				BasicAuth:          &BasicAuthDetails{
					Username:          "username",
					PasswordEncrypted: false,
					Password:          "password",
				},
			}

			AddUser(user)
			AddUser(user)
			UserExists("user")
			GetUser("user")
			UserNameExists("username")
			OutputUsers(output.Struct, true)
			OutputUsers(output.Struct, false)
			DeleteUser("user")
			DeleteUser("user2")

			addContext()
			GetContext("context")
			OutputContexts(output.Struct, true)
			OutputContexts(output.Struct, false)
			DeleteContext("context", true)
			addContext()
			addContext()
			SetCurrentContextName("context")
			GetContainerId()
			GetCurrentContext()
			RemoveCurrentContext()
			RemoveCurrentContext()
			Update("containerId",
				"imageName",
				1433,
				"user",
				"password",
				false,
				"context")
		})
		}
}

func addContext() {
	user := "user1"
	AddContext(Context{
		ContextDetails: ContextDetails{
			Endpoint: "endpoint",
			User:     &user,
		},
		Name: "context",
	})
}

func TestDeleteUser(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteUser(tt.args.name)
		})
	}
}

func TestFindUniqueUserName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name               string
		args               args
		wantUniqueUserName string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUniqueUserName := FindUniqueUserName(tt.args.name); gotUniqueUserName != tt.wantUniqueUserName {
				t.Errorf("FindUniqueUserName() = %v, want %v", gotUniqueUserName, tt.wantUniqueUserName)
			}
		})
	}
}

func TestGetUser(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name     string
		args     args
		wantUser User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotUser := GetUser(tt.args.name); !reflect.DeepEqual(gotUser, tt.wantUser) {
				t.Errorf("GetUser() = %v, want %v", gotUser, tt.wantUser)
			}
		})
	}
}

func TestOutputUsers(t *testing.T) {
	type args struct {
		formatter func(interface{})
		detailed  bool
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			OutputUsers(tt.args.formatter, tt.args.detailed)
		})
	}
}

func TestUserExists(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExists := UserExists(tt.args.name); gotExists != tt.wantExists {
				t.Errorf("UserExists() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func TestUserNameExists(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotExists := UserNameExists(tt.args.name); gotExists != tt.wantExists {
				t.Errorf("UserNameExists() = %v, want %v", gotExists, tt.wantExists)
			}
		})
	}
}

func Test_userOrdinal(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		wantOrdinal int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOrdinal := userOrdinal(tt.args.name); gotOrdinal != tt.wantOrdinal {
				t.Errorf("userOrdinal() = %v, want %v", gotOrdinal, tt.wantOrdinal)
			}
		})
	}
}
