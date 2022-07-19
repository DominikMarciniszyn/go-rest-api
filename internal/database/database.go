package database

import (
	"fmt"
	"github.com/spf13/viper"
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

	dbError := Database.AutoMigrate(&entities.Order{})

	if dbError != nil {
		return dbError
	}

	return nil
}

func prepareDatabaseUri() string {
	host := viper.Get("DATABASE_HOST")
	user := viper.Get("DATABASE_USER")
	password := viper.Get("DATABASE_PASSWORD")
	databaseName := viper.Get("DATABASE_NAME")
	databasePort := viper.Get("DATABASE_PORT")

	return fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v", host, user, password, databaseName, databasePort)
}
