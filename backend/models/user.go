package models

type User struct {
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

func (u *User) CreateUser() error {
	// Create the user
	if err := db.Create(u).Error; err != nil {
		return err
	}

	return nil
}
