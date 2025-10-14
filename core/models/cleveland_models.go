package models

import "strings"

type ClevelandArtMetadata struct {
	Data []ClevelandArtData `json:"data"`
}

type ClevelandArtData struct {
	ID              int               `json:"id"`
	AccessionNumber string            `json:"accession_number"`
	Title           string            `json:"title"`
	DidYouKnow      string            `json:"did_you_know"`
	CurrentLocation string            `json:"current_location"`
	Creators        []CreatorMetadata `json:"creators"`
	Images          ImageSet          `json:"images"`
}

type CreatorMetadata struct {
	Description            string `json:"description"`
	NameInOriginalLanguage string `json:"name_in_original_language"`
}

func (cm *ClevelandArtData) ExtractFields() string {
	var artists []string

	for _, creators := range cm.Creators {
		if creators.NameInOriginalLanguage != "" {
			artists = append(artists, creators.NameInOriginalLanguage)
		} else if creators.Description != "" {
			artists = append(artists, creators.Description)
		} else {
			artists = append(artists, "Unknown Artist")
		}
	}

	return strings.Join(artists, ",")
}

type ImageSet struct {
	Web ImageMetadata `json:"web"`
}

type ImageMetadata struct {
	URL    string `json:"url"`
	Width  string `json:"width"`
	Height string `json:"height"`
}
