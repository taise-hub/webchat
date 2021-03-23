package model

type Message struct {
	Id 	 int `gorm:"unique; not null"`
	Text string
	User *User `gorm:"not null"`
}

type Messages []Message