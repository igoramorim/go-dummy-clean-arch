package main

import (
	"go-dummy-clean-arch/api"
	"go-dummy-clean-arch/infrastructure/connection"
	"log"
	"os"
)

func main() {
	// load some config

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// connecting databases
	connectionsDB, err := connection.OpenConnections()
	if err != nil {
		log.Fatal("error connecting databases", err)
	}

	// map routes
	router := api.Mapping(connectionsDB)

	if err = router.Run(":" + port); err != nil {
		log.Fatal("error running server", err)
	}
}
