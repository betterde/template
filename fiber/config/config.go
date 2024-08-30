package config

import (
	"github.com/betterde/template/fiber/internal/journal"
	"github.com/spf13/viper"
	"os"
	"strings"
)

var Conf *Config

type Config struct {
	// Add config item to here
	Env     string  `yaml:"env"`
	HTTP    HTTP    `yaml:"http"`
	Logging Logging `yaml:"logging"`
}

type HTTP struct {
	Listen  string `yaml:"listen"`
	TLSKey  string `yaml:"tlsKey"`
	TLSCert string `yaml:"tlsCert"`
}

type Logging struct {
	Level string `yaml:"level"`
}

func Parse(file string, envPrefix string) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName(".config")
	}

	// read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix(envPrefix)

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil {
		journal.Logger.Errorf("Failed to read configuration file: %s", err)
		os.Exit(1)
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	err := viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
