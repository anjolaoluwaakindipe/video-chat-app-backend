package handler

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Handler struct {
	Gin *gin.Engine
}

func NewHandler() *Handler {
	ginHandler := gin.Default()
	return &Handler{Gin: ginHandler}
}

var Module = fx.Options(
	fx.Provide(NewHandler),
)
