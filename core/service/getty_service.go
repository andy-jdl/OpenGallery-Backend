package service

import (
	"api/core/models"
	"api/core/service/shared"
	"context"
	"fmt"
	"net/http"
)

type GettyService struct {
	*shared.ArtService
}

func NewGettyService(client *http.Client, baseURL string) *GettyService {
	return &GettyService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
	}
}

func (g *GettyService) FetchRawArtwork(ctx context.Context) (any, error) {
	// Debug path for now while I figure out the query for Getty
	path := "iiif/manifest/8116191d-5f05-4763-bcf6-87dfcd0bd178"

	var data models.GettyManifest
	if err := g.GetJSON(ctx, path, &data); err != nil {
		return models.ArtworkMetadata{}, err
	}

	return data, nil
}

func (g *GettyService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	// This will change when you have more entries
	manifest, ok := metadata.(models.GettyManifest)
	if !ok {
		return []models.ArtworkMetadata{}, fmt.Errorf("invalid metadata type for GettyManifest")
	}

	artist, title, accession := manifest.ExtractFields()

	var result []models.ArtworkMetadata
	meta := models.ArtworkMetadata{
		ID:          accession,
		ImageID:     accession,
		Title:       title,
		Artist:      artist,
		Description: manifest.Description[0],
		ImageURL:    manifest.Thumbnail.ID,
		Museum:      "The Getty",
		MuseumURL:   "https://www.getty.edu/",
		Attribution: "Courtesy of the J. Paul Getty Museum, Los Angeles",
		Related:     manifest.Related,
	}
	result = append(result, meta)
	return result, nil
}

func (g *GettyService) BuildResponse(m models.ArtworkMetadata) (models.ArtworkResponse, error) {
	return models.ArtworkResponse{
		ID:          m.ID,
		Title:       m.Title,
		Artist:      m.Artist,
		Description: m.Description,
		ImageID:     m.ImageID,
		ImageURL:    m.ImageURL,
		Museum:      m.Museum,
		Related:     m.Related,
		Attribution: "Courtesy of the J. Paul Getty Museum, Los Angeles",
	}, nil
}

func (g *GettyService) BuildIIIFImageURL(iiifURL string, imageID string) string {
	panic("unimplemented")
}

func (g *GettyService) Name() string {
	return "gty"
}
