package models

import "strings"

type LouvreMetadata struct {
	ObjectNumber    []LouvreObjectNumber `json:"objectNumber"`
	ID              string               `json:"arkId"`
	Title           string               `json:"title"`
	Image           []LouvreImage        `json:"image"`
	CurrentLocation string               `json:"currentLocation"`
	Creator         []LouvreCreators     `json:"creator"`
	Description     string               `json:"description"`
	Related         string               `json:"url"`
	Museum          string
	MuseumWebsite   string
}

type LouvreObjectNumber struct {
	Value string `json:"value"`
}

type LouvreCreators struct {
	Label string `json:"label"`
}

type LouvreImage struct {
	URLImage  string `json:"urlImage"`
	Copyright string `json:"copyright"`
}

func (l *LouvreMetadata) ExtractArtists() string {
	var names []string
	for _, c := range l.Creator {
		names = append(names, c.Label)
	}
	return strings.Join(names, ", ")
}

func (l *LouvreMetadata) ExtractObjectID() string {
	return l.ObjectNumber[0].Value
}

func (l *LouvreMetadata) ExtractImage() (string, string) {
	return l.Image[0].URLImage, l.Image[0].Copyright
}
