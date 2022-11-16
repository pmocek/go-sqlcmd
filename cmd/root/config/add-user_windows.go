package config

import "github.com/microsoft/go-sqlcmd/internal/helpers/cmd"

func (c *AddUser) encryptPasswordFlag() {
	c.AddFlag(cmd.FlagInfo{
		Bool: &c.encryptPassword,
		Name: "encrypt-password",
		Usage: "Encode the password",
	})
}
