// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"fmt"
	"github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"strconv"
)

func AddEndpoint(endpoint sqlconfig.Endpoint) {
	config.Endpoints = append(config.Endpoints, endpoint)
	Save()
}

func EndpointsExists() (exists bool) {
	if len(config.Endpoints) > 0 {
		exists = true
	}

	return
}

func EndpointNameExists(name string) (exists bool) {
	for _, v := range config.Endpoints {
		if v.Name == name {
			exists = true
			break
		}
	}

	return
}

func FindUniqueEndpointName(name string) (uniqueEndpointName string) {
	if !EndpointNameExists(name) {
		uniqueEndpointName = name
	} else {
		var postfixNumber = 2

		for {
			uniqueEndpointName = fmt.Sprintf(
				"%v%v",
				name,
				strconv.Itoa(postfixNumber),
			)
			if !EndpointNameExists(uniqueEndpointName) {
				break
			} else {
				postfixNumber++
			}
			if postfixNumber == 5000 {
				panic("Did not find an available endpoint name")
			}
		}
	}

	return
}

func GetContainerId() (containerId string) {
	currentContextName := config.CurrentContext

	if currentContextName == "" {
		panic("currentContextName must not be empty")
	}

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					containerId = e.ContainerDetails.Id

					if len(containerId) != 64 {
						panic("container id must be 64 characters")
					}

					return
				}
			}
		}
	}
	panic("Id not found")
}

func FindFreePortForTds() (portNumber int) {
	const startingPortNumber = 1433

	portNumber = startingPortNumber

	for {
		foundFreePortNumber := true
		for _, endpoint := range config.Endpoints {
			if endpoint.Port == portNumber {
				foundFreePortNumber = false
				break
			}
		}

		if foundFreePortNumber == true {
			// Check this port is actually available on the local machine
			if isLocalPortAvailableCallback(portNumber) {
				break
			} else {
				foundFreePortNumber = false
			}
		}

		portNumber++

		if portNumber == 5000 {
			panic("Did not find an available port")
		}
	}

	return
}

func EndpointExists(name string) (exists bool) {
	for _, c := range config.Endpoints {
		if name == c.Name {
			exists = true
			break
		}
	}
	return
}

func GetEndpoint(name string) (endpoint sqlconfig.Endpoint) {
	for _, e := range config.Endpoints {
		if name == e.Name {
			endpoint = e
			break
		}
	}
	return
}
