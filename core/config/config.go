package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	DB   PostGresConfig
	Port string
	APIs APIConfig
}

type APIConfig struct {
	ArticBaseURL     string
	MetBaseURL       string
	HarvardBaseURL   HarvardConfig
	GettyBaseURL     string
	ClevelandBaseURL string
}

type HarvardConfig struct {
	HarvardBaseURL string
	APIKey         string
}

type PostGresConfig struct {
	Username string
	Password string
	Name     string
	Port     string
	Host     string
}

func LoadConfig() (*Config, error) {
	cfg := &Config{
		DB: PostGresConfig{
			Username: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
		},
		Port: os.Getenv("PORT"),
		APIs: APIConfig{
			MetBaseURL:   "https://collectionapi.metmuseum.org/public/collection/v1/",
			GettyBaseURL: "https://media.getty.edu/iiif/manifest",
			HarvardBaseURL: HarvardConfig{
				HarvardBaseURL: "https://api.harvardartmuseums.org",
				APIKey:         os.Getenv("HARVARD_API_KEY"),
			},
			ArticBaseURL:     "https://api.artic.edu/api/v1/artworks/",
			ClevelandBaseURL: "https://openaccess-api.clevelandart.org/api/artworks/",
		},
	}

	return cfg, nil
}
