package service

import (
	"api/core/models"
	"api/core/service/shared"
	core "api/core/utils"
	"context"
	"fmt"
	"math/rand/v2"
	"net/http"
)

type LouvreService struct {
	*shared.ArtService
	data []string
}

func (l *LouvreService) BuildIIIFImageURL(iiifURL string, imageID string) string {
	panic("unimplemented")
}

func (l *LouvreService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		ImageID:     m.ImageID,
		ImageURL:    m.ImageURL,
		Description: m.Description,
		Museum:      m.Museum,
		Copyright:   m.Copyright,
		Attribution: m.Attribution,
	}, nil
}

func (l *LouvreService) FetchRawArtwork(ctx context.Context) (any, error) {
	limit := 10
	arkIDs := append([]string(nil), l.data...)
	rand.Shuffle(len(arkIDs), func(i, j int) {
		arkIDs[i], arkIDs[j] = arkIDs[j], arkIDs[i]
	})

	selectedArkIds := arkIDs[:min(limit, len(arkIDs))]

	var results []models.LouvreMetadata
	for _, arkId := range selectedArkIds {
		objectPath := fmt.Sprintf("%s.json", arkId)
		var object models.LouvreMetadata
		if err := l.GetJSON(ctx, objectPath, &object); err != nil {
			continue
		}
		results = append(results, object)
	}

	return results, nil
}

func (l *LouvreService) Name() string {
	return "lve"
}

func (l *LouvreService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	data, ok := metadata.([]models.LouvreMetadata)
	if !ok {
		return []models.ArtworkMetadata{}, &core.InvalidMetadata{Code: 422, Message: "LouvreService"}
	}

	var result []models.ArtworkMetadata
	for _, item := range data {
		artists := item.ExtractArtists()
		objectID := item.ExtractObjectID()
		imageURL, copyright := item.ExtractImage()

		meta := models.ArtworkMetadata{
			ID:          objectID,
			ImageID:     item.ID,
			Title:       item.Title,
			Artist:      artists,
			ImageURL:    imageURL,
			Description: item.Description,
			Related:     item.Related,
			Museum:      "The Louvre",
			Copyright:   copyright,
			Attribution: "Courtesy of The Louvre",
		}
		result = append(result, meta)
	}
	return result, nil
}

func NewLouvreService(client *http.Client, baseURL string) *LouvreService {
	data := core.Flatten(core.GetCSVDataFromFile("core/internal/data/louvre_data.csv"))
	return &LouvreService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
		data:       data,
	}
}
