package routes

import (
	"api/core/config"
	"api/core/controllers"
	"api/core/internal"
	"api/core/models"
	"api/core/registry"
	"api/core/repository"
	"api/core/service"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	client := &http.Client{Timeout: 10 * time.Second}
	cache := internal.NewCache[[]models.ArtworkMetadata](time.Minute * 120)

	registry := registry.NewIIIFRegistry(
		service.NewArticService(client, cfg.APIs.ArticBaseURL),
		service.NewClevelandService(client, cfg.APIs.ClevelandBaseURL),
		service.NewMetService(client, cfg.APIs.MetBaseURL),
		service.NewHarvardService(client, cfg.APIs.HarvardBaseURL.HarvardBaseURL, cfg.APIs.HarvardBaseURL.APIKey),
		service.NewLouvreService(client, cfg.APIs.LouvreBaseURL),
	)

	repository := repository.NewArtworkRepository(cache, registry)
	controller := controllers.NewArtworkController(repository)

	r.GET("/api/v1/artworks/:source/random", controller.GetRandomArtworkWithSource)
	r.GET("/api/v1/artworks/random", controller.GetRandomArtwork)
}

//Take reviews and turn into a personal profile
