package env

import (
	"fmt"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/logger"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type EnvVar struct {
	ServerPort       string `mapstructure:"SERVER_PORT"`
	SecretKey        string `mapstructure:"SECRET_KEY"`
	TokenExpireTime  int    `mapstructure:"TOKEN_EXPIRE_TIME"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDbname   string `mapstructure:"POSTGRES_DBNAME"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
}

func newEnvVar(logger *logger.Logger) *EnvVar {
	var enironmentVariables EnvVar
	viper.SetConfigFile(".env")
	readErr := viper.ReadInConfig()
	if readErr != nil {
		logger.Logger.Fatalln("Could not read environment variables")
	}
	unmarshallErr := viper.Unmarshal(&enironmentVariables)
	if unmarshallErr != nil {
		logger.Logger.Fatalln("Could not unmarshall environment variables")
	}
	fmt.Println(enironmentVariables.TokenExpireTime)
	return &enironmentVariables
}

var Module = fx.Options(
	fx.Provide(newEnvVar),
)
