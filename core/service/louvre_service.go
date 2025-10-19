package service

import (
	"api/core/models"
	"api/core/service/shared"
	"context"
	"encoding/csv"
	"net/http"
	"os"
	"path/filepath"
)

type LouvreService struct {
	*shared.ArtService
}

func (l *LouvreService) BuildIIIFImageURL(iiifURL string, imageID string) string {
	panic("unimplemented")
}

func (l *LouvreService) BuildResponse(metadata models.ArtworkMetadata) (models.ArtworkResponse, error) {
	panic("unimplemented")
}

func (l *LouvreService) FetchRawArtwork(ctx context.Context) (any, error) {
	arks := l.GetArkDataFromFile()
	return nil, nil
}

func (l *LouvreService) GetArkDataFromFile() [][]string {
	path, _ := filepath.Abs("core/internal/data/louvre_data.csv")
	file, _ := os.Open(path)

	defer file.Close()
	csvReader := csv.NewReader(file)
	data, _ := csvReader.ReadAll()
	return data
}

func (l *LouvreService) Name() string {
	return "lve"
}

func (l *LouvreService) NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error) {
	panic("unimplemented")
}

func NewLouvreService(client *http.Client, baseURL string) *LouvreService {
	return &LouvreService{
		ArtService: shared.NewArtServiceClient(client, baseURL),
	}
}
