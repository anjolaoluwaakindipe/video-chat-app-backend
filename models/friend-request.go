package models

import (
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/enums"
	"gorm.io/gorm"
)

type FriendRequest struct {
	gorm.Model
	ID         uint
	Status     enums.FriendRequestStatus
	SenderID   uint
	Sender     User `gorm:"foreignKey:SenderID"`
	ReceiverID uint
	Receiver   User `gorm:"foreignKey:ReceiverID"`
}

func NewFriendRequestModel() *FriendRequest {
	return &FriendRequest{}
}
