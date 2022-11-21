package install

import "github.com/microsoft/go-sqlcmd/internal/helper/cmd"

func (c *MssqlBase) encryptPasswordFlag(addFlag func(cmd.FlagOptions)) {
	addFlag(cmd.FlagOptions{
		Bool:  &c.encryptPassword,
		Name:  "encrypt-password",
		Usage: "Encode the generated password in the sqlconfig file",
	})
}
