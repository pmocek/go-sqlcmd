/*
 * // Copyright (c) Microsoft Corporation.
 * // Licensed under the MIT license.
 */

package sqlconfig

type EndpointDetails struct {
	Address string `mapstructure:"address"`
	Port int `mapstructure:"port"`
}

type DockerDetails struct {
	ContainerId string `mapstructure:"containerId"`
}

type Endpoint struct {
	DockerDetails   `mapstructure:",squash"`
	EndpointDetails `mapstructure:",squash"`
	Name            string          `mapstructure:"name"`
}

type ContextDetails struct {
	Endpoint string `mapstructure:"endpoint"`
	User string `mapstructure:"user"`
}

type Context struct {
	ContextDetails `mapstructure:",squash"`
	Name           string         `mapstructure:"name"`
}

type UserDetails struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type User struct {
	UserDetails `mapstructure:",squash"`
	Name        string      `mapstructure:"name"`

	// BUGBUG: Cannot get the nested struct (UserDetails) to
	// viper.Unmarshall (it comes back empty), putting password here
	// temporarily
	Password string `mapstructure:"password-temp"`
}

type Sqlconfig struct {
	ApiVersion string `mapstructure:"apiVersion"`
	Endpoints []Endpoint  `mapstructure:"endpoints"`
	Contexts []Context    `mapstructure:"contexts"`
	CurrentContext string `mapstructure:"current-context"`
	Kind string  `mapstructure:"kind"`
	Users []User `mapstructure:"users"`
}
