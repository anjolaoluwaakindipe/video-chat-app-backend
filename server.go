package main

import (
	"context"
	"net/http"
	"time"

	"github.com/anjolaoluwaakindipe/video-chatapp-golang/controllers"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/entities"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/handler"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/db"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/env"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/logger"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/lib/validation"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/repositories"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/routes"
	"github.com/anjolaoluwaakindipe/video-chatapp-golang/services"

	"go.uber.org/fx"
)

var Module = fx.Options(
	env.Module,
	repositories.Module,
	validation.Module,
	services.Module,
	routes.Module,
	controllers.Module,
	db.Module,
	logger.Module,
	handler.Module,
	entities.Module,
	fx.Invoke(bootstrap),
)

func main() {

	fx.New(Module).Run()

}

func bootstrap(lifecycle fx.Lifecycle, postgres *db.PGDB, handler *handler.Handler, routes *routes.AllRouters, logger *logger.Logger, envVars *env.EnvVar) {
	// "host=" + envVars.PostgresHost + " user=" + envVars.PostgresUser + " password=" + envVars.PostgresPassword + " dbname=" + envVars.PostgresDbname + " port=" + envVars.PostgresPort + " sslmode=disable TimeZone=Asia/Shanghai"
	postgresConfig := "postgres://"+envVars.PostgresUser+":"+envVars.PostgresPassword+"@"+envVars.PostgresHost+":"+envVars.PostgresPort+"/"+envVars.PostgresDbname+"?sslmode=disable"

	postgres.OpenNewPSQLDB(logger.Logger, postgresConfig)
	routes.SetUp()

	server := &http.Server{
		Addr:    ":" + envVars.ServerPort,
		Handler: handler.Gin,
	}

	lifecycle.Append(fx.Hook{

		OnStart: func(context.Context) error {

			logger.Logger.Println("Starting Application")
			logger.Logger.Println("---------------------")
			logger.Logger.Println("------- CLEAN -------")
			logger.Logger.Println("---------------------")

			go func() {
				err := server.ListenAndServe()
				if err != nil {
					logger.Logger.Println(err)
				}

			}()

			return nil
		},
		OnStop: func(context.Context) error {
			// shutDownChannel := make(chan os.Signal)

			// signal.Notify(shutDownChannel, os.Interrupt)
			// signal.Notify(shutDownChannel, os.Kill)
			// sig := <-shutDownChannel
			logger.Logger.Print("Gracefully Shutting Down")
			ctx, ctxCancelFunc := context.WithTimeout(context.Background(), 30*time.Second)

			defer ctxCancelFunc()
			server.Shutdown(ctx)

			logger.Logger.Print("Stopping Container")

			return nil
		},
	})

}
