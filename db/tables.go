package db

import "time"

type User struct {
	ID              uint `gorm:"primaryKey"`
	CreatedAt       time.Time
	LastKeyUpdateAt time.Time
	Key             string
	Email           string `gorm:"index"`
	Password        string
	Name            string
}

type Friendship struct {
	ID           uint `gorm:"primaryKey"`
	SenderID     uint `gorm:"index"`
	ReceiverID   uint `gorm:"index"`
	Accepted     bool
	AcceptedTime time.Time
}

type Message struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time `json:"created_at"`
	AuthorId  uint      `json:"author_id" gorm:"index"`
	Message   string    `json:"message"`
	IsPublic  bool      `json:"is_public"`
}
