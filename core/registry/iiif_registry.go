package registry

import (
	"api/core/models"
	"context"
	"fmt"
)

type IIIFProvider interface {
	FetchRawArtwork(ctx context.Context) (any, error)
	NormalizeMetadata(metadata any) ([]models.ArtworkMetadata, error)
	BuildResponse(metadata models.ArtworkMetadata) (models.ArtworkResponse, error)
	BuildIIIFImageURL(iiifURL, imageID string) string
	Name() string
}

type IIIFRegistry struct {
	providers map[string]IIIFProvider
}

func NewIIIFRegistry(providers ...IIIFProvider) *IIIFRegistry {
	reg := &IIIFRegistry{providers: make(map[string]IIIFProvider)}
	for _, provider := range providers {
		reg.providers[provider.Name()] = provider
	}
	return reg
}

func (r *IIIFRegistry) GetProvider(name string) (IIIFProvider, error) {
	p, ok := r.providers[name]
	if !ok {
		return nil, fmt.Errorf("IIIF provider not found for: %s", name)
	}
	return p, nil
}
