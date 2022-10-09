package config

import (
	"fmt"
	. "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"
	"strconv"
)

func ContextNameExists(name string) (exists bool) {
	for _, v := range config.Contexts {
		if v.Name == name {
			exists = true
			break
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

func SetCurrentContextName(name string) {
	if ContextExists(name) {
		config.CurrentContext = name
		Save()
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

func GetContext(name string) (context Context) {
	for _, c := range config.Contexts {
		if name == c.Name {
			context = c
			break
		}
	}
	return
}
