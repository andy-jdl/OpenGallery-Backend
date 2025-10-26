package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Port string
	APIs APIConfig
}

type APIConfig struct {
	ArticBaseURL     string
	MetBaseURL       string
	HarvardBaseURL   HarvardConfig
	GettyBaseURL     string
	LouvreBaseURL    string
	ClevelandBaseURL string
}

type HarvardConfig struct {
	HarvardBaseURL string
	APIKey         string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		APIs: APIConfig{
			MetBaseURL: "https://collectionapi.metmuseum.org/public/collection/v1/",
			HarvardBaseURL: HarvardConfig{
				HarvardBaseURL: "https://api.harvardartmuseums.org",
				APIKey:         os.Getenv("HARVARD_API_KEY"),
			},
			ArticBaseURL:     "https://api.artic.edu/api/v1/artworks/",
			ClevelandBaseURL: "https://openaccess-api.clevelandart.org/api/artworks/",
			LouvreBaseURL:    "https://collections.louvre.fr/en/ark:/53355/",
		},
	}

	return cfg, nil
}
