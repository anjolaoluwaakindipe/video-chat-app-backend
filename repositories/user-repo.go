package repositories

import (
	"errors"
	"fmt"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/db"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	// BeforeSave() error
	// Validate(action string) error

	// Saves User to the database being used
	SaveUser(user models.User) *errs.AppError
	
	//  Find a user using both email or password, 
	//  if either email or password is wrong an app error should occur (for logging in purpose)
	FindUserByEmailAndPassword(email string, password string) (*models.User, *errs.AppError)

	// Checks if the email given exists within the database being used
	CheckIfEmailExist (email string) (*errs.AppError)

	// Checks if the username given exists within the database being used
	CheckIfUsernameExist (username string) (*errs.AppError)
	// FindUserByID(id int) (models.User, error)
}

// user repo with pgsql implementation
type UserRepoPGSQL struct {
	pgdb *db.PGDB
}

// Creates a new psql repo implementation struct 
func NewUserRepoPGSQL(pgsqlDatabase *db.PGDB) *UserRepoPGSQL {
	return &UserRepoPGSQL{pgdb: pgsqlDatabase}
}

// Saves User to the database being used
func (ur *UserRepoPGSQL) SaveUser(user models.User) *errs.AppError {
	hashPassword, error := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if error != nil {
		return errs.NewUnexpectedError("Unexpected Error while hashing password")
	}
	user.Password = string(hashPassword)
	result := ur.pgdb.Db.Create(&user)

	if result.Error != nil {
		return errs.NewUnexpectedError("Unexpected Error when creating user")
	}
	return nil
}


//  Find a user using both email or password, 
//  if either email or password is wrong an app error should occur (for logging in purpose)
func (ur *UserRepoPGSQL) FindUserByEmailAndPassword(email string, password string) (*models.User, *errs.AppError) {
	var user models.User
	result := ur.pgdb.Db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, errs.NewContentNotFoundError("Invalid Email or Password")
		} else {
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}

	}
	return &user, nil
}

// Checks if user exist by giving the column to work with and the value
func (ur *UserRepoPGSQL) userExistenceChecker( field string, value string ) *errs.AppError{
	var user models.User;
	result := ur.pgdb.Db.Where(field + " = ?", value).First(&user)
	
	fmt.Println(result, errors.Is(result.Error, gorm.ErrRecordNotFound))
	fmt.Println(user)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound{
			return nil
		}else{
			return errs.NewUnexpectedError("Unexpected database error 1")
		}
	}

	return errs.NewConflictError( field +" already exist")
}
 

// Checks if the email given exists within the database being used
func (ur*UserRepoPGSQL) CheckIfEmailExist(email string) (*errs.AppError) {
	return ur.userExistenceChecker("email", email)
}


// Checks if the username given exists within the database being used
func (ur *UserRepoPGSQL) CheckIfUsernameExist (username string) (*errs.AppError){
	return ur.userExistenceChecker("username", username)
}

