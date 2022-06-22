package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"go-rest-api/internal/database/entities"
)

var database *gorm.DB
var DATABASE_URI string = "host=localhost user=admin password=admin dbname=postgres port=5433 sslmode=disable"

func Connect() error {
	var err error

	database, err = gorm.Open(postgres.Open(DATABASE_URI), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: true,
	})

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&entities.Order{})

	return nil
}
