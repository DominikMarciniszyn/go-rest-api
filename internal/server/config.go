package server

type Config struct {
	Port int `env:"PORT" default:"8080"`
}
