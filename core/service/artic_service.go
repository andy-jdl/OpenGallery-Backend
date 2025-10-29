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

type ArticService struct {
	*shared.ArtService
}

func NewArticService(client *http.Client, baseURL string) *ArticService {
	return &ArticService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
	}
}

func (s *ArticService) FetchRawArtwork(ctx context.Context) (any, error) {
	totalPages := 6095
	page := rand.Intn(totalPages)
	path := fmt.Sprintf("?page=%d&query[term][is_public_domain]=true&limit=10&fields=short_description,id,title,artist_titles,image_id,color", page)

	var data models.ArticMetadata
	if err := s.GetJSON(ctx, path, &data); err != nil {
		return models.ArtworkMetadata{}, err
	}

	return data, nil
}

func (s *ArticService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	data, ok := metadata.(models.ArticMetadata)
	if !ok {
		return []models.ArtworkMetadata{}, &core.InvalidMetadata{Code: 422, Message: "ArticService"}
	}

	var result []models.ArtworkMetadata
	for _, item := range data.Data {

		color, artists := item.ExtractFields()

		meta := models.ArtworkMetadata{
			ID:          strconv.Itoa(item.ID),
			ImageID:     item.ImageID,
			Title:       item.Title,
			Artist:      artists,
			Description: item.ShortDescription,
			Colors:      models.ColorSpectrum{Profile: "hls", Palette: *color},
			IIIFURL:     data.Config.IiifURL,
			Museum:      "Art Institute of Chicago",
			MuseumURL:   "https://www.artic.edu",
			Attribution: "Courtesy of The Art Institute of Chicago",
		}
		result = append(result, meta)
	}

	return result, nil
}

func (s *ArticService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		Description: m.Description,
		ImageID:     m.ImageID,
		ImageURL:    s.BuildIIIFImageURL(m.IIIFURL, m.ImageID),
		Colors:      m.Colors,
		Museum:      m.Museum,
		Attribution: m.Attribution,
		City:        "Chicago",
	}, nil
}

// TODO appropriate width and sizes
func (s *ArticService) BuildIIIFImageURL(iiifUrl, imageId string) string {
	imageSize := "full/600,/0/default.jpg"
	return fmt.Sprintf("%s/%s/%s", iiifUrl, imageId, imageSize)
}

func (s *ArticService) Name() string {
	return "aic"
}
