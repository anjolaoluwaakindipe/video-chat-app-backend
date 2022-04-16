package db

import (
	"log"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/models"
	"go.uber.org/fx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGDB struct {
	Db *gorm.DB
}

func NewDatabase() *PGDB {
	return &PGDB{}
}

func (d *PGDB) OpenNewPSQLDB(log *log.Logger, postgresDSN string) {
	// https://github.com/go-gorm/postgres
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:  postgresDSN               ,
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})

	if err != nil {
		log.Panicln(err)
	}

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.FriendRequest{})

	d.Db = db
}

var Module = fx.Options(
	fx.Provide(NewDatabase),
)
