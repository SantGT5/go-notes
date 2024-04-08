package database

import (
	"github.com/SantGT5/quintosgo/schemas"
)

func MigrateDB() error {
	// Get all structs from the models package
	models := []interface{}{&schemas.User{}}

	// Migrate each struct
	err := db.AutoMigrate(models...)

	if err != nil {
		return err
	}

	return nil
}
