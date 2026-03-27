package main

import (
	"kafka_module_2/internal/app"
	l "kafka_module_2/internal/pkg/logger"
)

func main() {

	logger := l.New()

	application, err := app.New(logger)
	if err != nil {
		logger.Error("Failed to create Application: %s", err)
		return
	}

	err = application.Start()
	if err != nil {
		logger.Error("Failed to start Application: %s", err)
		return
	}
}
