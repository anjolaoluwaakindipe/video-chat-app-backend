package services

import (
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/dto"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/models"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/repositories"
)

type UserService interface {
	SaveUser(signUpInfo dto.SignUp) *errs.AppError
}

type UserServiceImpl struct {
	UserRepo repositories.UserRepo
}

func NewUserService(userRepo *repositories.UserRepoPGSQL) UserService {
	return &UserServiceImpl{UserRepo: userRepo}
}

func (us *UserServiceImpl) SaveUser(signUpInfo dto.SignUp) *errs.AppError {
	newUser := models.User{Email: signUpInfo.Email, Password: signUpInfo.Password, Firstname: signUpInfo.Firstname, Lastname: signUpInfo.Lastname, Username: signUpInfo.Username}
	err := us.UserRepo.SaveUser(newUser)

	if err != nil {
		return err
	}
	return nil
}
