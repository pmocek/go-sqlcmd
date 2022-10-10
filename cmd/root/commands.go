package root

import . "github.com/microsoft/go-sqlcmd/cmd/commander"

var Commands = []Commander{
	// TODO: &Bulkcopy{},
	&Config{},
	&Query{},
	&Install{},
	&Uninstall{},
}
