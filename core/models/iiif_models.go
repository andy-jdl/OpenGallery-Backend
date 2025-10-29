package models

import "time"

type ArtworkResponse struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Artist      string        `json:"artist"`
	ImageID     string        `json:"image_id"`
	ImageURL    string        `json:"image_url"`
	Museum      string        `json:"museum"`
	Attribution string        `json:"attribution"`
	Colors      ColorSpectrum `json:"colors"`
	Related     string        `json:"related"`
	Copyright   string        `json:"copyright"`
	Description string        `json:"description"`
	City        string        `json:"city"`
}

type ArtworkMetadata struct {
	ID          string        `json:"id"`
	ImageID     string        `json:"image_id"`
	Title       string        `json:"title"`
	Artist      string        `json:"artist"`
	Date        string        `json:"date"`
	Medium      string        `json:"medium"`
	Dimensions  string        `json:"dimensions"`
	ImageURL    string        `json:"imageUrl"`
	IIIFURL     string        `json:"iiif_url"`
	Gallery     string        `json:"gallery"`
	Museum      string        `json:"museum"`
	MuseumURL   string        `json:"museumUrl"`
	Description string        `json:"description,omitempty"`
	Related     string        `json:"related"`
	Colors      ColorSpectrum `json:"colors"`
	Attribution string        `json:"attribution"`
	Copyright   string        `json:"copyright"`
	Explore     string        `json:"explore"`
}

type ArtworkBatch struct {
	Museum   string            `json:"museum"`
	Count    int               `json:"count"`
	Artworks []ArtworkResponse `json:"artworks"`
	Fetched  time.Time         `json:"fetched"`
}

type ColorSpectrum struct {
	Profile string `json:"profile"`
	Palette string `json:"palette"`
}
