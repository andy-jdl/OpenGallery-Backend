package controllers

import (
	"api/core/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

// How do you communicate progress
func (ac *ArtWorkController) GetRandomArtwork(c *gin.Context) {
	source := c.Param("source")

	artWork, err := ac.repository.GetRandomArtwork(c, source)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, artWork)
}
