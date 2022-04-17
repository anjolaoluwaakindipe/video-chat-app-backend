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
	AccessTokenExpireTime  int    `mapstructure:"ACCESS_TOKEN_EXPIRE_TIME"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDbname   string `mapstructure:"POSTGRES_DBNAME"`
	PostgresPort     string `mapstructure:"POSTGRES_PORT"`
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
}

func newEnvVar(logger *logger.Logger) *EnvVar {
	var enironmentVariables EnvVar
	viper.SetConfigFile(".env")
	readErr := viper.ReadInConfig()
	if readErr != nil {
		logger.Logger.Println("No .env file found")
		viper.AutomaticEnv()
		viper.BindEnv("SERVER_PORT")
		viper.BindEnv("SECRET_KEY")
		viper.BindEnv("ACCESS_TOKEN_EXPIRE_TIME")
		viper.BindEnv("POSTGRES_USER")
		viper.BindEnv("POSTGRES_PASSWORD")
		viper.BindEnv("POSTGRES_DBNAME")
		viper.BindEnv("POSTGRES_PORT")
		viper.BindEnv("POSTGRES_HOST")
	}
	unmarshallErr := viper.Unmarshal(&enironmentVariables)
		if unmarshallErr != nil {
			logger.Logger.Fatalln("Could not unmarshall environment variables")
		}
	fmt.Println(enironmentVariables.AccessTokenExpireTime)
	fmt.Println(enironmentVariables.PostgresDbname)
	fmt.Println(enironmentVariables.PostgresHost)
	fmt.Println(enironmentVariables.PostgresPort)
	fmt.Println(enironmentVariables.PostgresUser)
	
	
	return &enironmentVariables
}

var Module = fx.Options(
	fx.Provide(newEnvVar),
)
