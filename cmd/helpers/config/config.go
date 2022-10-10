// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
)

var config Sqlconfig

func Initialize(handler func(err error), configFile string) {
	if handler == nil {
		panic("Please provide an error handler")
	}
	errorHandlerCallback = handler
	configureViper(configFile)
	load()
}

func Update(
	id string,
	imageName string,
	portNumber int,
	username string,
	password string,
	contextName string,
) {
	if id == "" {
		panic("id must be provided")
	}
	if imageName == "" {
		panic("imageName must be provided")
	}
	if portNumber == 0 {
		panic("portNumber must be non-zero")
	}
	if password == "" {
		panic("password must be provided")
	}
	if contextName == "" {
		panic("contextName must be provided")
	}

	contextName = FindUniqueContextName(contextName, username)
	endPointName := FindUniqueEndpointName(contextName)
	userName := username + "@" + contextName // This is the name of the user config entry, not the mssql user 'sa'

	config.ApiVersion = "v1"
	config.Kind = "Config"
	config.CurrentContext =  contextName

	config.Endpoints = append(config.Endpoints, Endpoint{
		DockerDetails:   DockerDetails{
			ContainerId: id,
			Image: imageName},
		EndpointDetails: EndpointDetails{
			Address: "localhost",
			Port:    portNumber,
		},
		Name:            endPointName,
	})

	config.Contexts = append(config.Contexts, Context{
		ContextDetails: ContextDetails{
			Endpoint: endPointName,
			User:    username + "@" + contextName,
		},
		Name:           contextName,
	})

	config.Users = append(config.Users, User{
		UserDetails: UserDetails{
			Username: username,
			Password: secret.Encrypt(password),
		},
		Name:        userName,
	})

	Save()
}

func GetRedactedConfig(raw bool) (c Sqlconfig) {
	c = config
	for i, v := range c.Users {
		if raw {
			c.Users[i].UserDetails.Password = secret.Decrypt(
				v.UserDetails.Password)
		} else {
			c.Users[i].UserDetails.Password = "REDACTED"
		}
	}

	return
}

func GetConfig() Sqlconfig {
	return config
}

func GetUsers() []User {
	return config.Users
}

func GetEndpoints() []Endpoint {
	return config.Endpoints
}
