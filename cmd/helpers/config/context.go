// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import (
	"errors"
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"strconv"
)

// FindUniqueContextName finds a unique context name, that is both a
// unique context name, but also a unique sa@context name.  If the name passed
// in is unique then this is returned, else we look for the name with a numeral
// postfix, starting at 2
func FindUniqueContextName(name string, username string) (uniqueContextName string) {
	if !ContextExists(name) &&
		!UserNameExists(username + "@" + name) {
		uniqueContextName = name
	} else {
		var postfixNumber = 2
		for {
			uniqueContextName = fmt.Sprintf(
				"%v%v",
				name,
				strconv.Itoa(postfixNumber),
			)
			if !ContextExists(uniqueContextName) {
				if !UserNameExists(username + "@" + uniqueContextName) {
					break
				}
			} else {
				postfixNumber++
			}

			if postfixNumber == 5000 {
				panic("Did not an available context name")
			}
		}
	}

	return
}

func GetCurrentContextName() (name string) {
	name = config.CurrentContext

	return
}

func GetCurrentContextOrFatal() (currentContextName string) {
	currentContextName = GetCurrentContextName()
	if currentContextName == "" {
		checkErr(errors.New("No current context. To create a context use `sqlcmd install`, e.g. `sqlcmd install mssql`"))
	}
	return
}

func SetCurrentContextName(name string) {
	if ContextExists(name) {
		config.CurrentContext = name
		Save()
	}

	return
}

func RemoveCurrentContext() {
	currentContextName := config.CurrentContext

	for ci, c := range config.Contexts {
		if c.Name == currentContextName {
			for ei, e := range config.Endpoints {
				if e.Name == c.Endpoint {
					config.Endpoints = append(
						config.Endpoints[:ei],
						config.Endpoints[ei+1:]...)
					break
				}
			}

			for ui, u := range config.Users {
				if u.Name == c.User {
					config.Users = append(
						config.Users[:ui],
						config.Users[ui+1:]...)
					break
				}
			}

			config.Contexts = append(
				config.Contexts[:ci],
				config.Contexts[ci+1:]...)
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

func ContextExists(name string) (exists bool) {
	for _, c := range config.Contexts {
		if name == c.Name {
			exists = true
			break
		}
	}
	return
}

func GetCurrentContext() (endpoint Endpoint, user User){
	currentContextName := GetCurrentContextOrFatal()

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

func GetContext(name string) (context Context) {
	for _, c := range config.Contexts {
		if name == c.Name {
			context = c
			break
		}
	}
	return
}
