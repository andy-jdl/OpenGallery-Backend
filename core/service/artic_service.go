package service

import (
	"api/core/models"
	"api/core/service/shared"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
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
	path := fmt.Sprintf("?page=%d&query[term][is_public_domain]=true&limit=10&fields=short_description,id,title,artist_titles,image_id", page)

	var data models.ArticMetadata
	if err := s.GetJSON(ctx, path, &data); err != nil {
		return models.ArtworkMetadata{}, err
	}

	return data, nil
}

func (s *ArticService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	data, ok := metadata.(models.ArticMetadata)
	if !ok {
		return []models.ArtworkMetadata{}, fmt.Errorf("invalid metadata type for ArticService")
	}

	var result []models.ArtworkMetadata
	for _, item := range data.Data {
		// TODO make sure you're getting all artists
		var artists string
		if len(item.ArtistTitles) > 0 {
			artists = strings.Join(item.ArtistTitles, ", ")
		}

		meta := models.ArtworkMetadata{
			ID:          strconv.Itoa(item.ID),
			ImageID:     item.ImageID,
			Title:       item.Title,
			Artist:      artists,
			Description: item.ShortDescription,
			IIIFURL:     data.Config.IiifURL,
			Museum:      "Art Institute of Chicago",
			MuseumURL:   "https://www.artic.edu",
			Attribution: "Courtesy of the Art Institute of Chicago",
		}
		result = append(result, meta)
	}

	return result, nil
}

func (s *ArticService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:            m.ID,
		Title:         m.Title,
		ArtistDisplay: m.Artist,
		Description:   m.Description,
		ImageID:       m.ImageID,
		ImageURL:      s.BuildIIIFImageURL(m.IIIFURL, m.ImageID),
		Museum:        m.Museum,
		Attribution:   "Courtesy of the Art Institute of Chicago",
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
