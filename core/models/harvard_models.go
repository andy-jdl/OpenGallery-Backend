package models

import (
	"slices"
	"strings"
)

type HarvardRecords struct {
	Records []HarvardMetadata `json:"records"`
}

type HarvardMetadata struct {
	ID              int      `json:"id"`
	ObjectID        int      `json:"objectid"`
	Title           string   `json:"title"`
	Copyright       string   `json:"copyright"`
	Description     string   `json:"description"`
	PrimaryImageURL string   `json:"primaryimageurl"`
	People          []People `json:"people"`
	Colors          []Colors `json:"colors"`
}

type People struct {
	Name string `json:"name"`
}

type Colors struct {
	Color    string  `json:"color"`
	Spectrum string  `json:"spectrum"`
	Hue      string  `json:"hue"`
	Percent  float32 `json:"percent"`
	Css3     string  `json:"css3"`
}

func (h *HarvardMetadata) ExtractFields() (string, string) {
	return h.ExtractArtists(), h.ExtractColor()
}

func (h *HarvardMetadata) ExtractColor() string {
	var colors []string
	for _, css3 := range h.Colors {
		colors = append(colors, css3.Css3)
	}
	slices.Sort(colors)
	color := slices.Compact(colors)
	return strings.Join(color, ", ")
}

func (h *HarvardMetadata) ExtractArtists() string {
	var names []string
	for _, name := range h.People {
		names = append(names, name.Name)
	}
	return strings.Join(names, ", ")
}
