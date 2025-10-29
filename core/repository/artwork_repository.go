package repository

import (
	"api/core/internal"
	"api/core/models"
	"api/core/registry"
	"context"
	"fmt"
	"math/rand"
)

type ArtworkRepository struct {
	cache    *internal.Cache[[]models.ArtworkMetadata]
	registry *registry.IIIFRegistry
}

func NewArtworkRepository(cache *internal.Cache[[]models.ArtworkMetadata], registry *registry.IIIFRegistry) *ArtworkRepository {
	return &ArtworkRepository{
		cache:    cache,
		registry: registry,
	}
}

func (ar *ArtworkRepository) FetchArtwork(ctx context.Context, source string) error {

	service, err := ar.registry.GetProvider(source)
	if err != nil {
		return fmt.Errorf("service not found for source: %s", source)
	}

	raw, err := service.FetchRawArtwork(ctx)
	if err != nil {
		return err
	}

	normalized, err := service.NormalizeMetadata(raw)
	if err != nil {
		return fmt.Errorf("failed to normalize metadata %w", err)
	}

	ar.cache.Set(source, normalized)
	return nil
}

func (ar *ArtworkRepository) GetRandomArtwork(ctx context.Context, source string) (models.ArtworkResponse, error) {
	// call the service given source
	service, err := ar.registry.GetProvider(source)
	if err != nil {
		return models.ArtworkResponse{}, fmt.Errorf("unknown source: %s", source)
	}

	// Check cache first
	artworks, ok := ar.cache.Get(source)
	if !ok || len(artworks) == 0 {
		// make a new call to hydrate the cache
		if err := ar.FetchArtwork(ctx, source); err != nil {
			return models.ArtworkResponse{}, fmt.Errorf("failed to hydrate cache: %w", err)
		}
		artworks, _ = ar.cache.Get(source)
	}

	// select artwork from cache at random
	idx := rand.Intn(len(artworks))
	selected := artworks[idx]

	normalized, err := service.BuildResponse(selected)
	if err != nil {
		return models.ArtworkResponse{}, fmt.Errorf("failed to build response: %w", err)
	}

	return normalized, nil
}
