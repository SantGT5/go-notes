package models

import (
	"github.com/SantGT5/quintosgo/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	db = database.GetDatabase()
}
