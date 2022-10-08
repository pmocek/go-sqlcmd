// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"bytes"
	"fmt"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/file"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/folder"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/secret"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var config Sqlconfig

// Initialize reads in config file and ENV variables if set.
func Initialize(configFile string, handler errorHandlerService) {
	if handler == nil {
		panic("Please provide an error handler")
	}
	errorHandlerCallback = handler

	if configFile == "" {
		home, err := os.UserHomeDir()
		checkErr(err)

		configFile = filepath.Join(home, ".sqlcmd", "sqlconfig")
	}

	createEmptyFileIfNotExists(configFile)

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configFile)
	err := viper.ReadInConfig()
	checkErr(err)
	viper.AutomaticEnv() // read in environment variables that match
	err = viper.Unmarshal(&config)
	checkErr(err)

	fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
}

func createEmptyFileIfNotExists(filename string) {
	file := file.GetInstance()
	if !file.Exists(filename) {
		folder := folder.GetInstance()
		folder.MkdirAll(filepath.Base(filename))
		f, err := os.Create(filename)
		defer f.Close()
		checkErr(err)
	}
}

func Update(
	id string,
	imageName string,
	portNumber int,
	password string,
	contextName string,
) {
	if contextName == "" {
		panic("contextName must be provided")
	}

	contextName = FindUniqueContextName(contextName)
	endPointName := FindUniqueEndpointName(contextName)
	userName := "sa@" + contextName // This is the name of the user config entry, not the mssql user 'sa'

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
			User:     "sa@" + contextName,
		},
		Name:           contextName,
	})

	config.Users = append(config.Users, User{
		UserDetails: UserDetails{
			Username: "sa",
			Password: secret.Encrypt(password),
		},
		Name:        userName,
	})

	Save()
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
			if isLocalPortAvailable(portNumber) {
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

func EndpointNameExists(name string) (exists bool) {
	for _, v := range config.Endpoints {
		if v.Name == name {
			exists = true
			break
		}
	}

	return
}

func ContextNameExists(name string) (exists bool) {
	for _, v := range config.Contexts {
		if v.Name == name {
			exists = true
			break
		}
	}

	return
}

func UserNameExists(name string) (exists bool) {
	for _, v := range config.Users {
		if v.Name == name {
			exists = true
			break
		}
	}

	return
}

func FindUniqueEndpointName(contextName string) (uniqueEndpointName string) {
	var postfixNumber = 1

	for {
		uniqueEndpointName = fmt.Sprintf(
			"%v%v",
			contextName,
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

	return
}

// FindUniqueContextName finds a unique context name, that is both a
// unique context name, but also a unique sa@context name
func FindUniqueContextName(name string) (uniqueContextName string) {
	var postfixNumber = 1

	for {
		uniqueContextName = fmt.Sprintf(
			"%v%v",
			name,
			strconv.Itoa(postfixNumber),
		)
		if !ContextNameExists(uniqueContextName) {
			if !UserNameExists("sa@" + uniqueContextName) {
				break
			}
		} else {
			postfixNumber++
		}

		if postfixNumber == 5000 {
			panic("Did not an available context name")
		}
	}

	return
}

func Save() {
	b, err := yaml.Marshal(&config)
	checkErr(err)
	err = viper.ReadConfig(bytes.NewReader(b))
	checkErr(err)

	createEmptyFileIfNotExists(viper.ConfigFileUsed())

	err = viper.WriteConfig()
	checkErr(err)
}

func GetContainerId() (containerId string) {
	currentContextName := config.CurrentContext

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					containerId = e.DockerDetails.ContainerId
					return
				}
			}
		}
	}
	panic("ContainerId not found")
}

func GetContainerShortId() (containerId string) {
	containerId =  GetContainerId()
	containerId = containerId[len(containerId)-12:]

	return
}

func RemoveCurrentContext() {
	currentContextName := config.CurrentContext

	for i, c := range config.Contexts {
		if c.Name == currentContextName {
			for ei, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					config.Endpoints = append(
						config.Endpoints[:ei],
						config.Endpoints[ei+1:]...)
					break
				}
			}

			for ii, u := range config.Users {
				if u.Name == c.User {
					config.Users = append(
						config.Users[:ii],
						config.Users[ii+1:]...)
					break
				}
			}

			config.Contexts = append(
				config.Contexts[:i],
				config.Contexts[i+1:]...)
			break
		}
	}

	if len(config.Contexts) > 0 {
		config.CurrentContext = config.Contexts[0].Name
	} else {
		config.CurrentContext = ""
	}

	return
}

func GetCurrentContextName() (name string) {
	name = config.CurrentContext

	return
}

func EndpointsExists() (exists bool) {
	if len(config.Endpoints) > 0 {
		exists = true
	}

	return
}

func GetCurrentContext() (endpoint Endpoint, user User){
	currentContextName := config.CurrentContext

	for _, c := range config.Contexts {
		if c.Name == currentContextName {
			for _, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					endpoint = e
					break
				}
			}

			for _, u := range config.Users {
				if u.Name == c.User {
					user = u
					break
				}
			}
		}
	}
	return
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

func isLocalPortAvailable(port int) (portAvailable bool) {
	timeout := time.Second
	conn, err := net.DialTimeout(
		"tcp",
		net.JoinHostPort("localhost", strconv.Itoa(port)),
		timeout,
	)
	if err != nil {
		fmt.Println("Connecting error:", err)
		portAvailable = true
	}
	if conn != nil {
		conn.Close()
		fmt.Println("Opened", net.JoinHostPort("localhost", strconv.Itoa(port)))
	}

	return
}
