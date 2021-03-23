package model

type Message struct {
	id 	 int `gorm:"unique; not null"`
	text string
	user User `gorm:"not null"`
}

type Messages []Message