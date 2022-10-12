package config

import . "github.com/microsoft/go-sqlcmd/cmd/commander"

var Commands = []Commander{
	&ConnectionStrings{},
	&CurrentContext{},
	&GetContexts{},
	&GetEndpoints{},
	&GetUsers{},
	&UseContext{},
	&View{},
}
