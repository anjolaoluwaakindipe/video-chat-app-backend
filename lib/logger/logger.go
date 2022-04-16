package logger

import (
	"log"
	"os"

	"go.uber.org/fx"
)

type Logger struct {
	Logger *log.Logger
}

func NewLogger() *Logger {
	myLogger := log.New(os.Stdout, "video-app", log.LstdFlags)
	return &Logger{Logger: myLogger}
}

var Module = fx.Options(
	fx.Provide(NewLogger),
)
