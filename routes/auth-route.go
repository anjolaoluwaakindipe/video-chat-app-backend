package routes

import (
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/controllers"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/handler"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/logger"
)

type AuthRouter struct {
	AuthController *controllers.AuthController
	logger         *logger.Logger
	handler        *handler.Handler
}

func NewAuthRouter(router *handler.Handler, logger *logger.Logger, authController *controllers.AuthController) *AuthRouter {
	return &AuthRouter{handler: router, logger: logger, AuthController: authController}

}

func (ar *AuthRouter) SetUp() {
	routes := ar.handler.Gin.Group("/api/v1")
	{
		routes.POST("/login", ar.AuthController.LoginUser())
		routes.POST("/signup", ar.AuthController.SignUpUser())
	}
}
