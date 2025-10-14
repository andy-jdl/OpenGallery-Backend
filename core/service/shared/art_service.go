package shared

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ArtService struct {
	client  *http.Client
	baseURL string
}

func NewArtServiceClient(client *http.Client, baseURL string) *ArtService {
	return &ArtService{
		client:  client,
		baseURL: baseURL,
	}
}

func (c *ArtService) GetJSON(ctx context.Context, path string, model any) error {
	url := fmt.Sprintf("%s%s", c.baseURL, path)
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	response, err := c.client.Do(request)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(model)
}
