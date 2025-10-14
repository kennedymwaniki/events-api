package main

import (
	"log"

	"github.com/kennedymwaniki/events-api/internal/api"
	"github.com/kennedymwaniki/events-api/internal/config"
)

func main() {
	appconfig := config.InitializeDependencies()
	router := api.NewRouter(appconfig)


	log.Println("Starting server on :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}