package model

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Email 	 string `gorm:"unique; not null"`
	Name 	 string `gorm:"unique; not null"`
	Password string `gorm:"unique; not null"`
}