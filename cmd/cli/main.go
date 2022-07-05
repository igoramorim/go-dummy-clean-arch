package main

import (
	"context"
	"flag"
	"fmt"
	"go-dummy-clean-arch/infrastructure/connection"
	"go-dummy-clean-arch/infrastructure/repository"
	"go-dummy-clean-arch/usecase/artist"
	"log"
)

func main() {
	artistId := flag.Int64("artistId", 1, "the id of the artist")
	flag.Parse()

	// connecting databases
	connections, err := connection.OpenConnections()
	if err != nil {
		log.Fatal("error connecting databases", err)
	}

	// repositories
	artistRepo := repository.NewArtistMysql(connections.MySQL)

	// services
	artistService := artist.NewService(artistRepo)

	artist, err := artistService.GetByID(context.Background(), *artistId)
	if err != nil {
		log.Fatalf("%v", err)
	}

	fmt.Println(artist)
}
