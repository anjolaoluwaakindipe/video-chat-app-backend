package controllers

import (
	"log"
	"net/http"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RoomController struct {
	roomService services.RoomService
	websocket *websocket.Upgrader
}

func NewRoomController(roomService *services.RoomServiceImpl) *RoomController {
	return &RoomController{roomService: roomService, websocket: &websocket.Upgrader{CheckOrigin: func(r *http.Request)bool{return true}}}
}


// CreateRoom create a room and return RoomId
func (rc *RoomController) CreateRoom() gin.HandlerFunc{
	return func (ctx *gin.Context){
		roomID := rc.roomService.CreateRoom()
		body := make(map[string]string)
		body["room_id"] = roomID
		ctx.JSON(http.StatusOK, body)
	}
}


// JoinRoom join a room
func (rc *RoomController) JoinRoom() gin.HandlerFunc {
	return func (ctx *gin.Context){
		roomID :=ctx.Param("roomID")
		log.Println(roomID)

		ws , wsErr := rc.websocket.Upgrade(ctx.Writer, ctx.Request, nil)

		if  newWsErr := errs.NewUnexpectedError("An error occured when handling websocket"); wsErr != nil{
			ctx.JSON(newWsErr.Code, newWsErr.Message)
		}

		rc.roomService.JoinRoom(roomID, false, ws)
	}

	
}