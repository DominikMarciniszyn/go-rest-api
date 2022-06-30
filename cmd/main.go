package main

import (
	"go-rest-api/cmd/server"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := loadConfig()

	if err != nil {
		log.Fatal("Error:", err)
		os.Exit(1)
	}

	server.StartServer()
}

func loadConfig() error {
	return godotenv.Load()
}
