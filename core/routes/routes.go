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

// RESPONSE YOU GET WHEN YOU HAVE THE WRONG URL
// {
//   "id": 0,
//   "title": "",
//   "image_id": "",
//   "iiif_url": "",
//   "image_url": "//full/843,/0/default.jpg"
// }

// Notes from Artic API docs
/***
Requests are throttled to 60 requests per minute
You should only use public domain images from a legal perspective
Pagination: Listing shows 12 records per page by default
	params
	page: specify a page of results
	limit: 0 - 12 (but can go up to 100)
	Maybe contact the team to see what they think of this use case
	example: https://api.artic.edu/api/v1/artworks?page=2&limit=50
***/

// Normalize and make dynamic all APIs to return varying pieces of work
// Ensure they are public domain, give credit to all regardless if asked or not.

func RegisterRoutes(r *gin.Engine, cfg *config.Config) {
	client := &http.Client{Timeout: 10 * time.Second}
	cache := internal.NewCache[[]models.ArtworkMetadata](time.Minute * 120)

	registry := registry.NewIIIFRegistry(
		service.NewArticService(client, cfg.APIs.ArticBaseURL),
		service.NewClevelandService(client, cfg.APIs.ClevelandBaseURL),
		service.NewGettyService(client, cfg.APIs.GettyBaseURL),
		service.NewMetService(client, cfg.APIs.MetBaseURL),
		service.NewHarvardService(client, cfg.APIs.HarvardBaseURL.HarvardBaseURL, cfg.APIs.HarvardBaseURL.APIKey),
	)

	repository := repository.NewArtworkRepository(cache, registry)
	controller := controllers.NewArtworkController(repository)

	r.GET("/api/v1/artworks/:source/random", controller.GetRandomArtwork)
}

//Take reviews and turn into a personal profile
