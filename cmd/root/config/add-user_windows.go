package config

import "github.com/microsoft/go-sqlcmd/internal/helper/cmd"

func (c *AddUser) encryptPasswordFlag() {
	c.AddFlag(cmd.FlagOptions{
		Bool:  &c.encryptPassword,
		Name:  "encrypt-password",
		Usage: "Encode the password",
	})
}
