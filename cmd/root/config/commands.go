package config

import . "github.com/microsoft/go-sqlcmd/cmd/commander"

var Commands = []Commander{
	&CurrentContext{},
	&GetContexts{},
	&GetEndpoints{},
	&GetUsers{},
	&UseContext{},
	&View{},
}
