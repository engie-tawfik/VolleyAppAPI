package logger

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}
}

func StopLogger() {
	Logger.Sync()
}
