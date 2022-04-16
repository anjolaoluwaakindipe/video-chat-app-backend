package repositories

import (
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/db"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/errs"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepo interface {
	// BeforeSave() error
	// Validate(action string) error
	SaveUser(user models.User) *errs.AppError
	FindUserByEmailAndPassword(email string, password string) (*models.User, *errs.AppError)
	// FindUserByID(id int) (models.User, error)
}

// user repo with pgsql implementation
type UserRepoPGSQL struct {
	pgdb *db.PGDB
}

// constructor
func NewUserRepoPGSQL(pgsqlDatabase *db.PGDB) *UserRepoPGSQL {
	return &UserRepoPGSQL{pgdb: pgsqlDatabase}
}

// methods
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
