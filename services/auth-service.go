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
	//  Creates a Log in access token
	CreateToken(email string, password string) (string, *errs.AppError)

	// Checks if user already exist
	CheckIfUserExist (email string, username string) ([]*errs.AppError)
}

type AuthServiceImpl struct {
	envVars  *env.EnvVar
	userRepo repositories.UserRepo
}

// Creates a Auth Service implementation struct
func NewAuthService(env *env.EnvVar, userRepoPSQL *repositories.UserRepoPGSQL) AuthService {
	return &AuthServiceImpl{envVars: env, userRepo: userRepoPSQL}
}

//  Creates a Log in access token
func (as *AuthServiceImpl) CreateToken(email string, password string) (string, *errs.AppError) {

	user, userErr := as.userRepo.FindUserByEmailAndPassword(email, password)

	if userErr != nil {
		return "", userErr
	}

	idString := strconv.FormatUint(uint64(user.ID), 36)
	jwtToken := entities.Token{Id: idString, Email: user.Email, StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Add(time.Duration(as.envVars.AccessTokenExpireTime) * time.Second).Unix()}}

	tokenString, err := jwtToken.CreateAccessTokenString(jwt.SigningMethodHS256, as.envVars.SecretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// Checks if user already exist
func (as *AuthServiceImpl) CheckIfUserExist(email string, username string) ([]*errs.AppError){
	userErrs := make([]*errs.AppError, 0)
	emailErr := as.userRepo.CheckIfEmailExist(email)
	usernameErr:= as.userRepo.CheckIfUsernameExist( username)

	if emailErr != nil{
		userErrs = append(userErrs, emailErr)
	}

	if usernameErr != nil{
		userErrs = append(userErrs, usernameErr)
	}

	if len(userErrs) != 0 {
		return userErrs
	}
	return nil
}


