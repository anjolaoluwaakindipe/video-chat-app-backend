package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) getUserById(id int) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func (uc *UserController) getAllFriends() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
