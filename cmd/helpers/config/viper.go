package config

import (
	"bytes"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/file"
	"github.com/microsoft/go-sqlcmd/cmd/helpers/output"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

func configureViper(configFile string) {
	if configFile == "" {
		home, err := os.UserHomeDir()
		checkErr(err)

		configFile = filepath.Join(home, ".sqlcmd", "sqlconfig")
	}

	viper.SetConfigType("yaml")
	viper.SetEnvPrefix("SQLCMD")
	viper.SetConfigFile(configFile)

	file.CreateEmptyIfNotExists(configFile)
}

func load() {
	var err error

	err = viper.ReadInConfig()
	checkErr(err)
	err = viper.BindEnv("ACCEPT_EULA")
	checkErr(err)
	viper.AutomaticEnv() // read in environment variables that match
	err = viper.Unmarshal(&config)
	checkErr(err)

	output.Tracef("Config loaded from file: %v", viper.ConfigFileUsed())
}

func Save() {
	b, err := yaml.Marshal(&config)
	checkErr(err)
	err = viper.ReadConfig(bytes.NewReader(b))
	checkErr(err)
	file.CreateEmptyIfNotExists(viper.ConfigFileUsed())
	err = viper.WriteConfig()
	checkErr(err)
}

func GetConfigFileUsed() string {
	return viper.ConfigFileUsed()
}
