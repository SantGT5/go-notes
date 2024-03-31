package initializers

import (
	"github.com/SantGT5/notes/models"
)

func SyncDatabase() {
	// Get all structs from the models package
	structs := []interface{}{
		models.User{},
	}

	// Migrate each struct
	for _, s := range structs {
		DB.AutoMigrate(s)
	}
}
