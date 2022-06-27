package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func Initialize() {
	db, err := gorm.Open(sqlite.Open("social_media.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Conn = db

	if err := db.AutoMigrate(
		&User{},
		&Friendship{},
		&Message{}); err != nil {
		panic(err)
	}
}
