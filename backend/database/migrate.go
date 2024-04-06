package database

import (
	"github.com/SantGT5/quintosgo/models"
)

func MigrateDB() error {
	// Get all structs from the models package
	models := []interface{}{&models.User{}}

	// Migrate each struct
	err := db.AutoMigrate(models...)

	if err != nil {
		return err
	}

	return nil
}
