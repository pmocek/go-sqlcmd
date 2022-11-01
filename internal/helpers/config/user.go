// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import . "github.com/microsoft/go-sqlcmd/cmd/sqlconfig"

func AddUser(user User) {
	config.Users = append(config.Users, user)
	Save()
}

func DeleteUser(name string) {
	if UserExists(name) {
		ordinal := userOrdinal(name)
		config.Users = append(config.Users[:ordinal], config.Users[ordinal+1:]...)
		Save()
	}
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

func UserExists(name string) (exists bool) {
	for _, v := range config.Users {
		if name == v.Name {
			exists = true
			break
		}
	}
	return
}

func userOrdinal(name string) (ordinal int) {
	for i, c := range config.Users {
		if name == c.Name {
			ordinal = i
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
