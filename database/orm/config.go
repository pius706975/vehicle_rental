package orm

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// func NewDB() (*gorm.DB, error) {
// 	host := os.Getenv("DB_HOST")
// 	user := os.Getenv("DB_NAME")
// 	password := os.Getenv("DB_PASS")
// 	dbName := os.Getenv("DB_NAME")

// 	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

// 	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})

// 	if err != nil {
// 		return nil, err
// 	}

// 	return gormDb, nil 
// }

func NewDB() (*gorm.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return gormDb, nil
}