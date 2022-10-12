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

func (c *Mssql) GetCommand() *Command {
	const use = "mssql"
	const short = "Install/Create Sql Server"
	c.Command = &Command{
		Use:   use,
		Short: short,
		Long:  short,
		Example: `# Install SQL Server in a docker container
  sqlcmd install mssql`,
		Aliases: []string{"create"},

	}

	c.AddSubCommands()
	// c.setDefaultSubCommandIfNonePresent("server")

	return c.Command
}

func  (c *Mssql) setDefaultSubCommandIfNonePresent(defCmd string) {
	var cmdFound bool
	cmd :=c.Command.Commands()

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
		c.Command.SetArgs(args)
	}
	if err := c.Command.Execute(); err != nil {
		CheckErr(err)
	}
}
