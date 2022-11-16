package install

func (c *MssqlBase) encryptPasswordFlag(addFlag func(cmd.FlagInfo)) {
	// Linux OS doesn't have a native DPAPI or Keychain equivalent
}
