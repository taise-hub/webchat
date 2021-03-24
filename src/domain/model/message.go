package model

type Message struct {
	ID 	 	uint 	`gorm:"primaryKey; not null"`
	Text 	string	
	UserID 	uint 	`gorm:"not null"`
}
