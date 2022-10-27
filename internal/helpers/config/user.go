// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import . "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"

func UserNameExists(name string) (exists bool) {
	for _, v := range config.Users {
		if v.Name == name {
			exists = true
			break
		}
	}

	return
}

func UserExists(name string) (exists bool) {
	for _, v := range config.Users {
		if name == v.Name {
			exists = true
			break
		}
	}
	return
}

func GetUser(name string) (user User) {
	for _, v := range config.Users {
		if name == v.Name {
			user = v
			break
		}
	}
	return
}
