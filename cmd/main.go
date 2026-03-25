package main

import (
	"kafka_module_2/internal/app"
	"log"
)

func main() {

	application, err := app.New()
	if err != nil {
		log.Fatal("Failed to create Application: ", err)
		return
	}

	err = application.Start()
	if err != nil {
		log.Fatal("Failed to start Application: ", err)
		return
	}
}
