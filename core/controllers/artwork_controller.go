package controllers

import (
	"api/core/repository"
	"fmt"
	"math/rand"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	providers    = []string{"aic", "met", "clv", "hrv"}
	lastProvider string
	mu           sync.Mutex
)

type ArtWorkController struct {
	repository *repository.ArtworkRepository
}

func NewArtworkController(repository *repository.ArtworkRepository) *ArtWorkController {
	return &ArtWorkController{repository: repository}
}

func (ac *ArtWorkController) FetchArtWork(c *gin.Context) {
	source := c.Param("source")

	err := ac.repository.FetchArtwork(c, source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Artworks from '%s' fetched and cached successfully", source),
	})
}

func (ac *ArtWorkController) GetRandomArtworkWithSource(c *gin.Context) {
	source := c.Param("source")
	artWork, err := ac.repository.GetRandomArtwork(c, source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, artWork)
}

func (ac *ArtWorkController) GetRandomArtwork(c *gin.Context) {
	source := GetRandomProvider()
	artWork, err := ac.repository.GetRandomArtwork(c, source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"source": source, "artwork": artWork})
}

func GetRandomProvider() string {
	mu.Lock()
	defer mu.Unlock()

	var provider string
	for {
		provider = providers[rand.Intn(len(providers))]
		if lastProvider != provider {
			break
		}
	}

	lastProvider = provider
	return provider
}
