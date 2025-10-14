package models

import "strings"

type MetIds struct {
	ObjectIDs []int `json:"objectIDs"`
}

type MetMetadata struct {
	ObjectID          int            `json:"objectID"`
	Title             string         `json:"title"`
	Constituents      []Constituents `json:"constituents"`
	AccessionNumber   string         `json:"accessionNumber"`
	IsPublicDomain    bool           `json:"isPublicDomain"`
	PrimaryImageSmall string         `json:"primaryImageSmall"`
	Related           string         `json:"objectURL"`
}

type Constituents struct {
	Name string `json:"name"`
}

func (mm *MetMetadata) ExtractFields() string {
	var names []string
	for _, c := range mm.Constituents {
		names = append(names, c.Name)
	}
	return strings.Join(names, ", ")
}
