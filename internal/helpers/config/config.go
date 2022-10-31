// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
)

var config Sqlconfig

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
	userName := username + "@" + contextName

	config.ApiVersion = "v1"
	config.Kind = "Config"
	config.CurrentContext = contextName

	config.Endpoints = append(config.Endpoints, Endpoint{
		ContainerDetails: ContainerDetails{
			Id:    id,
			Image: imageName},
		EndpointDetails: EndpointDetails{
			Address: "localhost",
			Port:    portNumber,
		},
		Name: endPointName,
	})

	config.Contexts = append(config.Contexts, Context{
		ContextDetails: ContextDetails{
			Endpoint: endPointName,
			User:     username + "@" + contextName,
		},
		Name: contextName,
	})

	config.Users = append(config.Users, User{
		UserDetails: UserDetails{
			Username: username,
			Password: encryptCallback(password),
		},
		Name: userName,
	})

	Save()
}

func GetRedactedConfig(raw bool) (c Sqlconfig) {
	c = config
	for i, v := range c.Users {
		if raw {
			c.Users[i].UserDetails.Password = decryptCallback(v.UserDetails.Password)
		} else {
			c.Users[i].UserDetails.Password = "REDACTED"
		}
	}

	return
}

func OutputUsers(formatter func(interface{})) {
	formatter(config.Users)
}

func OutputEndpoints(formatter func(interface{})) {
	formatter(config.Endpoints)
}

func OutputContexts(formatter func(interface{})) {
	formatter(config.Contexts)
}
