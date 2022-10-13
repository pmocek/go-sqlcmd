// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.

package install

import (
	. "github.com/microsoft/go-sqlcmd/cmd/commander"
	. "github.com/spf13/cobra"
	"os"
)

type Mssql struct {
	AbstractBase

	tag string
	registry string
	repo string
	installType string
	acceptEula bool
	contextName string
	defaultDatabase string
}

func (c *Mssql) GetCommand() (command *Command) {
	const use = "mssql"
	const short = "Install/Create Sql Server"

	command = c.AddCommand(Command{
		Use:   use,
		Short: short,
		Long:  short,
		Example: `# Install SQL Server in a docker container
  sqlcmd install mssql`,
		Aliases: []string{"create"},
	})

	// BUG(stuartpa): Get this working: c.setDefaultSubCommandIfNonePresent(command, "server")

	return
}

func  (c *Mssql) setDefaultSubCommandIfNonePresent(command *Command, defCmd string) {
	var cmdFound bool
	cmd := command.Commands()

	for _,a:=range cmd{
		for _,b:=range os.Args[1:] {
			if a.Name()==b {
				cmdFound=true
				break
			}
		}
	}
	if cmdFound == false {
		args:=append([]string{defCmd}, os.Args[1:]...)
		command.SetArgs(args)
	}
	if err := command.Execute(); err != nil {
		CheckErr(err)
	}
}
