package routes

import (
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/controllers"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/handler"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/logger"
)

type UserRouter struct {
	UserController *controllers.UserController
	logger         *logger.Logger
	handler        *handler.Handler
}

func NewUserRouter(userController *controllers.UserController, router *handler.Handler, logger *logger.Logger) *UserRouter {
	return &UserRouter{UserController: userController, handler: router, logger: logger}

}

func (ur *UserRouter) SetUp() {
	ur.logger.Logger.Println("Setting up Auth Routes")
	userRoutes := ur.handler.Gin.Group("api/v1/user")
	{
		userRoutes.GET("/")
	}

}
