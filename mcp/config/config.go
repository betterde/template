package config

import (
	"errors"
	"os"
	"strings"

	"github.com/betterde/template/mcp/internal/journal"
	"github.com/spf13/viper"
	"github.com/subosito/gotenv"
)

var Conf *Config

type HTTP struct {
	Listen                     string `yaml:"listen" mapstructure:"LISTEN"`
	Stateless                  bool   `yaml:"stateless" mapstructure:"STATELESS"`
	DisableLocalhostProtection bool   `yaml:"disable_localhost_protection" mapstructure:"DISABLE_LOCALHOST_PROTECTION"`
}

type Logging struct {
	Level string `yaml:"level" mapstructure:"LEVEL"`
}

type Config struct {
	DSN      string  `yaml:"dsn" mapstructure:"DSN"`
	HTTP     HTTP    `yaml:"http" mapstructure:"HTTP"`
	Logging  Logging `yaml:"logging" mapstructure:"LOGGING"`
	ReadOnly bool    `yaml:"read_only" mapstructure:"READ_ONLY"`
}

func Parse(file string) {
	err := gotenv.Load()
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		journal.Logger.Sugar().Errorf("Unable to load .env file, %v", err)
	}

	if file != "" {
		viper.SetConfigFile(file)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yml")
		viper.SetConfigType("yaml")
		viper.SetConfigName(".config")
		viper.AddConfigPath("/etc/mcp-servers/mysql")
	}

	// read in environment variables that match
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("MYSQL_MCP_SERVER")

	viper.SetDefault("READ_ONLY", true)
	viper.SetDefault("HTTP.LISTEN", "0.0.0.0:8080")
	viper.SetDefault("LOGGING.LEVEL", "DEBUG")

	if err := viper.BindEnv("DSN", "MYSQL_MCP_SERVER_DSN"); err != nil {
		journal.Logger.Sugar().Error(err)
	}

	if err := viper.BindEnv("HTTP.LISTEN", "MYSQL_MCP_SERVER_HTTP_LISTEN"); err != nil {
		journal.Logger.Sugar().Error(err)
	}

	if err := viper.BindEnv("READ_ONLY", "MYSQL_MCP_SERVER_READ_ONLY"); err != nil {
		journal.Logger.Sugar().Error(err)
	}

	if err := viper.BindEnv("LOGGING.LEVEL", "MYSQL_MCP_SERVER_LOGGING_LEVEL"); err != nil {
		journal.Logger.Sugar().Error(err)
	}

	var notFoundError viper.ConfigFileNotFoundError

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err != nil && !errors.As(err, &notFoundError) {
		journal.Logger.Sugar().Errorf("Unable to read config file, %v", err)
		os.Exit(1)
	}

	// read in environment variables that match
	viper.AutomaticEnv()

	err = viper.Unmarshal(&Conf)
	if err != nil {
		journal.Logger.Sugar().Errorf("Unable to decode into config struct, %v", err)
		os.Exit(1)
	}
}
