package model

type User struct {
	ID 		 uint 		`gorm:"primaryKey"`
	Email 	 string 	`gorm:"unique; not null"`
	Name 	 string 	`gorm:"unique; not null"`
	Password string 	`gorm:"unique; not null"`
	Messages []Message
}