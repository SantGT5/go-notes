package repository

import (
	"github.com/SantGT5/quintosgo/internal/database"
)

type Signup struct {
	Email    string
	Password string
}

func (u *Signup) SignupUser() error {
	db := database.GetDatabase()

	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}
