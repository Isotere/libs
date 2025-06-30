package main

import (
	"context"
	"os"

	"github.com/Isotere/libs/logger"
)

func main() {
	ctx := context.Background()

	var log *logger.Logger

	logFile, _ := os.Create("example.log")

	log = logger.New(
		"example-service",
		logger.WithWriter(logFile),
		logger.WithLogLevel(logger.LevelDebug),
		logger.WithErrorEvent(func(ctx context.Context, r logger.Record) {
			log.Info(ctx, "******* SEND ALERT *******")
		}),
	)

	log.Debug(ctx, "Hello World")
	log.Info(ctx, "Hello World")
	log.Warn(ctx, "Hello World")
	log.Error(ctx, "Hello World")
}
