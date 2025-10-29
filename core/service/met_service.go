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

type MetService struct {
	*shared.ArtService
	idCache []int
}

func (m *MetService) BuildIIIFImageURL(iiifURL string, imageID string) string {
	panic("unimplemented")
}

func (ms *MetService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		ImageID:     m.ImageID,
		ImageURL:    m.ImageURL,
		Museum:      m.Museum,
		Related:     m.Related,
		Attribution: m.Attribution,
		City:        "New York",
	}, nil
}

func (m *MetService) FetchRawArtwork(ctx context.Context) (any, error) {
	const limit = 10
	queries := "q=hasImages=true"
	searchPath := fmt.Sprintf("search?%s", queries)

	if len(m.idCache) == 0 {
		var ids models.MetIds
		if err := m.GetJSON(ctx, searchPath, &ids); err != nil {
			return nil, err
		}
		m.idCache = ids.ObjectIDs
	}

	rand.Shuffle(len(m.idCache), func(i, j int) {
		m.idCache[i], m.idCache[j] = m.idCache[j], m.idCache[i]
	})

	selectedObjectIds := m.idCache[:min(limit, len(m.idCache))]

	var results []models.MetMetadata
	for _, objectId := range selectedObjectIds {
		objectPath := fmt.Sprintf("objects/%d", objectId)
		var object models.MetMetadata
		if err := m.GetJSON(ctx, objectPath, &object); err != nil {
			continue
		}
		results = append(results, object)
	}

	return results, nil
}

func (m *MetService) Name() string {
	return "met"
}

func (m *MetService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	data, ok := metadata.([]models.MetMetadata)
	if !ok {
		return []models.ArtworkMetadata{}, &core.InvalidMetadata{Code: 422, Message: "MetService"}
	}

	var result []models.ArtworkMetadata
	for _, item := range data {

		artists := item.ExtractFields()

		meta := models.ArtworkMetadata{
			ID:          strconv.Itoa(item.ObjectID),
			Title:       item.Title,
			Artist:      artists,
			ImageID:     item.AccessionNumber,
			ImageURL:    item.PrimaryImageSmall,
			Related:     item.Related,
			Museum:      "The Met",
			MuseumURL:   "https://www.metmuseum.org/",
			Attribution: "Courtesy of The Metropolitan Museum of Art, New York",
		}
		result = append(result, meta)
	}
	return result, nil
}

func NewMetService(client *http.Client, baseURL string) *MetService {
	return &MetService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
	}
}
