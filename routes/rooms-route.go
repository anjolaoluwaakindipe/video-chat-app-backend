package routes

import (
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/controllers"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/handler"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/logger"
)

type RoomRouter struct {
	handler *handler.Handler
	log *logger.Logger
	roomController *controllers.RoomController
}


func NewRoomRouter(handler *handler.Handler, log *logger.Logger, roomController *controllers.RoomController) *RoomRouter{
	return &RoomRouter{handler: handler, log: log, roomController: roomController}
}


func (rr *RoomRouter) SetUp(){

	rr.log.Logger.Println("Setting Up Room Routes")
	router := rr.handler.Gin.Group("/api/v1/rooms")

	{
		router.POST("/create", rr.roomController.CreateRoom())
		router.GET("/join/:roomID", rr.roomController.JoinRoom())
	}
}