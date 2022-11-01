// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package config

import . "github.com/microsoft/go-sqlcmd/cmd/commander"

var Commands = []Commander{
	&AddContext{},
	&AddEndpoint{},
	&AddUser{},
	&ConnectionStrings{},
	&CurrentContext{},
	&DeleteContext{},
	&DeleteEndpoint{},
	&DeleteUser{},
	&GetContexts{},
	&GetEndpoints{},
	&GetUsers{},
	&UseContext{},
	&View{},
}
