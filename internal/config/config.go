package config

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Port     int
	LogLevel string
	Pretty   bool
}

func InitServiceConfig() {
	pflag.Int("port", 8080, "TPC port to listen on")
	pflag.String("logLevel", "info", "Minimal log level (trace, debug, info, warn, error, fatal, panic)")
	pflag.Bool("pretty", false, "Pretty log formatting")

	pflag.Parsed()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		panic(err)
	}

	viper.SetEnvPrefix("playground")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	viper.AutomaticEnv()
}

func LoadConfig() *Config {
	return &Config{
		Port:     viper.GetInt("port"),
		LogLevel: viper.GetString("logLevel"),
		Pretty:   viper.GetBool("pretty"),
	}
}
