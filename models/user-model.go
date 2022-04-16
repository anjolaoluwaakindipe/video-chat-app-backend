package models

import (
	"time"

	"gorm.io/gorm"
)

// Class variables
type User struct {
	gorm.Model
	ID              uint
	Email           string `gorm:"not null;unique"`
	Password        string `gorm:"not null"`
	Firstname       string `gorm:"not null;size:100"`
	Lastname        string `gorm:"not null;size:100"`
	Username        string `gorm:"not null;size:100;unique"`
	CreatedAt       time.Time
	UpdatedAt       time.Time
	IsEmailVerified bool
	Friends         []User          `gorm:"many2many:user_friends"`
	FriendRequest   []FriendRequest `gorm:"foreignKey:ID"`
}




