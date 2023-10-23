package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreatePostgresConnection() (*gorm.DB, error) {
	dsn := "host= user= password= dbname= port= sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return database, nil
}
