package database

import (
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func Init() error {
	var err error

	// init Postgres DB
	db, err = ConnectToDb()

	if err != nil {
		return fmt.Errorf("error init postgres: %v", err.Error())
	}

	if err := MigrateDB(); err != nil {
		return fmt.Errorf("error migrating database: %v", err.Error())
	}

	return nil
}

func GetDatabase() *gorm.DB {
	return db
}
