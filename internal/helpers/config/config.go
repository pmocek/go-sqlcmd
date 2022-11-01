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
	encryptPassword bool,
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
		ContainerDetails: &ContainerDetails{
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

	user := User{
		AuthenticationType: "basic",
		BasicAuth: &BasicAuthDetails{
			Username: username,
			PasswordEncrypted: encryptPassword,
			Password: encryptCallback(password, encryptPassword),
		},
		Name: userName,
	}

	config.Users = append(config.Users, user)

	Save()
}

func GetRedactedConfig(raw bool) (c Sqlconfig) {
	c = config
	for i, _ := range c.Users {
		if raw {
			c.Users[i].BasicAuth.Password = decryptCallback(
				c.Users[i].BasicAuth.Password,
				c.Users[i].BasicAuth.PasswordEncrypted,
			)
		} else {
			c.Users[i].BasicAuth.Password = "REDACTED"
		}
	}

	return
}
