package install

import "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"

func (c *MssqlBase) encryptPasswordFlag(addFlag func(cmd.FlagInfo)) {
	addFlag(cmd.FlagInfo{
		Bool: &c.encryptPassword,
		Name: "encrypt-password",
		Usage: "Encode the generated password in the sqlconfig file",
	})
}
