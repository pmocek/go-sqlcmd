// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package sqlconfig

type EndpointDetails struct {
	Address string `mapstructure:"address"`
	Port    int    `mapstructure:"port"`
}

type ContainerDetails struct {
	Id    string `mapstructure:"id"`
	Image string `mapstructure:"image"`
}

type Endpoint struct {
	ContainerDetails `mapstructure:"container" yaml:"container"`
	EndpointDetails  `mapstructure:"endpoint" yaml:"endpoint"`
	Name            string `mapstructure:"name"`
}

type ContextDetails struct {
	Endpoint string `mapstructure:"endpoint"`
	User     string `mapstructure:"user"`
}

type Context struct {
	ContextDetails `mapstructure:"context" yaml:"context"`
	Name           string `mapstructure:"name"`
}

type UserDetails struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type User struct {
	UserDetails `mapstructure:"user" yaml:"user"`
	Name        string `mapstructure:"name"`
}

type Sqlconfig struct {
	ApiVersion     string     `mapstructure:"apiVersion"`
	Endpoints      []Endpoint `mapstructure:"endpoints"`
	Contexts       []Context  `mapstructure:"contexts"`
	CurrentContext string     `mapstructure:"currentcontext"`
	Kind           string     `mapstructure:"kind"`
	Users          []User     `mapstructure:"users"`
}
