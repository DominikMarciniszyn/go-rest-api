package internal

import (
	"fmt"
	"github.com/goava/di"
	"github.com/spf13/viper"
)

type postgresConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     int
}

type Config struct {
	Port     int
	LogLevel string
	Pretty   bool
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
