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

type HarvardService struct {
	*shared.ArtService
	APIKey string
}

func (h *HarvardService) BuildIIIFImageURL(iiifURL string, imageID string) string {
	panic("unimplemented")
}

func (h *HarvardService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		ImageID:     m.ImageID,
		ImageURL:    m.ImageURL,
		Museum:      m.Museum,
		Related:     m.Related,
		Attribution: "Courtesy of The Harvard Art Museums",
	}, nil
}

func (h *HarvardService) FetchRawArtwork(ctx context.Context) (any, error) {
	apiKey := h.APIKey
	totalPages := 24782
	randomPage := rand.Intn(totalPages)
	fields := "objectid,copyright,title,primaryimageurl,people,id,colors"
	objectPath := fmt.Sprintf("/object?fields=%s&hasimage:1&page=%d&apikey=%s", fields, randomPage, apiKey)

	var data models.HarvardRecords
	if err := h.GetJSON(ctx, objectPath, &data); err != nil {
		return models.ArtworkMetadata{}, err
	}

	return data, nil
}

func (h *HarvardService) Name() string {
	return "hrv"
}

func (h *HarvardService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	records, ok := metadata.(models.HarvardRecords)
	if !ok {
		return []models.ArtworkMetadata{}, &core.InvalidMetadata{Code: 422, Message: "HarvardService"}
	}
	data := records.Records

	var result []models.ArtworkMetadata
	for _, item := range data {
		artists, color := item.ExtractFields()
		meta := models.ArtworkMetadata{
			ID:          strconv.Itoa(item.ID),
			ImageID:     strconv.Itoa(item.ObjectID),
			Title:       item.Title,
			Artist:      artists,
			ImageURL:    item.PrimaryImageURL,
			Related:     color,
			Museum:      "Harvard Art Museum",
			MuseumURL:   "https://harvardartmuseums.org/",
			Attribution: "Courtesy of The Harvard Art Museums",
		}
		result = append(result, meta)
	}
	return result, nil
}

func NewHarvardService(client *http.Client, baseURL string, apiKey string) *HarvardService {
	return &HarvardService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
		APIKey:     apiKey,
	}
}
