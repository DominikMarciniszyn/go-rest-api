package internal

import (
	"fmt"
	"github.com/goava/di"
	"github.com/spf13/viper"
	"go-rest-api/internal/server"
)

type postgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

type Config struct {
	Server server.Config

	LogLevel string `env:"LOGLEVEL" default:"info"`
	Pretty   bool   `env:"PREETY" default:"false"`
	Postgres postgresConfig
}

func LoadConfig(container *di.Container) error {
	config := Config{}

	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	readConfigError := viper.ReadInConfig()

	if readConfigError != nil {
		return readConfigError
	}

	unmarshallError := viper.Unmarshal(&config)

	if unmarshallError != nil {
		return unmarshallError
	}

	fmt.Printf("%+v\n", config)

	return container.ProvideValue(&config)
}
