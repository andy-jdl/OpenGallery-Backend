package service

import (
	"api/core/models"
	"api/core/service/shared"
	core "api/core/utils"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
)

type ClevelandService struct {
	*shared.ArtService
}

func NewClevelandService(client *http.Client, baseURL string) *ClevelandService {
	return &ClevelandService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
	}
}

func (c *ClevelandService) FetchRawArtwork(ctx context.Context) (any, error) {
	totalAssets := 41562
	limit := 10

	totalPages := totalAssets / limit
	randomPage := rand.Intn(totalPages)

	skip := randomPage * limit
	path := fmt.Sprintf("?skip=%d&limit=%d&cc0=true&fields=id,title,images,creators,current_location,did_you_know", skip, limit)

	var data models.ClevelandArtMetadata
	if err := c.GetJSON(ctx, path, &data); err != nil {
		return models.ArtworkMetadata{}, err
	}

	return data, nil
}

func (c *ClevelandService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	data, ok := metadata.(models.ClevelandArtMetadata)

	if !ok {
		return []models.ArtworkMetadata{}, &core.InvalidMetadata{Code: 422, Message: "ClevelandService"}
	}

	var result []models.ArtworkMetadata
	for _, item := range data.Data {

		artists := item.ExtractFields()

		meta := models.ArtworkMetadata{
			ID:          strconv.Itoa(item.ID),
			ImageID:     item.AccessionNumber,
			Title:       item.Title,
			Artist:      artists,
			Description: item.DidYouKnow,
			ImageURL:    item.Images.Web.URL,
			Museum:      "Cleveland Museum of Art",
			MuseumURL:   "https://www.clevelandart.org/home",
			Gallery:     item.CurrentLocation,
			Attribution: "Courtesy of the The Cleveland Museum of Art",
		}

		result = append(result, meta)
	}

	return result, nil
}

func (c *ClevelandService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		Description: m.Description,
		ImageID:     m.ImageID,
		ImageURL:    m.ImageURL,
		Museum:      m.Museum,
		Attribution: "Courtesy of the The Cleveland Museum of Art",
	}, nil
}

func (c *ClevelandService) BuildIIIFImageURL(iiifURL string, imageID string) string {
	panic("unimplemented")
}

func (c *ClevelandService) Name() string {
	return "clv"
}
