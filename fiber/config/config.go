package config

import (
	"errors"
	"github.com/betterde/template/fiber/internal/build"
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

func Parse(file string) {
	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName("." + build.Name)
	}

	var notFoundError viper.ConfigFileNotFoundError

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil && errors.As(err, &notFoundError) {
		journal.Logger.Debugf("Config file not found, using defaults")
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	viper.SetEnvPrefix(build.Name)
	viper.SetEnvKeyReplacer(strings.NewReplacer("_", "."))

	err := viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
