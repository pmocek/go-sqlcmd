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
	*ContainerDetails `mapstructure:"container,omitempty" yaml:"container,omitempty"`
	EndpointDetails  `mapstructure:"endpoint" yaml:"endpoint"`
	Name            string `mapstructure:"name"`
}

type ContextDetails struct {
	Endpoint string `mapstructure:"endpoint"`
	User     *string `mapstructure:"user,omitempty"`
}

type Context struct {
	ContextDetails `mapstructure:"context" yaml:"context"`
	Name           string `mapstructure:"name"`
}

type BasicAuthDetails struct {
	Username string `mapstructure:"username"`
	PasswordEncrypted bool `mapstructure:"password-encrypted" yaml:"password-encrypted"`
	Password string `mapstructure:"password"`
}

// Use same names that the Azure SDK for GO uses
// Add Syntantaic sugar to command line that generates the parameters
//   -  Azure Active Directory
//   0.   Default
//   1.   Service Principal
//   2.   Managed Identiy
//   3.   Username and Password
//   4.   Interactive
//   5.   Device code
type OtherAuthDetails struct {
	Parameters map[string]string  `mapstructure:"parameters"`
}

type User struct {
	Name        string `mapstructure:"name"`
	AuthenticationType string `mapstructure:"authentication-type" yaml:"authentication-type"`
	BasicAuth* BasicAuthDetails `mapstructure:"basic-auth,omitempty" yaml:"basic-auth,omitempty"`
	OtherAuth* OtherAuthDetails `mapstructure:"other-auth,omitempty" yaml:"other-auth,omitempty"`
}

type Sqlconfig struct {
	ApiVersion     string     `mapstructure:"apiVersion"`
	Endpoints      []Endpoint `mapstructure:"endpoints"`
	Contexts       []Context  `mapstructure:"contexts"`
	CurrentContext string     `mapstructure:"currentcontext"`
	Kind           string     `mapstructure:"kind"`
	Users          []User     `mapstructure:"users"`
}
