package config

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-rest-api/internal/database/entities"
)

var Database *gorm.DB

func Connect() error {
	var err error
	databaseUri := os.Getenv("DATABASE_URI")

	Database, err = gorm.Open(postgres.Open(databaseUri), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Order{})

	return nil
}
