package model

type User struct {
	Email 	 string `gorm:"unique; not null"`
	Name 	 string `gorm:"unique; not null"`
	Password string `gorm:"unique; not null"`
}