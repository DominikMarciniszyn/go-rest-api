package main

import (
	"github.com/joho/godotenv"
	"go-rest-api/cmd/server"
	"log"
)

func main() {
	err := loadConfig()

	if err != nil {
		log.Fatal("Error:", err)
	}

	server.StartServer()
}

func loadConfig() error {
	return godotenv.Load()
}
