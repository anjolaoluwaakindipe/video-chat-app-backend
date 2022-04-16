package services

import (
	"strconv"
	"time"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/entities"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/env"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/repositories"
	"github.com/dgrijalva/jwt-go"
)

type AuthService interface {
	CreateToken(email string, password string) (string, *errs.AppError)
}

type AuthServiceImpl struct {
	envVars  *env.EnvVar
	userRepo repositories.UserRepo
}

func NewAuthService(env *env.EnvVar, userRepoPSQL *repositories.UserRepoPGSQL) AuthService {
	return &AuthServiceImpl{envVars: env, userRepo: userRepoPSQL}
}

func (as *AuthServiceImpl) CreateToken(email string, password string) (string, *errs.AppError) {

	user, userErr := as.userRepo.FindUserByEmailAndPassword(email, password)

	if userErr != nil {
		return "", userErr
	}

	idString := strconv.FormatUint(uint64(user.ID), 36)
	jwtToken := entities.Token{Id: idString, Email: user.Email, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(as.envVars.TokenExpireTime) * time.Second).Unix()}}

	tokenString, err := jwtToken.CreateAccessTokenString(jwt.SigningMethodHS256, as.envVars.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
