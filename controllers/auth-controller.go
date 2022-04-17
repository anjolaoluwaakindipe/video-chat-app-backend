package controllers

import (
	"net/http"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/dto"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/env"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/validation"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/services"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
)

type AuthController struct {
	ValidationHandler *validation.ValidationHandler
	AuthService       services.AuthService
	UserService       services.UserService
	envVars           *env.EnvVar
}

func NewAuthController(validationHandler *validation.ValidationHandler, authService services.AuthService, userService services.UserService, env *env.EnvVar) *AuthController {
	return &AuthController{ValidationHandler: validationHandler, AuthService: authService, UserService: userService, envVars: env}
}

func (ac *AuthController) SignUpUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var signupInfo *dto.SignUp
		ctx.BindJSON(&signupInfo)

		jsonErr := ac.ValidationHandler.ValidateStruct(signupInfo)

		if jsonErr != nil {
			ctx.JSON(http.StatusBadRequest, map[string]interface{}{
				"message":jsonErr,
			})
			return
		}

		userCheckErrs := ac.AuthService.CheckIfUserExist(signupInfo.Email, signupInfo.Username)

		if userCheckErrs != nil {
			ctx.JSON(http.StatusConflict, userCheckErrs)
			return 
		}


		error := ac.UserService.SaveUser(*signupInfo)

		if error != nil {
			ctx.JSON(error.Code, error)
			return
		}

		ctx.JSON(http.StatusOK, map[string]interface{}{"message": "User Registration Successful"})

	}
}

func (ac *AuthController) LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginInfo *dto.Login
		ctx.BindJSON(&loginInfo)
		jsonErr := ac.ValidationHandler.ValidateStruct(loginInfo)

		if jsonErr != nil {
			ctx.JSON(http.StatusBadRequest, jsonErr)
			return
		}

		token, tokenErr := ac.AuthService.CreateToken(loginInfo.Email, loginInfo.Password)

		if tokenErr != nil {
			ctx.JSON(tokenErr.Code, tokenErr)
			return
		}
		s := securecookie.New([]byte("adfsadad"), nil)
		encodedCookie, encodingErr := s.Encode("token", token)

		if encodingErr != nil{
			appEncodingErr := errs.NewUnexpectedError("Error occured when encoding token as cookie")
			ctx.JSON(appEncodingErr.Code, appEncodingErr )
			return
		}

		ctx.SetCookie("token", encodedCookie, 30000,"/api/v1/login", "recto.com",false, true)
	
		

		ctx.JSON(http.StatusOK, map[string]interface{}{
			"accessToken":          token,
			"accessTokenExpiresTime": ac.envVars.AccessTokenExpireTime,
		})
		return

	}
}
