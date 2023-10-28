package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreatePostgresConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password= dbname=holis-database port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return database, nil
}
