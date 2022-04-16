package services

import (
	"log"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/entities"
	"github.com/gorilla/websocket"
)

type RoomService interface {
	CreateRoom() string
	JoinRoom(roomID string, host bool , conn *websocket.Conn)
}

type RoomServiceImpl struct {
	AllRooms *entities.RoomMap
}

func NewRoomServiceImpl( allRooms *entities.RoomMap) *RoomServiceImpl {
	return &RoomServiceImpl{allRooms}
}

func (rs *RoomServiceImpl) CreateRoom()  string{
	roomID := rs.AllRooms.CreateRoom()
	log.Println(rs.AllRooms.Map)
	return roomID
}

func (rs *RoomServiceImpl) JoinRoom(roomID string, host bool , conn *websocket.Conn) {
	rs.AllRooms.InsertIntoRoom(roomID, host, conn)
}