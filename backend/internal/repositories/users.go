package repositories

import (
	"github.com/SantGT5/quintosgo/internal/database"
)

type User struct {
	Email    string
	Password string
}

func (u *User) CreateUser() error {
	db := database.GetDatabase()

	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}
