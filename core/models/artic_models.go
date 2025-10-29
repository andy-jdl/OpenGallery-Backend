package models

import (
	"fmt"
	"strings"
)

type ArticMetadata struct {
	Data   []ArticData `json:"data"`
	Config ArticConfig `json:"config"`
}

type ArticData struct {
	ID               int      `json:"id"`
	Title            string   `json:"title"`
	ShortDescription string   `json:"short_description"`
	ArtistTitles     []string `json:"artist_titles"`
	ImageID          string   `json:"image_id"`
	Color            *Color   `json:"color,omitempty"`
}

type ArticInfo struct {
	LicenseLinks string `json:"license_links,omitempty"`
}

type ArticConfig struct {
	IiifURL string `json:"iiif_url"`
}

type Color struct {
	Hue        float32 `json:"h"`
	Light      float32 `json:"l"`
	Saturation float32 `json:"s"`
}

func (a *ArticData) ExtractFields() (*string, string) {
	return a.ExtractColor(), a.ExtractArtists()
}

func (a *ArticData) ExtractColor() *string {
	if a.Color == nil {
		empty := ""
		return &empty
	}
	colorStr := fmt.Sprintf("%f, %f, %f", a.Color.Hue, a.Color.Light, a.Color.Saturation)
	return &colorStr
}

func (a *ArticData) ExtractArtists() string {
	return strings.Join(a.ArtistTitles, ", ")
}
