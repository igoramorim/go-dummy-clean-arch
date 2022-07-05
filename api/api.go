package api

import (
	"github.com/gin-gonic/gin"
	"go-dummy-clean-arch/api/handler"
	"go-dummy-clean-arch/api/middleware"
	"go-dummy-clean-arch/infrastructure/client"
	"go-dummy-clean-arch/infrastructure/connection"
	"go-dummy-clean-arch/infrastructure/repository"
	"go-dummy-clean-arch/usecase/album"
	"go-dummy-clean-arch/usecase/artist"
)

func Mapping(connections *connection.DBConnections) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.HandleError())

	// repositories
	albumRepo := client.NewSpotifyAlbumClient()
	//albumRepo := client.NewVagalumeAlbumCLient()
	//albumRepo := repository.NewAlbumMysql(connections.MySQL)

	artistRepo := repository.NewArtistMysql(connections.MySQL)
	//artistRepo := repository.NewArtistMemoryKVS()

	// services
	albumService := album.NewService(albumRepo)
	artistService := artist.NewService(artistRepo)

	// routes
	handler.MapHealthCheckRoutes(router)
	handler.MapAlbumRoutes(albumService, router)
	handler.MapArtistRoutes(artistService, router)

	return router
}
