package main

import (
	"go-dummy-clean-arch/infrastructure/connection"
	"go-dummy-clean-arch/infrastructure/queue"
	"go-dummy-clean-arch/infrastructure/repository"
	"go-dummy-clean-arch/usecase/artist"
	"log"
	"os"
)

func main() {
	topic := os.Getenv("TOPIC")

	// connecting databases
	connections, err := connection.OpenConnections()
	if err != nil {
		log.Fatal("error connecting databases", err)
	}

	// repositories
	artistRepo := repository.NewArtistMysql(connections.MySQL)

	// services
	artistService := artist.NewService(artistRepo)

	artistConsumer := queue.NewKafkaConsumer(topic, artistService)
	artistConsumer.Consume()
}
