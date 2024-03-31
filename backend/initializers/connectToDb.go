package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() error {
	var err error

	dbURL := os.Getenv("POSTGRES_URL")

	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Println("Error to connect to DB")
		return err
	}

	return nil
}
