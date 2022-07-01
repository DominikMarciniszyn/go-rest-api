package config

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-rest-api/internal/database/entities"
)

var Database *gorm.DB

func Connect() error {
	var err error
	databaseUri := prepareDatabaseUri()

	Database, err = gorm.Open(postgres.Open(databaseUri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Order{})

	return nil
}

func prepareDatabaseUri() string {
	host := os.Getenv("DATABASE_HOST")
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	databaseName := os.Getenv("DATABASE_NAME")
	databasePort := os.Getenv("DATABASE_PORT")

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, databaseName, databasePort)
}
